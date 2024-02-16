<script setup lang="ts">
const openGithub = () => {
    window.open("https://github.com/XljBear/HoneyHollow", "_blank")
}
</script>
<template>
    <div class="footer">
        <div>copyright © StupidBear Studio 2024</div>
        <div class="github">
            <el-button @click="openGithub" link>Project On Github</el-button>
        </div>
    </div>
</template>
<style scoped lang="scss">
@import '../styles/global.scss';

.footer {
    margin-top: 30px;
    color: $google-grey-700;
    font-size: 14px;
    font-weight: bold;

    .github {
        font-size: 12px;
        font-weight: normal;
        margin-top: 5px;
    }
}
</style>