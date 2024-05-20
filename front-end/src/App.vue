<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import { onMounted } from 'vue';
import { mainStore } from './store/mainStore'

const store = mainStore();

const checkColorScheme = () => {
  document.body.classList.add('body-bg-attr');
  const colorScheme = window.matchMedia('(prefers-color-scheme: dark)');
  if(colorScheme.matches) {
    document.body.classList.remove('light-bg');
    document.body.classList.add('dark-bg');
    return
  }
  document.body.classList.remove('dark-bg');
  document.body.classList.add('light-bg')
}

onMounted(() => {
  document.title = import.meta.env.VITE_TITLE;
  store.setUrl("http://artl-app05:8100/")
  // if(import.meta.env.VITE_MODE === "dev") {
  //   store.setUrl("http://localhost:8000/")
  // }
  checkColorScheme();
  window.matchMedia('(prefers-color-scheme:dark)').addEventListener("change", () => {
    checkColorScheme();
  });
});
</script>

<template>
  <router-view></router-view>
</template>

<style>
.light-bg {
  background: url('./assets/light.svg')
}
.dark-bg {
  background: url('./assets/dark.svg');
}
.body-bg-attr {
  background-position: bottom right;
  background-repeat: no-repeat;
}
</style>