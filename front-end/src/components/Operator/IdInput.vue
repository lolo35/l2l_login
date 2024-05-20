<script lang="ts" setup>
import { ref, computed } from 'vue';
import { mainStore } from '../../store/mainStore';
import { areaStore } from '../../store/Areas';
import { machinesStore } from '../../store/Machines';
import { storeToRefs } from 'pinia';
import { login } from '../../requests/loginRequest';
import { Machines } from '../../store/Machines/index'

const store = mainStore();
const area_store = areaStore();
const machine_store = machinesStore();
const { selectedLines } = storeToRefs(area_store);
const { selectedMachines } = storeToRefs(machine_store);
const { url, loginid } = storeToRefs(store);
const isLoading = ref<boolean>(false);

// const loginid = ref<string>("");
const lineIds = computed(() => {
    let line_ids = "";
    selectedLines.value.forEach(element => {
        line_ids += `${element.id},`;
    })
    return line_ids;
});

const machineIds = computed(() => {
    let machine_ids = "";
    const unique = [...new Set(selectedMachines.value)];
    unique.forEach(element => {
        machine_ids += `${element.id},`;
    });
    return machine_ids;
})

const checkMachinesForLines = computed(() => {
    let machines:string[] = [];
    selectedMachines.value.forEach(element => {
        machines.push(element.linecode);
    });
    const uniq = [...new Set(machines)]
    return uniq;
})

const handleLogin = async () => {
    isLoading.value = true;
    if(checkMachinesForLines.value.length !== selectedLines.value.length) {
        machine_store.setMachinesNotSelected(true);
        setTimeout(() => {
            machine_store.setMachinesNotSelected(false);
        }, 1000);
        return
    }
    const resp = await login(url.value, lineIds.value, machineIds.value, loginid.value)
    if(resp.success) {
        store.setPopupMessage("Success");
        store.setShowPopupMessage(true);
        loginid.value = "";
        const machinesJ = JSON.stringify(selectedMachines.value);
        const machines:Machines[] = JSON.parse(machinesJ);
        machines.forEach(element => {
            console.log(element);
            machine_store.removeSelectedMachine(element)
        });
        store.setResetStations(true);
        setTimeout(() => {
            store.setResetStations(false);
        }, 100);
        isLoading.value = false;
    }
    if(!resp.success) {
        if(resp.error.includes("request.Lines")) {
            area_store.setLinesNotSelected(true);
            setTimeout(() => {
                area_store.setLinesNotSelected(false);
            }, 1000);
            isLoading.value = false;
            return;
        }
        if(resp.error.includes("request.Machines")) {
            machine_store.setMachinesNotSelected(true);
            setTimeout(() => {
                machine_store.setMachinesNotSelected(false);
            }, 1000);
            isLoading.value = false;
            return;
        }
    }
    isLoading.value = false;
}


</script>
<template>
    <form @submit.prevent="handleLogin()" class="flex flex-col w-full space-y-2">
        <input 
            v-model="loginid" 
            type="text" 
            class="w-full px-3 py-2 rounded dark:bg-gray-500 bg-slate-200 dark:outline-sky-300 outline-blue-500 dark:text-gray-200" 
            placeholder="Introduceti ID-ul de operator Ex: 7513" 
            required
        >
        <button type="submit" :disabled="isLoading" class="px-3 py-2 text-white dark:bg-slate-900 dark:hover:bg-slate-800 font-bold rounded bg-blue-500 hover:bg-blue-600">Logheaza-ma</button>
    </form>
</template>