<script setup lang="ts">
import { ref } from "vue";

import { Sparkles } from "lucide-vue-next";

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
        class="object-cover rounded-md"
      />
      <div
        v-if="
          images[currentIndex]?.accessibility_caption !==
          'No photo description available.'
        "
        class="absolute bottom-2 left-1/2 transform -translate-x-1/2 bg-black bg-opacity-70 text-white text-sm px-4 py-2 rounded-md flex items-center"
      >
        <Sparkles class="h-4 w-4 mr-2" />
        {{ images[currentIndex].accessibility_caption }}
      </div>
    </div>

    <div class="flex gap-2 mt-4 overflow-x-auto w-full justify-center p-1">
      <div v-for="(img, index) in images" :key="index" class="flex-shrink-0">
        <img
          :src="img.image.uri"
          :alt="`Thumbnail ${index + 1}`"
          @click="setImage(index)"
          class="w-16 h-16 object-cover rounded-md cursor-pointer border-2 transition-transform duration-200"
          :class="{
            'scale-110': currentIndex === index,
          }"
        />
      </div>
    </div>
  </div>
</template>
