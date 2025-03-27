<template>

  <article>
    <a href="#top" class="top-link">
      <IconUp />
      Back to the Top
    </a>
    <footer id="footer">
      <section class="footer-links">
        <div v-for="section in footerSections" :key="section.title" class="footer-column">
          <h5>{{ section.title }}</h5>
          <ul class="list-unstyled">
            <li v-for="link in section.links" :key="link.id">
              <a :href="link.link">{{ link.title }}</a>
            </li>
          </ul>
        </div>
      </section>

      <section class="credits">
        <p>&copy; Some Rights Reserved
          <a href="https://turnerkendall.com">Turner Kendall</a> 2025
        </p>
      </section>

    </footer>
  </article>
</template>

<script setup>
import { footerSections } from './footerData'
import IconUp from '../../icons/IconUp.vue';
</script>

<style scoped>
#footer {
  padding: 25px 0;
}

h5 {
  margin: 0 0 5px 0 !important;
  padding: 0;
  font-size: .9rem !important;
  font-weight: 600;
}

.credits {
  padding: 20px 0 0 0;
  text-align: center;
}

.footer-column {
  text-align: left;
  padding: 10px;
}

.footer-column ul {
  list-style-type: none;
  margin: 0;
  padding: 0 0 0 5px;
}

.footer-column ul li {
  list-style-type: none;
  margin: 0;
  padding: 0;
  font-size: .8rem;
}

.footer-links {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.footer-links div:last-child {
  margin-right: 2rem;
}

.footer-links div:first-child {
  margin-left: 1rem;
}

@media (max-width: 640px) {
  .footer-links div:last-child {
    margin-right: 0;
    margin-top: 0;
  }

  .footer-links div:first-child {
    margin-left: 0;
  }

  .footer-column {
    text-align: left;
    padding-left: 0;
  }
}

@media (max-width: 468px) {
  .footer-links {
    display: grid;
    grid-template-columns: 1fr;
    text-align: center;
  }

}
</style>
