<script lang="ts" setup>
import { storeToRefs } from 'pinia';
import { defineProps, ref, PropType, watch, onMounted } from 'vue';
import { machinesStore } from '../../store/Machines';
import { Machines } from '../../store/Machines';
import { mainStore } from '../../store/mainStore';
import localforage from 'localforage';

interface UserMachine {
    machine: number;
}

const machine_store = machinesStore();
const store = mainStore();
const { resetStations } = storeToRefs(store);
const { loginid } = storeToRefs(store);

const props = defineProps({
    station: {
        type: Object as PropType<Machines>,
        required: true
    },
    selectedAll: {
        type: Boolean,
        required: true
    }
});

const value = ref<boolean>(false);
watch(() => resetStations.value, () => {
    if(resetStations.value) {
        value.value = false;
    }
});

watch(() => props.selectedAll, () => {
    setSelected();
});

watch(() => loginid.value, () => {
    checkForMachines();
});

const setMachine = async () => {
    if(!value.value) {
        console.log(`settings machine`);
        machine_store.setSelectedMachines(props.station)
        if (loginid.value !== "") {
            const check = await localforage.getItem<UserMachine[]|null>(loginid.value);
            if (check) {
                const checkMachine = check.find(element => element.machine === props.station.id);
                if (!checkMachine) {
                    const userMachine: UserMachine = {
                        machine: props.station.id
                    }
                    check.push(userMachine);
                    await localforage.setItem(loginid.value, check);
                }
            } else {
                const userMachine: UserMachine = {
                    machine: props.station.id
                }
                const userMachines: UserMachine[] = [];
                userMachines.push(userMachine);
                await localforage.setItem(loginid.value, userMachines);
            }
        }
        return;
    } 
    if(value.value) {
        console.log(`removing machine`);
        machine_store.removeSelectedMachine(props.station);
        if (loginid.value !== "") {
            const check = await localforage.getItem<UserMachine[]|null>(loginid.value);
            if (check) {
                const checkMachine = check.find(element => element.machine === props.station.id);
                if (checkMachine) {
                    const index = check.indexOf(checkMachine);
                    check.splice(index, 1);
                    await localforage.setItem(loginid.value, check);
                }
            }
        }
    }
}

const setSelected = async () => {
    if(props.selectedAll && !props.station.code.includes("General")) {
        value.value = true;
        machine_store.setSelectedMachines(props.station);
        
    }
    if(!props.selectedAll && !props.station.code.includes("General")) {
        value.value = false;
        machine_store.removeSelectedMachine(props.station)
    }
}

const checkForMachines = async () => {
    if (loginid.value !== "") {
        const check = await localforage.getItem<UserMachine[]|null>(loginid.value);
        if (check) {
            console.log(check);
            const checkMachine = check.find(element => element.machine === props.station.id);
            if (checkMachine) {
                value.value = false;
                setMachine();
                value.value = true;
            }
        }
    }
}
</script>
<template>
    <div class="flex flex-row items-center justify-between">
        <input :id="station.id.toString()" type="checkbox" v-model="value" @click="setMachine()">
        <label :for="station.id.toString()" 
        class="dark:text-gray-200 font-semibold text-sm cursor-pointer dark:hover:text-sky-300 hover:text-blue-500">
            <div class="flex flex-col items-end px-3 py-2">
                <p class="font-bold text-lg">{{ station.code }}</p> 
                <p class="text-sm italic">{{ station.description }}</p>
            </div>
        </label>
    </div>
</template>