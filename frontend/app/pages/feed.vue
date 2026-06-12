<script setup lang="ts">
import { storeToRefs } from 'pinia';
import { useAuthStore } from '~/stores/auth';
import type { Post } from '~/types/post';

const { posts, loading, error } = usePost();
const { fetchPosts, createPost, updatePost, deletePost } = usePost();
const { user } = storeToRefs(useAuthStore());

definePageMeta({
  middleware: 'auth'
});

// Compose logic
const newPostContent = ref('');
const isSubmitting = ref(false);
const MAX_CHARS = 280;

// Per-post states
const editingId = ref<number | null>(null);
const editContent = ref('');
const deleteConfirmId = ref<number | null>(null);
const processingIds = ref(new Set<number>());

onMounted(() => {
  fetchPosts();
});

const handleCreate = async () => {
  if (!newPostContent.value || newPostContent.value.length > MAX_CHARS) return;
  isSubmitting.value = true;
  const post = await createPost(newPostContent.value);
  if (post) newPostContent.value = '';
  isSubmitting.value = false;
};

const startEdit = (post: Post) => {
  editingId.value = post.id;
  editContent.value = post.content.String;
  deleteConfirmId.value = null;
};

const cancelEdit = () => {
  editingId.value = null;
  editContent.value = '';
};

const handleUpdate = async (id: number) => {
  processingIds.value.add(id);
  const updated = await updatePost(id, editContent.value);
  if (updated) cancelEdit();
  processingIds.value.delete(id);
};

const handleDelete = async (id: number) => {
  processingIds.value.add(id);
  await deletePost(id);
  processingIds.value.delete(id);
  deleteConfirmId.value = null;
};

const icons = {
  reply: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 20.25c4.97 0 9-3.694 9-8.25s-4.03-8.25-9-8.25S3 7.444 3 12c0 2.104.859 4.023 2.273 5.48.432.447.74 1.04.586 1.641a4.483 4.483 0 0 1-.923 1.785 0.5 0 0 0 .445.79c1.023-.033 1.99-.368 2.805-.937a0.5 0 0 1 .45-.03c1.17.47 2.443.731 3.774.731Z" /></svg>`,
  repost: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 12c0-1.232-.046-2.453-.138-3.662a4.006 4.006 0 0 0-3.7-3.7 48.678 48.678 0 0 0-7.324 0 4.006 4.006 0 0 0-3.7 3.7c-.017.22-.032.441-.046.662M19.5 12l3-3m-3 3-3-3m-12 3c0 1.232.046 2.453.138 3.662a4.006 4.006 0 0 0 3.7 3.7 48.656 48.656 0 0 0 7.324 0 4.006 4.006 0 0 0 3.7-3.7c.017-.22.032-.441.046-.662M4.5 12l3 3m-3-3-3 3" /></svg>`,
  like: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z" /></svg>`,
  share: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M7.217 10.907a2.25 2.25 0 1 0 0 2.186m0-2.186c.18.324.283.696.283 1.093s-.103.77-.283 1.093m0-2.186 9.566-5.314m-9.566 7.5 9.566 5.314m0 0a2.25 2.25 0 1 0 3.935 2.186 2.25 2.25 0 0 0-3.935-2.186Zm0-12.814a2.25 2.25 0 1 0 3.933-2.185 2.25 2.25 0 0 0-3.933 2.185Z" /></svg>`,
  dots: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M6.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM18.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" /></svg>`
};

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr);
  return new Intl.DateTimeFormat('en-US', { month: 'short', day: 'numeric' }).format(date);
};
</script>

