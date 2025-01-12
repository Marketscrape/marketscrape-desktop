<script setup lang="ts">
import { ref, computed } from "vue";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";

import { MapPin, Tag, Clock, Image as ImageIcon } from "lucide-vue-next";

import ImageGallery from "./ImageGallery.vue";
import { cn, toTitleCase, capitalizeFirstLetter } from "@/lib/utils";
import { formatDistanceToNow } from "date-fns";

const props = defineProps<{
  listing: any;
}>();

const creationTime = computed(() =>
  formatDistanceToNow(new Date(props.listing.target.creation_time * 1000), {
    addSuffix: true,
  }),
);
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle class="space-y-1">
        <div class="text-xl font-semibold leading-none tracking-tight">
          {{
            listing.marketplace_listing_renderable_target
              .marketplace_listing_title
          }}
        </div>
        <div>
          {{ listing.target.listing_price?.formatted_amount_zeros_stripped || "N/A" }}
        </div>
      </CardTitle>
      <CardDescription class="space-x-1">
        <Badge
          v-for="category in listing.marketplace_listing_renderable_target
            ?.seo_virtual_category?.taxonomy_path || []"
          :key="category.id"
          class="text-xs"
        >
          {{ toTitleCase(category.seo_info?.seo_url?.replace("-", " ") || "Unknown") }}
        </Badge>
      </CardDescription>
    </CardHeader>
    <CardContent class="space-y-4">
      <p>
        {{ listing.target.redacted_description?.text || "No description available." }}
      </p>
      <Separator />
      <div class="grid grid-flow-row text-sm text-muted-foreground gap-1">
        <div class="flex items-center space-x-1">
          <Tag class="h-4 w-4" />
          <span>{{ listing.target.attribute_data?.[0]?.label || "N/A" }}</span>
        </div>
        <div class="flex items-center space-x-1">
          <Clock class="h-4 w-4" />
          <span>{{ capitalizeFirstLetter(creationTime) }}</span>
        </div>
        <div class="flex items-center space-x-1">
          <MapPin class="h-4 w-4" />
          <span>{{ listing.target.location_text?.text || "Location not specified." }}</span>
        </div>
      </div>
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
          <ImageGallery :images="listing.target.listing_photos || []" />
        </DialogContent>
      </Dialog>
    </CardFooter>
  </Card>
</template>
