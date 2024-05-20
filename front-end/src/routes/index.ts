import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import HomeView from '../views/HomeView.vue';

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Home",
        component: HomeView,
    }
];

const router = createRouter({
    history: createWebHistory(import.meta.env.VITE_APP_BASE_PATH as string),
    routes,
})

export default router;