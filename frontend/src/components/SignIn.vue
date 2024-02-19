<template>
  <div id="app">
    <h1>SignIn</h1>
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
      style="margin: 20px 100px; width: 90%"
    >
      <a-form-item ref="user_name" label="UserName" name="user_name">
        <a-input v-model:value="formState.user_name" style="width: 200px" />
      </a-form-item>
      <a-form-item ref="password" label="Password" name="password">
        <a-input-password
          v-model:value="formState.password"
          style="width: 200px"
        />
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 18 }">
        <a-button type="primary" style="width: 100%" @click="onSubmit"
          >Signin</a-button
        >
      </a-form-item>
    </a-form>
    <p @click="SignupRoute"><u>New User? Signup</u></p>
  </div>
</template>

<script setup>
import { reactive, ref, toRaw } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";
import { notification } from "ant-design-vue";
import { getbaseURI} from "./../utils"

const router = useRouter();
const formRef = ref();
const labelCol = {
  span: 5,
};
const wrapperCol = {
  span: 13,
};
const formState = reactive({
  password: "",
  user_name: "",
});

const rules = {
  user_name: [
    {
      required: true,
      message: "Please input symbol",
      trigger: "change",
    },
  ],
  password: [
    {
      required: true,
      message: "Please enter password",
      trigger: "change",
    },
  ],
};
const onSubmit = () => {
  formRef.value
    .validate()
    .then(() => {
      Login(toRaw(formState));
    })
    .catch((error) => {
      console.log("error", error);
    });
};

const Login = (formData) => {
  
  const url = getbaseURI()+"/users/login";
  axios
    .post(url, formData)
    .then((res) => {
      localStorage.setItem("token", res?.data?.data?.token);
      localStorage.setItem("isUserLoggedIn", true);
      notification["success"]({
        message: "Login Successfull",
      });
      setTimeout(() => {
        router.push("/dashboard");
      }, 1000);
    })
    .catch((err) => {
      console.error(err);
      notification["error"]({
        message: "Login Failed",
        description: err?.response?.data?.error,
      });
    });
};

const SignupRoute = () => {
  router.push("/signup");
};
</script>

<style scoped>
#app {
  height: 300px;
  width: 450px;
  border: 2px solid rgb(175, 174, 179);
  border-radius: 8px;
  margin: auto;
}
.text-box {
  height: 30px;
  width: 200px;
  margin: 10px;
}
.input-data {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  margin: 10px;
}
.label {
  font-weight: bold;
}
p:hover {
  color: blue;
  cursor: pointer;
}
</style>
