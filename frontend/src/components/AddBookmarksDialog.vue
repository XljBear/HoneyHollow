<script setup lang="ts">
import { ref, reactive, onMounted, watch } from 'vue';
import { api } from '../apis/bookmarks/api';
import { ElMessage } from "element-plus";
import { createBookmarks, updateBookmarks } from '../apis/bookmarks/request';
import type { FormRules, FormInstance } from 'element-plus'

const props = defineProps<{
    bookmarks?: api.bookmarks.response.bookmarks
    isEdit: boolean
}>()

const emits = defineEmits(['onDirty'])

const dialogOpen = ref(false)
interface BookmarksForm {
    id: string
    url: string
    title: string
    remark: string
    sort: number
}
const form = reactive<BookmarksForm>({
    id: "",
    url: "",
    title: "",
    remark: "",
    sort: 1
})
const openDialog = () => {
    dialogOpen.value = true
}
const closeDialog = () => {
    dialogOpen.value = false
}

const formRef = ref<FormInstance>()
const inSave = ref(false)
const confirmSubmit = async () => {
    formRef.value?.validate((valid) => {
        if (valid) {
            if (props.isEdit) {
                inSave.value = true
                updateBookmarks(form).then(() => {
                    formRef.value?.resetFields()
                    closeDialog()
                    ElMessage({
                        message: '好嘞，收藏更新成功～',
                        type: 'success',
                    })
                    emits("onDirty")
                }).catch(() => {
                }).finally(() => {
                    inSave.value = false
                })
            } else {
                inSave.value = true
                createBookmarks(form).then(() => {
                    formRef.value?.resetFields()
                    closeDialog()
                    ElMessage({
                        message: '吼吼，创建好咯～',
                        type: 'success',
                    })
                    emits("onDirty")
                }).catch(() => {
                }).finally(() => {
                    inSave.value = false
                })
            }
        }
    })
}
onMounted(() => {
    if (props.isEdit) {
        form.id = props.bookmarks?.id || ""
        form.url = props.bookmarks?.url || ""
        form.title = props.bookmarks?.title || ""
        form.remark = props.bookmarks?.remark || ""
        form.sort = props.bookmarks?.sort || 1
    }
})
watch(() => props.bookmarks, () => {
    form.id = props.bookmarks?.id || ""
    form.url = props.bookmarks?.url || ""
    form.title = props.bookmarks?.title || ""
    form.remark = props.bookmarks?.remark || ""
    form.sort = props.bookmarks?.sort || 1
})
const rules = reactive<FormRules<BookmarksForm>>({
    url: [
        { required: true, message: '请输入网址', trigger: 'blur' },
    ],
})
defineExpose({ openDialog })
</script>
<template>
    <el-dialog v-model="dialogOpen" :title="isEdit ? '修改收藏' : '创建收藏'" width="500">
        <el-form ref="formRef" :model="form" :rules="rules">
            <el-form-item label="网址：" prop="url" :label-width="68">
                <el-input placeholder="输入需要添加收藏的网址" v-model="form.url" autocomplete="off" />
            </el-form-item>
            <el-form-item label="标题：" :label-width="60">
                <el-input v-model="form.title" autocomplete="off" />
            </el-form-item>
            <el-form-item label="备注：" :label-width="60">
                <el-input v-model="form.remark" autocomplete="off" />
            </el-form-item>
            <el-form-item label="优先级：" :label-width="70">
                <el-select v-model="form.sort" placeholder="请选择">
                    <el-option label="普通" :value="1" />
                    <el-option label="有意思" :value="0" />
                    <el-option label="非常有意思" :value="-1" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
                <el-button :disabled="inSave" @click="closeDialog">取消</el-button>
                <el-button :loading="inSave" type="primary" @click="confirmSubmit()">
                    {{ isEdit ? "更新" : "创建" }}
                </el-button>
            </div>
        </template>
    </el-dialog>
</template>
<style scoped lang="scss"></style>