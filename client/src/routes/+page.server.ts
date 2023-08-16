import { generateUUID } from '$lib/genereateUid';
import { addTodo, deleteTodo, getTodos, updateTodo } from '$lib/server/Api';
import { URL_SERVER } from "$env/static/private";
import type { Actions } from './$types';
import type { Todo } from './interface';


export async function load() {
  const data = await getTodos( URL_SERVER);
  return {
    todos: data == undefined ? [] : data
  }
}


export const actions = {
  addTodo: async ( {request} ) => {
    const data = await request.formData();

    console.log(data.get('title'));

    // create a todo
    const todo: Todo = {
      id: generateUUID(),
      title: data.get('title')?.toString(),
      description: data.get('description')?.toString(),
      completed: false
    }
    // send the data to the server
    const res = addTodo(todo, URL_SERVER);

    // return the response
    return res
  },
  
  updateTodo: async ( {request} ) => {
    console.log('updateTodo');
    const data = await request.formData();
    console.log(data);
    // create a todo
    const todo: Todo = {
      id: data.get('id')?.toString(),
      title: data.get('title')?.toString(),
      description: data.get('description')?.toString(),
      completed: true
    }

    console.log("entra aca ?",todo);

    const res = await updateTodo(todo.id, todo, URL_SERVER);
    console.log(res);
    return res

  },
  deleteTodo: async ( {request} ) => {
    const data = await request.formData();
    console.log(data);
    const id = data.get('id')?.toString();
    const res = await deleteTodo(id, URL_SERVER);
    console.log(res);
  }
} 