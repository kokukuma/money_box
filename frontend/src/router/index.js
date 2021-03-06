import Vue from 'vue'
import VueRouter from 'vue-router'
import DApp from '../views/DApp.vue'
import Wallet from '../views/Wallet.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'DApp',
    component: DApp
  },
  {
    path: '/wallet',
    name: 'Wallet',
    component: Wallet
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
