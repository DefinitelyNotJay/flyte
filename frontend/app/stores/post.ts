import { defineStore } from 'pinia';
import type { Post } from '~/types/post';

export const usePostStore = defineStore('post', () => {
  const posts = ref<Post[]>([]);
  const currentPost = ref<Post | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const api = useApi();

  const fetchPosts = async (page_id = 1, page_size = 10) => {
    loading.value = true;
    error.value = null;
    try {
      const data = await api<Post[]>('/posts', {
        query: { page_id, page_size }
      });
      posts.value = data;
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const fetchPost = async (id: number) => {
    loading.value = true;
    error.value = null;
    try {
      const data = await api<Post>(`/posts/${id}`);
      currentPost.value = data;
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const createPost = async (content: string) => {
    loading.value = true;
    error.value = null;
    try {
      const data = await api<Post>('/posts', {
        method: 'POST',
        body: { content }
      });
      posts.value.unshift(data);
      return data;
    } catch (e: any) {
      error.value = e.message;
      return null;
    } finally {
      loading.value = false;
    }
  };

  const updatePost = async (id: number, content: string) => {
    loading.value = true;
    error.value = null;
    try {
      const data = await api<Post>(`/posts`, {
        method: 'PATCH',
        body: { id, content: { String: content, Valid: true } }
      });
      const index = posts.value.findIndex(p => p.id === id);
      if (index !== -1) posts.value[index] = data;
      return data;
    } catch (e: any) {
      error.value = e.message;
      return null;
    } finally {
      loading.value = false;
    }
  };

  const deletePost = async (id: number) => {
    loading.value = true;
    error.value = null;
    try {
      await api(`/posts/${id}`, { method: 'DELETE' });
      posts.value = posts.value.filter(p => p.id !== id);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  return {
    posts,
    currentPost,
    loading,
    error,
    fetchPosts,
    fetchPost,
    createPost,
    updatePost,
    deletePost
  };
});
