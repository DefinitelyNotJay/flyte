import type { Config } from 'tailwindcss'

export default <Partial<Config>>{
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        muji: {
          bg: 'rgb(var(--muji-bg) / <alpha-value>)',
          ink: 'rgb(var(--muji-ink) / <alpha-value>)',
          accent: 'rgb(var(--muji-accent) / <alpha-value>)',
        }
      },
      fontFamily: {
        sans: ['DM Sans', 'sans-serif'],
        serif: ['Noto Serif', 'serif'],
      }
    }
  }
}
