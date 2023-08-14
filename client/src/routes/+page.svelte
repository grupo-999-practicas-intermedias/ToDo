<script>
	import TodoForm from "./components/TodoForm.svelte";
	import TodoItem from "./components/TodoItem.svelte";

  export let data;
</script>


<svelte:head>
  <title>Todo-App</title>
</svelte:head>



<!-- Createa container -->
<div class="px-20 bg-slate-300 full-height">
  <!-- ... -->
  <p class="font-mono text-4xl pt-16 pb-5"  >Practicas intermedias - Grupo 999</p>
  <p class="font-mono text-3xl pb-5" >Github actions</p>
  <p class="font-mono text-2xl pb-10" >ToDo List - Svelte + REST API Go + Testing</p>
  <hr class="w-full border-2 border-gray-500 mb-4" />
  <div  class="flex items-start justify-between">
    <p class="font-mono text-2xl " >Agrega un elemento a la lista</p>
    <p class="font-mono text-2xl " >Lista de tareas pendientes:</p>
  </div>
  <div class="grid gap-4 grid-cols-2">
    <div >
        <TodoForm />
    </div>
    <div class="flex flex-col  ">
      <!-- List -->
      <div class="grow w-full">
        {#each data.todos as todo (todo.id)}
        {#if !todo.completed}
        <TodoItem {...todo} />
        {/if}
        {/each}
      </div>
      <!-- Line of todos done -->
      <hr class="w-full border-2 border-gray-500" />
      <!-- Lista de completados -->
      <div class="grow w-full">
        <p class="text-right pt-2 font-mono text-2xl " >Lista de tareas completadas:</p>
        {#each data.todos as todo (todo.id)}
        {#if todo.completed}
        <TodoItem {...todo} />
        {/if}
        {/each}
      </div>
    </div>
  </div>
</div>

<style>
  .full-height {
    height: 100vh;
  }
</style>