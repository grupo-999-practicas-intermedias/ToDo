

<form  method="POST" action="?/addTodo">
  <div>
    <div class="mt-2">
      <input id="title" 
      type="text" 
      name="title"
      required 
      placeholder="Ingrese el titulo"
      class="block w-full rounded-md border-0 py-1.5 bg-slate-100 text-gray-900 pl-2 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-800 focus:ring-2 focus:ring-inset focus:ring-indigo-200 sm:text-lg "
      >
    </div>
  </div>

  <div>
    <div class="flex items-center justify-between">
    </div>
    <div class="mt-2">
      <label for="password" class="block font-medium leading-6 py-2  text-gray-900 font-mono text-2xl ">Descripción</label>
      <textarea name="description"
      class="block w-full h-40 rounded-md border-0 bg-slate-100 pl-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-200 sm:text-lg"></textarea>
    </div>
  </div>

  <!-- Buttom -->
  <div class="flex items-center justify-end mt-4">
    <button type="submit" class=" inline-flex items-center  px-4 py-2 bg-indigo-600 border border-transparent rounded-md font-semibold text-xs text-white uppercase tracking-widest hover:bg-indigo-500 active:bg-indigo-900 focus:outline-none focus:border-indigo-900 focus:ring ring-indigo-300 disabled:opacity-25 transition ease-in-out duration-150">
      Agregar
    </button>
  </div>

</form>