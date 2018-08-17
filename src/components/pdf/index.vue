<template>
    <div class="PDF">
        <button @click="openPDF1">PDF1</button>
        <button @click="openPDF2">PDF2</button>
    </div>
</template>

<script>
    import pdf1 from './raw/pdf-light.txt';
    import pdf2 from './raw/pdf-heavy.txt';

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
            openPDF1() {
                window.open(generatePDFFileUrl(pdf1));
            },
            openPDF2() {
                window.open(generatePDFFileUrl(pdf2));
            }
        }
    };
</script>

<style scoped>

</style>
