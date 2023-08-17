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

export async function getTodos( url: string) {
  const res = await fetch( url + "/todos",{
    method: 'GET',
  })

  if (res.status !== 200) {
    console.log(res.status)
    return null
  }

  const data = await res.json()

  return data

}


export async function addTodo(todo: Todo, url : string) {

  const res = await request( url + "/todos",'POST', JSON.stringify(todo))

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

export async function updateTodo(id: any, todo: Todo, url : string) {
  
  const res = await request(url +"/todos/"+id, 'PUT', JSON.stringify(todo))

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

export async function deleteTodo(id: any, url : string) {

  const res = await fetch(url + "/todos/"+id, {
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