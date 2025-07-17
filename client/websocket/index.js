const express = require("express");
const http = require("http");
const socketIo = require("socket.io");
const net = require("net");

const {
  handleTcpData,
  setTcpSocket,
  markTcpDisconnected,
  handleFileUpload, 
} = require("./handler/handler");

const app = express();
const server = http.createServer(app);
const io = socketIo(server, {
  cors: { origin: "*" },
  maxHttpBufferSize: 1e9,
});

const TCP_SERVER_HOST = "127.0.0.1";
const TCP_SERVER_PORT = 9000;

const tcpClient = new net.Socket();

function connectTcp() {
  tcpClient.connect(TCP_SERVER_PORT, TCP_SERVER_HOST, () => {
    setTcpSocket(tcpClient);
    console.log("✅ TCP bağlantısı kuruldu.");
  });

  tcpClient.on("data", handleTcpData);
}

function setupWebSocket() {
  io.on("connection", (socket) => {
    socket.on("upload-file", (payload) => handleFileUpload(socket, payload));
  });
}

connectTcp();
setupWebSocket();

server.listen(3000, () => {
  console.log("🚀 WebSocket sunucusu 3000 portunda çalışıyor.");
});