export default {
  async fetch(request, env, ctx) {
    if (request.method !== "POST") {
      return new Response("Method not allowed", { status: 405 });
    }

    // Daftar multi-tujuan (silakan sesuaikan URL server)
    const targetURLs = [
      "https://server1.com/path",
      "https://server2.com/path",
      "https://server3.com/path"
    ];

    try {
      // Mapping setiap URL menjadi sebuah fetch promise
      const fetchPromises = targetURLs.map(url => {
        // Body request HTTP hanya bisa dibaca sekali, sehingga kita harus menggunakan .clone()
        const newRequest = new Request(url, new Request(request.clone()));
        return fetch(newRequest);
      });

      // Eksekusi semua request secara paralel tanpa saling memblokir
      const results = await Promise.allSettled(fetchPromises);

      // Cek apakah ada request yang gagal secara jaringan (ditolak)
      const failed = results.filter(r => r.status === 'rejected');
      if (failed.length > 0) {
        return new Response(`Relay partially failed. ${failed.length} requests rejected.`, { status: 502 });
      }

      return new Response("Relay to multiple destinations initiated successfully", { status: 200 });

    } catch (error) {
      return new Response("Relay failed: " + error.message, { status: 502 });
    }
  }
};