<template>
  <article>
    <form @submit="handleSubmit" class="login-form">

      <section class="pairs">
        <input type="text" name="username" placeholder="Username" id="username" aria-label="Username"
          autocomplete="username" required />

        <input type="password" name="password" placeholder="Password" id="password" aria-label="Password" required />
      </section>

      <button type="submit">User Login</button>

      <section class="pairs">
        <fieldset>
          <label for="terms">
            <input type="checkbox" role="switch" id="terms" name="terms" />
            Remember Me
          </label>
        </fieldset>
        <div>
          <p class="register-link">
            Don't have an account? <a href="/register">Register</a>
          </p>
        </div>
      </section>

    </form>
  </article>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from "vue-router";
let message = ref('')
const router = useRouter();

const handleSubmit = async (e) => {
  e.preventDefault();
  const form = e.target;
  const formData = new FormData(form);
  const data = {
    username: formData.get('username'),
    password: formData.get('password')
  };

  console.log('Request Payload:', data); // Log the request payload

  try {
    const response = await fetch('http://localhost:8080/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    console.log('Response Status:', response.status); // Log the response status

    if (!response.ok) {
      const errorText = await response.text(); // Get the error response text
      console.error('Error Response:', errorText); // Log the error response
      throw new Error('Network response was not ok');
    }

    const result = await response.json();
    const token = result.token; // Extract token
    localStorage.setItem("authToken", token); // Store token for future requests

    console.log('Token:', result.token);
    message.value = 'Success';
    router.push("/profile"); // Redirect after login
  } catch (error) {
    console.error('Error:', error);
    message.value = 'Error';
  }
};

</script>

<style scoped>
/* Keep this is the component */
article {
  min-height: 200px;
}
</style>
