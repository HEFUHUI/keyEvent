<template>
  <div class="home">
    <div class="keyBored">
        <h2 >按键排行榜</h2>
      <key-statistics :keyEvent="i" class="item" v-for="i in keyData" :key="i.name">
        <h2>
          <el-tag>{{i.name}}</el-tag>
        </h2>
        <span>{{i.count}}</span>
      </key-statistics>
    </div>
    <key-bored-vue :data="keyData"></key-bored-vue>
    <div class="total">
      总敲击次数：<b>{{ total }}</b>
    </div>
    
  </div>
</template>

<script>
import KeyBoredVue from '../components/KeyBored.vue';
import KeyStatistics from "../components/KeyStatistics.vue"

export default {
  name: 'HomeView',
  components:{
    KeyBoredVue,
    KeyStatistics
  },
  data() {
    return {
      keyData: []
    }
  },
  computed:{
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
      this.keyData = Object.values(data).sort((i, o) => {
        return o.count - i.count
      });
    }
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
*{
  transition: 0.2s ease;
}
.item{
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}
.item:hover{
  background: rgba(0, 0, 0, 0.229);
  padding: 10px;
  cursor: pointer;
  color: #fff;
}
.keyBored{
  width: 400px;
  box-sizing: border-box;
  height: 100vh;
  padding: 10px 20px;
  overflow-x: scroll;
  box-shadow: 1px 0 10px #ddd;
}
.keyBored::-webkit-scrollbar{
  width: 5px;
}
.keyBored::-webkit-scrollbar-thumb{
  background: #ddd;
  border-radius: 10px;
}
.home{
  display: flex;
}
.total{
  padding: 10px;
  color: var(--el-color-primary);
}
</style>
