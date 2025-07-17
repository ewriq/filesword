<script>
  import { onMount } from "svelte";
  import { io } from "socket.io-client";

  import FileUpload from "../components/add.svelte";
  import FileList from "../components/list.svelte";

  let socket;
  let showFileUpload = false;

  onMount(() => {
    socket = io("http://localhost:3000");
  });

  function toggleFileUpload() {
    showFileUpload = !showFileUpload;
  }
</script>

<div class="max-w-[900px] mx-auto p-6 flex flex-col gap-6">
  <button
    class="self-start flex items-center gap-2 px-4 py-2 bg-zinc-600 text-white rounded-full shadow-lg hover:bg-zinc-900 transition"
    on:click={toggleFileUpload}
    aria-expanded={showFileUpload}
    aria-controls="file-upload-section"
  >
    <span class="text-lg font-bold">+</span>
    <span class="text-sm font-medium">Dosya Ekle</span>
  </button>

  <div class="flex gap-6 w-full min-h-[320px] transition-all duration-500">
    {#if showFileUpload}
      <section
        id="file-upload-section"
        class="flex-[0_0_40%] p-6 bg-gray-50 rounded-lg shadow-lg transition-all duration-500"
        tabindex="-1"
      >
        <FileUpload {socket} />
      </section>
    {/if}

    <section
      class={`p-6 bg-white rounded-lg shadow-lg overflow-auto max-h-[400px] transition-all duration-500 ${
        showFileUpload ? "flex-[0_0_60%]" : "flex-1"
      }`}
    >
      <FileList />
    </section>
  </div>
</div>
