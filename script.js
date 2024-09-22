document.getElementById('textForm').addEventListener('submit', async function (event) {
    event.preventDefault();

    const inputText = document.getElementById('textInput').value;

    // Sending the input text to the Go backend via a POST request
    try {
        const response = await fetch('http://localhost:6969/parse-text', {
            method: 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body: inputText // Sending plain text
        });

        const data = await response.text();
        document.getElementById('textOutput').innerText = data; // Display the parsed response
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('textOutput').innerText = "Error communicating with the server.";
    }
});

// async function fetchGoString() {
//     try {
//         const response = await fetch('http://localhost:6969/get-string');
//         const text = await response.text();
//         document.getElementById('result').innerText += "\n" + text;
//     } catch (error) {
//         console.error('Error fetching the string:', error);
//     }
// }
