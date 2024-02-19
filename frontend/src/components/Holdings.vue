<template>
  <div class="holdings">
    <Card heading="Total Investment" :value="totalAvgPrice" />
    <Card heading="Current Value" :value="totalCurrVal" />
    <Card heading="Days P&L" :value="totalDayPL" />
    <Card heading="Total P&L" :value="totalPL" />
  </div>
  <a-table
    :columns="columns"
    :data-source="data"
    :pagination="false"
    size="small"
    style="width: 80%; margin: auto"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.dataIndex === 'day_change_percentage'">
        <span :class="getClassName(record.day_change_percentage, 'cell')">
          {{
            record.day_change_percentage > 0
              ? `+${record.day_change_percentage}`
              : record.day_change_percentage
          }}%
        </span>
      </template>
      <template v-if="column.dataIndex === 'pnl'">
        <span :class="getClassName(record.pnl, 'cell')">
          {{ record.pnl > 0 ? `+${record.pnl}` : record.pnl }}
        </span>
      </template>
    </template>
    <template #summary>
      <a-table-summary-row>
        <a-table-summary-cell :col-span="2" style="font-weight: bold"
          >Total</a-table-summary-cell
        >
        <a-table-summary-cell>
          <span style="text-align: center; font-weight: 600">
            {{ totalAvgPrice }}
          </span>
        </a-table-summary-cell>
        <a-table-summary-cell>
          <span :class="getClassName(totalCurrVal, 'footer')">
            {{ totalCurrVal }}
          </span>
        </a-table-summary-cell>
        <a-table-summary-cell>
          <span :class="getClassName(totalCurrVal, 'footer')">
            {{ totalDayPL }}
          </span>
        </a-table-summary-cell>
      </a-table-summary-row>
    </template>
  </a-table>
</template>

<script setup>
import axios from "axios";
import Card from "./Card.vue";
import { onMounted, ref } from "vue";
import { notification } from "ant-design-vue";
import { getbaseURI } from "./../utils"
const totalDayPL = ref(0);
const totalCurrVal = ref(0);
const totalAvgPrice = ref(0);
const totalPL = ref(0);
const data = ref([]);
const columns = ref([
  {
    title: "Instrument Name",
    dataIndex: "tradingsymbol",
  },
  {
    title: "Quantity",
    dataIndex: "quantity",
  },
  {
    title: "Avg Price",
    dataIndex: "average_price",
  },

  {
    title: "Cur Price",
    dataIndex: "close_price",
    sorter: (a, b) => a.close_price - b.close_price,
  },
  {
    title: "P & L",
    dataIndex: "pnl",
    sorter: (a, b) => a.pnl - b.pnl,
  },
  {
    title: "Day chg.",
    dataIndex: "day_change_percentage",
    sorter: (a, b) => a.day_change_percentage - b.day_change_percentage,
  },
]);
onMounted(() => {
  FetchHoldings();
});
const FetchHoldings = () => {
  const apiURL = getbaseURI()+"/portfolio/holdings";

  axios
    .get(apiURL, {
      headers: {
        "Content-Type": "multipart/form-data",
        token: localStorage.getItem("token"),
      },
    })
    .then((response) => {
      data.value = response?.data?.data?.holdings?.map((val) => val);
      if (data.value) {
        data.value.forEach((share) => {
          totalCurrVal.value += share.close_price;
          totalDayPL.value += share.day_change;
          totalPL.value += share.pnl;
          totalAvgPrice.value += share.average_price;
        });
      }
    })
    .catch((error) => {
      console.error("Error uploading file:", error);
      notification["error"]({
        message: "Could not fetch holdings",
        description: error?.response?.data?.error,
      });
    });
};

const getClassName = (num, type) => {
  if (num > 0) return "positive-" + type;
  return "negative-" + type;
};
</script>

<style scoped>
.positive-footer {
  color: green;
  text-align: center;
  font-weight: bold;
}
.negative-footer {
  color: red;
  text-align: center;
  font-weight: bold;
}
.positive-cell {
  color: green;
  text-align: center;
  font-weight: 500;
}
.negative-cell {
  color: red;
  text-align: center;
  font-weight: 500;
}
.holdings {
  display: flex;
  flex-direction: row;
  width: 80%;
  margin: 20px auto;
  justify-content: space-evenly;
}
</style>
