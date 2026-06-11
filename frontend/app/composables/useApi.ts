export const useApi = () => {
  const token = useCookie('auth_token');
  const config = useRuntimeConfig();
  
  // Base URL should ideally be in runtimeConfig, but I'll use your backend default for now
  const baseURL = 'http://localhost:8000';

  return $fetch.create({
    baseURL,
    onRequest({ options }) {
      if (token.value) {
        options.headers = {
          ...options.headers,
          Authorization: `Bearer ${token.value}`,
        };
      }
    },
    onResponseError({ response }) {
      // Standardize error extraction from your backend's {"error": "..."} format
      const errorMsg = response._data?.error || 'An unexpected error occurred';
      throw new Error(errorMsg);
    }
  });
};
