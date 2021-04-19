import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import axios from 'axios'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
//  {
//    path: '/register',
//    name: 'Register',
//    // route level code-splitting
//    // this generates a separate chunk (about.[hash].js) for this route
//    // which is lazy-loaded when the route is visited.
//    component: () => import(/* webpackChunkName: "register" */ '../views/Register.vue')
//  },
 {
   path: '/user',
   name: 'User',
   component: () => import(/* webpackChunkName: "user" */ '../views/User.vue')
 },
 {
  path: '/email/change/:token',
  name: 'PasswordReset',
  beforeEnter: (to, from, next) => {
    let token = to.params["token"]
    axios.put("http://localhost:8081/email/change/" + token)
      .finally(function() {
        next({name: "Home"});
      });
  },
  component: () => import(/* webpackChunkName: "passwordReset" */ '../views/PasswordReset.vue')
 },
 {
   path: '/forgotpassword',
   name: 'ForgotPassword',
   component: () => import(/* webpackChunkName: "forgotPassword" */ '../views/ForgotPassword.vue')
 },
 {
   path: '/email/activate/:token',
   name: 'ActivateEmail',
   beforeEnter: (to, from, next) => {
    let token = to.params["token"]
    axios.put("http://localhost:8081/email/activate/" + token)
      .finally(function(){
        next({ name: "Home"});
      })
   },
  //  component: () => import(/* webpackChunkName: "activate" */ '../views/ActivateEmail')
 }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

// TODO(Jovan): Authentication
// router.beforeEach((to, from, next) => {

// });

export default router
