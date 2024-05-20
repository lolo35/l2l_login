<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { mainStore } from '../store/mainStore';
import { storeToRefs } from 'pinia';
import autoAnimate from '@formkit/auto-animate';

const store = mainStore();
const { popupMessage } = storeToRefs(store)
const popup = ref();
onMounted(() => {
    setTimeout(() => {
        autoAnimate(popup.value);
    }, 500)
});
</script>
<template>
    <div class="absolute top-0 bottom-0 right-0 left-0 bg-black opacity-30 z-10"></div>
    <div class="absolute top-0 bottom-0 right-0 left-0 z-20">
        <div class="flex flex-col w-full h-full items-center justify-center" ref="popup">
            <div class="dark:bg-gray-700 bg-white px-5 py-3 shadow rounded border dark:border-slate-700 flex flex-col space-y-2 w-1/5">
                <div class="flex flex-row">
                    <h3 class="dark:text-gray-200 text-gray-800 font-bold text-2xl">{{ popupMessage }}</h3>
                </div>
                <div class="flex flex-row-reverse">
                    <button class="text-white dark:bg-sky-500 px-5 py-3 rounded w-20 dark:hover:bg-sky-600 bg-blue-500 hover:bg-blue-600" @click="store.setShowPopupMessage(false)">Ok</button>
                </div>
            </div>
        </div>
    </div>
</template>