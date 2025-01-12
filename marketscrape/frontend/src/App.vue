<script lang="ts" setup>
import { ref } from "vue";

import { Button } from "./components/ui/button";

import { ArrowLeft } from "lucide-vue-next";

import ListingInputCard from "./components/ListingInputCard.vue";
import ListingInfoCard from "./components/ListingInfoCard.vue";

const listing = ref<string | null>(null);

const handleListing = (data: string) => {
  listing.value = JSON.parse(data).data;
};

const resetListing = () => {
  listing.value = null;
};
</script>

<template>
  <div>
    <ListingInputCard v-if="!listing" @listing="handleListing" />
    <div v-else class="space-y-8">
      <Button @click="resetListing">
        <ArrowLeft />
        Go Back
      </Button>
      <ListingInfoCard :listing="listing" />
    </div>
  </div>
</template>
