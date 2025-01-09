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
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form/index.js";

import { Send } from "lucide-vue-next";

import { GetMarketplaceListing } from "../../wailsjs/go/main/App.js";

const res = ref("");

const formSchema = toTypedSchema(
  z.object({
    listingId: z
      .string()
      .length(15, { message: "Listing ID must be exactly 15 characters long" })
      .regex(/^[0-9]{15}$/, {
        message: "Listing ID must only contain 15 digits",
      }),
  }),
);

const form = useForm({
  validationSchema: formSchema,
});

const onSubmit = form.handleSubmit(async (values) => {
  const { listingId } = values;

  GetMarketplaceListing(listingId).then((resp) => {
    res.value = JSON.stringify(resp, null, 2);
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
            Enter the listing ID to retrieve detailed information and insights
            about the product.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <FormField v-slot="{ componentField }" name="listingId">
            <FormItem v-auto-animate>
              <FormLabel>Listing ID</FormLabel>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Enter 15-digit Listing ID"
                  v-bind="componentField"
                />
              </FormControl>
              <FormDescription>
                The unique 15-digit ID of the listing.
              </FormDescription>
              <FormMessage />
            </FormItem>
          </FormField>
        </CardContent>
        <CardFooter class="flex justify-end">
          <Button type="submit">
            <Send />
            Submit
          </Button>
        </CardFooter>
      </Card>
    </form>
  </div>
</template>
