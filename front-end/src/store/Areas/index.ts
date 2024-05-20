import { defineStore } from 'pinia';

export interface State {
    areas: Areas[],
    selectedArea: Areas,
    areaSelected: boolean,
    lines: Lines[],
    selectedLines: Lines[],
    linesSelected: boolean,
    linesNotSelected: boolean,
}

export interface Areas {
    id: number,
    code: string,
    description: string,
}

export interface Lines {
    id: number,
    code: string,
    area: number,
    description: string,
}

export const areaStore = defineStore({
    id: "areaStore",
    state: ():State => {
        return {
            areas: [],
            selectedArea: {
                id: 0,
                code: "",
                description: ""
            },
            lines: [],
            selectedLines: [],
            areaSelected: false,
            linesSelected: false,
            linesNotSelected: false,
        }
    },
    actions: {
        setAreas(value: Areas[]) {
            this.$state.areas = value;
        },
        setSelectedArea(value: Areas) {
            this.$state.selectedArea = value;
        },
        setAreaSelected(value: boolean) {
            this.$state.areaSelected = value;
        },
        setLines(value: Lines[]) {
            this.$state.lines = value;
        },
        setSelectedLines(value: Lines) {
            let isSelected = false;
            this.$state.selectedLines.forEach((element, index) => {
                if(element.id === value.id) {
                    console.log(element);
                    isSelected = true;
                    this.$state.selectedLines.splice(index, 1);
                }
            });
            if(!isSelected) {
                this.$state.selectedLines.push(value);
            }
        },
        setLinesSelected(value: boolean) {
            this.$state.linesSelected = value;
        },
        setLinesNotSelected(value: boolean) {
            this.$state.linesNotSelected = value;
        }
    }
})