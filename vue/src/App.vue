<script setup lang="ts">
import { RouterView } from "vue-router";
import IgntHeader from "./components/IgntHeader.vue";

const items = [
  {
    label: "Hot Storage Browser",
    to: "/",
  },
  {
    label: "Wallet",
    to: "/wallet",
  },
  
];
</script>

<template>
  <IgntHeader :nav-items="items"></IgntHeader>
  <RouterView />
</template>
