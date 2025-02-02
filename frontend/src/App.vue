<!-- filepath: /home/jawdypus/Development/homeworkProject/frontend/src/App.vue -->
<template>
  <div id="app" class="min-h-screen bg-gray-900 text-white flex flex-col items-center justify-center">
    <div :class="{'flex-grow': !homework, 'mt-4': homework}" class="flex items-center justify-center w-full mb-4">
      <Button @homework-fetched="handleHomeworkFetched" @loading="setLoading" />
    </div>
    <div v-if="loading" class="absolute inset-0 flex items-center justify-center w-full h-full">
      <div class="loading-gradient flex items-center justify-center">
        <span class="loading-text">Loading homework data...</span>
      </div>
    </div>
    <div v-if="homework && !loading" class="homework-container w-full flex items-center justify-center mt-4">
      <Homework :homework="homework" />
    </div>
  </div>
</template>

<script>
import Button from './components/Button.vue';
import Homework from './components/Homework.vue';

export default {
  name: 'App',
  components: {
    Button,
    Homework
  },
  data() {
    return {
      homework: null,
      loading: false
    }
  },
  methods: {
    handleHomeworkFetched(data) {
      this.homework = data;
      this.loading = false;
    },
    setLoading(isLoading) {
      this.loading = isLoading;
    }
  }
};
</script>

<style scoped>
.loading-gradient {
  width: 70%;
  height: 60%;
  background: linear-gradient(45deg, #1f2937, #4b5563, #6b7280, #9ca3af);
  background-size: 300% 300%;
  animation: gradientAnimation 3s ease infinite;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.loading-text {
  color: white;
  font-size: 1.2rem;
  text-align: center;
}

@keyframes gradientAnimation {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.homework-container {
  width: 70%; /* Set a fixed width */
  overflow: auto; /* Add scroll if content overflows */
}
</style>