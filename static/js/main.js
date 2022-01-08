
function loadImage(element) {
    if (window.matchMedia &&
        window.matchMedia('(prefers-color-scheme: dark)').matches &&
        (element.dataset.darked === undefined || element.dataset.darked == 0)) {
        element.src = element.dataset.src.replace("/static/images/", "/static/images/dark/");
        element.dataset.darked = 1;
    } else {
        element.src = element.dataset.src;
    }
}
