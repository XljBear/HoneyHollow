<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { api } from '../apis/bookmarks/api';
import { getBookmarksList, getBookmarksById } from '../apis/bookmarks/request'
import BookmarksBar from './BookmarksBar.vue';
const props = defineProps<{
  isViewed: boolean
}>()
const emits = defineEmits(['onDirty'])
const onDirty = () => {
  refreshBookmarksData()
  emits("onDirty")
}
const listData = ref<api.bookmarks.response.bookmarks[]>([])
const needReCheckList = ref<api.bookmarks.response.bookmarks[]>([])
const refreshTimer = ref(0)
const inLoad = ref(false)
const refreshBookmarksData = () => {
  inLoad.value = true
  if (refreshTimer.value !== 0) {
    clearTimeout(refreshTimer.value)
  }
  getBookmarksList(props.isViewed).then((resultList) => {
    listData.value = resultList
    listData.value.forEach((bookmarks: api.bookmarks.response.bookmarks) => {
      if (bookmarks.inProcess) {
        needReCheckList.value.push(bookmarks)
      }
    })
    refreshTimer.value = setTimeout(reCheckListData, 2000)
  }).catch(() => { }).finally(() => {
    inLoad.value = false
  })
}
const reCheckListData = () => {
  needReCheckList.value.forEach(async (bookmarks, index) => {
    await getBookmarksById(bookmarks.id).then((bookmarks) => {
      if (!bookmarks.inProcess) {
        for (var i = 0; i < listData.value.length; i++) {
          if (listData.value[i].id === bookmarks.id) {
            listData.value[i] = bookmarks
            break
          }
        }
        needReCheckList.value.splice(index, 1)
      }
    }).catch(() => {
      needReCheckList.value.splice(index, 1)
    })
  })
  if (needReCheckList.value.length > 0) {
    if (refreshTimer.value !== 0) {
      clearTimeout(refreshTimer.value)
    }
    refreshTimer.value = setTimeout(reCheckListData, 2000)
  }
}
onMounted(() => {
  refreshBookmarksData()
})
defineExpose({ refreshBookmarksData })
</script>

<template>
  <el-row class="maxWidth">
    <el-col v-for="(bookmarks, index) in listData" class="bookmarks" :span="24">
      <BookmarksBar :key="bookmarks.id" :index="index + 1" :is-viewed="isViewed" :bookmarks="bookmarks"
        @on-dirty="onDirty" />
    </el-col>
    <div class="noRecord" v-if="listData.length === 0">
      {{ inLoad ? "正在载入..." : "- 暂无记录 -" }}
    </div>
  </el-row>
</template>

<style scoped lang="scss">
@import '../styles/global.scss';

.maxWidth {
  width: 100%;

  .bookmarks+.bookmarks {
    margin-top: 20px;
  }
}

.noRecord {
  font-size: 20px;
  font-weight: bold;
  width: 100%;
  color: $google-grey-500;
}
</style>
