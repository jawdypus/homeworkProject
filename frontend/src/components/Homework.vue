<template>
    <div class="bg-gray-800 rounded p-2 m-2">
      <div v-if="homework" class="tabs">
        <div class="tab" v-for="(value, key) in homework" :key="key" @click="selectTab(key)">
          {{ key }}
        </div>
        <div v-if="selectedTab" class="tab-content">
          <pre>{{ homework[selectedTab] }}</pre>
        </div>
      </div>
    </div>
</template>
  
<script>
  export default {
    props: {
      homework: {
        type: Object,
        required: true
      }
    },
    data() {
      return {
        selectedTab: null
      }
    },
    watch: {
      homework: {
        immediate: true,
        handler(newHomework) {
          this.selectedTab = Object.keys(newHomework)[0]; // Select the first tab by default
        }
      }
    },
    methods: {
      selectTab(key) {
        this.selectedTab = key;
      }
    }
  }
</script>