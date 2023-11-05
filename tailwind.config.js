/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html'],
  theme: {
    extend: {
      colors: {
        accent: {
          1: "rgb(var(--color-accent1) / <alpha-value>)",
          2: "rgb(var(--color-accent2) / <alpha-value>)",
          3: "rgb(var(--color-accent3) / <alpha-value>)"
        },
        content: "rgb(var(--color-content) / <alpha-value>)"
      }
    },
  },
  plugins: [],
}

