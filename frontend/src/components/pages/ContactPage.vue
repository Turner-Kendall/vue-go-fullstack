<template>

  <article>

    <header>
      <h6>Thanks for stopping by. I'm happy to hear from everyone - so - feel free to
        drop me a line.</h6>
    </header>

    <form class="contact-form" action="https://api.web3forms.com/submit" method="POST">

      <input type="hidden" name="access_key" value="ac31ec4b-b4e6-4f1a-9dcc-3b0701992210" />
      <input type="hidden" name="subject" value="New Contact Form Submission from weblog network" />
      <input type="checkbox" name="botcheck" class="hidden" style="display: none;">
      <input type="hidden" name="from_name" value="gÄuth" />

      <div class="form-group-container">
        <div class="form-group">
          <label for="name" class="form-label">Name</label>
          <input id="name" name="name" class="form-input" placeholder="Your name" type="text" />
        </div>

        <div class="form-group-inline">
          <div class="form-group">
            <label for="email" class="form-label">Email</label>
            <input id="email" name="email" class="form-input" placeholder="Your email" type="email" />
          </div>
          <div class="form-group">
            <label for="phone" class="form-label">Phone</label>
            <input id="phone" name="phone" class="form-input" placeholder="+1 (234) 56789" type="text" />
          </div>
        </div>

        <div class="form-group">
          <label for="message" class="form-label">Message</label>
          <textarea class="form-textarea" id="message" name="message" placeholder="Your message"></textarea>
        </div>
      </div>
      <button class="form-submit" type="submit">Send Message</button>
    </form>

  </article>
</template>


<style scoped>
.form-group-container {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.form-group-inline {
  display: flex;
  gap: 15px;
}

.form-group {
  flex: 1;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  margin-bottom: 0.5rem;
}

.form-input,
.form-textarea {
  padding: 0.5rem;
  display: flex;
  height: 2rem;
  width: 100%;
  font-size: 0.875rem;
  line-height: 1.25rem;
}

.form-textarea {
  min-height: 120px;
}
</style>
