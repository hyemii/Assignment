import Vue from 'vue'
import Router from 'vue-router'
import Main from '../components/Main.vue'
import NotFound from '../components/NotFound.vue'
import InventoryList from '../components/InventoryList.vue'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/',
            name: 'main',
            component: Main
        },
        {
            path: '*',
            name: 'notFound',
            component: NotFound
        },
        {
            path: '/signout',
            redirect: '/'
        },
        {
            path: '/inventoryList',
            name: 'inventoryList',
            component: InventoryList
        }
    ]
})
