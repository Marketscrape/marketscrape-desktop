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

import {
  Clock,
  Image as ImageIcon,
  MapPin,
  Tag,
  Car,
  Fuel,
  Gauge,
  Palette,
  Cog,
} from "lucide-vue-next";

import {
  capitalizeFirstLetter,
  toTitleCase,
  useListingUtils,
} from "@/lib/utils";
import ImageGallery from "./ImageGallery.vue";
import { main } from "./../../wailsjs/go/models";

const getColourStyle = (color: string) => {
  return {
    backgroundColor: color.toLowerCase(),
    width: "1rem",
    height: "1rem",
    display: "inline-block",
    borderRadius: "9999px",
    border: "1px solid #262626",
    marginLeft: "0.25rem",
  };
};

const props = defineProps<{
  listing: main.Root;
}>();

const vehicleData = computed(() => props.listing.target.vehicle_data);

const formattedOdometer = computed(() => {
  const odometerData = vehicleData.value?.vehicle_odometer_data;
  return odometerData
    ? `${odometerData.value.toLocaleString()} ${odometerData.unit}`
    : "N/A";
});

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
        class="space-y-2 space-x-1"
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

      <div class="grid grid-cols-2 gap-4 text-sm">
        <div class="space-y-2">
          <div class="flex items-center space-x-2">
            <Car class="h-4 w-4 text-muted-foreground" />
            <span>
              {{ vehicleData.vehicle_make_display_name }}
              {{ vehicleData.vehicle_model_display_name }}
              {{ vehicleData.vehicle_trim_display_name }}
            </span>
          </div>
          <div class="flex items-center space-x-2">
            <Gauge class="h-4 w-4 text-muted-foreground" />
            <span>{{ formattedOdometer }}</span>
          </div>
          <div class="flex items-center space-x-2">
            <Fuel class="h-4 w-4 text-muted-foreground" />
            <span>{{ vehicleData.vehicle_fuel_type || "N/A" }}</span>
          </div>
          <div class="flex items-center space-x-2">
            <Cog class="h-4 w-4 text-muted-foreground" />
            <span>{{ vehicleData.vehicle_transmission_type || "N/A" }}</span>
          </div>
        </div>
        <div class="space-y-2">
          <div class="flex items-center space-x-2">
            <Palette class="h-4 w-4 text-muted-foreground" />
            <span class="flex items-center">
              Exterior:
              <span
                v-if="vehicleData.vehicle_exterior_color"
                :style="getColourStyle(vehicleData.vehicle_exterior_color)"
              >
              </span>
            </span>
          </div>
          <div class="flex items-center space-x-2">
            <Palette class="h-4 w-4 text-muted-foreground" />
            <span class="flex items-center">
              Interior:
              <span
                v-if="vehicleData.vehicle_interior_color"
                :style="getColourStyle(vehicleData.vehicle_interior_color)"
              >
              </span>
            </span>
          </div>
          <div
            class="flex items-center space-x-2"
            v-if="vehicleData.vehicle_identification_number"
          >
            <Tag class="h-4 w-4 text-muted-foreground" />
            <span>VIN: {{ vehicleData.vehicle_identification_number }}</span>
          </div>
          <div class="flex items-center space-x-2">
            <Clock class="h-4 w-4 text-muted-foreground" />
            <span>{{ capitalizeFirstLetter(creationTime) }}</span>
          </div>
        </div>
      </div>

      <Separator />

      <div
        v-if="
          vehicleData.vehicle_features && vehicleData.vehicle_features.length
        "
        class="space-y-2"
      >
        <h4 class="font-semibold">Features</h4>
        <div class="flex flex-wrap gap-1">
          <Badge
            v-for="feature in vehicleData.vehicle_features"
            :key="feature"
            variant="secondary"
          >
            {{ feature }}
          </Badge>
        </div>
      </div>
    </CardContent>
    <CardFooter class="flex justify-between">
      <div class="flex items-center space-x-2 text-sm text-muted-foreground">
        <MapPin class="h-4 w-4" />
        <span>{{ listing.target.location_text.text }}</span>
      </div>
      <Dialog>
        <DialogTrigger asChild>
          <Button variant="outline" size="sm">
            <ImageIcon class="h-4 w-4 mr-2" />
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