<template>
  <div class="w-full max-w-[600px] flex flex-col mx-auto">
    
    <!-- Compose Area -->
    <div class="px-8 pb-12 border-b border-stone-200/40 dark:border-stone-800/40">
      <div class="flex gap-6">
        <div class="flex-shrink-0">
          <div class="w-9 h-9 rounded-full bg-muji-accent/10 animate-pulse" v-if="!user"></div>
          <img v-else :src="user.avatar_url.String || 'https://api.dicebear.com/7.x/initials/svg?seed=' + user.username" class="w-9 h-9 rounded-full grayscale opacity-80" />
        </div>
        <div class="flex-grow">
          <textarea 
            v-model="newPostContent"
            placeholder="What is on your mind?"
            class="w-full bg-transparent border-none focus:ring-0 text-[15px] leading-[1.7] resize-none min-h-[100px] placeholder:text-stone-400 dark:placeholder:text-stone-600"
            :disabled="isSubmitting"
          ></textarea>
          <div class="flex items-center justify-between mt-4">
            <span 
              class="text-[12px] font-mono tracking-widest transition-colors duration-300"
              :class="newPostContent.length > MAX_CHARS ? 'text-red-400' : 'text-stone-400/60'"
            >
              {{ newPostContent.length }} / {{ MAX_CHARS }}
            </span>
            <button 
              @click="handleCreate"
              :disabled="!newPostContent || newPostContent.length > MAX_CHARS || isSubmitting"
              class="text-[13px] font-medium text-muji-accent hover:opacity-100 disabled:opacity-20 transition-all duration-300"
            >
              {{ isSubmitting ? 'Posting...' : 'Post' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Timeline -->
    <div class="space-y-0 min-h-[400px]">
      <!-- Loading Skeleton -->
      <template v-if="loading && posts.length === 0">
        <div v-for="i in 3" :key="i" class="py-12 px-8 border-b border-stone-200/40 dark:border-stone-800/40 animate-pulse">
          <div class="flex gap-6">
            <div class="w-9 h-9 rounded-full bg-stone-200 dark:bg-stone-800/50"></div>
            <div class="flex-grow space-y-4">
              <div class="h-4 w-32 bg-stone-200 dark:bg-stone-800/50 rounded"></div>
              <div class="space-y-2">
                <div class="h-3 w-full bg-stone-200 dark:bg-stone-800/50 rounded"></div>
                <div class="h-3 w-4/5 bg-stone-200 dark:bg-stone-800/50 rounded"></div>
              </div>
            </div>
          </div>
        </div>
      </template>

      <!-- Post List -->
      <article 
        v-for="post in posts" 
        :key="post.id"
        class="group relative py-12 px-8 transition-all duration-300 border-b border-stone-200/40 dark:border-stone-800/40"
        :class="{ 'opacity-50 pointer-events-none': processingIds.has(post.id) }"
      >
        <!-- Bookmark Hover Accent -->
        <div class="absolute left-0 top-0 bottom-0 w-[2px] bg-muji-accent scale-y-0 group-hover:scale-y-100 transition-transform duration-300 origin-top"></div>

        <div class="flex gap-6">
          <!-- Avatar -->
          <div class="flex-shrink-0">
            <img 
              :src="'https://api.dicebear.com/7.x/avataaars/svg?seed=' + post.author_id" 
              class="w-9 h-9 rounded-full grayscale hover:grayscale-0 transition-all duration-500 opacity-80" 
            />
          </div>

          <!-- Content Area -->
          <div class="flex-grow">
            <!-- Meta & Actions Dropdown -->
            <div class="flex items-center gap-2 mb-4">
              <span class="font-medium text-[14px] text-stone-700 dark:text-stone-300">User #{{ post.author_id }}</span>
              <span class="text-[13px] text-stone-400/60 dark:text-stone-600/60">{{ formatDate(post.created_at) }}</span>
              
              <!-- Owner Actions -->
              <div v-if="user && post.author_id === user.id" class="ml-auto flex items-center gap-4">
                 <button 
                  v-if="editingId !== post.id && deleteConfirmId !== post.id"
                  @click="startEdit(post)"
                  class="text-[12px] text-stone-400 hover:text-muji-accent transition-colors opacity-0 group-hover:opacity-100"
                >edit</button>
                <button 
                  v-if="editingId !== post.id && deleteConfirmId !== post.id"
                  @click="deleteConfirmId = post.id"
                  class="text-[12px] text-stone-400 hover:text-red-400 transition-colors opacity-0 group-hover:opacity-100"
                >delete</button>
              </div>
            </div>

            <!-- Body / Edit Mode -->
            <div v-if="editingId === post.id" class="space-y-4 animate-in fade-in slide-in-from-top-1 duration-300">
              <textarea 
                v-model="editContent"
                class="w-full bg-transparent border-b border-muji-accent/20 focus:border-muji-accent focus:ring-0 text-[15px] leading-[1.7] resize-none min-h-[80px] p-0"
              ></textarea>
              <div class="flex justify-end gap-4">
                <button @click="cancelEdit" class="text-[13px] text-stone-400">Cancel</button>
                <button @click="handleUpdate(post.id)" class="text-[13px] text-muji-accent font-medium">Save</button>
              </div>
            </div>

            <!-- Delete Confirm Mode -->
            <div v-else-if="deleteConfirmId === post.id" class="py-4 bg-red-50/5 dark:bg-red-900/5 rounded px-4 -mx-4 flex items-center justify-between animate-in zoom-in-95 duration-200">
              <span class="text-[13px] text-red-400 font-serif italic">ลบโพสต์นี้?</span>
              <div class="flex gap-4">
                <button @click="deleteConfirmId = null" class="text-[13px] text-stone-400">Cancel</button>
                <button @click="handleDelete(post.id)" class="text-[13px] text-red-400 font-medium underline underline-offset-4">Confirm</button>
              </div>
            </div>

            <!-- Normal Content Display -->
            <p v-else class="text-[15px] leading-[1.7] text-stone-600 dark:text-stone-400 font-normal whitespace-pre-wrap">
              {{ post.content.String }}
            </p>

            <!-- Bottom Actions -->
            <div v-if="editingId !== post.id && deleteConfirmId !== post.id" class="flex items-center gap-8 mt-8 opacity-25 group-hover:opacity-80 transition-opacity duration-300">
              <button class="hover:text-muji-accent transition-colors flex items-center gap-2">
                <span v-html="icons.reply"></span>
                <span v-if="post.reply_count > 0" class="text-[11px] font-mono">{{ post.reply_count }}</span>
              </button>
              <button class="hover:text-muji-accent transition-colors flex items-center gap-2">
                <span v-html="icons.repost"></span>
                <span v-if="post.repost_count > 0" class="text-[11px] font-mono">{{ post.repost_count }}</span>
              </button>
              <button class="hover:text-muji-accent transition-colors flex items-center gap-2">
                <span v-html="icons.like"></span>
                <span v-if="post.like_count > 0" class="text-[11px] font-mono">{{ post.like_count }}</span>
              </button>
              <button class="hover:text-muji-accent transition-colors" v-html="icons.share"></button>
            </div>
          </div>
        </div>
      </article>

      <div v-if="!loading && posts.length === 0" class="py-24 text-center">
        <span class="font-serif italic text-stone-400 text-sm">Silence fills the air.</span>
      </div>
    </div>

    <!-- Footer -->
    <footer v-if="posts.length > 0" class="py-24 text-center opacity-20">
      <span class="font-serif italic text-sm">Quietly reaching the end.</span>
    </footer>
  </div>
</template>
