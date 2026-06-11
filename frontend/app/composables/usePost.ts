import { storeToRefs } from 'pinia';
import { usePostStore } from '~/stores/post';

export const usePost = () => {
  const store = usePostStore();
  const { posts, currentPost, loading, error } = storeToRefs(store);

  const { fetchPosts, fetchPost, createPost, updatePost, deletePost } = store;

  return {
    // Reactive State
    posts,
    currentPost,
    loading,
    error,

    // Actions
    fetchPosts,
    fetchPost,
    createPost,
    updatePost,
    deletePost,

    /**
     * usePostForm provides logic for post creation/editing forms
     */
    usePostForm(initialContent = '') {
      const content = ref(initialContent);
      const MAX_CHARS = 280;

      const charCount = computed(() => content.value.length);
      const isOverLimit = computed(() => charCount.value > MAX_CHARS);
      const isValid = computed(() => charCount.value > 0 && !isOverLimit.value);

      const submit = async () => {
        if (!isValid.value) return null;
        return await createPost(content.value);
      };

      return {
        content,
        charCount,
        isOverLimit,
        isValid,
        submit
      };
    }
  };
};
