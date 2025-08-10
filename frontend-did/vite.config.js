import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      '/ping': 'http://localhost:8080',
      '/v1/api/did': 'http://localhost:8080'
    }
  }
});
