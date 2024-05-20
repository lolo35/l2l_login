import { createApp } from 'vue'
import App from './App.vue'
import router from './routes/index';
import './assets/css/main.css';
import { createPinia } from 'pinia';
import './assets/css/fontawesome-free-6.1.2-web/css/all.css';
import localforage from 'localforage';

localforage.config({
    name: "L2L Login"
});

const app = createApp(App);
app.use(router);
app.use(createPinia());
app.mount('#app');
