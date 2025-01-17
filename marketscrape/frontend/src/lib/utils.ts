import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { computed } from "vue";
import { formatDistanceToNow } from "date-fns";

import { main } from "./../../wailsjs/go/models";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function toTitleCase(str: string): string {
  return str.replace(
    /\w\S*/g,
    (text: string) =>
      text.charAt(0).toUpperCase() + text.substring(1).toLowerCase(),
  );
}

export function capitalizeFirstLetter(val: string) {
  return String(val).charAt(0).toUpperCase() + String(val).slice(1);
}

export function useListingUtils(props: { listing: main.Root }) {
  const creationTime = computed(() =>
    formatDistanceToNow(new Date(props.listing.target.creation_time * 1000), {
      addSuffix: true,
    }),
  );

  const filteredCategories = computed(() => {
    const taxonomyPath =
      props.listing?.marketplace_listing_renderable_target?.seo_virtual_category
        ?.taxonomy_path;

    return Array.isArray(taxonomyPath)
      ? taxonomyPath.filter(
          (category: main.TaxonomyPathItem) =>
            category?.seo_info?.seo_url?.trim() !== "",
        )
      : [];
  });

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

  return {
    creationTime,
    filteredCategories,
    formatPrice,
  };
}

export function textToNumber(text: string): number | null {
  const numberMap: { [key: string]: number } = {
    ONE: 1,
    TWO: 2,
    THREE: 3,
    FOUR: 4,
    FIVE: 5,
    SIX: 6,
    SEVEN: 7,
    EIGHT: 8,
    NINE: 9,
    TEN: 10,
  };

  const upperCaseText = text.toUpperCase();
  return numberMap[upperCaseText] || null;
}
