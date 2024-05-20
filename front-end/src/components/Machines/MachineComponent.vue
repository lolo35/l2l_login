<script lang="ts" setup>
import { defineProps, PropType, onMounted, ref, computed } from 'vue';
import { Machines } from '../../store/Machines/index';
import autoAnimate from '@formkit/auto-animate';
import StationComponentVue from './StationComponent.vue';
import { machinesStore } from '../../store/Machines/index';
import { storeToRefs } from 'pinia';

const machine_store = machinesStore();
const { machinesNotSelected } = storeToRefs(machine_store);

const props = defineProps({
    machines: {
        type: Object as PropType<Machines[]>,
        required: true,
    },
    height: {
        type: Number,
        required: true,
    }
});

const container = ref();
const searchTerm = ref('');
const selectAll = ref(false);

const filteredMachines = computed(() => {
    return props.machines.filter(machine => {
        return machine.code.toLowerCase().includes(searchTerm.value.toLowerCase()) || machine.description.toLowerCase().includes(searchTerm.value.toLowerCase());
    });
});

onMounted(() => {
    autoAnimate(container.value)
});

</script>
<template>
    <div ref="container" 
        :class="[{'dark:bg-gray-700 bg-white': !machinesNotSelected, 'dark:bg-red-300 bg-red-500 animate-pulse': machinesNotSelected}]"
        :style="[`height: calc(${$props.height}vh - 2rem);`]"
        class="px-3 py-2 rounded shadow border dark:border-gray-800 flex flex-col space-x-2 space-y-2 transition-colors ease-in-out duration-500 
        overflow-auto relative
        scrollbar-thin dark:scrollbar-track-slate-50 dark:scrollbar-thumb-slate-800 scrollbar-track-slate-300 scrollbar-thumb-slate-50">
        <div class="flex flex-row items-center space-x-2">
            <p class="dark:text-gray-200 font-bold">Linia</p>
            <p class="dark:text-gray-200 font-semibold italic">{{ machines[0].linecode }}</p>
        </div>
        <!-- <div class=" h-36">&nbsp;</div> -->
        <div class="sticky right-5 left-5 top-7 flex flex-col dark:bg-gray-700 bg-white py-2 border-gray-200 border dark:border-gray-600 rounded px-3">
            <input v-model="searchTerm" type="text" placeholder="Cautare..." class="w-full px-3 py-2 rounded dark:bg-gray-500 bg-slate-200 dark:outline-sky-300 outline-blue-500 dark:text-gray-200">
            <div class="flex flex-row-reverse space-x-reverse space-x-2">
                <label for="select_all" class="text-xs font-semibold dark:text-gray-200 cursor-pointer hover:text-sky-400">Alege toate</label>
                <input type="checkbox" id="select_all" v-model="selectAll">
            </div>
        </div>
        <station-component-vue v-for="machine in filteredMachines" :key="machine.id" :station="machine" :selected-all="selectAll"></station-component-vue>
    </div>
</template>