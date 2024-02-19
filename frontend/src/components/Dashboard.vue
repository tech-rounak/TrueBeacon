<template>
  <div>
    <h1>Welcome to True Beacon</h1>
    <a-tabs
      v-model:activeKey="activeKey"
      size="large"
      centered
      class="dashboard-tab"
    >
      <a-tab-pane key="1" tab="Holdings"><Holdings /></a-tab-pane>
      <a-tab-pane key="2" tab="Chart"><StockHistory /></a-tab-pane>
      <a-tab-pane key="3" tab="Profile"><Profile /></a-tab-pane>
      <a-tab-pane key="4" tab="Trade"><Trade /></a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup>
import Holdings from "./Holdings.vue";
import StockHistory from "./StockHistory.vue";
import Profile from "./Profile.vue";
import Trade from "./Trade.vue";
</script>

<style lang="css" scoped>
.dashboard-tab {
  width: 1000px;
  margin: auto;
}
</style>
