import { createRouter, createWebHistory } from "vue-router";

import  SignUp from "./components/SignUp"
import SignIn from "./components/SignIn"
import Profile from "./components/Profile"
import Dashboard from "./components/Dashboard"

const routes = [
      {
        name: "SignIn",
        path: "/signin",
        meta: {
          title: "Signin",
        },
        component:SignIn,
      },
      {
        name: "Signup",
        path: "/signup",
        meta: {
          title: "SignUp",
        },
        component:SignUp,
      },
      { 
        name: "Profile",
        path: "/profile",
        meta: {
         title: "Profile",
        },
        component:Profile,

      },
      { 
        name: "Dashboard",
        path: "/dashboard",
        meta: {
         title: "Dashboard",
        },
        component:Dashboard,

      }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})
export default router