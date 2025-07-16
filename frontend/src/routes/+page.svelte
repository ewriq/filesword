<script>
    import { onMount } from "svelte";
    import { io } from "socket.io-client";
  
    let socket;
    let status = "Bağlantı bekleniyor...";
    let selectedFile;
  
    onMount(() => {
      socket = io("http://localhost:3000");
  
      socket.on("connect", () => {
        status = "Socket.IO bağlantısı açıldı";
      });
  
      socket.on("upload-success", (msg) => {
        status = msg;
      });
  
      socket.on("upload-error", (msg) => {
        status = "Hata: " + msg;
      });
  
      socket.on("disconnect", () => {
        status = "Socket.IO bağlantısı kapandı";
      });
    });
  
    function sendFile() {
      if (!selectedFile) {
        status = "Lütfen dosya seçin";
        return;
      }
  
      const reader = new FileReader();
  
      reader.onload = () => {
        const base64Data = reader.result.split(",")[1]; // data:<mime>;base64,... kısmını ayır
        const now = new Date();
  
        // Metadata olarak şu bilgileri gönderebiliriz:
        const meta = {
          fileName: selectedFile.name,
          mode: 0o644, // default unix izin (burada stat alamıyoruz tarayıcıda, sabit verdik)
          atimeMs: now.getTime(),
          mtimeMs: now.getTime(),
          fileData: base64Data,
        };
  
        socket.emit("upload-file", meta);
        status = "Dosya gönderiliyor: " + selectedFile.name;
      };
  
      reader.onerror = () => {
        status = "Dosya okunamadı!";
      };
  
      reader.readAsDataURL(selectedFile); // base64 için DataURL
    }
  </script>
  
  <h2>Dosya Yükleme</h2>
  <input type="file" on:change="{e => selectedFile = e.target.files[0]}" />
  <button on:click="{sendFile}">Gönder</button>
  
  <p>{status}</p>
  