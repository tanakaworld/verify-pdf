<template>
    <div class="PDF">
        <button class="row" @click="openPDFInNewTab(pdfSmall)">Open PDF Small</button>
        <button class="row" @click="openPDFInNewTab(pdfLarge)">Open PDF Large</button>
        <button class="row" @click="openImgInNewTab(imgDog)">Open imgDog</button>
        <button class="row" @click="openImgInNewTab(imgExLarge)">Open imgExLarge</button>

        <div id="canvasContainer"></div>
    </div>
</template>

<script>
    import pdfSmall from '../../base64/sources/sample.pdf.txt';
    import pdfLarge from '../../base64/sources/pdf-large.pdf.txt';

    import imgDog from '../../base64/sources/dog.jpg.txt';
    import imgExLarge from '../../base64/sources/ex-large.jpg.txt';

    // https://stackoverflow.com/questions/12092633/pdf-js-rendering-a-pdf-file-using-a-base64-file-source-instead-of-url
    // https://mozilla.github.io/pdf.js/examples/
    import PDFJS from 'pdfjs-dist';

    export default {
        name: "PDF",
        computed: {
            pdfSmall: () => pdfSmall,
            pdfLarge: () => pdfLarge,
            imgDog: () => imgDog,
            imgExLarge: () => imgExLarge
        },
        methods: {
            openImgInNewTab(base64Str) {
                const image = new Image();
                image.src = `data:image/png;base64,${base64Str}`;
                image.style['max-width'] = '100%';

                this.openNewTabWith(image);
            },
            openPDFInNewTab(base64Str) {
                const options = options || {scale: 1};
                const canvasContainer = document.createElement('div');

                function renderPage(page) {
                    const viewport = page.getViewport(options.scale);
                    const canvas = document.createElement('canvas');
                    const ctx = canvas.getContext('2d');
                    const renderContext = {
                        canvasContext: ctx,
                        viewport: viewport
                    };

                    canvas.height = viewport.height;
                    canvas.width = viewport.width;
                    canvas.style.display = 'block';
                    canvasContainer.appendChild(canvas);

                    page.render(renderContext);
                }

                function renderPages(pdfDoc) {
                    for (let num = 1; num <= pdfDoc.numPages; num++) {
                        pdfDoc.getPage(num).then(renderPage);
                    }
                }

                PDFJS.disableWorker = true;
                PDFJS.getDocument({data: atob(base64Str)})
                    .then(renderPages)
                    .then(() => this.openNewTabWith(canvasContainer));
            },
            openNewTabWith(content) {
                const win = window.open("");
                win.document.body.appendChild(content);
            }
        }
    };
</script>

<style scoped>
    .row {
        margin-top: 24px;
        display: block;
    }
</style>
