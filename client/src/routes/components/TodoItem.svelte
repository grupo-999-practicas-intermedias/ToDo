<script>
	import { enhance } from "$app/forms";
	import { fade, fly } from "svelte/transition";
  /**
	 * @type {any}
	 */
  export let id;
  /**
	 * @type {any}
	 */
   export let title;
  /**
	 * @type {any}
	 */
   export let description ;
  /**
	 * @type {boolean}
	 */
   export let completed ;

  

</script>

<ul role="list" class="grid mb-4 " in:fly={{x:200, duration:2000}} out:fade  >
  <form class="col-span-1 rounded-md divide-yrounded-lg  shadow {completed ? 'line-through divide-gray-300 bg-slate-200': 'divide-gray-200 bg-slate-100'}" method="POST" use:enhance >
    <div class="flex w-full items-center justify-between space-x-6 p-6">
      <input type="hidden" name="id" value={id} />
      <input type="hidden" name="description" value={description} />
      <input type="hidden" name="title" value={title} />
      <div class="flex-1 truncate">
        <div class="flex items-center space-x-3">
          <h3 class="truncate text-lg font-medium text-gray-900">{title}</h3>
        </div>
        <p class="mt-1 truncate text-sm text-gray-500">{description}</p>
      </div>
    </div>
    <div>
      <div class="-mt-px flex divide-x divide-gray-200">
        <div class="flex w-0 flex-1"  >
          <button  formaction="?/updateTodo" disabled={completed}
          class="relative -mr-px inline-flex w-0 flex-1 items-center justify-center gap-x-3 rounded-bl-lg border border-transparent py-4 text-sm font-semibold text-gray-900 {completed ? '' : 'hover:bg-slate-300'}">
            <svg class="h-7 w-7 text-gray-400" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle opacity="0.5" cx="12" cy="12" r="10" stroke="#1C274C" stroke-width="1.5"/>
              <path d="M8.5 12.5L10.5 14.5L15.5 9.5" stroke="#1C274C" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            Completada
          </button>
        </div>
        <div class="-ml-px flex w-0 flex-1">
          <button  formaction="?/deleteTodo" 
          class="relative inline-flex w-0 flex-1 items-center justify-center gap-x-3 rounded-br-lg border border-transparent py-4 text-sm font-semibold text-gray-900 hover:bg-slate-300">
            <svg class="h-7 w-7 text-gray-400" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M9.17065 4C9.58249 2.83481 10.6937 2 11.9999 2C13.3062 2 14.4174 2.83481 14.8292 4" stroke="#1C274C" stroke-width="1.5" stroke-linecap="round"/>
              <path d="M20.5 6H3.49988" stroke="#1C274C" stroke-width="1.5" stroke-linecap="round"/>
              <path d="M18.3735 15.3991C18.1965 18.054 18.108 19.3815 17.243 20.1907C16.378 21 15.0476 21 12.3868 21H11.6134C8.9526 21 7.6222 21 6.75719 20.1907C5.89218 19.3815 5.80368 18.054 5.62669 15.3991L5.16675 8.5M18.8334 8.5L18.6334 11.5" stroke="#1C274C" stroke-width="1.5" stroke-linecap="round"/>
              <path d="M9.5 11L10 16" stroke="#1C274C" stroke-width="1.5" stroke-linecap="round"/>
              <path d="M14.5 11L14 16" stroke="#1C274C" stroke-width="1.5" stroke-linecap="round"/>
              </svg>
            Eliminar
          </button>
        </div>
      </div>
    </div>
  </form>
</ul>
