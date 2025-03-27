<template>
  <article class="page-wrapper">
    <div v-if="pageRoute == 'about'">
      <About />
    </div>
    <div v-if="pageRoute == 'contact'">
      <Contact />
    </div>
  </article>
</template>

<script setup>
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';
import About from '@/components/pages/AboutPage.vue'
import Contact from '@/components/pages/ContactPage.vue'
const route = useRoute();
const pageRoute = ref(route.params.name);

watch(route, (newRoute) => {
  pageRoute.value = newRoute.params.name;
});
</script>
