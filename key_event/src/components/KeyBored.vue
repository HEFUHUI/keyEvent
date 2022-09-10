<template>
  <div class="key-bored" :style="{backgroundColor: light ? 'black' : 'white',color: (light ? 'white' : 'black' ) + '!important'}" ref="keyBored">
    <el-form inline>
      <el-form-item label="灯光效果">
        <el-switch v-model="light"></el-switch>
      </el-form-item>
      <el-form-item label="灯光速度">
        <el-slider style="width: 400px" :min="1" v-model="lightSpeed" :max="1000"></el-slider>
      </el-form-item>
    </el-form>
    <div>
      <div class="key-row" v-for="(value,key) in keyLayout.keys" :key="`x=${key}`">
        <template v-for="(v,k) in value" :key="v.name + v.code + k">
          <div @click="showDetail(v)" class="key-item" :class="v.code"
               :style="{width: (itemWidth * (v.w_scale || 1)) + 'rem',height: itemWidth-0.43 + 'rem',fontSize: (0.8 * ((itemWidth) / 4)).toFixed(2)+'rem'}">
            <h3>{{ mapper["keyName"][v.code] }}</h3>
            <small v-if="v.code">{{ getCount(v) }}</small>
            <div class="color" :style="{backgroundColor: v.color}"></div>
          </div>
        </template>
      </div>
    </div>
    <div style="display: flex;margin: 10px">
     <el-slider :max="6" style="width: 30rem" :show-tooltip="false" :step="0.1" v-model="itemWidth" :min="2"></el-slider>
    </div>
    <el-dialog destroy-on-close append-to-body v-model="visible" :title="`按键${mapper.keyName[curr_key.code]}的使用统计`">
      <div id="container"></div>
    </el-dialog>
  </div>
</template>
<script>
import {Chart} from "@antv/g2";

