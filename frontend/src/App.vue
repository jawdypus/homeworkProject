<template>
  <div id="app" class="min-h-screen bg-gray-900 text-white flex flex-col items-center justify-center">
    <div :class="{'flex-grow': !homework, 'mt-4': homework}" class="flex items-center justify-center w-full">
      <Button @homework-fetched="handleHomeworkFetched" />
    </div>
    <div v-if="homework" class="w-full flex items-center justify-center mt-4">
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
      homework: null
    }
  },
  methods: {
    handleHomeworkFetched(data) {
      this.homework = data;
    }
  }
};
</script>