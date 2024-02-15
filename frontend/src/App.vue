<script setup lang="ts">
import { ref, onMounted } from 'vue';

import BookmarksList from './components/BookmarksList.vue'
import { StarFilled, Finished } from '@element-plus/icons-vue';

import { getUnViewBookmarksCount } from './apis/bookmarks/request';

const nowTabName = ref("unView")
const unViewBookmarksCount = ref(0)
const refreshUnViewBookmarksCount = () => {
  getUnViewBookmarksCount().then((num: number) => {
    unViewBookmarksCount.value = num
  })
}
onMounted(() => {
  refreshUnViewBookmarksCount()
})

</script>

<template>
  <el-config-provider namespace="hh">
    <div>
      <el-tabs v-model="nowTabName" type="border-card">
        <el-tab-pane name="unView">
          <template #label>
            <el-icon>
              <StarFilled />
            </el-icon>
            待浏览！({{ unViewBookmarksCount }})
          </template>
          <BookmarksList @on-dirty="refreshUnViewBookmarksCount" v-if="nowTabName === 'unView'" :is-viewed="false" />
        </el-tab-pane>
        <el-tab-pane name="viewed">
          <template #label>
            <el-icon>
              <Finished />
            </el-icon>
            已浏览～
          </template>
          <BookmarksList @on-dirty="refreshUnViewBookmarksCount" v-if="nowTabName === 'viewed'" :is-viewed="true" />
        </el-tab-pane>
      </el-tabs>
    </div>
  </el-config-provider>
</template>

<style scoped lang="scss"></style>
