<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { api } from '../apis/bookmarks/api';
import { getBookmarksList } from '../apis/bookmarks/request'
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
const refreshBookmarksData = () => {
  getBookmarksList(props.isViewed).then((resultList) => {
    listData.value = resultList
  })
}
onMounted(() => {
  refreshBookmarksData()
})
</script>

<template>
  <el-row class="maxWidth">
    <el-col v-for="(bookmarks, index) in listData" class="bookmarks" :span="24">
      <BookmarksBar :index="index + 1" :is-viewed="isViewed" :bookmarks="bookmarks" @on-dirty="onDirty" />
    </el-col>
    <div class="noRecord" v-if="listData.length === 0">
      - 暂无记录 -
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
