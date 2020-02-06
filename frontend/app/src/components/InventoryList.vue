<template>
    <div class="inventoryList">
        <div class="main" style="padding:2%">
            <v-card>
                <v-card-text style="color: #fa8f1d">
                    <h1>Inventory List</h1>
                </v-card-text>
            </v-card><br>

            <div style="border: solid gainsboro 1px">
                <v-row>
                    <v-col cols="12" sm="2" style="margin: auto;">
                        <div align="center">Filter</div>
                    </v-col>
                    <v-col cols="12" sm="7" class="center">
                        <v-text-field label="Model" v-model="search.model" class="filter"></v-text-field>
                        <v-text-field label="Make" v-model="search.make" class="filter"></v-text-field>
                        <v-text-field label="Year" v-model="search.year" class="filter"></v-text-field>
                        <v-text-field label="MSRP" v-model="search.msrp" class="filter"></v-text-field>
                        <v-text-field label="Status" v-model="search.status" class="filter"></v-text-field>
                        <v-select :items="['Y', 'N']" label="Booked" v-model="search.booked" clearable class="filter"></v-select>
                        <v-select :items="['Y', 'N']" label="Listed" v-model="search.listed" clearable class="filter"></v-select>
                    </v-col>
                    <v-col cols="12" sm="3" class="center" style="margin: auto">
                        <v-btn outlined rounded color="grey" @click="updateFilter">Update Filter</v-btn>
                    </v-col>
                </v-row>
            </div>

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
                    <tbody v-if="inventories.length > 0">
                    <tr v-for="(inventory, index) in filterInventories" :key="index">
                        <td><v-checkbox v-model="delvins" :value="inventory.vin" color="orange"></v-checkbox></td>
                        <td>{{ index + 1 }}</td>
                        <td>{{ inventory.vin }}</td>
                        <td>{{ inventory.model }}</td>
                        <td>{{ inventory.make }}</td>
                        <td>{{ inventory.year }}</td>
                        <td>{{ inventory.msrp | currency}}</td>
                        <td>{{ inventory.status }}</td>
                        <td>{{ inventory.booked }}</td>
                        <td>{{ inventory.listed }}</td>
                    </tr>
                    </tbody>
                    <tbody v-if="inventories.length < 1 || filterInventories.length < 1">
                        <td colspan="10" style="font-size: 2.5em; padding:5%">
                            No Data
                        </td>
                    </tbody>
                </table>
            </div>

            <AddDialog v-model="showAddForm"/>

            <v-btn class="ma-2" outlined fab color="orange" v-on:click="deleteInventory">
                <v-icon>mdi-minus</v-icon>
            </v-btn>

            <UploadFile v-model="showUploadForm" />
        </div>
    </div>
</template>

<script>
import AddDialog from '../components/AddDialog'
import UploadFile from '../components/UploadFile'

export default {
    name: "inventoryList",
    components: {
        AddDialog,
        UploadFile
    },
    data() {
        return {
            inventories: [],
            filterInventories: [],
            search: {
                model: '',
                make: '',
                year: '',
                msrp: '',
                status: '',
                booked: '',
                listed: ''
            },
            delvins: [],
            showAddForm: false,
            showUploadForm: false,
            updateFlag: false
        }
    },
    computed: {
        selectAll: {
            get: function () {
                return this.inventories ? this.delvins.length == this.inventories.length : false;
            },
            set: function (value) {
                var delvins = [];
                if (value) {
                    this.inventories.forEach(function (inventory) {
                        delvins.push(inventory.vin);
                    });
                }
                this.delvins = delvins;
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
                    this.filterInventories = this.inventories
                }
            }).catch(reason => {
                console.log('list error', reason)
                alert('Fail!')
            })
        },
        deleteInventory () {
            if (this.delvins.length < 1) {
                alert("Choose a list to delete")
                return false
            }

            var msg = "Are you sure you want to delete?\n\n"
            this.delvins.forEach(function (delvin) {
                msg = msg + 'VIN# = ' + delvin + '\n'
            });

            if (!confirm(msg)) {
                return false;
            }

            var urlVal = "http://localhost:8090"

            this.$http.delete(
                urlVal + '/inventories/' + this.delvins
            ).then(result => {
                if (result.status === 200) {
                    this.inventories = result.data
                    this.filterInventories = this.inventories
                    this.delvins = []
                    alert('Deleted!')
                }
            }).catch(reason => {
                console.log('list error', reason)
                alert('Fail!')
            })
        },
        updateFilter () {
            if (typeof(this.search.booked) === 'undefined') {this.search.booked = ''}
            if (typeof(this.search.listed) === 'undefined') {this.search.listed = ''}

            if (!this.search) { this.filterInventories = this.inventories }
            this.delvins = []
            this.filterInventories = this.inventories.filter(inventory => {
                return inventory.model.toLowerCase().includes(this.search.model.toLowerCase()) &&
                        inventory.make.toLowerCase().includes(this.search.make.toLowerCase()) &&
                        inventory.year.toString().includes(this.search.year.toString()) &&
                        inventory.msrp.toString().includes(this.search.msrp.toString()) &&
                        inventory.status.toLowerCase().includes(this.search.status.toLowerCase()) &&
                        inventory.booked.includes(this.search.booked) &&
                        inventory.listed.includes(this.search.listed)
            })
        }
    },
    filters: {
        currency: function (value) {
            var num = new Number(value)
            return num.toFixed(0).replace(/(\d)(?=(\d{3})+(?:\.\d+)?$)/g, "$1,")
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

.filter {
    width: 90px;
    float: left;
    margin-left: 2%;
    margin-right: 2%;
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