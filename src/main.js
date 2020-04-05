import Vue from "vue"
import App from "./App.vue"
import vuetify from "./plugins/vuetify"
import Router from "vue-router"
import Meta from "vue-meta"

import routes from './router/routers'

Vue.use(Router)
Vue.use(Meta)

// Create a new router
const router = new Router({
  mode: "hash",
  routes
})


new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount("#app")
