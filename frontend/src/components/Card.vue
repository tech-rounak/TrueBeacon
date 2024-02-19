<template>
  <div>
    <a-card :bordered="false">
      <div>
        <span class="card-title">{{ props.heading }}</span>
      </div>
      <div>
        <span
          v-if="props.heading === 'Days P&L' || props.heading === 'Total P&L'"
          :class="getClassName(props.value)"
        >
          {{ props.value }}
        </span>
        <span v-else style="font-size: large; font-weight: bolder">
          {{ props.value }}
        </span>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { defineProps } from "vue";
const props = defineProps({
  heading: {
    type: String,
    required: true,
  },
  value: {
    type: Number,
    required: true,
  },
});
const getClassName = (num) => {
  if (num > 0) return "positive";
  else return "negative";
};
</script>

<style scoped>
.card-title {
  font-size: medium;
  font-weight: 500;
  color: rgb(143, 155, 157);
}
.positive {
  font-size: large;
  font-weight: bolder;
  color: green;
}
.negative {
  font-size: larger;
  font-weight: bolder;
  color: red;
}
</style>
