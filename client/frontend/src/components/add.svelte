<script>
  import { onMount } from "svelte";
  import { io } from "socket.io-client";
  import { Button } from "bits-ui";

  let socket;
  let status = "Bağlantı bekleniyor...";
  let selectedFile;
  import toast, { Toaster } from "svelte-french-toast";

  onMount(() => {
    socket = io("http://localhost:3000", {
      withCredentials: true,
      transports: ["websocket", "polling"],
    });

    socket.on("connect", () => {
      status = "Socket.IO bağlantısı açıldı";
      toast.success(status);
    });

    socket.on("upload-success", (msg) => {
      toast.success(msg);
      status = msg;
    });

    socket.on("upload-error", (msg) => {
      status = "Hata: " + msg;
      toast.error("Hata: " + msg);
    });

    socket.on("disconnect", () => {
      toast.error("Socket.IO bağlantısı kapandı.");
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
      const base64Data = reader.result.split(",")[1];
      const now = new Date();

      const meta = {
        fileName: selectedFile.name,
        mode: 0o644,
        atimeMs: now.getTime(),
        mtimeMs: now.getTime(),
        fileData: base64Data,
        username: "admin",
        password: "admin",
      };

      socket.emit("upload-file", meta);

      status = "Dosya gönderiliyor: " + selectedFile.name;
      toast.success(status);
      window.location.reload();
    };

    reader.onerror = () => {
      status = "Dosya okunamadı!";
      toast.success(status);
    };

    reader.readAsDataURL(selectedFile);
  }
</script>

<h2 class="text-xl font-semibold mb-4">Dosya Yükleme</h2>

<div class="flex flex-col gap-4 w-full">
  <input
    type="file"
    class="border border-zinc-300 p-2 rounded"
    on:change={(e) => (selectedFile = e.target.files[0])}
  />

  <button
    class="px-4 py-2 bg-zinc-600 hover:bg-zinc-900 text-white rounded shadow"
    on:click={sendFile}
  >
    Yükle
  </button>
</div>

<Toaster
  position="top-right"
  toastOptions={{
    style: {
      background: "#333",
      color: "#fff",
      fontSize: "18px",
      padding: "16px 24px",
      boxShadow: "0 4px 12px rgba(0,0,0,0.3)",
      borderRadius: "8px",
      minWidth: "300px",
      maxWidth: "400px",
    },
    success: {
      duration: 4000,
      theme: {
        primary: "green",
        secondary: "black",
      },
    },
    error: {
      duration: 5000,
      theme: {
        primary: "red",
        secondary: "black",
      },
    },
  }}
/>
