import { defineStore } from 'pinia';
import type { User } from '~/types/user';

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null);
  const loading = ref(false);
  const token = useCookie('auth_token');
  const api = useApi();

  const fetchUser = async () => {
    if (!token.value) return;
    loading.value = true;
    try {
      // In a real app, you'd have a /me or /users/profile endpoint
      // For now, I'll assume we fetch the user by username if it was stored in the token payload
      // or simply rely on the fact that we have a token.
      // Assuming a generic user for the sake of the UI demo if no profile endpoint exists yet.
      const data = await api<User>('/users/1'); // Fetching a profile
      user.value = data;
    } catch (e) {
      user.value = null;
    } finally {
      loading.value = false;
    }
  };

  const logout = () => {
    token.value = null;
    user.value = null;
    navigateTo('/login');
  };

  const login = async (identifier: string, password: string) => {
    loading.value = true;
    try {
      const data = await api<{ access_token: string, user: User }>('/users/login', {
        method: 'POST',
        body: { username: identifier, password }
      });
      token.value = data.access_token;
      user.value = data.user;
      navigateTo('/feed');
      return true;
    } catch (e: any) {
      throw e;
    } finally {
      loading.value = false;
    }
  };

  const register = async (username: string, email: string, password: string) => {
    loading.value = true;
    try {
      await api('/users', {
        method: 'POST',
        body: { username, email, password }
      });
      // After registration, usually we login or tell them to login.
      // Let's just login them automatically if you want, but for now let's return true.
      return true;
    } catch (e: any) {
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return {
    user,
    loading,
    fetchUser,
    login,
    register,
    logout
  };
});
