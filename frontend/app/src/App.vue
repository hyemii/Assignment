<template>
    <v-app>
        <v-navigation-drawer :clipped="$vuetify.breakpoint.lgAndUp" v-model="drawer" app>
            <SideMenu/>
        </v-navigation-drawer>

        <v-app-bar :clipped-left="$vuetify.breakpoint.lgAndUp" app>
            <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
            <v-toolbar-title style="width: 300px" class="ml-0 pl-4">
                <router-link to="/"><img src="./assets/testlogo.png"></router-link>
            </v-toolbar-title>
        </v-app-bar>

        <v-content>
            <v-container fluid>
                <router-view></router-view>
            </v-container>
        </v-content>
    </v-app>
</template>

<script>
import SideMenu from "./components/SideMenu";

export default {
    name: "App",
    components: {
        SideMenu
    },
    data() {
        return {
            drawer: null
        };
    }
};
</script>

<style>
</style>