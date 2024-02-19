<template>
  <div style="display: flex; flex-direction: column">
    <div style="display: flex; flex-direction: column">
      <div
        style="
          display: flex;
          flex-direction: row;
          margin: 10px;
          justify-content: space-evenly;
        "
      >
        <a-select v-model:value="symbol" size="large" style="width: 120px">
          <a-select-option value="NIFTY-BANK">NIFTY BANK</a-select-option>
          <a-select-option value="NIFTY-50">NIFTY 50</a-select-option>
        </a-select>
        <a-range-picker v-model:value="range" size="large" />
        <a-button type="primary" @click="GetStockHistory">Fetch Data</a-button>
      </div>
      <div style="height: 600px; margin: 10px" size="large">
        <canvas style="width: 100%" ref="lineChartCanvas"></canvas>
      </div>
    </div>
  </div>
</template>

<script setup>
import axios from "axios";
import { ref, onMounted } from "vue";
import dayjs from "dayjs";
import Chart from "chart.js/auto";
import { notification } from "ant-design-vue";

const symbol = ref("NIFTY-50");
const range = ref([]);
const lineChartCanvas = ref(null);
const lineChart = ref(null);
onMounted(() => {
  range.value.push(dayjs("2015-01-01"));
  range.value.push(dayjs(new Date()));
  GetStockHistory();
});
const setupLineChart = (val) => {
  const ctx = lineChartCanvas.value.getContext("2d");
  if (lineChart.value) {
    lineChart.value.destroy();
  }
  const labels = val?.map((val) => val.date.substring(0, 10));
  const data = val?.map((val) => val.price);

  lineChart.value = new Chart(ctx, {
    type: "line",
    data: {
      labels: labels,
      datasets: [
        {
          label: symbol.value,
          data: data,
          fill: false,
          borderColor: "rgb(75, 192, 192)",
          tension: 0.1,
        },
      ],
    },
    options: {
      scales: {
        x: {
          title: {
            display: true,
            text: "Date",
            font: {
              size: 18,
            },
          },
        },
        y: {
          beginAtZero: true,
          title: {
            display: true,
            text: "Amount in rupees",
            font: {
              size: 18,
            },
          },
        },
      },
    },
  });
};
const GetStockHistory = () => {
  if (range.value?.length !== 2) {
    notification["error"]({
      message: "Invalid Request",
      description: "select the start and end date in range",
    });
    return;
  }
  const startDate = dayjs(range.value[0]).format("YYYY-MM-DD");
  const endDate = dayjs(range.value[1]).format("YYYY-MM-DD");
  const apiURL = "http://localhost:8000/stock_history";

  axios
    .get(apiURL, {
      params: {
        symbol: symbol.value,
        startDate,
        endDate,
      },
      headers: {
        "Content-Type": "multipart/form-data",
        token: localStorage.getItem("token"),
      },
    })
    .then((response) => {
      if (response?.data) {
        setupLineChart(response?.data?.data);
      }
    })
    .catch((error) => {
      console.error("Error while fetching stock history", error);
      notification["error"]({
        message: "Could not fetch data",
        description: error?.response?.data?.error,
      });
    });
};
</script>

<style lang="scss" scoped></style>
