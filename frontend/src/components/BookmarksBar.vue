<script setup lang="ts">
import { api } from '../apis/bookmarks/api';
import dayjs from "dayjs";
import { Delete, Link, Timer, StarFilled, EditPen } from '@element-plus/icons-vue';
import { viewBookmarks, unViewBookmarks, deleteBookmarks } from "../apis/bookmarks/request";
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import AddBookmarksDialog from './AddBookmarksDialog.vue';
const props = defineProps<{
    bookmarks: api.bookmarks.response.bookmarks
    isViewed: boolean
    index: number
}>()
const emits = defineEmits(['onDirty'])

const onBookmarksClick = () => {
    if (props.bookmarks.remark !== "") {
        ElMessageBox.alert(props.bookmarks.remark, '康康备注', {
            confirmButtonText: '好哩',
            showClose: false,
            callback: () => {
                doViewBookmarks()
            },
        })
    } else {
        doViewBookmarks()
    }
}
const doViewBookmarks = () => {
    window.open(props.bookmarks.url, "_blank")
    viewBookmarks(props.bookmarks.id).then(() => {
        emits("onDirty")
    }).catch(() => {

    })
}
const inUnView = ref(false)
const onUnViewBookmarksClick = (e: Event) => {
    inUnView.value = true
    unViewBookmarks(props.bookmarks.id).then(() => {
        emits("onDirty")
        ElMessage({
            message: '好嘞～',
            type: 'success',
        })
    }).catch(() => {

    }).finally(() => {
        inUnView.value = false
    })
    e.stopPropagation()
}
const stopPropagation = (e: Event) => {
    e.stopPropagation()
}
const inDelete = ref(false)
const doDeleteBookmarks = () => {
    inDelete.value = true
    deleteBookmarks(props.bookmarks.id).then(() => {
        emits("onDirty")
        ElMessage({
            message: '好了～没啦！',
            type: 'success',
        })
    }).catch(() => {

    }).finally(() => {
        inDelete.value = false
    })
}
const getRandomVal = (minVal: number, maxVal: number): number => {
    return Math.round(Math.random() * (maxVal - minVal)) + minVal;
}
const addBookmarksDialog = ref<typeof AddBookmarksDialog>()
const editBookmarks = (e: Event) => {
    addBookmarksDialog.value?.openDialog()
    e.stopPropagation()
}

const getSortLevelName = (level: number): string => {
    switch (level) {
        case -1:
            return "非常有趣"
        case 0:
            return "有意思"
        case 1:
            return "普通"
        default:
            return "??"
    }
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
                <div class="label">{{ bookmarks.title === "" ? bookmarks.url : bookmarks.title }}</div>
            </div>
            <div v-else class="title inProcess">
                <el-icon class="icon">
                    <Timer />
                </el-icon>
                获取页面信息中...
            </div>
            <div class="moreInfo">
                <span :class="['sortTag', 'level' + bookmarks.sort]">
                    {{ getSortLevelName(bookmarks.sort) }}
                </span>
                <span class="createTime">
                    {{ dayjs(bookmarks.createAt).format('YY年MM月DD日 HH时mm分') }}
                </span>
            </div>
        </div>
        <div class="action">
            <span>
                <el-tooltip content="编辑" placement="top">
                    <el-button :disabled="bookmarks.inProcess" @click="editBookmarks" color="#616161" :icon="EditPen"
                        circle></el-button>
                </el-tooltip>
            </span>
            <span v-if="isViewed">
                <el-tooltip content="丢回待浏览" placement="top">
                    <el-button :loading="inUnView" @click="onUnViewBookmarksClick" type="info" :icon="StarFilled"
                        circle></el-button>
                </el-tooltip>
            </span>
            <span>
                <el-popconfirm @confirm="doDeleteBookmarks" confirm-button-text="删除" cancel-button-text="取消" width="250"
                    title="要删除这条小小的收藏吗？">
                    <template #reference>
                        <span @click="stopPropagation">
                            <el-tooltip content="删除" placement="top">
                                <el-button :loading="inDelete" type="primary" :icon="Delete" circle></el-button>
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
    <AddBookmarksDialog @on-dirty="emits('onDirty')" ref="addBookmarksDialog" :is-edit="true" :bookmarks="bookmarks" />
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
            display: flex;
            align-items: center;
            margin-bottom: 5px;

            .label {
                color: $bilibili-pink;
                font-size: 20px;
                font-weight: bold;
                text-align: left;
            }

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

        .moreInfo {
            text-align: left;

            .sortTag {
                padding: 3px 6px;
                border-radius: 6px;
                margin-right: 8px;
                font-size: 12px;
                color: white;

                &.level-1 {
                    background-color: $bilibili-pink;
                }

                &.level0 {
                    background-color: $bilibili-blue;
                }

                &.level1 {
                    background-color: $google-grey-500;
                }
            }

            .createTime {
                color: $google-grey-500;
                text-align: left;
                font-size: 12px;
            }
        }
    }

    .action {
        display: flex;
        align-items: center;
        margin-left: 10px;

        span+span {
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

@media only screen and (max-width: 600px) {
    .block {
        padding: 3px 15px;

        .info {
            .title {
                .label {
                    font-size: 14px;
                    word-break: break-all;
                    overflow: hidden;
                    display: -webkit-box;
                    -webkit-line-clamp: 1;
                    -webkit-box-orient: vertical;
                }
            }

            .moreInfo {
                .sortTag {
                    padding: 3px 6px;
                    border-radius: 6px;
                    margin-right: 6px;
                    font-size: 10px;
                    color: white;
                }
            }
        }

        .action {
            span+span {
                margin-left: 6px;
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
            font-size: 10px;
            height: 50px;
            width: 50px;
            align-items: center;
            border-radius: 60px;
            user-select: none;
            -webkit-user-select: none;

            span {
                font-size: 10px;
                font-weight: bold;
            }
        }
    }
}
</style>