const express = require("express");
const http = require("http");
const socketIo = require("socket.io");
const path = require("path");
const fs = require("fs");
const cors = require("cors");

const app = express();
const server = http.createServer(app);

const io = socketIo(server, {
  maxHttpBufferSize: 1e9,
  cors: {
    origin: "http://localhost:5173",      
        methods: ["GET", "POST", "DELETE", "PUT"],
    credentials: true
  }
});


const websocket = require("./app/init");

const PORT = 4000;

app.use(cors({
  origin: "http://localhost:5173",
  methods: ["GET", "POST", "DELETE"],
  credentials: true
}));

app.use('/files', express.static(path.join(__dirname, './app/snap')));

app.get('/file', (req, res) => {
  fs.readdir(path.join(__dirname, './app/snap'), (err, files) => {
    if (err) {
      return res.status(500).json({ error: 'Snap klasörü okunamadı' });
    }
    res.json(files);
  });
});

app.delete("/file", (req, res) => {
  const fileName = req.query.name;
  if (!fileName) {
    return res.status(400).json({ error: "Dosya adı belirtilmeli" });
  }

  const filePath = path.join(path.join(__dirname, './app/snap'), fileName);

  fs.access(filePath, fs.constants.F_OK, (err) => {
    if (err) {
      return res.status(404).json({ error: "Dosya bulunamadı" });
    }


    fs.unlink(filePath, (err) => {
      if (err) {
        return res.status(500).json({ error: "Dosya silinemedi" });
      }

      res.json({ message: "Dosya başarıyla silindi" });
    });
  });
});

websocket.connectTcp();
websocket.setupWebSocket(io);

server.listen(PORT, () => {
  console.log(`🚀 Sunucu ${PORT} portunda çalışıyor.`);
});
