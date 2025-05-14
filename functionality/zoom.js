let scale = 1;
const zoomAmmount = 0.25;

document.getElementById('zoom-in').addEventListener('click', () => {
    scale = Math.min(scale + zoomAmmount, 5);
    window.textOutput.style.transform = `scale(${scale})`;
});

document.getElementById('zoom-out').addEventListener('click', () => {
    scale = Math.max(scale - zoomAmmount, 0.25);
    window.textOutput.style.transform = `scale(${scale})`;
});

document.getElementById('zoom-reset').addEventListener('click', () => {
    scale = 1;
    window.textOutput.style.transform = `scale(${scale})`;
});
