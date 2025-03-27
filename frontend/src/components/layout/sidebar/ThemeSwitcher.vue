<template>
  <button @click="toggleTheme">
    <Iconbright />
    {{ theme === 'dark' ? 'Light Mode' : 'Dark Mode' }}
  </button>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import Iconbright from '../../icons/Iconbright.vue';
const theme = ref(localStorage.getItem('theme') || 'dark');

const toggleTheme = () => {
  theme.value = theme.value === 'dark' ? 'light' : 'dark';
  document.documentElement.setAttribute('data-theme', theme.value);
  localStorage.setItem('theme', theme.value);
};

onMounted(() => {
  document.documentElement.setAttribute('data-theme', theme.value);
});
</script>

<style scoped>
button {
  padding: 0 .5rem;
  margin: 0 0 2rem 1rem;
  display: block;
  border: none;
  border-radius: 0.25rem;
  cursor: pointer;
  height: 2rem;
  width: 90%;
  font-size: 1rem;
}

@media (max-width: 600px) {
  button {
    margin: 0 0 2rem 0;
    width: 100%;
  }
}
</style>
