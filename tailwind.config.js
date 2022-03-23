module.exports = {
  content: ["./index.html", "./dashboard/**/*.{vue,ts}"],
  darkMode: "class",
  mode: "jit",
  theme: {
    extend: {
      fontFamily: {
        sans: ["InterVariable"],
        mono: ["JetBrains MonoVariable"],
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/line-clamp"),
    require('@tailwindcss/typography'),
  ],
}