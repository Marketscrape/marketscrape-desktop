<script setup lang="ts">
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";
import { Sparkles } from "lucide-vue-next";
import { ref } from "vue";

defineProps(["images"]);

const currentIndex = ref(0);

const setImage = (index: number) => {
  currentIndex.value = index;
};
</script>

<template>
  <div class="flex flex-col items-center p-4">
    <div class="relative w-full max-w-lg">
      <img
        :src="images[currentIndex].image.uri"
        :alt="`Image ${currentIndex + 1}`"
        class="w-full h-auto object-cover rounded-md"
      />
      <div
        v-if="
          images[currentIndex]?.accessibility_caption &&
          images[currentIndex].accessibility_caption !==
            'No photo description available.'
        "
        class="absolute bottom-2 left-1/2 transform -translate-x-1/2 bg-black bg-opacity-70 text-white text-sm px-4 py-2 rounded-md flex items-center"
      >
        <Sparkles class="h-4 w-4 mr-2" />
        {{ images[currentIndex].accessibility_caption }}
      </div>
    </div>

    <ScrollArea class="w-full max-w-lg mt-4 pb-4">
      <div class="flex w-max space-x-4 overflow-y-hidden">
        <div v-for="(img, index) in images" :key="index" class="flex-shrink-0">
          <img
            :src="img.image.uri"
            :alt="`Thumbnail ${index + 1}`"
            @click="setImage(index)"
            class="w-16 h-16 object-cover rounded-md cursor-pointer border-2 transition-transform duration-200"
            :class="{
              'scale-110 border-primary': currentIndex === index,
              'border-transparent': currentIndex !== index,
            }"
          />
        </div>
      </div>
      <ScrollBar orientation="horizontal" />
    </ScrollArea>
  </div>
</template>
