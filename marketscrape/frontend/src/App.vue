<script lang="ts" setup>
import { ref } from "vue";

import { Button } from "./components/ui/button";

import { ArrowLeft } from "lucide-vue-next";

import ListingInputCard from "./components/ListingInputCard.vue";
import ProductListingCard from "./components/ProductListingCard.vue";
import VehicleListingCard from "./components/VehicleListingCard.vue";

import { main } from "./../wailsjs/go/models";

enum Categories {
  Product = "FOR_SALE_ITEM",
  Vehicle = "AUTOS_VEHICLE",
}

const listing = ref<main.Root | null>(null);
const listingType = ref<string | null>(null);
const error = ref<string | null>(null);

const handleListing = (data: string) => {
  try {
    const parsedListing = main.Root.createFrom(data);
    listing.value = parsedListing;
    listingType.value = parsedListing.product_details_type;
  } catch (err) {
    error.value = "Ran into a problem parsing data";
  }
};

const resetListing = () => {
  listing.value = null;
  listingType.value = null;
  error.value = null;
};
</script>

<template>
  <div>
    <ListingInputCard v-if="!listing" @listing="handleListing" />
    <div v-else class="space-y-8">
      <div v-if="error">
        {{ error }}
      </div>
      <div v-else>
        <Button @click="resetListing">
          <ArrowLeft />
          Go Back
        </Button>
        <ProductListingCard
          v-if="listingType === Categories.Product"
          :listing="listing"
        />
        <VehicleListingCard
          v-if="listingType === Categories.Vehicle"
          :listing="listing"
        />
      </div>
    </div>
  </div>
</template>
