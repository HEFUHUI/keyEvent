import './plugins/axios'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import installElementPlus from './plugins/element'
import VueAxios from 'vue-axios'
import axios from './plugins/axios'


const app = createApp(App)
installElementPlus(app)
app.use(store)
.use(VueAxios, axios)
.use(router).mount('#app')