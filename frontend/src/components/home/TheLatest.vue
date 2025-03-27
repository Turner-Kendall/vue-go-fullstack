<template>
  <article>
    <h3><IconDoc /> The Full Stack Starter Kit:</h3>
    <div class="container">
      <section class="column">
        <LeftColumn />
      </section>
      <section class="column">
        <RightColumn />
      </section>
    </div>
  </article>
</template>

<script setup>
import LeftColumn from './LeftColumn.vue'
import RightColumn from './RightColumn.vue'
import IconDoc from '../icons/IconHeart.vue';
</script>

