<!-- filepath: /home/jawdypus/Development/homeworkProject/frontend/src/components/Homework.vue -->
<template>
  <div class="bg-gray-800 rounded p-4 m-4 flex">
    <div class="tabs w-1/4">
      <div
        class="tab p-2 m-2 rounded cursor-pointer"
        :class="{'bg-gray-600': selectedTab !== key, 'bg-gray-500': selectedTab === key}"
        v-for="(value, key) in homework"
        :key="key"
        @click="selectTab(key)"
        style="min-width: 150px; white-space: normal;"
      >
        {{ key }}
      </div>
    </div>
    <div class="content-container flex w-3/4 bg-gray-500 rounded ml-4 p-4">
      <div class="tab-content w-1/2 p-4">
        <div v-if="selectedTab">
          <h1 class="text-xl font-bold mb-4">{{ homework[selectedTab].class_theme }}</h1>
          <div v-if="homework[selectedTab].presence">
            <h2 class="text-lg font-semibold mb-2">Присутність</h2>
            <div v-for="(present, student) in homework[selectedTab].presence" :key="student" class="mb-2">
              <label class="flex items-center">
                <input type="checkbox" :checked="present" @change="togglePresence(student)" class="mr-2" />
                {{ student }}
              </label>
            </div>
          </div>
          <h2 class="text-lg font-semibold mb-2">Відгук про урок</h2>
          <textarea v-model="feedback" @input="updateBoilerplateText" class="w-full p-4 border-2 border-gray-300 rounded bg-white text-black" rows="4"></textarea>
          <h2 class="text-lg font-semibold mb-2">Завдання</h2>
          <div v-for="(task, index) in homework[selectedTab].tasks" :key="index" class="mb-2">
            <label class="flex items-center">
              <input type="checkbox" :checked="checkedTasks.includes(task)" @change="toggleTask(task)" class="mr-2" />
              {{ task }}
            </label>
          </div>
          <input v-model="newTask" @keyup.enter="addTask" placeholder="Add a new task" class="w-full p-2 border-2 border-gray-300 rounded bg-white text-black mt-2" />
        </div>
      </div>
      <div class="vertical-line bg-gray-600 w-px mx-4"></div>
      <div class="boilerplate-text w-1/2 p-4">
        <p v-html="boilerplateText"></p>
        <button @click="copyToClipboard" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-4">Copy to Clipboard</button>
      </div>
    </div>
    <div v-if="showCopyMessage" class="copy-message">
      Boilerplate text copied to clipboard!
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
      selectedTab: null,
      boilerplateText: '',
      feedback: 'На уроці ми',
      checkedTasks: [],
      newTask: '',
      showCopyMessage: false
    }
  },
  watch: {
    homework: {
      immediate: true,
      handler(newHomework) {
        this.selectedTab = Object.keys(newHomework)[0]; // Select the first tab by default
        this.updateBoilerplateText();
      }
    },
    selectedTab() {
      this.updateBoilerplateText();
    }
  },
  methods: {
    selectTab(key) {
      this.selectedTab = key;
      this.checkedTasks = []; // Reset checked tasks when a new tab is selected
      this.updateBoilerplateText();
    },
    togglePresence(student) {
      this.homework[this.selectedTab].presence[student] = !this.homework[this.selectedTab].presence[student];
      this.updateBoilerplateText();
    },
    toggleTask(task) {
      const index = this.checkedTasks.indexOf(task);
      if (index > -1) {
        this.checkedTasks.splice(index, 1);
      } else {
        this.checkedTasks.push(task);
      }
      this.updateBoilerplateText();
    },
    addTask() {
      if (this.newTask.trim() !== '') {
        const task = this.newTask.trim();
        this.homework[this.selectedTab].tasks.push(task);
        this.checkedTasks.push(task); // Add the new task to checkedTasks
        this.newTask = '';
        this.updateBoilerplateText();
      }
    },
    updateBoilerplateText() {
      if (!this.selectedTab) return '';

      const absentStudents = Object.entries(this.homework[this.selectedTab].presence)
        .filter(([student, present]) => !present)
        .map(([student]) => `🎓 ${student}`)
        .join('<br>');

      const absentText = absentStudents
        ? `На занятті були відсутні:<br>${absentStudents}<br><br>${this.getMakeupClassText()}<br><br>`
        : 'На уроці були присутні всі. Дякую за активність!<br><br>';

      const tasksText = this.checkedTasks.length
        ? `📕ДОМАШНЯ ПРАКТИКА:<br>- ${this.checkedTasks.join('<br>- ')}<br><br>`
        : '';

      this.boilerplateText = `Добрий день, шановні батьки! ☀️<br><br>${absentText}💡 ${this.feedback}<br><br>${tasksText}З повагою викладач Міжнародної школи Logika - Максим 💜`;
    },
    getMakeupClassText() {
      const groupName = this.selectedTab;
      const parts = groupName.split('_');
      const dayAbbreviation = parts[2].toUpperCase();
      const time = parts[3];

      const dayMap = {
        'ПН': 'понеділок',
        'ВТ': 'вівторок',
        'СР': 'середу',
        'ЧТ': 'четвер',
        'ПТ': 'п’ятницю',
        'СБ': 'суботу',
        'НД': 'неділю'
      };

      const day = dayMap[dayAbbreviation] || 'день';
      const normalizedTime = time.replace('.', ':');
      const [hours, minutes] = normalizedTime.split(':').map(Number);
      const makeupTime = new Date();
      makeupTime.setHours(hours);
      makeupTime.setMinutes(minutes - 30);

      const makeupTimeString = `${makeupTime.getHours()}:${makeupTime.getMinutes().toString().padStart(2, '0')}`;

      return `Запрошуємо на відпрацювання в ${day} о ${makeupTimeString}`;
    },
    copyToClipboard() {
      const el = document.createElement('textarea');
      el.value = this.boilerplateText.replace(/<br>/g, '\n').replace(/<\/?[^>]+(>|$)/g, ""); // Convert HTML to plain text
      document.body.appendChild(el);
      el.select();
      document.execCommand('copy');
      document.body.removeChild(el);
      this.showCopyMessage = true;
      setTimeout(() => {
        this.showCopyMessage = false;
      }, 2000); // Hide the message after 2 seconds
    }
  }
}
</script>

<style scoped>
.vertical-line {
  height: auto;
}

.copy-message {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 20px;
  border-radius: 10px;
  text-align: center;
  z-index: 1000;
}
</style>