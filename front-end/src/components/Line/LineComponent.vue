<script lang="ts" setup>
import { storeToRefs } from 'pinia';
import { PropType, computed } from 'vue';
import { Lines } from '../../store/Areas/index';
import { areaStore } from '../../store/Areas/index';

const area_store = areaStore();
const { selectedLines } = storeToRefs(area_store);

const props = defineProps({
    line: {
        type: Object as PropType<Lines>,
        required: true,
    }
});

const setLine = () => {
    area_store.setSelectedLines(props.line);
}

const isSelected = computed(() => {
    return selectedLines.value.filter(element => {
        const id = element.id;
        const lineid = props.line.id;

        return lineid === id ? true : false;
    })
});
</script>
<template>
    <div class="flex flex-row px-2 py-1 group hover:bg-sky-500 rounded space-x-2 items-center" @click="setLine()">
        <i class="fa-solid fa-circle-check text-green-500" v-if="isSelected.length > 0"></i>
        <p class="dark:text-gray-200">{{ line.code }} - {{ line.description }}</p>
    </div>
</template>