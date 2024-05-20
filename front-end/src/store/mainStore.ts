import { defineStore } from "pinia";
interface State {
    url: string
    popupMessage: string
    showPopupMessage: boolean,
    resetStations: boolean,
    loginid: string,
}

export const mainStore = defineStore({
    id: "mainStore",
    state: ():State => {
        return {
            url: "",
            popupMessage: "lorem ipsum dolorem etc etc etc etc etc etc",
            showPopupMessage: false,
            resetStations: false,
            loginid: "",
        }
    },
    actions: {
        setUrl(value:string) {
            this.$state.url = value
        },
        setPopupMessage(value: string) {
            this.$state.popupMessage = value;
        },
        setShowPopupMessage(value: boolean) {
            this.showPopupMessage = value;
        },
        setResetStations(value: boolean) {
            this.$state.resetStations = value;
        },
        setLoginid(value: string) {
            this.$state.loginid = value;
        },
    }
});