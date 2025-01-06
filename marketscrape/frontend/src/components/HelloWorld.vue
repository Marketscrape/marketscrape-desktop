<script setup lang="ts">
import { ref } from "vue";
import { Input } from "./ui/input/index.js";
import { Button } from "./ui/button/index.js";

import { GetMarketplaceListing } from "../../wailsjs/go/main/App.js";

const id = ref("");
const res = ref("");

async function getMarketplaceListing() {
  if (!id.value) {
    return;
  }
  GetMarketplaceListing(id.value).then((resp) => (res.value = resp));
}
</script>

<template>
  <Input v-model="id" />
  <Button @click="getMarketplaceListing">Submit</Button>
  {{ res }}
</template>
