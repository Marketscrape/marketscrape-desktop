<script setup lang="ts">
import { ref, onMounted } from "vue";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Skeleton } from "@/components/ui/skeleton";
import { SendIcon, Loader2, Sparkles } from "lucide-vue-next";

import OllamaIcon from "./OllamaIcon.vue";
import { GetOllamaModels, PostOllamaModel } from "../../wailsjs/go/main/App.js";
import { main } from "./../../wailsjs/go/models";

const models = ref<main.Model[] | null>(null);
const error = ref<string | null>(null);
const selectedModel = ref<string | undefined>(undefined);
const message = ref("");
const conversation = ref<{ role: string; content: string }[]>([]);
const isMessageLoading = ref<boolean>(false);

const props = defineProps<{
  listing: main.Root;
}>();

onMounted(async () => {
  try {
    const res = await GetOllamaModels();
    models.value = res;
  } catch (err) {
    error.value = "Failed to fetch Ollama models";
  }
});

const askModel = async ({ message }: { message: string }) => {
  if (!selectedModel.value || !message) return;

  isMessageLoading.value = true;

  // Add a temporary loading message from the assistant
  conversation.value.push({ role: "assistant", content: "" });

  try {
    // Use chat history as context
    const chatHistory = conversation.value
      .map((msg) => `${msg.role}: ${msg.content}`)
      .join("\n");

    const formattedMessage = `Here is the chat so far:\n\n${chatHistory}\n\nUser's new query:\n${message}`;
    const modelRes = await PostOllamaModel(
      selectedModel.value,
      formattedMessage,
    );

    // Replace the temporary loading message with the actual response
    conversation.value[conversation.value.length - 1] = modelRes.message;
  } catch {
    // Update the loading message with the error
    conversation.value[conversation.value.length - 1] = {
      role: "assistant",
      content: "Failed to send message.",
    };
  } finally {
    isMessageLoading.value = false;
  }
};

const sendMessage = async () => {
  const trimmedMessage = message.value.trim();
  if (!trimmedMessage) return;

  conversation.value.push({ role: "user", content: trimmedMessage });
  message.value = "";

  askModel({ message: trimmedMessage });
};

const askModelsOpinion = () => {
  if (!props.listing) return;

  const formattedMessage = `Using the following JSON object, provide a short and concise opinion (maximum 2 sentences) on whether this item is a good deal. Evaluate factors like price, condition, features, and any other relevant details. Suggest a better price if necessary:\n\n${JSON.stringify(props.listing, null, 2)}`;

  conversation.value.push({
    role: "user",
    content: `What do you think of the ${props.listing.target.marketplace_listing_title}?`,
  });

  askModel({ message: formattedMessage });
};
</script>

<template>
  <div class="flex justify-end p-4">
    <Popover>
      <PopoverTrigger as-child>
        <div
          class="h-12 w-12 p-2 rounded-full bg-white flex items-center justify-center shadow cursor-pointer"
        >
          <OllamaIcon class="h-7 w-7" />
        </div>
      </PopoverTrigger>
      <PopoverContent align="end" class="w-[600px] p-0">
        <div class="flex flex-col h-[500px]">
          <div class="p-4 border-b grid grid-cols-2 gap-4">
            <Select v-if="models" v-model="selectedModel">
              <SelectTrigger>
                <SelectValue placeholder="Select a model" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectLabel>Models</SelectLabel>
                  <SelectItem
                    v-for="model in models"
                    :key="model.digest"
                    :value="model.name"
                    >{{ model.name }}</SelectItem
                  >
                </SelectGroup>
              </SelectContent>
            </Select>

            <Button
              variant="outline"
              v-if="listing && selectedModel"
              @click="askModelsOpinion"
            >
              <Sparkles />
              Ask about the
              <p class="truncate">
                {{ listing.target.marketplace_listing_title }}
              </p>
            </Button>
          </div>

          <ScrollArea class="relative flex-grow p-4">
            <div
              v-if="conversation.length === 0"
              class="absolute inset-0 flex justify-center items-center text-muted-foreground text-sm"
            >
              Select a model and start chatting!
            </div>
            <div v-else class="space-y-4">
              <div
                v-for="(msg, index) in conversation"
                :key="index"
                class="flex"
              >
                <div
                  :class="[
                    'max-w-[80%] rounded-md p-3 text-sm break-words',
                    msg.role === 'user' ? 'bg-accent ml-auto' : '',
                  ]"
                >
                  <div class="leading-relaxed">
                    <span
                      v-if="
                        isMessageLoading &&
                        index === conversation.length - 1 &&
                        msg.role === 'assistant'
                      "
                    >
                      <Skeleton class="h-9 w-[200px] rounded-md" />
                    </span>
                    <span v-else>{{ msg.content }}</span>
                  </div>
                </div>
              </div>
            </div>
          </ScrollArea>

          <div class="p-4 border-t">
            <form @submit.prevent="sendMessage" class="flex space-x-2">
              <Input
                v-model="message"
                placeholder="Type your message..."
                :disabled="!selectedModel"
                class="flex-grow"
              />
              <Button
                class="h-9 w-9"
                type="submit"
                :disabled="!selectedModel || !message.trim()"
              >
                <SendIcon class="h-4 w-4" />
              </Button>
            </form>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  </div>
</template>
