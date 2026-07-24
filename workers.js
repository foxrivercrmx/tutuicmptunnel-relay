export default {
  async fetch(request, env, ctx) {
    if (request.method !== "POST") {
      return new Response("Method not allowed", { status: 405 });
    }

    const targetURL = "https://yourserver.com/path";

    const newRequest = new Request(targetURL, new Request(request));
    
    try {
      const response = await fetch(newRequest);
      return response;
    } catch (error) {
      return new Response("Relay failed: " + error.message, { status: 502 });
    }
  }
};