const net = require('net');
const { handleTcpData, setTcpSocket, markTcpDisconnected, handleFileUpload } = require('./handler/handler');

const TCP_SERVER_HOST = "127.0.0.1";
const TCP_SERVER_PORT = 9000;

const tcpClient = new net.Socket();

function connectTcp() {
  tcpClient.connect(TCP_SERVER_PORT, TCP_SERVER_HOST, () => {
    setTcpSocket(tcpClient);
    console.log("✅ TCP bağlantısı kuruldu.");
  });

  tcpClient.on("data", handleTcpData);
  tcpClient.on("error", () => markTcpDisconnected());
  tcpClient.on("close", () => markTcpDisconnected());
}

function setupWebSocket(io) {
  io.on("connection", (socket) => {
    socket.on("upload-file", (payload) => handleFileUpload(socket, payload));
  });
}

module.exports = {
  connectTcp,
  setupWebSocket,
};
