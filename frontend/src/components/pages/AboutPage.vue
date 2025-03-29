<template>
  <div class="about">
    <h2>About The Vue-Go-Full Stack Front End</h2>
    <p>
      😼 Hey Now!
    </p>
    <p>
      This is what I use as a starting point for new full stack projects.
      It provides a working frontend that is instended to be used with the
      backend
      <a href="https://github.com/Turner-Kendall/vue-go-fullstack/blob/main/README.md">in the same repository</a>
      - with user auth, (postgres) database connection, etc.
    </p>

    <h3>Vue3 and Pico.css</h3>
    <p>
      The front end is built with Vue3 with the composition API and Pico.css.
      The goal is to provide a simple, fast, and responsive user interface that is easy to use
      and navigate. It's built with vite and simple to get up and running:
    </p>


    <ul>
      <li>
        <pre><code>npm install</code></pre>
      </li>
      <li>
        <pre><code>npm run dev</code></pre>
      </li>
      <li>
        <pre><code>Open <a href=" http://localhost:5173/"> http://localhost:5173/</a> in your browser</code></pre>
      </li>
      <li>
        <pre><code><a href="Open http://localhost:5173/__devtools__/<">Open http://localhost:5173/__devtools__/</a></code></pre>
      </li>
    </ul>

    <h3>Pico overides</h3>
    <p>
      Like most frameworks Pico uses generic class names. To override them, I have attempted increase specificity:
    </p>
    <pre>
      <code>
  /* Framework has: */
  .btn { margin:0; }

  /* You can override with: */
  .custom-btn { margin:0 1rem; }  /* Preferred */
  button.btn { margin:0 1rem; }   /* More specific but not foolproof */
      </code>
    </pre>

    <p>
      If necessary, <strong>use element + class</strong> combos or nested selectors.</p>
    <p>I have used !important sparingly, in the notes to try and remove alltogether. The most aggresive approach ( which
      I have removed now ) is
    </p>

    <pre>
      <code>
  .btn {
  all: unset;
  display: inline-block;
  padding: 10px 20px;
  }
      </code>
    </pre>

  <h3>
    Todos
  </h3>

  <ul>
    <li>make catch all / 404 page</li>
    <li>Remove all !important in the CSS file where possible</li>
    <li>Remember Me Implimentation</li>
    </ul>
  </div>
</template>

<script setup>
import { useHead } from '@unhead/vue';

useHead({
  title: 'About the Weblog Network',
  meta: [
    { name: 'description', content: 'The Weblog Network is a collection of blogs and websites that are dedicated to providing an alternative to the mainstream media' },
    { property: 'og:title', content: 'About The Weblog.Network' },
    { property: 'og:description', content: 'The Weblog Network is a collection of blogs and websites that are dedicated to providing an alternative to the mainstream media' },
    { property: 'og:image', content: 'https://example.com/image.jpg' },
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:title', content: 'About The Weblog Network' },
    { name: 'twitter:description', content: 'The Weblog Network is a collection of blogs and websites that are dedicated to providing an alternative to the mainstream media' },
    { name: 'twitter:image', content: 'https://example.com/image.jpg' },
  ],
});
</script>
