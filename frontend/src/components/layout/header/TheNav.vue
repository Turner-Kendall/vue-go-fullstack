<template>
  <nav class="menu">
    <Slide>
      <a id="home" href="/"> <span class="bm-link">The Home Page</span> </a>

      <a id="about" href="/page/about">
        <span class="bm-link">About The Weblog.Network</span>
      </a>

      <a id="contact" href="/page/contact">
        <span class="bm-link">Contact The Weblog.Network</span>
      </a>

      <a id="contact" href="/page/contact">
        <span class="bm-link">Contact The Weblog.Network</span>
      </a>

    </Slide>
  </nav>
</template>

<script>
import { Slide } from 'vue3-burger-menu';

export default {
  components: {
    Slide
  }
}
</script>

<style>
.bm-menu,
nav {
  display: block;
}

.bm-link {
  line-height: 1;
}

html[data-theme="light"] .bm-menu,
html[data-theme="light"] nav {
  background-color: #b2b2b3;
}

html[data-theme="light"] .bm-item-list a:hover {
  background-color: #9a9a9b !important;
  color: #fff !important;
}

html[data-theme="dark"] .bm-menu,
html[data-theme="dark"] nav {
  background-color: #0d1016;
}

html[data-theme="dark"] .bm-item-list a:hover {
  background-color: #000 !important;
  color: #d24e69 !important;
}

</style>
