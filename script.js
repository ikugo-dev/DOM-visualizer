// BACKEND PARSING
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

// ZOOM CONTROLS
const textOutput = document.getElementById('text-output');
let scale = 1;
const zoomAmmount = 0.25;

document.getElementById('zoom-in').addEventListener('click', () => {
    scale = Math.min(scale + zoomAmmount, 5);
    textOutput.style.transform = `scale(${scale})`;
});

document.getElementById('zoom-out').addEventListener('click', () => {
    scale = Math.max(scale - zoomAmmount, 0.25);
    textOutput.style.transform = `scale(${scale})`;
});

document.getElementById('zoom-reset').addEventListener('click', () => {
    scale = 1;
    textOutput.style.transform = `scale(${scale})`;
});

let isPanning = false;
let mouseStartX, mouseStartY;
let currentPosX = 0, currentPosY = 0;

window.addEventListener('keydown', (e) => {
    if (e.code === 'Space') {
        e.preventDefault();
        isPanning = true;
    }
});

window.addEventListener('keyup', (e) => {
    if (e.code === 'Space') {
        e.preventDefault();
        isPanning = false;
    }
});

window.addEventListener('mousemove', (e) => {
    if (!isPanning) {
        mouseStartX = e.clientX;
        mouseStartY = e.clientY;
        currentPosX = parseInt(textOutput.style.left || "0", 10);
        currentPosY = parseInt(textOutput.style.top || "160", 10);
        return;
    }
    let offsetX = currentPosX + e.clientX - mouseStartX;
    let offsetY = currentPosY + e.clientY - mouseStartY;
    textOutput.style.left = `${offsetX}px`;
    textOutput.style.top = `${offsetY}px`;
});

textOutput.addEventListener('mouseleave', () => {
    isPanning = false;
});
