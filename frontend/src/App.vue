<template>
  <router-view/>
</template>

<script setup>
import {useRouter} from "vue-router"

const router = useRouter();
  if(localStorage.getItem("isUserLoggedIn")){
    router.push("/dashboard")
  }else{
    router.push("/signin")
  }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
