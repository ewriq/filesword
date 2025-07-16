const express = require("express");
const http = require("http");
const socketIo = require("socket.io");
const net = require("net");
const fs = require("fs");
const path = require("path");
const app = express();
const server = http.createServer(app);
const io = socketIo(server, { cors: { origin: "*" } });

const TCP_SERVER_HOST = "127.0.0.1";
const TCP_SERVER_PORT = 9000;
const FILE_FOLDER = path.join(__dirname, "file");

const tcpClient = new net.Socket();
let tcpConnected = false;
let buffer = "";
let currentRequestId = 0;
const pendingResponses = new Map();

if (!fs.existsSync(FILE_FOLDER)) fs.mkdirSync(FILE_FOLDER);


function connectTcp() {
  tcpClient.connect(TCP_SERVER_PORT, TCP_SERVER_HOST, () => {
    tcpConnected = true;
  });

  tcpClient.on("data", handleTcpData);
  tcpClient.on("error", () => (tcpConnected = false));
  tcpClient.on("close", () => (tcpConnected = false));
}

function handleTcpData(data) {
  buffer += data.toString();
  const parts = buffer.split("\n");

  while (parts.length > 1) {
    const raw = parts.shift();
    if (!raw.trim()) continue;

    let msg;
    try {
      msg = JSON.parse(raw);
    } catch {
      continue;
    }

    const pending = pendingResponses.get(msg.requestId);
    if (!pending) continue;

    pendingResponses.delete(msg.requestId);
    if (msg.status === "OK") handleApprovedMeta(pending);
    else handleDeniedMeta(pending);
  }

  buffer = parts[0];
}

function handleApprovedMeta({ socket, meta, tempPath, resolve, reject }) {
  const finalPath = path.join(FILE_FOLDER, meta.fileName);
  try {
    fs.copyFileSync(tempPath, finalPath);
    fs.chmodSync(finalPath, meta.mode);
    fs.utimesSync(finalPath, new Date(meta.atimeMs), new Date(meta.mtimeMs));
    fs.unlinkSync(tempPath);
    socket.emit("upload-success", meta.fileName);
    resolve("OK");
  } catch (err) {
    socket.emit("upload-error", err.message);
    reject(err);
  }
}

function handleDeniedMeta({ socket, meta, tempPath, reject }) {
  fs.unlinkSync(tempPath);
  socket.emit("upload-error", "TCP reddetti");
  reject(new Error("Reddedildi"));
}

function sendMetaToTcp(meta, socket, tempPath) {
  return new Promise((resolve, reject) => {
    if (!tcpConnected) return reject(new Error("TCP bağlantı yok"));

    const requestId = ++currentRequestId;
    meta.requestId = requestId;

    pendingResponses.set(requestId, { socket, meta, tempPath, resolve, reject });
    tcpClient.write(JSON.stringify(meta) + "\n");
  });
}

function handleFileUpload(socket, payload) {
  const { fileName, mode, atimeMs, mtimeMs, fileData } = payload;
  if (!fileName || !fileData) {
    socket.emit("upload-error", "eksik veri");
    return;
  }

  const buffer = Buffer.from(fileData, "base64");
  const tempPath = path.join(FILE_FOLDER, "temp_" + fileName);
  fs.writeFileSync(tempPath, buffer);

  const meta = { fileName, mode, atimeMs, mtimeMs };
  sendMetaToTcp(meta, socket, tempPath).catch(() => {});
}

function setupWebSocket() {
  io.on("connection", (socket) => {
    socket.on("upload-file", (payload) => handleFileUpload(socket, payload));
  });
}

connectTcp();
setupWebSocket();

server.listen(3000);
