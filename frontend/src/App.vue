<script setup>

import {ref} from 'vue';

const homework = ref({});


const fetchHomework = () => {
  fetch('http://127.0.0.1:8080/get-homework')
  .then(response => {
    return response.json();
  })
  .then(data => {
    homework.value = data
    console.log(data);
  })
  .catch(error => {
    console.error('Parsing error:', error);
  });
};

</script>

<template>
  <button @click="fetchHomework">Get homework</button>
  <p>{{ homework }}</p>
</template>


