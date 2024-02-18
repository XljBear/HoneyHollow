<script setup lang="ts">
import logo from '../assets/logo.png'
</script>
<template>
    <div class="header">
        <div class="title">
            <div :style="`background-image: url(${logo});`" class="logo"></div> Honey Hollow
        </div>
        <div class="subtitle">蠢熊的一个收藏树洞</div>
    </div>
</template>
<style scoped lang="scss">
@import '../styles/global.scss';

.header {
    margin-bottom: 30px;

    .title {
        color: $bilibili-blue;
        font-size: 40px;
        font-weight: bold;
        display: flex;
        align-items: center;
        justify-content: center;

        .logo {
            width: 80px;
            height: 80px;
            background-size: contain;
            background-repeat: no-repeat;
            margin-right: 20px;
        }
    }

    .subtitle {
        color: $google-grey-700;
        font-size: 16px;
        font-weight: bold;
    }
}

@media only screen and (max-width: 600px) {
    .header {
        .title {
            font-size: 30px;

            .logo {
                width: 60px;
                height: 60px;
            }
        }

        .subtitle {
            font-size: 12px;
        }
    }
}</style>