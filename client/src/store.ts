
import { writable } from "svelte/store";
import type { Todo } from "./routes/interface";


export const TodoData = writable<Todo>({
  id: '',
  title: '',
  description: '',
  completed: false
})