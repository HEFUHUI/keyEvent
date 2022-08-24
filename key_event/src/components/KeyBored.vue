<template>
  <div style="flex: 1">
    <div id="view" style="width: 100%"></div>
  </div>
  <div>
    总敲击次数：<b>{{ total }}</b>
  </div>
</template>
<script>
import { Chart,Util } from '@antv/g2';

export default {
  props: {
    data:{
      type: Array,
      default: () => []
    }
  },
  computed:{
    total(){
      const countList = this.data.map(i=>i.count);
      let t = 0;
      countList.forEach(count => t+=count)
      return t;
    }
  },
  data() {
    return {};
  },
  mounted() {
    this.renderView();
  },
  methods:{
    renderView(){
      const chart = new Chart({
        container: 'view',
        autoFit: true,
        height: 800,
        width: 800
      });
      const data = this.data.map(item=>new Object({
        type:item.name,
        value: item.count
      }));
      chart.data(data); // 加载数据


      chart.coordinate('theta', {
        radius: 0.75
      });
      chart.tooltip({
        showMarkers: true
      });

      chart
          .interval()
          .adjust('stack')
          .position('value')
          .color('type', ['#063d8a', '#1770d6', '#47abfc', '#38c060'])
          .style({ opacity: 0.4 })
          .state({
            active: {
              style: (element) => {
                const shape = element.shape;
                return {
                  matrix: Util.zoom(shape, 1.1),
                }
              }
            }
          })
          .label('type', (val) => {
            const opacity = val === '四线及以下' ? 1 : 0.5;
            return {
              offset: -30,
              style: {
                opacity,
                fill: 'white',
                fontSize: 12,
                shadowBlur: 2,
                shadowColor: 'rgba(0, 0, 0, .45)',
              },
              content: (obj) => {
                return obj.type + '\n' + obj.value + '%';
              },
            };
          });


      chart.interaction('element-single-selected');
      chart.render()
    }
  },
};
</script>
<style scoped>
div{
  transition: .2s;
}

</style>