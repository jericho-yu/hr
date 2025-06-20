<template>
  <q-layout view="hHh lpR fFf">
    <q-header bordered class="bg-primary text-white" height-hint="98">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="leftDrawerOpen = !leftDrawerOpen" />

        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg" />
          </q-avatar>
          人事：🤠💩2V
        </q-toolbar-title>

        <!-- <q-btn dense flat round icon="menu" @click="rightDrawerOpen = !rightDrawerOpen" /> -->
      </q-toolbar>

      <q-tabs align="center"> </q-tabs>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" side="left" overlay bordered>
      <q-list link dense separator style="min-width: 200px">
        <q-item dense clickable v-ripple @click="fnToHref(router, '/')">
          <q-item-section avatar>考勤预处理</q-item-section>
        </q-item>

        <q-item dense clickable v-ripple @click="fnToHref(router, '/timeCalculator')">
          <q-item-section avatar>时间计算器</q-item-section>
        </q-item>

        <q-item dense clickable v-ripple @click="fnToHref(router, '/rezip')">
          <q-item-section avatar>重新压缩</q-item-section>
        </q-item>
      </q-list>
      <!-- drawer content -->
    </q-drawer>

    <!-- <q-drawer v-model="rightDrawerOpen" side="right" overlay bordered> -->
    <!-- drawer content -->
    <!-- </q-drawer> -->

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer bordered class="bg-grey-8 text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg" />
          </q-avatar>
          <div></div>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from "vue-router"

const router = useRouter()

const leftDrawerOpen = ref(false)

const fnToHref = (router, url) => router.push(url)

// const rightDrawerOpen = ref(false)
</script>
