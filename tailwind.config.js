/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './templates/**/*.html'],
  theme: {
    extend: {
      colors: {
        accent: {
          1: "rgb(var(--color-accent1) / <alpha-value>)",
          2: "rgb(var(--color-accent2) / <alpha-value>)",
          3: "rgb(var(--color-accent3) / <alpha-value>)",
          4: "rgb(var(--color-accent4) / <alpha-value>)",
        },
        content: "rgb(var(--color-content) / <alpha-value>)"
      },
      gridTemplateColumns: {
        "form-layout": "repeat(auto-fill, minmax(120px, 1fr))"
      }
    },
  },
  plugins: [],
}

