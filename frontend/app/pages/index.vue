<script setup lang="ts">
// Automatically redirect root path to the feed
definePageMeta({
  middleware: [
    function (to, from) {
      return navigateTo('/feed');
    },
  ],
});
</script>

<template>
  <div></div>
</template>
