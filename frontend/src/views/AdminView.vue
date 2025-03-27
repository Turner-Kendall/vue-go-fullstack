<template>
  <article>
    <Button @click="open">
      Open Modal
    </Button>
    <ModalsContainer />
  </article>
</template>

<script setup>
import { ModalsContainer, useModal } from 'vue-final-modal'
import ModalConfirmPlainCss from './ModalConfirmPlainCss.vue'

const mssg = '<p>The content of the modal does the modal grow, or...</p> <p>What about Multiple paragraphs? </p>'

const { open, close } = useModal({
  component: ModalConfirmPlainCss,
  attrs: {
    title: 'Hello World! Big Titles!',
    onConfirm() {
      close()
    },
  },
  slots: {
    default: mssg,
  },
})
</script>
