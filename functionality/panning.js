let isPanning = false;
let mouseStartX, mouseStartY;
let currentPosX = 0, currentPosY = 0;

window.addEventListener('keydown', (e) => {
    if (e.code === 'Space') {
        const form = document.getElementById('text-form');
        if (form.contains(document.activeElement)) {
            return;
        }
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
        currentPosX = parseInt(window.textOutput.style.left || "0", 10);
        currentPosY = parseInt(window.textOutput.style.top || "160", 10);
        return;
    }
    let offsetX = currentPosX + e.clientX - mouseStartX;
    let offsetY = currentPosY + e.clientY - mouseStartY;
    window.textOutput.style.left = `${offsetX}px`;
    window.textOutput.style.top = `${offsetY}px`;
});

window.textOutput.addEventListener('mouseleave', () => {
    isPanning = false;
});
