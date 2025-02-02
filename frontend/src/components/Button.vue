<!-- filepath: /home/jawdypus/Development/homeworkProject/frontend/src/components/Button.vue -->
<template>
  <button @click="fetchHomework" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
    Get Homework
  </button>
</template>

<script>
export default {
  methods: {
    async fetchHomework() {
      this.$emit('loading', true);
      try {
        const response = await fetch('http://127.0.0.1:8080/get-homework');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        this.$emit('homework-fetched', data);
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        this.$emit('loading', false);
      }
    }
  }
}
</script>

<style scoped>
/* Add any additional styles here */
</style>