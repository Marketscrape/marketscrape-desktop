<script setup lang="ts">
import { ref } from "vue";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { vAutoAnimate } from "@formkit/auto-animate/vue";

import { Input } from "./ui/input/index.js";
import { Button } from "./ui/button/index.js";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card/index.js";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form/index.js";

import { Send, Loader2 } from "lucide-vue-next";

import { GetMarketplaceListing } from "../../wailsjs/go/main/App.js";

const extractListingIdFromUrl = (url: string) => {
  const regex = /\/item\/([0-9]{15,16})/;
  const match = url.match(regex);
  return match ? match[1] : null;
};

const isLoading = ref(false);
const emit = defineEmits(["listing"]);
const formSchema = toTypedSchema(
  z.object({
    listingId: z.string().refine(
      (value) => {
        const isUrl = value.startsWith("http");
        const isId = /^[0-9]{15,16}$/.test(value);
        return isUrl || isId;
      },
      {
        message: "Listing must be a valid URL or 15-16 digit ID",
      },
    ),
  }),
);

const form = useForm({
  validationSchema: formSchema,
});

const onSubmit = form.handleSubmit(async (values) => {
  isLoading.value = true;
  const { listingId } = values;

  let idToUse = listingId;
  if (listingId.startsWith("http")) {
    // Extract the listing ID if URL is provided
    idToUse = extractListingIdFromUrl(listingId) || listingId;
  }

  GetMarketplaceListing(idToUse).then((resp) => {
    emit("listing", JSON.stringify(resp, null, 2));
  });
});
</script>

<template>
  <div class="flex h-[calc(100vh-4rem)] items-center">
    <form class="mx-auto" @submit="onSubmit">
      <Card>
        <CardHeader>
          <CardTitle>Analyze a Marketplace Listing</CardTitle>
          <CardDescription>
            Enter the listing URL or ID to retrieve detailed information and
            insights about the product.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <FormField v-slot="{ componentField }" name="listingId">
            <FormItem v-auto-animate>
              <FormLabel>Listing ID or URL</FormLabel>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Enter Marketplace Listing URL or ID"
                  v-bind="componentField"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </CardContent>
        <CardFooter class="flex justify-end">
          <Button type="submit" :disabled="isLoading">
            <template v-if="!isLoading">
              <Send />
              Submit
            </template>
            <template v-else>
              <Loader2 class="animate-spin" />
              Loading...
            </template>
          </Button>
        </CardFooter>
      </Card>
    </form>
  </div>
</template>
