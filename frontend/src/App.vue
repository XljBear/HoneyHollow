<script setup lang="ts">
import { ref, onMounted } from 'vue';

import BookmarksList from './components/BookmarksList.vue'
import { StarFilled, Finished } from '@element-plus/icons-vue';

import { getUnViewBookmarksCount } from './apis/bookmarks/request';

import AddBookmarksBar from './components/AddBookmarksBar.vue';
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';

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

const unViewBookmarksListRef = ref<typeof BookmarksList>()
const viewedBookmarksListRef = ref<typeof BookmarksList>()
const onBookmarksModify = () => {
  if (nowTabName.value !== 'unView') {
    nowTabName.value = 'unView'
  }
  refreshUnViewBookmarksCount()
  unViewBookmarksListRef.value?.refreshBookmarksData()
}
</script>

<template>
  <el-config-provider namespace="hh">
    <div>
      <Header />
      <AddBookmarksBar @on-dirty="onBookmarksModify" />
      <el-tabs v-model="nowTabName" type="border-card">
        <el-tab-pane name="unView">
          <template #label>
            <el-icon>
              <StarFilled />
            </el-icon>
            待浏览！({{ unViewBookmarksCount }})
          </template>
          <BookmarksList ref="unViewBookmarksListRef" @on-dirty="refreshUnViewBookmarksCount"
            v-if="nowTabName === 'unView'" :is-viewed="false" />
        </el-tab-pane>
        <el-tab-pane name="viewed">
          <template #label>
            <el-icon>
              <Finished />
            </el-icon>
            已浏览～
          </template>
          <BookmarksList ref="viewedBookmarksListRef" @on-dirty="refreshUnViewBookmarksCount"
            v-if="nowTabName === 'viewed'" :is-viewed="true" />
        </el-tab-pane>
      </el-tabs>
      <Footer />
    </div>
  </el-config-provider>
</template>

<style scoped lang="scss"></style>
