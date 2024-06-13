<script setup lang="ts">
import { ref } from "vue";

interface Item {
  name: string;
  time: number;
}

const items = ref<Item[]>([
  { name: "お腹空いた", time: 100 },
  { name: "ねむいよー", time: 160 },
]);

const newTweet = ref("");
const newTime = ref();

const addItem = () => {
  items.value.unshift({ name: newTweet.value, time: newTime.value });
};
</script>

<template>
  <div>
    <div>
      <label>
        ツイート
        <input v-model="newTweet" type="text" />
      </label>
      <label>
        時間
        <input v-model="newTime" type="number" />
      </label>
      <button @click="addItem(), (newTweet = ''), (newTime = ref())">
        tweet
      </button>
      <button @click="items = []">ツイート全部消す</button>

      <div>ツイート一覧</div>
      <ul>
        <li v-for="item in items.slice()" :key="item.name">
          <div>
            tweet: {{ item.name }} &emsp;&emsp;&emsp;&emsp;&emsp;time
            {{ item.time }}
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<style></style>
