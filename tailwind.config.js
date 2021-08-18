module.exports = {
  mode: "jit",
  purge: ["./public/**/*.html", "./src/views/**/*.{js,ts,tsx,hbs}"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
