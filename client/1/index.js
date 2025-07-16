const net = require("net");
const readline = require("readline");

// Sunucu bilgileri
const HOST = "127.0.0.1";
const PORT = 9000;

// TCP bağlantısı oluştur
const client = new net.Socket();
client.connect(PORT, HOST, () => {
  console.log(`✅ Bağlantı kuruldu: ${HOST}:${PORT}`);
});

// Gelen mesajları ekrana yaz
client.on("data", (data) => {
  console.log("📩 Sunucudan gelen:", data.toString().trim());
});

// Bağlantı kapanırsa
client.on("close", () => {
  console.log("❌ Bağlantı kapandı.");
});

// Hata olursa
client.on("error", (err) => {
  console.error("⚠️ Hata:", err.message);
});

// Kullanıcıdan giriş almak için arayüz
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  prompt: "> "
});

rl.prompt();

rl.on("line", (line) => {
  client.write(line + "\n"); // Her satır yeni mesaj
  rl.prompt();
});
