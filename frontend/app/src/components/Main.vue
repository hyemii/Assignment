<template>
    <div class="main" style="padding:2%">
        <v-card>
            <v-card-text style="color: #fa8f1d">
                <h1>Home</h1>
            </v-card-text>
        </v-card><br>

        <v-row>
            <v-col>
                <v-card>
                    <v-card-text>
                        <h1>WebApp	Assignment</h1><br><br>
                        <h2>Task Summary</h2>
                        <p style="margin:1%;">
                            • You need to create single page web application with ajax (front-end/backend)<br>
                            • Use,<br>
                            &nbsp;– Vue.Js (Front End) https://vuejs.org/ and Vuetify (https://vuetifyjs.com/)<br>
                            &nbsp;– Go (Back end) https://golang.org/<br>
                            &nbsp;– (Optional) MySQL database	and ORM you choose (May try beego, https://beego.me/)<br>
                            • Create web app according to following slides and submit<br>
                            &nbsp;– Full source code of both Front End and BackEnd<br>
                            &nbsp;– (Optional) Run the service on AWS micro instance (Free tier) and provide the link to see
                        </p><br>
                        <h2>Must</h2>
                        <p style="margin:1%;">
                            • (FrontEnd) Must use vue-router. (https://router.vuejs.org/)<br>
                            • (BackEnd) Must use router. You can pick whichever router you prefer.<br>
                            &nbsp;– Some references are<br>
                            &nbsp;– http://www.gorillatoolkit.org/<br>
                            &nbsp;– https://github.com/julienschmidt/httprouter
                        </p>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </div>
</template>

<script>
export default {
    name: "Main"
}
</script>
