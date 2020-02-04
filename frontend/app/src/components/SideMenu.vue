<template>
    <v-list>
        <v-list-item-group>
            <v-list-item v-for="(item, index) in items" :key="index" :to="item.link" link>
                <v-list-item-content>
                    <v-list-item-title v-text="item.text"></v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list-item-group>
    </v-list>
</template>

<script>
export default {
    name: 'SideMenu',
    data() {
        return {
            items: [
                { text: 'Home', link: '/'},
                { text: 'Inventory', link: '/inventoryList'},
                { text: 'Commission', link: '/commission'},
                { text: 'Manage Market', link: '/market'},
                { text: 'Manage Customer', link: '/customer'},
                { text: 'Report Setting', link: '/setting'},
                { text: 'Sign Out', link: '/signout'}
            ]
        }
    }
}
</script>

<style>
</style>
