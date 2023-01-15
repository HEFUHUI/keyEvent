<template>
  <span @click="visible = true">
    <el-dialog destroy-on-close	 v-model="visible" :title="`按键${keyEvent.name}的使用统计`">
      <div id="container"></div>
    </el-dialog>
    <slot></slot>
  </span>
</template>
<script>
import { Chart } from "@antv/g2";
export default {
  name: "KeyStatistcsVue",
  props: ["keyEvent"],
  data() {
    return {
      data: {},
      visible: false,
    };
  },
  watch: {
    visible(val) {
      if (val) {
        this.fetchData();
        const itv = setInterval(() => {
          if (window.container) {
            this.draw();
            clearInterval(itv);
          }
        }, 200);
      }
    },
  },
  mounted() {},
  methods: {
    async fetchData() {
      this.dir = (
        await this.axios.get(``)
      ).data;
      this.data = await window.hzfui.readKeyDetail(this.dir, this.keyEvent.raw_code);
    },
    draw() {
      const data = Object.keys(this.data).map(date=> new Object({year: date, value: this.data[date]}))
      const chart = new Chart({
        container: "container",
        autoFit: true,
        height: 500,
        width: 500,
      });

      chart.data(data);

      chart.tooltip({
        showMarkers: false,
      });

      chart.interval().position("year*value");

      chart.interaction("element-active");

      chart.render();
    },
  },
};
</script>