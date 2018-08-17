<template>
    <div class="PDF">
        <button @click="onClickOpenPDF">Open PDF</button>
    </div>
</template>

<script>
    import pdfLight from './raw/pdf-light.txt';

    const generatePDFFileUrl = (base64String) => {
        const file = new Blob([atob(base64String)], {type: 'application/pdf'});
        const fileURL = URL.createObjectURL(file);
        // eslint-disable-next-line
        console.log(fileURL);
        return fileURL;
    };

    export default {
        name: "PDF",
        methods: {
            onClickOpenPDF() {
                window.open(generatePDFFileUrl(pdfLight));
            }
        }
    };
</script>

<style scoped>

</style>
