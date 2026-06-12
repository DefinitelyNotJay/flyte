<script setup lang="ts">
import { useAuthStore } from '~/stores/auth';

const authStore = useAuthStore();
const isLoginMode = ref(true);
const errorMsg = ref<string | null>(null);

// Form fields
const form = reactive({
  identifier: '', // used as username or email in login
  username: '',   // used in register
  email: '',      // used in register
  password: ''
});

const toggleMode = () => {
  isLoginMode.value = !isLoginMode.value;
  errorMsg.value = null;
  // Clear fields on toggle if desired, or keep password
};

const handleSubmit = async () => {
  errorMsg.value = null;
  try {
    if (isLoginMode.value) {
      await authStore.login(form.identifier, form.password);
    } else {
      const success = await authStore.register(form.username, form.email, form.password);
      if (success) {
        // Automatically login after register
        await authStore.login(form.username, form.password);
      }
    }
  } catch (e: any) {
    errorMsg.value = e.message;
  }
};
</script>

<template>
  <div class="flex flex-col items-center justify-center min-h-[60vh] px-8">
    <div class="w-full max-w-[360px] space-y-12">
      <!-- Title -->
      <header class="text-center space-y-4">
        <h1 class="font-serif text-3xl font-light tracking-tight text-stone-800 dark:text-stone-200">
          {{ isLoginMode ? 'Welcome back' : 'Create Account' }}
        </h1>
        <div class="h-[1px] w-8 bg-muji-accent/40 mx-auto"></div>
      </header>

      <!-- Error Message -->
      <div v-if="errorMsg" class="animate-in fade-in slide-in-from-top-1 duration-300">
        <p class="text-[13px] text-muji-accent font-serif italic text-center">
          {{ errorMsg }}
        </p>
      </div>

      <!-- Form -->
      <form @submit.prevent="handleSubmit" class="space-y-8" :class="{ 'opacity-40 pointer-events-none': authStore.loading }">
        <!-- Login Fields -->
        <template v-if="isLoginMode">
          <div class="space-y-1">
            <label class="text-[11px] uppercase tracking-[0.2em] text-stone-400 font-medium">Identifier</label>
            <input 
              v-model="form.identifier"
              type="text" 
              placeholder="Username or Email"
              class="w-full bg-transparent border-b border-stone-200 dark:border-stone-800 focus:border-muji-accent focus:ring-0 p-0 py-2 text-[15px] transition-colors"
              required
            />
          </div>
        </template>

        <!-- Register Fields -->
        <template v-else>
          <div class="space-y-1">
            <label class="text-[11px] uppercase tracking-[0.2em] text-stone-400 font-medium">Username</label>
            <input 
              v-model="form.username"
              type="text" 
              placeholder="Unique username"
              class="w-full bg-transparent border-b border-stone-200 dark:border-stone-800 focus:border-muji-accent focus:ring-0 p-0 py-2 text-[15px] transition-colors"
              required
            />
          </div>
          <div class="space-y-1">
            <label class="text-[11px] uppercase tracking-[0.2em] text-stone-400 font-medium">Email</label>
            <input 
              v-model="form.email"
              type="email" 
              placeholder="your@email.com"
              class="w-full bg-transparent border-b border-stone-200 dark:border-stone-800 focus:border-muji-accent focus:ring-0 p-0 py-2 text-[15px] transition-colors"
              required
            />
          </div>
        </template>

        <div class="space-y-1">
          <label class="text-[11px] uppercase tracking-[0.2em] text-stone-400 font-medium">Password</label>
          <input 
            v-model="form.password"
            type="password" 
            placeholder="Min. 6 characters"
            class="w-full bg-transparent border-b border-stone-200 dark:border-stone-800 focus:border-muji-accent focus:ring-0 p-0 py-2 text-[15px] transition-colors"
            required
            minlength="6"
          />
        </div>

        <!-- Actions -->
        <div class="pt-4 flex flex-col items-center space-y-6">
          <button 
            type="submit"
            class="text-[14px] font-medium text-muji-accent hover:opacity-80 transition-opacity underline underline-offset-8"
          >
            {{ isLoginMode ? 'Sign In' : 'Register' }}
          </button>
          
          <button 
            type="button"
            @click="toggleMode"
            class="text-[12px] text-stone-400 hover:text-stone-600 dark:hover:text-stone-200 transition-colors"
          >
            {{ isLoginMode ? "Don't have an account? Join us." : "Already a member? Sign in." }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
