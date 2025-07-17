<script>
    import axios from "axios";

    export let fileName = "";

    let loading = false;
    let message = "";
    let error = "";

    async function deleteFile() {
        if (!fileName) {
            error = "Dosya adı boş olamaz";
            return;
        }
        loading = true;
        message = "";
        error = "";

        try {
            const res = await axios.delete("http://localhost:3000/file", {
                params: { name: fileName },
            });
            message = res.data.message || "Dosya başarıyla silindi";
            window.location.reload();
        } catch (e) {
            error = e.response?.data?.error || "Dosya silinemedi";
        } finally {
            loading = false;
        }

    }
</script>

<div>
    <button on:click={deleteFile} disabled={loading || !fileName}>
        {loading ? "Siliniyor..." : "❌"}
    </button>

</div>
