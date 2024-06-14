<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { store } from "./store";
</script>

<template>
  <div>
    <h1>Hello App!</h1>
    <p><strong>Current route path:</strong> {{ $route.fullPath }}</p>
    <div><RouterLink to="/">Home</RouterLink></div>
    <div>
      <label for="userId">User ID:</label>
      <input
        id="userId"
        type="text"
        v-model="store.user.id"
        placeholder="Enter user ID"
      />
    </div>
    <div>
      <RouterLink :to="{ name: 'User', params: { id: store.user.id } }"
        >User</RouterLink
      >
    </div>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
