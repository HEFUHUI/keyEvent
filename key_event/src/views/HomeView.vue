<template>
  <div class="home">
    <div class="keyBored">
        <h2 >按键排行榜</h2>
      <div class="item" v-for="i in keyData" :key="i.name">
        <h2>
          <el-tag>{{i.name}}</el-tag>
        </h2>
        <span>{{i.count}}</span>
      </div>
    </div>
    <key-bored-vue :data="keyData"></key-bored-vue>
  </div>
</template>

<script>
import KeyBoredVue from '../components/KeyBored.vue';

export default {
  name: 'HomeView',
  components:{
    KeyBoredVue
  },
  data() {
    return {
      keyData: []
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
</style>
