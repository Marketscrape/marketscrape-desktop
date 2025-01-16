<script setup lang="ts">
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import { Separator } from "@/components/ui/separator";
import { computed } from "vue";

import { Clock, Image as ImageIcon, MapPin, Tag } from "lucide-vue-next";

import { capitalizeFirstLetter, toTitleCase } from "@/lib/utils";
import { formatDistanceToNow } from "date-fns";
import ImageGallery from "./ImageGallery.vue";
import { main } from "./../../wailsjs/go/models";

const props = defineProps<{
  listing: any;
}>();

const creationTime = computed(() =>
  formatDistanceToNow(new Date(props.listing.target.creation_time * 1000), {
    addSuffix: true,
  }),
);

const filteredCategories = computed(() =>
  props.listing.marketplace_listing_renderable_target.seo_virtual_category.taxonomy_path.filter(
    (category: main.TaxonomyPathItem) =>
      category.seo_info.seo_url.trim() !== "",
  ),
);

const formatPrice = computed(() => {
  const { amount, currency } = props.listing.target.listing_price;

  if (!amount || !currency) {
    return "";
  }

  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency,
    minimumFractionDigits: 0,
    maximumFractionDigits: 2,
  }).format(parseFloat(amount));
});
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle class="space-y-1">
        <div class="text-xl font-semibold leading-none tracking-tight">
          {{ listing.target.marketplace_listing_title }}
        </div>
        <div>
          {{ formatPrice }}
        </div>
      </CardTitle>
      <CardDescription class="space-x-1">
        <Badge
          v-for="category in filteredCategories"
          :key="category.id"
          class="text-xs"
        >
          {{ toTitleCase(category.seo_info.seo_url.replace("-", " ")) }}
        </Badge>
      </CardDescription>
    </CardHeader>
    <CardContent class="space-y-4">
      <p class="whitespace-pre-wrap">
        {{ listing.target.redacted_description.text }}
      </p>

      <Separator />

      <div class="grid grid-flow-row text-sm text-muted-foreground gap-1">
        <div class="flex items-center space-x-1">
          <Tag class="h-4 w-4" />
          <span>{{ listing.target.attribute_data[0].label }}</span>
        </div>
        <div class="flex items-center space-x-1">
          <Clock class="h-4 w-4" />
          <span>{{ capitalizeFirstLetter(creationTime) }}</span>
        </div>
        <div class="flex items-center space-x-1">
          <MapPin class="h-4 w-4" />
          <span>{{ listing.target.location_text.text }}</span>
        </div>
      </div>

      <Separator />
    </CardContent>
    <CardFooter class="flex justify-end">
      <Dialog>
        <DialogTrigger asChild>
          <Button variant="outline" size="sm">
            <ImageIcon class="h-4 w-4" />
            Image Gallery
          </Button>
        </DialogTrigger>
        <DialogContent class="sm:max-w-[700px] p-6">
          <ImageGallery :images="listing.target.listing_photos" />
        </DialogContent>
      </Dialog>
    </CardFooter>
  </Card>
</template>
