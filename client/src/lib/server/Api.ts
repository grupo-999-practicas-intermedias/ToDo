import { fail } from "@sveltejs/kit";
import type { Todo } from "../../routes/interface";

async function request(endpoint: string, method: string, body: string = "") {
  return await fetch(endpoint, {
    method,
    headers: {
      "Content-Type": "application/json"
    },
    body: body
  })
}

export async function getTodos() {
  const res = await fetch("http:///127.0.0.1:3000/todos",{
    method: 'GET',
  })

  if (res.status !== 200) {
    console.log(res.status)
    return null
  }

  const data = await res.json()

  return data

}


export async function addTodo(todo: Todo) {

  const res = await request("http://127.0.0.1:3000/todos",'POST', JSON.stringify(todo))

  if (res.status !== 200) return {
    status: res.status,
    id: null
  }

  const response = await res.json()
  return {
    status: res.status,
    id: response.id
  }
}

export async function updateTodo(id: any, todo: Todo) {


  const url = "http://127.0.0.1:3000/todos/"+id
  
  const res = await request(url, 'PUT', JSON.stringify(todo))

  if (res.status !== 200) return {
    status: res.status,
    id: null
  }

  const response = await res.json()
  return {
    status: res.status,
    id: response.id
  }
}

export async function deleteTodo(id: any) {
  const url = "http://127.0.0.1:3000/todos/"+id

  const res = await fetch(url, {
    method: 'DELETE',
  })

  if (res.status !== 200) return {
    status: res.status,
    id: null
  }

  const response = await res.json()

  console.log(response);
  return response
}