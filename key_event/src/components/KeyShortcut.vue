<template>
  <span>
    <el-drawer v-model="visible" title="快捷键按键统计">
      <el-table :data="keyData" :height="props.height">
        <el-table-column type="index" label="序号"></el-table-column>
        <el-table-column label="按键名称" prop="name">
          <template #default="{row}">
            <el-tag>{{props.mapper["keyName"][row.raw_code] || row.name}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="编码" prop="raw_code"></el-table-column>
        <el-table-column label="总次数" prop="count"></el-table-column>
      </el-table>
    </el-drawer>
    <el-button style="margin: 0 10px" @click="visible = true">快捷键排行</el-button>
  </span>
</template>

<script setup>

import {ref, defineProps, inject} from "vue";

const keyData = ref([]);
const visible = ref(false);
const props = defineProps(['mapper', 'height']);
const axios = inject("$axios");
async function fetchData(){
  const dir = (await axios.get("")).data;
  window.hzfui.getShortcutKey(dir).then((data) => {
    Object.keys(data).forEach((key) => {
      keyData.value.push({
        name: key.split("+").map((k) => props.mapper["keyName"][k] || k).join("+"),
        count: data[key].count,
      });
    });
  });
}

fetchData();
</script>

<style scoped>

</style>