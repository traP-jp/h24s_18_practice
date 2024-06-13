<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount  } from "vue";

interface Item {
  name: string;
  time: string;
}



const newTweet = ref("");
const newTime = ref("");
let startTime = "";


const interval = ref()
onMounted(() => {
  setInterval(() => {
    let time = new Date();
    let hour = time.getHours();
    let minute = time.getMinutes();
    let second = time.getSeconds();
    newTime.value = hour + "時" + minute + "分" + second + "秒";
  }, 1000);
  
});

onBeforeUnmount(() => {
    if(interval.value){
        clearInterval(interval.value);
    }
})

const items = ref<Item[]>([
  { name: "お腹空いた", time:  "1時15分33秒"},
  { name: "ねむいよー",time: "1時15分33秒"},
]);

const addItem = () => {
  items.value.unshift({ name: newTweet.value, time:newTime.value});
};


</script>

<template>
  <div>
    <div>
      <label>
        ツイート
        <input v-model="newTweet" type="text" />
      </label>

      <button @click="addItem(), (newTweet = '')">
        tweet
      </button>
      <button @click="items = []">ツイート全部消す</button>

      <div>ツイート一覧</div>
      <ul>
        <li v-for="item in items.slice()" :key="item.name">
          <div>
            tweet: {{ item.name }} &emsp;&emsp;&emsp;&emsp;&emsp;time
            {{ item.time}}
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<style></style>
