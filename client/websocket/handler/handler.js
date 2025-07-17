const net = require("net");
const fs = require("fs");
const path = require("path");
const { FILE_FOLDER, VALID_USERNAME, VALID_PASSWORD } = require("./shared");

let tcpBuffer = ""; 
let currentRequestId = 0;
const pendingResponses = new Map();
let tcpClient;
let tcpConnected = false;


function handleDeniedMeta({ socket, meta, tempPath, reject }) {
  fs.unlinkSync(tempPath);
  socket.emit("upload-error", "TCP reddetti");
  reject(new Error("Reddedildi"));
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


function sendMetaToTcp(meta, socket, tempPath) {
  return new Promise((resolve, reject) => {
    if (!tcpConnected) return reject(new Error("TCP bağlantı yok"));

    const requestId = ++currentRequestId;
    meta.requestId = requestId;

    pendingResponses.set(requestId, { socket, meta, tempPath, resolve, reject });
    tcpClient.write(JSON.stringify(meta) + "\n");
  });
}

// --- INDEX.JS TARAFINDAN KULLANILACAK DIŞA AKTARILAN FONKSİYONLAR ---

function handleFileUpload(socket, payload) {
  const { username, password, fileName, mode, atimeMs, mtimeMs, fileData } = payload;

  if (username !== VALID_USERNAME || password !== VALID_PASSWORD) {
    socket.emit("upload-error", "giriş reddedildi");
    return;
  }

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

function handleTcpData(data) {
  tcpBuffer += data.toString();
  const parts = tcpBuffer.split("\n");

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

  tcpBuffer = parts[0];
}


function setTcpSocket(socket) {
  tcpClient = socket;
  tcpConnected = true;
}


function markTcpDisconnected() {
  tcpConnected = false;
}

module.exports = {
  handleFileUpload,
  handleTcpData,
  setTcpSocket,
  markTcpDisconnected,
};