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
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { computed } from "vue";

import {
  Car,
  Clock,
  Cog,
  Fuel,
  Gauge,
  Image as ImageIcon,
  MapPin,
  Palette,
  User,
  Users,
  Zap,
} from "lucide-vue-next";

import {
  capitalizeFirstLetter,
  textToNumber,
  toTitleCase,
  useListingUtils,
} from "@/lib/utils";
import { main } from "./../../wailsjs/go/models";
import ImageGallery from "./ImageGallery.vue";

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

const vehicleData = computed(() => props.listing.target.vehicle_data!);

const numberOfOwners = computed(() => {
  const ownersText = vehicleData.value?.vehicle_number_of_owners;
  if (!ownersText) return null;
  return textToNumber(ownersText);
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
      <blockquote className="border-l-2 pl-6 italic whitespace-pre-wrap">
        {{ listing.target.redacted_description.text }}
      </blockquote>

      <Separator />

      <div class="grid grid-cols-2 gap-y-1 text-sm">
        <div class="flex items-center">
          <Car class="h-4 w-4 text-muted-foreground mr-2" />
          <span>
            {{ vehicleData.vehicle_make_display_name }}
            {{ vehicleData.vehicle_model_display_name }}
            {{ vehicleData.vehicle_trim_display_name }}
          </span>
        </div>
        <div class="flex items-center">
          <Gauge class="h-4 w-4 text-muted-foreground mr-2" />
          <span>
            {{
              vehicleData.vehicle_odometer_data.value.toLocaleString("en-US")
            }}
            {{ vehicleData.vehicle_odometer_data.unit.toLowerCase() }}
          </span>
        </div>
        <div v-if="vehicleData.vehicle_fuel_type" class="flex items-center">
          <Fuel class="h-4 w-4 text-muted-foreground mr-2" />
          <span>{{ toTitleCase(vehicleData.vehicle_fuel_type) }}</span>
        </div>
        <div
          v-if="vehicleData.vehicle_transmission_type"
          class="flex items-center"
        >
          <Cog class="h-4 w-4 text-muted-foreground mr-2" />
          <span>{{ toTitleCase(vehicleData.vehicle_transmission_type) }}</span>
        </div>
        <div class="flex items-center">
          <Palette class="h-4 w-4 text-muted-foreground mr-2" />
          <div class="flex items-center space-x-2">
            <span class="flex items-center">
              Exterior:
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger as-child>
                    <span
                      v-if="vehicleData.vehicle_exterior_color"
                      :style="
                        getColourStyle(vehicleData.vehicle_exterior_color)
                      "
                      class="cursor-help"
                    >
                    </span>
                  </TooltipTrigger>
                  <TooltipContent>
                    {{ toTitleCase(vehicleData.vehicle_exterior_color) }}
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            </span>

            <Separator orientation="vertical" class="h-4" />

            <span class="flex items-center">
              Interior:
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger as-child>
                    <span
                      v-if="vehicleData.vehicle_interior_color"
                      :style="
                        getColourStyle(vehicleData.vehicle_interior_color)
                      "
                      class="cursor-help"
                    >
                    </span>
                  </TooltipTrigger>
                  <TooltipContent>
                    {{ toTitleCase(vehicleData.vehicle_interior_color) }}
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            </span>
          </div>
        </div>
        <div
          v-if="vehicleData.vehicle_specifications.horse_power"
          class="flex items-center"
        >
          <Zap class="h-4 w-4 text-muted-foreground mr-2" />
          <span>
            {{ vehicleData.vehicle_specifications.horse_power.value }}
            {{
              vehicleData.vehicle_specifications.horse_power.units.toUpperCase()
            }}
          </span>
        </div>
        <div class="flex items-center" v-if="numberOfOwners">
          <User
            v-if="numberOfOwners === 1"
            class="h-4 w-4 text-muted-foreground mr-2"
          />
          <Users v-else class="h-4 w-4 text-muted-foreground mr-2" />
          <span>
            {{ numberOfOwners === 1 ? "1 owner" : `${numberOfOwners} owners` }}
          </span>
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
