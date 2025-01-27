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

import { Clock, Image as ImageIcon, MapPin, Tag } from "lucide-vue-next";

import {
  capitalizeFirstLetter,
  toTitleCase,
  useListingUtils,
} from "@/lib/utils";
import ImageGallery from "./ImageGallery.vue";
import { main } from "./../../wailsjs/go/models";

const props = defineProps<{
  listing: main.Root;
}>();

const { creationTime, filteredCategories, formatPrice } =
  useListingUtils(props);
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
      <CardDescription
        v-if="filteredCategories?.length"
        class="space-y-2 space-x-2"
      >
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

      <div class="grid grid-flow-row text-sm gap-1">
        <div class="flex items-center">
          <Tag class="h-4 w-4 text-muted-foreground mr-2" />
          <span>{{ listing.target.attribute_data[0].label }}</span>
        </div>
        <div class="flex items-center">
          <Clock class="h-4 w-4 text-muted-foreground mr-2" />
          <span>{{ capitalizeFirstLetter(creationTime) }}</span>
        </div>
      </div>

      <Separator />
    </CardContent>
    <CardFooter class="flex justify-between">
      <div class="flex items-center text-sm text-muted-foreground">
        <MapPin class="h-4 w-4 mr-2" />
        <span>{{ listing.target.location_text.text }}</span>
      </div>
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
