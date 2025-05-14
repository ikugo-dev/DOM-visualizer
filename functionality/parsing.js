document.getElementById('text-form').addEventListener('submit', async function(event) {
    event.preventDefault();
    const inputText = document.getElementById('text-input').value;

    // Sending the input text to the Go backend via a POST request
    try {
        const response = await fetch('http://localhost:8080/parse-text', {
            method: 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body: inputText // Must be plain text
        });
        const data = await response.text();
        window.textOutput.innerText = data;
    } catch (error) {
        console.error('Error:', error);
        window.textOutput.innerText = "Error communicating with the server.";
    }
});