export default {
  props: {
    data: {
      type: Array,
      default: () => []
    },
    mapper:{
      type: Object,
      default:()=>new Object({
        keyName:{},
        keyCode:{}
      })
    },
    // light:{
    //   type:Boolean,
    //   default:false
    // },
    // lightSpeed:{
    //   type:Number,
    //   default:500
    // }
  },
  data() {
    return {
      keyData: [],
      light:false,
      lightSpeed: 500,
      visible: false,
      keyDetail: {},
      curr_key: {},
      lightItv:0,
      keyLayout: {
        padding: [
          {direction: "vertical", start: 0, end: 1, size: "10"}
        ],
        keys: [
          [
            {code: "27"}, {code: ""}, {code: "112"},
            {code: "113"}, {code: "114"}, {code: "115"},
            {code: "116"}, {code: "117"}, {code: "118"},
            {code: "119"}, {code: "120"}, {code: "121"},
            {code: "122"}, {code: "123"},
            // 14: {code: ""},
            // 15: {code: '44'},
            // 16: {code: "145"}},
            // 17: {code: "19"}},
            // 18: {code: '174'},
            // 19: {code: "181"}},
            // 20: {code: "179"}},
            // 21: {code: "175"}}
          ],
          [
            {code: "192"}, {code: "49"}, {code: "50"}, {code: "51"},
            {code: "52"}, {code: "53"}, {code: "54"}, {code: "55"},
            {code: "56"}, {code: "57"}, {code: "48"}, {code: "189"},
            {code: "187"}, {code: '8', w_scale: 1.3}
          ],
          [
            {code: "9", w_scale: 1.3}, {code: "81"}, {code: "87"}, {code: "69"},
            {code: "82"}, {code: "84"}, {code: "89"}, {code: "85"}, {code: "73"},
            {code: "79"}, {code: "80"}, {code: "219"}, {code: "221"}, {code: '220'}
          ],
          [{code: "20", w_scale: 1.5}, {code: "65"}, {code: "83"}, {code: "68"}, {code: "70"},
            {code: "71"}, {code: "72"}, {code: "74"}, {code: "75"}, {code: "76"}, {code: "186"},
            {code: "222"}, {code: "13", w_scale: 1.85},
          ],
          [
            {code: "16-1", w_scale: 1.8}, {code: "90"}, {code: "88"}, {code: "67"},
            {code: "86"}, {code: "66"}, {code: "78"}, {code: "77"}, {code: "188"},
            {code: "190"}, {code: "191"}, {code: "16-2", w_scale: 2.65}
          ],
          [
            {code: "17-1"}, {code: "91-1"}, {code: "18-1"}, {code: "32", w_scale: 6.5},
            {code: "18-2", w_scale: 2}, {code: "93",}, {code: "17-2", w_scale: 2.2},
          ]
        ],
      },
      itemWidth: 3.0
    };
  },
  computed: {},
  watch:{
    lightSpeed(){
      // 更改灯光速度时先暂停灯光
      this.light = false;
      setTimeout(()=>{
        this.light = true;
      },100)
    },
    light(val){
      if(val){
        this.startLight()
      }else{
        this.stopLight()
      }
    }
  },
  async mounted() {
    window.onkeyup = this.up;
    window.onkeydown = this.down;
  },
  methods: {
    stopLight(){
      if(this.lightItv !== 0){
        clearInterval(this.lightItv)
        this.lightItv = 0
      }
      // 恢复按键颜色
      for(let i=0;i<this.keyData.length;i++){
        let item = this.keyData[i]
        let el = this.$refs.keyBored.querySelector(`.${item.code}`)
        if(el){
          el.style.backgroundColor = "#fff"
        }
      }
    },
    startLight(){
      // 设置定时器
      this.lightItv = setInterval(() => {
        // 获取所有按键
        let keys = document.querySelectorAll('.key-item .color')
        // 随机获取一个按键
        let key = keys[Math.floor(Math.random() * keys.length)]
        // 设置按键的颜色为随机
        key.style.backgroundColor = `rgb(${Math.floor(Math.random() * 255)},${Math.floor(Math.random() * 255)},${Math.floor(Math.random() * 255)})`
      }, this.lightSpeed)
    },
    up(e) {
      let code = e.keyCode;
      if(e.location){
        code += "-" + e.location;
      }
      const className = document.getElementsByClassName(code);
      if (className.length < 1) {
        this.$message.warning("没有找到对应的按键:" + e.key);
      } else {
        for (let i = 0; i < className.length; i++) {
          className[i].style.backgroundColor = "#fff";
        }
      }
    },
    async showDetail(key) {
      if(key.code === ""){
        return;
      }
      this.curr_key = key;
      await this.fetchKeyDetail(this.mapper.keyCode[key.code] || key.code);
      this.visible = true;
      const itv = setInterval(() => {
        if (window.container) {
          this.draw();
          clearInterval(itv);
        }
      })
    },
    async fetchKeyDetail(raw_code) {
      this.keyDetail = (
          await this.axios.get(`/statistics?key=${raw_code}`)
      ).data;
    },
    draw() {
      const data = Object.keys(this.keyDetail).map(date => new Object({year: date, value: this.keyDetail[date]}))
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
      this.isDraw = true;
    },
    down(e) {
      let code = e.keyCode;
      if(e.location){
        code += "-" + e.location;
      }
      const className = document.getElementsByClassName(code);
      if (className.length >= 1) {
        for (let i = 0; i < className.length; i++) {
          className[i].style.backgroundColor = "#ccc";
        }
      }
    },
    getCount(v) {
      const filter = this.data.filter(item => {
        return Number.parseInt(item.raw_code) === (this.mapper.keyCode[v.code] || Number.parseInt(v.code))
      })[0];
      if (filter) {
        return filter.count;
      } else {
        return ""
      }
    }
  },
};
</script>
<style scoped>
.key-bored {
  flex: 1;
  padding: 10px;
  border-radius: 10px;
  box-shadow: 1px 1px 10px #ccc;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.key-item {
  position: relative;
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 0 5px #ccc;
  margin: 0.1rem;
  box-sizing: border-box;
  padding: 0.1rem 0.2rem;
  color: #333;
  user-select: none;
  transition: 0.2s color, 0.2s background-color;
}
.key-item .color{
  position: absolute;
  top: 0.1rem;
  right: 0.1rem;
  width: .6rem;
  height: .6rem;
  border-radius: 50%;
/*  灯光效果*/
  box-shadow: 0 0 5px #fff;
  background-color: #fff;
}

.key-item:hover {
  background-color: #eee;
  cursor: pointer;
}

.key-item:active {
  background-color: #ddd;
}

.key-row {
  display: flex;
}

</style>