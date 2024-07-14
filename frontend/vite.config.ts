import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import {adorableCSS} from "adorable-css/vite"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [adorableCSS(), react()],
})
