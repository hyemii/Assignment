<template>
    <div class="inventoryList">
        <div class="main" style="padding:2%">
            <v-card>
                <v-card-text style="color: #fa8f1d">
                    <h1>Inventory List</h1>
                </v-card-text>
            </v-card><br>

            <v-row>
                <v-col cols="12" sm="8">
                    <v-card>
                        <v-card-text>Filter</v-card-text>
                    </v-card>
                </v-col>
                <v-col cols="12" sm="4" class="center">
                    <v-card>
                        <v-card-text><button>Update Filter</button></v-card-text>
                    </v-card>
                </v-col>
            </v-row>

            <div class="scroll">
                <table class="table">
                    <thead>
                    <tr>
                        <th><v-checkbox v-model="selectAll" color="orange"></v-checkbox></th>
                        <th>No</th>
                        <th>Vin#</th>
                        <th>Model</th>
                        <th>Make</th>
                        <th>Year</th>
                        <th>MSRP</th>
                        <th>Status</th>
                        <th>Booked</th>
                        <th>Listed</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr v-for="(inventory, index) in inventories" :key="index">
                        <td><v-checkbox v-model="selected" :value="inventory.vin" color="orange"></v-checkbox></td>
                        <td>{{ index + 1 }}</td>
                        <td>{{ inventory.vin }}</td>
                        <td>{{ inventory.model }}</td>
                        <td>{{ inventory.make }}</td>
                        <td>{{ inventory.year }}</td>
                        <td>{{ inventory.msrp }}</td>
                        <td>{{ inventory.status }}</td>
                        <td>{{ inventory.booked }}</td>
                        <td>{{ inventory.listed }}</td>
                    </tr>
                    </tbody>
                </table>
            </div>

            <AddDialog v-model="showAddForm"/>

            <v-btn class="ma-2" outlined fab color="orange" v-on:click="deleteInventory">
                <v-icon>mdi-minus</v-icon>
            </v-btn>
            <v-btn class="ma-2" rounded outlined color="orange" large>Upload File</v-btn>
        </div>
    </div>
</template>

<script>
import AddDialog from '../components/AddDialog'

export default {
    name: "inventoryList",
    components: {
        AddDialog
    },
    data() {
        return {
            inventories: [],
            selected: [],
            showAddForm: false
        }
    },
    computed: {
        selectAll: {
            get: function () {
                return this.inventories ? this.selected.length == this.inventories.length : false;
            },
            set: function (value) {
                var selected = [];

                if (value) {
                    this.inventories.forEach(function (inventory) {
                        selected.push(inventory.vin);
                    });
                }

                this.selected = selected;
            }
        }
    },
    methods: {
        getInventories () {
            var urlVal = "http://localhost:8090"

            this.$http.get(
                urlVal + '/inventories'
            ).then(result => {
                if (result.status === 200) {
                    this.inventories = result.data
                }
            }).catch(reason => {
                console.log('list error', reason)
                alert('Fail!')
            })
        },
        deleteInventory () {
            if (this.selected.length < 1) {
                alert("Choose a list to delete")
                return false
            }

            var msg = "Are you sure you want to delete?\n\n"
            this.selected.forEach(function (select) {
                msg = msg + 'VIN# = ' + select + '\n'
            });

            if (!confirm(msg)) {
                return false;
            }
            var urlVal = "http://localhost:8090"

            console.log(this.selected)
            var list = this.selected

            this.$http.delete(
                urlVal + '/inventories', list
            ).then(result => {
                if (result.status === 200) {
                    alert('Deleted!')
                }
            }).catch(reason => {
                console.log('list error', reason)
                alert('Fail!')
            })
        }
    },
    created() {
        this.getInventories()
    }
}
</script>

<style scoped>
.center {
    text-align: center !important;
}

.scroll {
    height: 50vh;
    overflow: scroll;
    margin: 2% 0;
    border: solid gainsboro 1px;
}

table{
    text-align: center !important;
    color: gray;
    width: 100%;
    border-collapse: collapse;
}

th {
    font-size: 1.2em;
    padding: 0 16px;
}

td {
    font-size: 1.05em;
    padding: 0 16px;
    border-top: solid gray 1px;
}
</style>