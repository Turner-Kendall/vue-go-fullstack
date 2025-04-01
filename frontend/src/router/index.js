import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import PageView from "@/views/PageView.vue";
import ProfileView from "@/views/ProfileView.vue";
import LoginView from "@/views/LoginView.vue";
import AdminView from "@/views/AdminView.vue";
import RegisterView from "@/views/RegisterView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: "/", name: "home", component: HomeView },
    { path: "/page/:name", name: "PageView", component: PageView },
    { path: "/login", name: "login", component: LoginView },
    { path: "/register", name: "register", component: RegisterView },
    {
      path: "/profile",
      name: "profile",
      component: ProfileView,
      meta: { requiresAuth: true }
    },
    {
      path: "/admin",
      name: "admin",
      component: AdminView,
      meta: { requiresAuth: true }
    }

  ],
  scrollBehavior(to, from, savedPosition) {
    return { top: 0 };
  },
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("authToken");

  // console.log("Navigation Triggered:", to.path);
  // console.log("Token Exists?", !!token);
  // console.log("Route Requires Auth?", to.meta.requiresAuth); // Debug meta field

  if (to.meta.requiresAuth && !token) {
    console.log("Redirecting to /");
    next("/"); // Redirect to home if not logged in
  } else {
    console.log("Access granted");
    next();
  }
});

export default router;
