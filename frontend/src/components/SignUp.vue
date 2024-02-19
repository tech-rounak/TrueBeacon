<template>
  <div>
    <div id="app">
      <h1>Sign Up</h1>
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
        style="margin: 20px 100px; width: 90%"
      >
        <a-form-item ref="name" label="Name" name="Name">
          <a-input v-model:value="formState.name" style="width: 200px" />
        </a-form-item>
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
      <a href="/signin">Already a user, SignIn</a>
    </div>
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
  user_name: "",
  name: "",
  password: "",
});
const rules = {
  user_name: [
    {
      required: true,
      message: "Please input symbol",
      trigger: "change",
    },
  ],
  name: [
    {
      required: true,
      message: "Please input name",
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
      SignUp(toRaw(formState));
    })
    .catch((error) => {
      console.log("error", error);
    });
};

const SignUp = (formData) => {
  const url = getbaseURI()+"/users/signup";
  axios
    .post(url, formData)
    .then(() => {
      notification["success"]({
        message: "Signup Successfull",
      });
      setTimeout(() => {
        router.push("/signin");
      }, 1000);
    })
    .catch((err) => {
      notification["error"]({
        message: "Signup Failed",
        description: err?.response?.data?.error,
      });
    });
};
</script>

<style scoped>
#app {
  height: 600px;
  width: 450px;
  margin: auto;
  border: 2px solid rgb(175, 174, 179);
  border-radius: 8px;
}
</style>
