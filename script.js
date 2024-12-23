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
            body: inputText // Sending plain text
        });

        const data = await response.text();
        document.getElementById('text-output').innerText = data; // Display the parsed response
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('text-output').innerText = "Error communicating with the server.";
    }
});


const textOutput = document.getElementById('text-output');
let scale = 1;
const zoomAmmount = 0.25;

document.getElementById('zoom-in').addEventListener('click', () => {
    scale = Math.min(scale + zoomAmmount, 5); // Max zoom level
    textOutput.style.transform = `scale(${scale})`;
});

document.getElementById('zoom-out').addEventListener('click', () => {
    scale = Math.max(scale - zoomAmmount, 0.5); // Min zoom level
    textOutput.style.transform = `scale(${scale})`;
});

document.getElementById('zoom-reset').addEventListener('click', () => {
    scale = 1;
    textOutput.style.transform = `scale(${scale})`;
});


// Panning functionality
let isPanning = false;
let startX = 0, startY = 0, offsetX = 0, offsetY = 0;
let mouseStartX, mouseStartY;
let currentPosX, currentPosY;

window.addEventListener('keydown', (e) => {
    if (e.code === 'Space') {
        e.preventDefault();
        mouseStartX = e.client
        isPanning = true;
    }
});

window.addEventListener('keyup', (e) => {
    if (e.code === 'Space') {
        isPanning = false;
    }
});

// Panning functionality

window.addEventListener('mousemove', (e) => {
    if (!isPanning) {
        mouseStartX = e.clientX;
        mouseStartY = e.clientY;
        currentPosX = textOutput.style.left;
        currentPosY = textOutput.style.top;
        return;
    }
    offsetX = currentPosX + e.clientX - mouseStartX;
    offsetY = currentPosY + e.clientY - mouseStartY;
    textOutput.style.left = `${offsetX}px`;
    textOutput.style.top = `${offsetY}px`;
});

// textOutput.addEventListener('mouseleave', () => {
//     isPanning = false;
// });
