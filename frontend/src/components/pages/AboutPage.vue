<template>
  <div class="about">
    <h2>About the Weblog.Network</h2>
    <p>
      The Weblog.Network is a collection of blogs and websites that are dedicated to providing an alternative to the
      mainstream media. Our goal is to provide a platform for independent journalists, bloggers, and content creators to
      share their work with the world. This site is the flagship site of the Weblog.Network and is dedicated to
      providing news, analysis, and commentary on a wide range of topics.
    </p>
    <h3>Our Mission</h3>
    <p>
      Our mission is to provide a platform for independent journalists, bloggers, and content creators to share their
      work with the world. We believe that the mainstream media is failing to provide the public with the information
      they need to make informed decisions. We aim to fill that gap by providing a platform for independent voices to be
      heard.
    </p>
    <h3>Our Values</h3>
    <p>
      We believe in the power of independent journalism to hold the powerful accountable and to shine a light on
      important issues that are often ignored by the mainstream media. We believe in the importance of free speech and
      open debate, and we are committed to providing a platform for a wide range of voices and perspectives.
    </p>
    <h3>Join Us!</h3>
    <p>
      If you are an independent journalist, blogger, or content creator who is interested in joining the Weblog.Network,
      please contact. We are always looking for new voices to add to our network, and we would love to hear from you!
    </p>
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
