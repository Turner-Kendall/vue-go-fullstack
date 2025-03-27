## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```


### Pico overides 
###### Use More Specific Selectors
Like most frameworks Pico uses generic class names. To override them, I have attempted increase specificity:

```css
/* Framework has: */
.btn { margin:0; }

/* You override with: */
.custom-btn { margin:0 1rem; }  /* Preferred */
button.btn { margin:0 1rem; }   /* More specific */
```
If necessary, use **element + class combos** or **nested selectors**.

I have used !important sparingly, in the notes to try and remove alltogether. The most aggresive approach ( which I have removed now ) is 

```css
.btn {
  all: unset;
  display: inline-block;
  padding: 10px 20px;
}
```

---

### Todos

-- make catch all / 404 page 
-- try to remove al !important
-- Remember Me Implimentation