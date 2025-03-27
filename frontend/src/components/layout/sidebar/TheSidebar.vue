<template>
  <aside id="sidebar">
    <ThemeSwitcher />
    <a href="https://go.nordvpn.net/aff_c?offer_id=15&aff_id=93756&url_id=902 "><img src="/nord.webp"
        alt="Nord VPN is our sponsor" /></a>
    <h5>Network Sites & Sponsors:</h5>
    <ul>
      <li v-for="item in asideData" :key="item.id" class="aside-item">
        <a :href="item.link">{{ item.title }}</a>
        <small>{{ item.description }}</small><br>
      </li>
    </ul>
  </aside>
</template>

<script setup>
import { asideData } from './asideData'
import ThemeSwitcher from './ThemeSwitcher.vue'
</script>

<style scoped>
#sidebar {
  text-align: right;
}

@media (max-width: 740px) {
  #sidebar {
    text-align: center;
  }
}

h5 {
  margin-top: 2rem !important;
}

.aside-item {
  margin-bottom: .5rem;
}

#sidebar ul {
  padding: 0;
  margin-top: 0;
}
</style>
