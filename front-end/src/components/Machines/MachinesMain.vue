<script lang="ts" setup>
import { machinesStore } from '../../store/Machines';
import { storeToRefs } from 'pinia';
import MachineComponentVue from './MachineComponent.vue';
import { computed } from "vue";

const machine_store = machinesStore();
const { machines } = storeToRefs(machine_store);

const machineElementHeight = computed(() => {
    return Math.floor(100 / Math.ceil(machines.value.length / 2));
});
</script>
<template>
    <div class="grid grid-cols-2 gap-3 w-full overflow-auto scrollbar-thin 
        dark:scrollbar-track-slate-50 dark:scrollbar-thumb-slate-800 scrollbar-track-slate-100 scrollbar-thumb-slate-300">
        <machine-component-vue v-for="machine in machines" :key="machine[0].id" :machines="machine" :height="machineElementHeight">{{ machine[0].id }}</machine-component-vue>
    </div>
</template>