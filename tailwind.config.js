/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["dim","night"],
  },
  plugins: [require("daisyui")],
}

