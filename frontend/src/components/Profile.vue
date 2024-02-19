<template>
  <div class="app">
    <div class="image">
      <img src="./../../account-logo.png" style="height: 180px; width: 180px" />
    </div>
    <div class="profile">
      <div class="profile-row">
        <a-typography-text strong>User Name : </a-typography-text>
        <a-typography-text type="secndary">{{ UserName }}</a-typography-text>
      </div>
      <div class="profile-row">
        <a-typography-text strong>User Id : </a-typography-text>
        <a-typography-text type="secndary">{{ UserId }}</a-typography-text>
      </div>
      <div class="profile-row">
        <a-typography-text strong>User Type : </a-typography-text>
        <a-typography-text type="secndary">{{ UserType }}</a-typography-text>
      </div>
      <div class="profile-row">
        <a-typography-text strong>Email : </a-typography-text>
        <a-typography-text type="secndary">{{ Email }}</a-typography-text>
      </div>
      <div class="profile-row">
        <a-typography-text strong>Broker : </a-typography-text>
        <a-typography-text type="secndary">{{ Broker }}</a-typography-text>
      </div>
      <a-button @click="Logout">Logout</a-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";
import { notification } from "ant-design-vue";
import { getbaseURI} from "./../utils"
const router = useRouter();

const UserName = ref("");
const UserId = ref("");
const UserType = ref("");
const Email = ref("");
const Broker = ref("");

onMounted(() => {
  GetProfileData();
});
const GetProfileData = () => {
  const apiURL = getbaseURI()+"/users/profile";


  axios
    .get(apiURL, {
      headers: {
        "Content-Type": "multipart/form-data",
        token: localStorage.getItem("token"),
      },
    })
    .then((response) => {
      Email.value = response?.data?.data?.user?.email;
      UserName.value = response?.data?.data?.user?.user_name;
      UserType.value = response?.data?.data?.user?.user_type;
      UserId.value = response?.data?.data?.user?.user_id;
      Broker.value = response?.data?.data?.user?.broker;
    })
    .catch((error) => {
      console.error("Error while fetching profile:", error);
      notification["error"]({
        message: "Could Not get profile data",
        description: error?.response?.data?.error,
      });
    });
};
const Logout = () => {
  localStorage.clear();
  router.push("/signin");
};
</script>
<style scoped>
.app {
  height: auto;
  display: flex;
  flex-direction: row;
}
.image {
  margin: 0px 100px;
}
.profile-row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.profile {
  width: 30%;
  margin: auto 0px;
  display: flex;
  flex-direction: column;
  align-content: center;
}
</style>
