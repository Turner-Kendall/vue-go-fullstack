<template>
  <article>
    <h3>User Registration</h3>
    <p>Please fill out all of the fields below to sign up.</p>
    <form @submit="handleSubmit" class="registration-form">

      <div class="pairs">
        <input type="text" name="firstname" placeholder="First Name" id="firstname" aria-label="First Name"
          autocomplete="firstname" required />
        <input type="text" name="lastname" placeholder="Last Name" id="lastname" aria-label="Last Name"
          autocomplete="lastname" required />
      </div>

      <div class="pairs">
        <input type="text" name="username" placeholder="Username" id="username" aria-label="Username"
          autocomplete="username" required />
        <input type="email" name="email" placeholder="Email Address" id="email" aria-label="Email Address"
          autocomplete="email" required />
      </div>

      <div class="pairs">
        <input type="password" name="password" placeholder="Password" id="password" aria-label="Password" required />
        <input type="password" name="confirmPassword" placeholder="Confirm Password" id="confirmPassword"
          aria-label="Confirm Password" required />
      </div>

      <div class="pairs">
        <input type="phone" name="phone" placeholder="Phone" id="phone" aria-label="Phone Number" required />
        <input type="text" name="dob" placeholder="Date of Birth" id="dob" aria-label="Date of Birth" required />
      </div>

      <button type="submit">Register</button>

      <section class="pairs">
        <fieldset>

        </fieldset>
        <div>
          <p class="register-link">
            Have an account? <a href="/login">Login</a>
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
    const response = await fetch('http://localhost:8080/auth/signup', {
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
