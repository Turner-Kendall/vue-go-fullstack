<template>
  <article>
    <Suspense>
      <template #default>
        <div>

          <section class="grid-header">
            <h5>User Profile</h5>
            <h6><a href="#" @click.prevent="handleLogout">Logout</a></h6>
          </section>

          <ul>
            <li>ID: {{ userProfile.id }}</li>
            <li>Username: {{ userProfile.username }}</li>
            <li>Email: {{ userProfile.email }}</li>
          </ul>

        </div>
      </template>
      <template #fallback>
        <div>Loading...</div>
      </template>
    </Suspense>
  </article>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from "vue-router";
const userProfile = ref({});
const router = useRouter();

const handleLogout = () => {
  localStorage.removeItem("authToken");
  router.push("/");
};

const token = localStorage.getItem("authToken");

onMounted(async () => {
  if (token) {
    try {
      const response = await fetch('http://localhost:8080/user/profile', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();
      userProfile.value = data.user;

    } catch (error) {
      console.error("Error fetching profile:", error);
    }
  }
});
</script>
