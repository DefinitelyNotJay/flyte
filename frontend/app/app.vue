<script setup lang="ts">
/**
 * Flyte - Minimal Social Feed
 * Aesthetic: Japanese/Muji Minimal
 */

interface Post {
  id: string;
  username: string;
  handle: string;
  avatar: string;
  content: string;
  timestamp: string;
}

const posts = ref<Post[]>([
  {
    id: '1',
    username: 'Haruki Murakami',
    handle: '@murakami_h',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Haruki',
    content: 'The world is a metaphor. We are all searching for a doorway to another reality, one that is quieter, more deliberate. Today, I watched a single leaf fall in the garden. It took exactly six seconds.',
    timestamp: '2h ago'
  },
  {
    id: '2',
    username: 'Issey Miyake',
    handle: '@issey_official',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Issey',
    content: 'Design is not for philosophy—it\'s for life. The fold is as important as the fabric itself. Minimalist does not mean empty; it means intentional.',
    timestamp: '5h ago'
  },
  {
    id: '3',
    username: 'Ryuichi Sakamoto',
    handle: '@skmt_r',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Ryu',
    content: 'Listen to the silence between the notes. That is where the music truly lives. Today\'s recording was just the sound of rain hitting a zinc roof.',
    timestamp: '12h ago'
  },
  {
    id: '4',
    username: 'Muji Global',
    handle: '@muji_official',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Muji',
    content: 'Quality of life is found in the objects we use every day. Simple, functional, and without ego. Our new porcelain collection is inspired by the fog over the Seto Inland Sea.',
    timestamp: '1d ago'
  }
]);

