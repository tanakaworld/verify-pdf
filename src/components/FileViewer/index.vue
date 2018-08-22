<template>
    <div class="PDF">
        <button class="row" @click="open(generatePDFFileUrl(sample))">___ sample</button>
        <button class="row" @click="open(generatePDFFileUrl(sampleMac))">sampleMac</button>

        <button class="row" @click="open(generatePDFFileUrl(pdfLarge))">pdfLarge</button>
        <button class="row" @click="open(generatePDFFileUrl(pdfLargeMac))">pdfLargeMac</button>
        <button class="row" @click="open(generatePDFFileUrl(highResolution))">highResolution</button>
        <button class="row" @click="open(generatePDFFileUrl(kana))">kana</button>
        <button class="row" @click="open(generatePDFFileUrl(kanaWImg))">kanaWImg</button>
        <button class="row" @click="open(generatePDFFileUrl(kanaSingle))">kanaSingle</button>
        <button class="row" @click="open(generatePDFFileUrl(enSingle))">enSingle</button>
        <button class="row" @click="open(generatePDFFileUrl(enSingleSob))">enSingleSob</button>

        <hr>
        <button class="row" @click="generatePDFFileByPDFJS(sample)">sample</button>
        <button class="row" @click="generatePDFFileByPDFJS2(sample)">sample2</button>
        <button class="row" @click="generatePDFFileByPDFJS3(sample)">___ sample2</button>

        <button class="row" @click="generatePDFFileByPDFJS(pdfLarge)">pdfLarge</button>
        <button class="row" @click="generatePDFFileByPDFJS2(pdfLarge)">pdfLarge2</button>
        <canvas id="the-canvas"></canvas>

        <img class="row" height="200px" :src="'data:image/png;base64,' + imgDog"/>
        <img class="row" height="200px" :src="'data:image/png;base64,' + imgExLarge"/>
    </div>
</template>

<script>
    import sample from '../../base64/sources/sample.pdf.txt';
    import sampleMac from '../../base64/sources/sample.pdf.mac.txt';
    import pdfLarge from '../../base64/sources/pdf-large.pdf.txt';
    import pdfLargeMac from '../../base64/sources/pdf-large.pdf.mac.txt';
    import highResolution from '../../base64/sources/pdf-high-resolution.pdf.txt';
    import kana from '../../base64/sources/kana.pdf.txt';
    import kanaWImg from '../../base64/sources/kana_w_img.pdf.txt';
    import kanaSingle from '../../base64/sources/kana_single.pdf.txt';
    import enSingle from '../../base64/sources/en_single.pdf.txt';
    import enSingleSob from '../../base64/sources/en_single_sob.pdf.txt';

    import imgDog from '../../base64/sources/dog.jpg.txt';
    import imgExLarge from '../../base64/sources/ex-large.jpg.txt';

    // https://stackoverflow.com/questions/12092633/pdf-js-rendering-a-pdf-file-using-a-base64-file-source-instead-of-url
    // https://mozilla.github.io/pdf.js/examples/
    import PDFJS from 'pdfjs-dist';

    const BASE64_MARKER = ';base64,';
    const convertDataURIToBinary = (dataURI) => {
        const base64Index = dataURI.indexOf(BASE64_MARKER) + BASE64_MARKER.length;
        const base64 = dataURI.substring(base64Index);
        const raw = window.atob(base64);
        const rawLength = raw.length;
        const array = new Uint8Array(new ArrayBuffer(rawLength));

        for (let i = 0; i < rawLength; i++) {
            array[i] = raw.charCodeAt(i);
        }
        return array;
    };

    export default {
        name: "PDF",
        computed: {
            sample: () => sample,
            sampleMac: () => sampleMac,
            pdfLarge: () => pdfLarge,
            pdfLargeMac: () => pdfLargeMac,
            highResolution: () => highResolution,
            kana: () => kana,
            kanaWImg: () => kanaWImg,
            kanaSingle: () => kanaSingle,
            enSingle: () => enSingle,
            enSingleSob: () => enSingleSob,
            imgDog: () => imgDog,
            imgExLarge: () => imgExLarge
        },
        methods: {
            generatePDFFileUrl(base64String) {
                const file = new Blob([atob(base64String)], {
                    type: 'application/pdf'
                    // type: 'application/octet-stream'
                });
                const fileURL = URL.createObjectURL(file);
                // eslint-disable-next-line
                console.log(fileURL);
                return fileURL;
            },
            generatePDFFileByPDFJS(base64String) {
                const pdfAsDataUri = `data:application/pdf;base64,${base64String}`;
                const pdfAsArray = convertDataURIToBinary(pdfAsDataUri);
                PDFJS.getDocument(pdfAsArray).then((pdf) => {
                    return pdf.getPage(1);
                }).then((page) => {
                    const scale = 1.5;
                    const viewport = page.getViewport(scale);
                    const canvas = document.getElementById('the-canvas');
                    const context = canvas.getContext('2d');
                    canvas.height = viewport.height;
                    canvas.width = viewport.width;
                    const renderContext = {
                        canvasContext: context,
                        viewport: viewport
                    };
                    page.render(renderContext);
                });
            },
            // render to canvas
            generatePDFFileByPDFJS2(base64String) {
                // eslint-disable-next-line
                console.log(PDFJS);
                // eslint-disable-next-line
                // debugger;
                PDFJS.getDocument({data: atob(base64String)}).then((pdf) => {
                    return pdf.getPage(1);
                }).then((page) => {

                    // const pdfLinkService = new PDFJS.PDFLinkService({
                    //     externalLinkTarget: 2
                    // });
                    // eslint-disable-next-line
                    // console.log(pdfLinkService);
                    // eslint-disable-next-line
                    // debugger;

                    const scale = 1.5;
                    const viewport = page.getViewport(scale);
                    const canvas = document.getElementById('the-canvas');
                    const context = canvas.getContext('2d');
                    canvas.height = viewport.height;
                    canvas.width = viewport.width;
                    const renderContext = {
                        canvasContext: context,
                        viewport: viewport
                    };
                    page.render(renderContext);
                });
            },
            // open in new tab
            generatePDFFileByPDFJS3(base64String) {
                PDFJS.getDocument({data: atob(base64String)}).then((pdf) => {
                    return pdf.getPage(1);
                }).then((page) => {
                    const scale = 1.5;
                    const viewport = page.getViewport(scale);
                    const canvas = document.getElementById('the-canvas');
                    const context = canvas.getContext('2d');
                    canvas.height = viewport.height;
                    canvas.width = viewport.width;
                    const renderContext = {
                        canvasContext: context,
                        viewport: viewport
                    };
                    page.render(renderContext);
                });
            },
            open(url) {
                window.open(url);
            }
        }
    };
</script>

<style scoped>
    .row {
        margin-top: 24px;
        display: block;
    }

    .the-canvas {
        border: solid 1px black;
    }
</style>
