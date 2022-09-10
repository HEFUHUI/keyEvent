<template>
  <div class="home">
    <el-drawer v-model="showRanking">
      <template #title>
        <h2>
          按键排行榜
          <small style="font-size: 12px;font-weight: normal">总数：{{ total }}</small>
        </h2>
      </template>
      <el-table :data="keyData" :height="height">
        <el-table-column type="index" label="序号"></el-table-column>
        <el-table-column label="按键名称" prop="name">
          <template #default="{row}">
            <el-tag>{{mapper["keyName"][row.raw_code] || row.name}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="编码" prop="raw_code"></el-table-column>
        <el-table-column label="总次数" prop="count"></el-table-column>
      </el-table>
    </el-drawer>
    <div class="container">
      <div>
        <el-button @click="showRanking = !showRanking">排行榜</el-button>
        <el-divider></el-divider>
        <key-bored-vue :mapper="mapper" :lightSpeed="lightSpeed" :light="light" :data="keyData"></key-bored-vue>
      </div>
    </div>
  </div>
</template>

<script>
import KeyBoredVue from '../components/KeyBored.vue';

export default {
  name: 'HomeView',
  components:{
    KeyBoredVue,
  },
  data() {
    return {
      keyData: [],

      keyDetail: {},
      curr_key: {},
      showRanking: false,

      mapper: {
        keyName: {
          192: "`", 49: "1", 50: "2", 51: "3", 52: "4", 53: "5", 54: "6", 55: "7", 56: "8", 57: "9", 48: "0",
          189: "-", 187: "=", 8: "退格", 9: "制表", 81: "Q", 87: "W", 69: "E", 82: "R", 84: "T", 89: "Y", 85: "U",
          73: "I", 79: "O", 80: "P", 219: "[", 221: "]", 220: "\\", 20: "大写锁定", 65: "A", 83: "S", 68: "D",
          70: "F", 71: "G", 72: "H", 74: "J", 75: "K", 76: "L", 186: ";", 222: "'", 13: "回车", "16-1": "Shift(L)",
          "16-2": "Shift(R)",
          90: "Z", 88: "X", 67: "C", 86: "V", 66: "B", 78: "N", 77: "M", 188: ",", 190: ".", 191: "/",
          "17-1": "Ctrl(L)","17-2": "Ctrl(R)", "91-1": "Win", "18-1": "Alt(L)","18-2":"Alt(R)", 32: "空格", 93: "菜单", 27: "Esc", 112: "F1", 113: "F2", 114: "F3",
          115: "F4", 116: "F5", 117: "F6", 118: "F7", 119: "F8", 120: "F9", 121: "F10", 122: "F11", 123: "F12",
          145: "滚动锁定", 19: "暂停", 174: "音量-", 181: "🎶", 179: "▶", 175: "音量+"
        },
        keyCode: {
          192: 192, 49: 49,50: 50, 51: 51, 52:52, 53: 53,"17-1": 162,"17-2": 163,"18-1": 164,"18-2": 165,
          "16-1": 160,
          "16-2": 161,
          "91-1": 91,
        }
      }
    }
  },
  computed:{
    height(){
      return window.innerHeight - 100;
    },
    total(){
      const countList = this.keyData.map(i=>i.count);
      let t = 0;
      countList.forEach(count => t+=count)
      return t;
    }
  },
  methods:{
    async fetchData(){
      const data = (await this.axios.get("")).data
      // golang记录的时候，没有记录到raw_code，所以这里需要手动添加
      Object.keys(data).forEach(key => {
        data[key].raw_code = key
      })
      this.keyData = Object.values(data).sort((i, o) => {
        return o.count - i.count
      });
    },

  },
  mounted(){
    setInterval(() => {
      this.fetchData();
    }, 1000);
  },
  created(){
    this.fetchData();
  }
}
</script>
<style>
.container{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80vh;
}
.key-ranking{
  height: calc(100vh - 100px);
}
</style>