// Stroke icons as components/functions to keep the template clean
const icons = {
  reply: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 20.25c4.97 0 9-3.694 9-8.25s-4.03-8.25-9-8.25S3 7.444 3 12c0 2.104.859 4.023 2.273 5.48.432.447.74 1.04.586 1.641a4.483 4.483 0 0 1-.923 1.785 0.5 0 0 0 .445.79c1.023-.033 1.99-.368 2.805-.937a0.5 0 0 1 .45-.03c1.17.47 2.443.731 3.774.731Z" /></svg>`,
  repost: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 12c0-1.232-.046-2.453-.138-3.662a4.006 4.006 0 0 0-3.7-3.7 48.678 48.678 0 0 0-7.324 0 4.006 4.006 0 0 0-3.7 3.7c-.017.22-.032.441-.046.662M19.5 12l3-3m-3 3-3-3m-12 3c0 1.232.046 2.453.138 3.662a4.006 4.006 0 0 0 3.7 3.7 48.656 48.656 0 0 0 7.324 0 4.006 4.006 0 0 0 3.7-3.7c.017-.22.032-.441.046-.662M4.5 12l3 3m-3-3-3 3" /></svg>`,
  like: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z" /></svg>`,
  share: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="w-5 h-5"><path stroke-linecap="round" stroke-linejoin="round" d="M7.217 10.907a2.25 2.25 0 1 0 0 2.186m0-2.186c.18.324.283.696.283 1.093s-.103.77-.283 1.093m0-2.186 9.566-5.314m-9.566 7.5 9.566 5.314m0 0a2.25 2.25 0 1 0 3.935 2.186 2.25 2.25 0 0 0-3.935-2.186Zm0-12.814a2.25 2.25 0 1 0 3.933-2.185 2.25 2.25 0 0 0-3.933 2.185Z" /></svg>`
};
</script>

<template>
  <div class="muji-root min-h-screen selection:bg-stone-200 dark:selection:bg-stone-800">
    
    <!-- Sidebar Navigation -->
    <nav class="fixed left-0 top-0 h-full w-16 md:w-24 flex flex-col items-center py-8 border-r border-stone-200/40 dark:border-stone-800/40 z-50">
      <div class="mb-12 text-muji-accent">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 21a9.004 9.004 0 0 0 8.716-6.747M12 21a9.004 9.004 0 0 1-8.716-6.747M12 21c2.485 0 4.5-4.03 4.5-9S14.485 3 12 3m0 18c-2.485 0-4.5-4.03-4.5-9s2.015-9 4.5-9m0 0a9.015 9.015 0 0 1 0 18" />
        </svg>
      </div>
      
      <div class="flex flex-col space-y-8 opacity-40">
        <button class="hover:opacity-100 transition-opacity p-2">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="m2.25 12 8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" /></svg>
        </button>
        <button class="hover:opacity-100 transition-opacity p-2">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" /></svg>
        </button>
        <button class="hover:opacity-100 transition-opacity p-2">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0" /></svg>
        </button>
        <button class="hover:opacity-100 transition-opacity p-2">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M17.982 18.725A7.488 7.488 0 0 0 12 15.75a7.488 7.488 0 0 0-5.982 2.975m11.963 0a9 9 0 1 0-11.963 0m11.963 0A8.966 8.966 0 0 1 12 21a8.966 8.966 0 0 1-5.982-2.275M15 9.75a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" /></svg>
        </button>
      </div>
    </nav>

    <!-- Main Timeline -->
    <main class="ml-16 md:ml-24 flex justify-center py-12 px-6">
      <div class="w-full max-w-[600px] flex flex-col">
        
        <!-- Header -->
        <header class="mb-16">
          <h1 class="font-serif text-3xl md:text-4xl font-light tracking-tight text-stone-800 dark:text-stone-200">Timeline</h1>
          <div class="h-[1px] w-12 bg-muji-accent/40 mt-6"></div>
        </header>

        <!-- Feed -->
        <div class="space-y-0">
          <article 
            v-for="post in posts" 
            :key="post.id"
            class="group relative py-12 px-8 transition-all duration-300 border-b border-stone-200/40 dark:border-stone-800/40 first:pt-0"
          >
            <!-- Bookmark Hover Accent -->
            <div class="absolute left-0 top-0 bottom-0 w-[2px] bg-muji-accent scale-y-0 group-hover:scale-y-100 transition-transform duration-300 origin-top"></div>

            <div class="flex gap-6">
              <!-- Avatar -->
              <div class="flex-shrink-0">
                <img :src="post.avatar" class="w-9 h-9 rounded-full grayscale hover:grayscale-0 transition-all duration-500 opacity-80" :alt="post.username" />
              </div>

              <!-- Content Area -->
              <div class="flex-grow">
                <!-- Meta -->
                <div class="flex items-center gap-2 mb-4">
                  <span class="font-medium text-[14px] text-stone-700 dark:text-stone-300">{{ post.username }}</span>
                  <span class="text-[13px] text-stone-400 dark:text-stone-600">{{ post.handle }}</span>
                  <span class="text-[13px] text-stone-400/60 dark:text-stone-600/60 ml-auto">{{ post.timestamp }}</span>
                </div>

                <!-- Body -->
                <p class="text-[15px] leading-[1.7] text-stone-600 dark:text-stone-400 font-normal">
                  {{ post.content }}
                </p>

                <!-- Actions -->
                <div class="flex items-center gap-8 mt-8 opacity-25 group-hover:opacity-80 transition-opacity duration-300">
                  <button class="hover:text-muji-accent transition-colors" v-html="icons.reply"></button>
                  <button class="hover:text-muji-accent transition-colors" v-html="icons.repost"></button>
                  <button class="hover:text-muji-accent transition-colors" v-html="icons.like"></button>
                  <button class="hover:text-muji-accent transition-colors" v-html="icons.share"></button>
                </div>
              </div>
            </div>
          </article>
        </div>

        <!-- End of Feed -->
        <footer class="py-24 text-center opacity-20">
          <span class="font-serif italic text-sm">Quietly reaching the end.</span>
        </footer>
      </div>
    </main>
  </div>
</template>
