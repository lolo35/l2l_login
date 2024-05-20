import { defineStore } from 'pinia';

export const machinesStore = defineStore({
    id: "machinesStore",
    state: ():Store => {
        return {
            machines: [],
            machinesLoaded: false,
            selectedMachines: [],
            machinesNotSelected: false,
        }
    },
    actions: {
        setMachines(value:Machines[][]) {
            this.$state.machines = value;
        },
        setMachinesLoaded(value: boolean) {
            this.$state.machinesLoaded = value;
        },
        setSelectedMachines(value: Machines) {
            this.$state.selectedMachines.push(value);
        },
        removeSelectedMachine(value: Machines) {
            this.$state.selectedMachines.forEach((element, index) => {
                if(element.id === value.id) {
                    this.$state.selectedMachines.splice(index, 1);
                    return;
                }
            })
        },
        setMachinesNotSelected(value: boolean) {
            this.$state.machinesNotSelected = value;
        }
    }
});

interface Store {
    machines: Machines[][],
    machinesLoaded: boolean,
    selectedMachines: Machines[],
    machinesNotSelected: boolean,
}

export interface Machines {
    id: number,
    code: string,
    description: string,
    line: number,
    linecode: string
}