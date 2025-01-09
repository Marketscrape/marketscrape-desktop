<script lang="ts" setup>
import { ref } from "vue";

import { Button } from "./components/ui/button";

import ListingInputCard from "./components/ListingInputCard.vue";

const listingData = ref<string | null>(null);

const handleListing = (data: string) => {
  listingData.value = data;
};

const resetListing = () => {
  listingData.value = null;
};
</script>

<template>
  <div>
    <ListingInputCard v-if="!listingData" @listing="handleListing" />
    <div v-else class="overflow-y-auto">
      <div>
        <h2 class="text-lg font-semibold mb-2">Listing Details</h2>
        <pre>
          {{ listingData }}
        </pre>
        <Button @click="resetListing">Analyze Another Listing</Button>
      </div>
    </div>
  </div>
</template>
