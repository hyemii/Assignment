import Vue from 'vue'
import Vuetify from 'vuetify'

import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import axios from 'axios'

Vue.config.productionTip = false

Vue.use(Vuetify)
Vue.prototype.$http = axios

new Vue({
    router,
    vuetify,
    render: h => h(App)
}).$mount('#app')
