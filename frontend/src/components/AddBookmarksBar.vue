<script setup lang="ts">
import { IceCream } from '@element-plus/icons-vue';
import { createBookmarks } from '../apis/bookmarks/request';
import { api } from '../apis/bookmarks/api';
import { ref } from 'vue';
import { ElMessage } from "element-plus";
import AddBookmarksDialog from './AddBookmarksDialog.vue';
const urlInput = ref("")
const emits = defineEmits(['onDirty'])
const inCreate = ref(false)
const doCreateBookmarks = () => {
    if (urlInput.value !== "") {
        inCreate.value = true
        var bookmarks: api.bookmarks.request.bookmarks = {
            url: urlInput.value,
            title: "",
            remark: "",
            sort: 1
        };
        createBookmarks(bookmarks).then(() => {
            urlInput.value = ""
            emits("onDirty")
            ElMessage({
                message: '吼吼，创建好咯～',
                type: 'success',
            })
        }).catch(() => {

        }).finally(() => {
            inCreate.value = false
        })
    } else {
        addBookmarksDialog.value?.openDialog()
    }
}
const addBookmarksDialog = ref<typeof AddBookmarksDialog>()
</script>
<template>
    <div class="addBookmarksBar">
        <el-card shadow="hover">
            <el-row :gutter="20">
                <el-col :xs="16" :sm="19">
                    <el-input :disabled="inCreate" v-model="urlInput" placeholder="这里写地址进行快捷创建！" />
                </el-col>
                <el-col :xs="8" :sm="5">
                    <div class="action">
                        <el-button :loading="inCreate" @click="doCreateBookmarks" :icon="IceCream"
                            type="primary">创建</el-button>
                    </div>
                </el-col>
            </el-row>
        </el-card>
    </div>
    <AddBookmarksDialog @on-dirty="emits('onDirty')" ref="addBookmarksDialog" :is-edit="false" />
</template>
<style scoped lang="scss">
.addBookmarksBar {
    width: 100%;
    margin-bottom: 10px;

    .action {
        display: flex;
        width: 100%;

        button {
            width: 100%;
        }
    }
}

@media only screen and (max-width: 600px) {
    .action {
        button {
            font-size: 16px;
        }
    }

}
</style>