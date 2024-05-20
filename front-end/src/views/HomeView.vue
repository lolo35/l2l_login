<script lang="ts" setup>
import SelectAreaVue from '../components/Areas/SelectArea.vue';
import SelectLineVue from '../components/Line/SelectLine.vue';
import MachinesMainVue from '../components/Machines/MachinesMain.vue';
import IdInputVue from '../components/Operator/IdInput.vue';
import SuccessMessageVue from '../components/SuccessMessage.vue';
import { areaStore } from '../store/Areas/index';
import { machinesStore } from '../store/Machines';
import { mainStore } from '../store/mainStore';
import { storeToRefs } from 'pinia';
import autoAnimate from '@formkit/auto-animate';
import { ref, onMounted } from 'vue';

const store = mainStore();
const { showPopupMessage } = storeToRefs(store);
const area_store = areaStore();
const { areaSelected } = storeToRefs(area_store);
const machine_store = machinesStore();
const { machinesLoaded } = storeToRefs(machine_store);
const dropdown = ref();
const key = ref<number>(0);

onMounted(() => {
    autoAnimate(dropdown.value);
})

</script>
<template>
    <success-message-vue v-if="showPopupMessage"></success-message-vue>
    <div class="grid grid-cols-2 w-full gap-2 h-[100vh] overflow-y-auto scrollbar-thin scrollbar-track-slate-300 scrollbar-thumb-slate-200">
        <div class="flex flex-col items-center justify-center px-3 py-2">
            <div class="flex flex-col bg-white rounded w-full px-3 py-2 dark:bg-gray-700 border dark:border-gray-800 shadow space-y-2" ref="dropdown">
                <select-area-vue v-on:reload="key++" :key="key"></select-area-vue>
                <select-line-vue :key="key" v-if="areaSelected"></select-line-vue>
                <id-input-vue></id-input-vue>
            </div>
        </div>
        <div class="flex flex-col w-full items-center justify-center px-3 py-2">
            <machines-main-vue v-if="machinesLoaded"></machines-main-vue>
        </div>      
    </div>
</template>