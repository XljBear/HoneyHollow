<script setup lang="ts">
import { api } from '../apis/bookmarks/api';
import dayjs from "dayjs";
import { Delete, Link, Timer, StarFilled, EditPen } from '@element-plus/icons-vue';
import { viewBookmarks, unViewBookmarks, deleteBookmarks } from "../apis/bookmarks/request";
import { ElMessage } from 'element-plus'

const props = defineProps<{
    bookmarks: api.bookmarks.response.bookmarks
    isViewed: boolean
    index: number
}>()
const emits = defineEmits(['onDirty'])
const onBookmarksClick = () => {
    window.open(props.bookmarks.url, "_blank")
    viewBookmarks(props.bookmarks.id).then(() => {
        emits("onDirty")
    })
}
const onUnViewBookmarksClick = (e: Event) => {
    unViewBookmarks(props.bookmarks.id).then(() => {
        emits("onDirty")
        ElMessage({
            message: '好嘞～',
            type: 'success',
        })
    })
    e.stopPropagation()
}
const stopPropagation = (e: Event) => {
    e.stopPropagation()
}
const doDeleteBookmarks = () => {
    deleteBookmarks(props.bookmarks.id).then(() => {
        emits("onDirty")
        ElMessage({
            message: '好了～没啦！',
            type: 'success',
        })
    })
}
const getRandomVal = (minVal: number, maxVal: number): number => {
    return Math.round(Math.random() * (maxVal - minVal)) + minVal;
}
</script>
<template>
    <div @click="onBookmarksClick" :class="['block', isViewed ? 'viewed' : '']">
        <div class="number">#{{ index }}</div>
        <div class="info">
            <div v-if="!bookmarks.inProcess" class="title">
                <div v-if="bookmarks.icon !== ''" class="icon" :style="`background-image: url(${bookmarks.icon});`"></div>
                <el-icon class="icon" v-else>
                    <Link />
                </el-icon>
                {{ bookmarks.title === "" ? bookmarks.url : bookmarks.title }}
            </div>
            <div v-else class="title inProcess">
                <el-icon class="icon">
                    <Timer />
                </el-icon>
                获取页面信息中...
            </div>
            <div class="remark">
                {{ dayjs(bookmarks.createAt).format('YYYY年MM月DD日 HH时mm分') }}
                <span v-if="bookmarks.remark !== ''"> - {{ bookmarks.remark }}</span>
            </div>
        </div>
        <div class="action">
            <span>
                <el-tooltip content="编辑" placement="top">
                    <el-button color="#616161" :icon="EditPen" circle></el-button>
                </el-tooltip>
            </span>
            <span v-if="isViewed">
                <el-tooltip content="丢回待浏览" placement="top">
                    <el-button @click="onUnViewBookmarksClick" type="info" :icon="StarFilled"
                        circle></el-button>
                </el-tooltip>
            </span>
            <span>
                <el-popconfirm @confirm="doDeleteBookmarks" confirm-button-text="删除" cancel-button-text="取消" width="250"
                    title="要删除这条小小的收藏吗？">
                    <template #reference>
                        <span @click="stopPropagation">
                            <el-tooltip content="删除" placement="top">
                                <el-button type="primary" :icon="Delete" circle></el-button>
                            </el-tooltip>
                        </span>
                    </template>
                </el-popconfirm>
            </span>
        </div>
        <div v-if="bookmarks.viewTime !== null" class="viewedTag"
            :style="`transform: rotate(${getRandomVal(10, 720)}deg);left:${getRandomVal(25, 75)}%;opacity:${getRandomVal(20, 70)}%;`">
            <span>阅</span>{{ dayjs(bookmarks.viewTime).format('MM月DD日 HH时mm分') }}
        </div>
    </div>
</template>
<style scoped lang="scss">
@import '../styles/global.scss';

.block {
    padding: 5px 20px;
    border-radius: 5px;
    border: 1px solid $google-grey-400;
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    position: relative;
    margin-left: 5px;

    &.viewed {
        .info {
            .title {
                color: $google-grey-500;
            }
        }
    }

    &:hover {
        border: 1px solid $bilibili-pink;

        .number {
            border: 1px dashed $bilibili-pink;
            border-left: none;
            border-right: none;
            color: $bilibili-pink;
        }
    }

    .number {
        position: absolute;
        left: -10px;
        font-size: 10px;
        top: 35%;
        width: 20px;
        background-color: white;
        color: $google-grey-400;
        border: 1px dashed $google-grey-400;
        border-left: none;
        border-right: none;
    }

    .info {
        .title {
            color: $bilibili-pink;
            font-size: 20px;
            font-weight: bold;
            text-align: left;
            display: flex;
            align-items: center;
            margin-bottom: 5px;

            .icon {
                width: 16px;
                height: 16px;
                background-size: contain;
                background-repeat: no-repeat;
                background-position: center center;
                margin-right: 8px;
            }

            &.inProcess {
                color: $google-grey-500;
                font-size: 15px;
            }
        }

        .remark {
            color: $google-grey-500;
            text-align: left;
            font-size: 12px;
        }
    }

    .action {
        display: flex;
        align-items: center;

        span + span {
            margin-left: 12px;
        }
    }

    .viewedTag {
        color: $bilibili-blue;
        display: flex;
        flex-direction: column;
        top: -10px;
        position: absolute;
        border: 3px solid $bilibili-blue;
        padding: 8px;
        font-size: 12px;
        height: 60px;
        width: 60px;
        align-items: center;
        border-radius: 60px;
        user-select: none;
        -webkit-user-select: none;

        span {
            font-size: 14px;
            font-weight: bold;
        }
    }
}
</style>