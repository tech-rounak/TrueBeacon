import { createApp } from 'vue'
import App from './App.vue'
import router from "./router"
import Antd from "ant-design-vue";
import dayjs from "dayjs";
import customParseFormat from "dayjs/plugin/customParseFormat";
const app = createApp(App);
dayjs.extend(customParseFormat)
app.use(router);
app.use(Antd);
app.mount('#app')
