/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class", // Enable dark mode via a class
  content: [
    "./src/**/*.{js,jsx,ts,tsx}", // Paths to all your components
  ],
  theme: {
    extend: {
      colors: {
        customColor: '#ff6347', // Custom color
      },
      fontFamily: {
        sans: ['Inter', 'sans-serif'], // Adding the Inter font
      },
    },
  },
  plugins: [],
};
