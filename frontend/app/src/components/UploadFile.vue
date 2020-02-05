<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <v-dialog v-model="dialog" persistent max-width="400px">
        <template v-slot:activator="{ on }">
            <v-btn class="ma-2" rounded outlined color="orange" large v-on="on">Upload File</v-btn>
        </template>
        <v-card>
            <div align="center" style="padding: 3%">
                <span class="headline">Upload File</span><br><br>
                <v-file-input label="File input"></v-file-input>
            </div>

            <div align="center" style="padding: 3%">
                <v-btn color="orange" text @click="dialog = false">Upload</v-btn>
                <v-btn color="orange" text @click="dialog = false">Close</v-btn>
            </div>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    name: "UploadFile",
    data() {
        return {
            dialog: false
        }
    }
}
</script>

<style scoped>
</style>