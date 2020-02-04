<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <v-dialog v-model="dialog" persistent max-width="400px">
        <template v-slot:activator="{ on }">
            <v-btn class="ma-2" outlined fab color="orange" v-on="on">
                <v-icon>mdi-plus</v-icon>
            </v-btn>
        </template>
        <v-card>
            <div align="center" style="padding: 3%">
                <span class="headline">Add Inventory</span>
            </div>
            <v-card-text>
                <v-container style="text-align: center">
                    <v-text-field label="VIN#" v-model="inventory.vin" required></v-text-field>
                    <v-text-field label="Model" v-model="inventory.model" required></v-text-field>
                    <v-text-field label="Make" v-model="inventory.make" required></v-text-field>
                    <v-text-field label="Year" v-model="inventory.year" required></v-text-field>
                    <v-text-field label="MSRP" v-model="inventory.msrp" required></v-text-field>
                    <v-text-field label="Status" v-model="inventory.status" required></v-text-field>
                    <v-row>
                        <v-col cols="12" sm="6"><v-select :items="['Y', 'N']" label="Booked" v-model="inventory.booked" required></v-select></v-col>
                        <v-col cols="12" sm="6"><v-select :items="['Y', 'N']" label="Listed" v-model="inventory.listed" required></v-select></v-col>
                    </v-row>
                </v-container>
            </v-card-text>

            <div align="center" style="padding: 3%">
                <v-btn color="orange" text @click="saveInventory">Save</v-btn>
                <v-btn color="orange" text @click="dialog = false">Close</v-btn>
            </div>
        </v-card>
    </v-dialog>
</template>

<script>
    export default {
        name: "AddDialog",
        data() {
            return {
                dialog: false,
                inventory: {
                    vin: '',
                    model: '',
                    make: '',
                    year: '',
                    msrp: '',
                    status: '',
                    booked: '',
                    listed: ''
                }
            }
        },
        methods: {
            saveInventory () {
                if (!confirm("Do you want to save?")) {
                    return false
                }

                // 저장
                var urlVal = "http://localhost:8090"

                this.$http.post(
                    urlVal + '/inventories', this.inventory
                ).then(result => {
                    if (result.status === 200) {
                        alert("Saved!")
                    }
                }).catch(reason => {
                    console.log('post error', reason)
                    alert('Fail!')
                })

                this.dialog = false
                this.$router.go()
            }
        }
    }
</script>

<style scoped>

</style>