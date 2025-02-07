<!-- filepath: /home/jawdypus/Development/homeworkProject/frontend/src/components/Button.vue -->
<template>
  <button @click="fetchHomework" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
    Get Homework
  </button>
</template>

<script>
import axios from '../axiosConfig';

export default {
  methods: {
    async fetchHomework() {
      this.$emit('loading', true);
      try {
        const response = await axios.get('/api/homework');
        const data = response.data;
        this.$emit('homework-fetched', data);
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        this.$emit('loading', false);
      }
    }
  }
}
</script>