<script>
    import axios from "axios";
    import { onMount } from "svelte";
    import Delete from "./delete.svelte";

    let files = [];
    let loading = true;
    let error = null;

    onMount(async () => {
        try {
            const res = await axios.get("http://localhost:3000/file");
            files = res.data;
        } catch (e) {
            error = "Dosyalar alınamadı";
        } finally {
            loading = false;
        }
    });
</script>

{#if loading}
    <p class="text-gray-600 italic">Yükleniyor...</p>
{:else if error}
    <p class="text-red-600 font-semibold">{error}</p>
{:else}
    <div class="overflow-x-auto rounded-lg shadow-lg border border-gray-200">
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-100">
                <tr>
                    <th
                        scope="col"
                        class="px-6 py-3 text-left text-xs font-medium text-gray-700 uppercase tracking-wider"
                    >
                        Dosya Adı
                    </th>
                    <th
                        scope="col"
                        class="px-6 py-3 text-center text-xs font-medium text-gray-700 uppercase tracking-wider"
                        style="width: 100px;"
                    >
                        İşlem
                    </th>
                </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-100">
                {#each files as fileName (fileName)}
                    <tr>
                        <td
                            class="px-6 py-4 whitespace-nowrap text-sm text-gray-900"
                        >
                            {fileName}
                        </td>
                        <td class="px-6 py-4 whitespace-nowrap text-center">
                            <Delete {fileName} />
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
{/if}
