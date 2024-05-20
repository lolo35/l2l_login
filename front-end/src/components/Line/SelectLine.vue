<script lang="ts" setup>
import { ref, onMounted, computed, watch } from 'vue';
import autoAnimate from '@formkit/auto-animate';
import { fetchLines } from '../../requests/linesRequest';
import { mainStore } from '../../store/mainStore';
import { areaStore } from '../../store/Areas/index';
import { storeToRefs } from 'pinia';
import LineComponent from './LineComponent.vue';
import localforage from 'localforage';
import { Lines } from '../../store/Areas/index';
import { fetchMachines } from '../../requests/machinesRequest';
import { machinesStore } from '../../store/Machines/index';

const store = mainStore();
const { url } = store;

const area_store = areaStore();
const { selectedArea } = area_store;
const { lines, selectedLines, linesNotSelected } = storeToRefs(area_store);

const machine_store = machinesStore();

const showOptions = ref<boolean>(false);
const search = ref<string>("");
const dropdown = ref();

onMounted(async () => {
    autoAnimate(dropdown.value)
    if(selectedArea) {
        const lines = await fetchLines(url, selectedArea?.id);
        if(lines.success) {
            area_store.setLines(lines.data);
        }
    }
    checkSavedLines();
});

// watch(() => selectedArea.id, async () => {
//     const lines = await fetchLines(url, selectedArea?.id)
//     if(lines.success) {
//         area_store.setLines(lines.data);
//     }
// });

const filteredLines = computed(() => {
    return lines.value.filter(element => {
        const description = element.description.toString().toLowerCase();
        const filter = search.value.toString().toLocaleLowerCase();
        return description.includes(filter);
    })
});

const selectedLine = computed(() => {
    let lines = "Alege linia/liniile";
    if(selectedLines.value.length > 0) {
        lines = "";
        selectedLines.value.forEach(element => {
            lines += `${element.description.substring(0, 5)}..., `
        });
    }
    return lines;
});

const saveSelectedLines = async () => {
    try {
        await localforage.setItem(`selectedLines`, JSON.stringify(selectedLines.value));
        showOptions.value = false;
        area_store.setLinesSelected(true);
        await fetch_machines();
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.VITE_MODE === "dev") {
                throw new Error(exception.message);
            }
        }
    }
}

const fetch_machines = async () => {
    let lineIDs  = "";
        selectedLines.value.forEach(element => {
            lineIDs += `${element.id},`;
        })

    const machines = await fetchMachines(url, lineIDs);
    if(machines.success) {
        machine_store.setMachines(machines.data);
        machine_store.setMachinesLoaded(true);
    }
}

const checkSavedLines = async () => {
    try {
        const lines = await localforage.getItem<string>(`selectedLines`);
        if(lines) {
            
            const sLines = JSON.parse(lines);
            sLines.forEach((element:Lines) => {
                area_store.setSelectedLines(element);
            });
            await fetch_machines();
            area_store.setLinesSelected(true);
        }
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.VITE_MODE === "dev") {
                throw new Error(exception.message);
            }
        }
    }
}

</script>

<template>
    <div class="flex flex-col space-y-2">
        <label for="line" class="font-semibold dark:text-gray-200 text-gray-800">
            <i class="fa-solid fa-gears text-blue-500 dark:text-sky-500"></i>
            Linia
        </label>
        <div 
            :class="{'dark:bg-gray-800 hover:dark:bg-gray-600 bg-slate-200 hover:bg-slate-300': !linesNotSelected, 'dark:bg-red-300 animate-pulse': linesNotSelected}"
            class="flex flex-row items-center justify-between px-3 py-2 rounded cursor-pointer relative transition-colors ease-in-out duration-500" ref="dropdown">
            <div class="absolute top-0 bottom-0 left-0 right-0 z-10" @click="showOptions = !showOptions"></div>
            <div class="absolute left-0 right-0 top-12 z-20" v-if="showOptions">
                <div class="flex flex-col px-3 py-2 bg-gray-200 rounded dark:bg-gray-600 shadow border dark:border-gray-800 max-h-96 overflow-y-scroll relative">
                    <input v-model="search" type="text" class="dark:bg-gray-700 px-3 py-2 rounded dark:text-gray-200" placeholder="Cautare...">
                    <!-- <div class="flex flex-row px-2 py-1 group hover:bg-sky-500 rounded" v-for="line in filteredLines" :key="line.id">
                        <p class="dark:text-gray-200">{{ line.description }}</p>
                    </div> -->
                    <line-component v-for="line in filteredLines" :key="line.id" :line="line"></line-component>
                    <div class="sticky bottom-0 right-0 left-0">
                        <button class="dark:bg-sky-500 px-3 py-2 rounded text-white dark:hover:bg-sky-700 mt-2 w-full bg-blue-500 hover:bg-blue-400" @click="saveSelectedLines()">Salveaza</button>
                    </div>
                </div>
            </div>
            <p class="dark:text-gray-200">{{ selectedLine }}</p>
            <div class="flex flex-col">
                <i class="fa-solid fa-sort dark:text-sky-500 text-blue-500"></i>
            </div>
        </div>
    </div>
</template>