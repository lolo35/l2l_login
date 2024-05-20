<script lang="ts" setup>
import { ref, onMounted, computed } from 'vue'
import { fetchAreas } from '../../requests/areasRequest';
import { mainStore } from '../../store/mainStore';
import { areaStore }from '../../store/Areas/index';
import { storeToRefs } from 'pinia';
import { Areas } from '../../store/Areas/index';
import localforage from 'localforage';
import autoAnimate from '@formkit/auto-animate';

const store = mainStore();
const { url } = store;

const area_store = areaStore();
const { areas } = storeToRefs(area_store);

const showOptions = ref<boolean>(false);
const areasLoaded = ref<boolean>(false);
const search = ref<string>("");
const selectedArea = ref<string>("Alege zona");
const dropdown = ref();
const emits = defineEmits(['reload']);

const filteredAreas = computed(() => {
    return areas.value.filter(element => {
        const description = element.description.toString().toLowerCase();
        const filter = search.value.toString().toLowerCase();
        return description.includes(filter);
    })
});

onMounted(async () => {
    const areas = await fetchAreas(url);
    if(areas.success) {
        area_store.setAreas(areas.data);
        areasLoaded.value = true;
    }
    checkForSelected();
    autoAnimate(dropdown.value);
});

const checkForSelected = async () => {
    try {
        const localArea = await localforage.getItem<string>(`area`);
        if(localArea) {
            const area:Areas = JSON.parse(localArea);
            area_store.setSelectedArea(area);
            area_store.setAreaSelected(true);
            selectedArea.value = area.description;
        }
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.VITE_MODE === "dev") throw new Error(exception.message);
        }
    }
}

const setSelectedArea = async (area:Areas) => {
    area_store.setAreaSelected(false);
    area_store.setSelectedArea(area);    
    selectedArea.value = area.description;
    const localArea = JSON.stringify(area);
    await localforage.setItem(`area`, localArea);
    showOptions.value = false;
    area_store.setAreaSelected(true);
}

const reset = async () => {
    await localforage.removeItem(`area`);
    await localforage.removeItem(`selectedLines`);
    area_store.$reset();
    emits('reload');
}

</script>
<template>
    <div class="flex flex-col space-y-2">
        <div class="flex flex-row items-center justify-between">
            <label for="area" class="dark:text-gray-200 text-gray-800 font-semibold">
                <i class="fa-solid fa-industry text-blue-500 dark:text-sky-500"></i>
                Zona
            </label>
            <button @click="reset()" class="text-blue-500 hover:text-blue-400 hover:underline">Reset</button>
        </div> 
        <div class="flex flex-row items-center justify-between px-3 py-2 dark:bg-gray-800 bg-slate-200 hover:bg-slate-300 rounded hover:dark:bg-gray-600 cursor-pointer relative" ref="dropdown">
            <div class="absolute top-0 bottom-0 left-0 right-0 z-10" @click="showOptions = !showOptions"></div>
            <div class="absolute left-0 right-0 top-12 z-20" v-if="showOptions && areasLoaded">
                <div class="flex flex-col px-3 py-2 bg-gray-200 rounded dark:bg-gray-600 shadow border dark:border-gray-800 max-h-96 overflow-y-scroll">
                    <input v-model="search" type="text" class="dark:bg-gray-700 px-3 py-2 rounded dark:text-gray-200" placeholder="Cautare...">
                    <div v-for="area in filteredAreas" :key="area.id" class="flex flex-row px-2 py-1 group hover:bg-sky-500 rounded group" @click="setSelectedArea(area)">
                        <p class="dark:text-gray-200 group-hover:text-white">{{ area.description }}</p>
                    </div>
                </div>
                
            </div>
            <p class="dark:text-gray-200">{{ selectedArea }}</p>
            <div class="flex flex-col">
                <i class="fa-solid fa-sort dark:text-sky-500 text-blue-500"></i>
            </div>
        </div>
    </div>
</template>