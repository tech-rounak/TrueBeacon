<template>
  <div>
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
      style="margin: 20px 100px; width: 90%"
    >
      <a-form-item ref="symbol" label="Symbol" name="symbol">
        <a-input v-model:value="formState.symbol" style="width: 200px" />
      </a-form-item>
      <a-form-item ref="quantity" label="Quantity" name="quantity">
        <a-input-number
          v-model:value="formState.quantity"
          style="width: 200px"
        />
      </a-form-item>
      <a-form-item ref="price" label="Price" name="price">
        <a-input-number
          v-model:value="formState.price"
          style="width: 200px"
          :min="2"
        />
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 14, offset: 3 }">
        <a-button type="primary" style="width: 100%" @click="onSubmit"
          >Place Order</a-button
        >
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { reactive, ref, toRaw } from "vue";
import { notification } from "ant-design-vue";
import axios from "axios";
import { getbaseURI} from "./../utils"

const formRef = ref();
const labelCol = {
  span: 5,
};
const wrapperCol = {
  span: 13,
};
const formState = reactive({
  symbol: null,
  quantity: null,
  price: null,
});
const rules = {
  symbol: [
    {
      required: true,
      message: "Please input symbol",
      trigger: "change",
    },
  ],
  price: [
    {
      required: true,
      message: "Please enter a number",
      trigger: "blur",
    },
    {
      type: "number",
      min: 0.01,
      message: "Price must be greater than 0.01",
      trigger: "blur",
    },
  ],
  quantity: [
    {
      required: true,
      message: "Please enter a number",
      trigger: "blur",
    },
    {
      type: "number",
      min: 1,
      message: "Quantity must be greater than 1",
      trigger: "blur",
    },
  ],
};
const onSubmit = () => {
  formRef.value
    .validate()
    .then(() => {
      PlaceOrder(toRaw(formState));
    })
    .catch((error) => {
      console.log("error", error);
    });
};

const resetForm = () => {
  formRef.value.resetFields();
};

const PlaceOrder = (formData) => {
  const url = getbaseURI()+"/place/order";
  axios
    .post(url, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
        token: localStorage.getItem("token"),
      },
    })
    .then(() => {
      notification["success"]({
        message: "Order Placed Successfully",
      });
      resetForm();
    })
    .catch((err) => {
      notification["error"]({
        message: "Place Order Failed",
        description: err?.response?.data?.error,
      });
    });
};
</script>
<style lang="css" scoped>
.label-class {
  font-size: 18px;
  font-weight: 500;
}
.input-row {
  margin: 8px;
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
}
</style>
