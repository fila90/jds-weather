<script lang="ts">
  import { searchLocation } from '../helpers/api'
  import { debounce } from '../helpers/util'
  import LocationItem from './LocationItem.svelte'

  let location = ''
  let promise = Promise.resolve([])

  const debounced = debounce((l: string) => {
    promise = searchLocation(l)
  }, 300)

  function handleLocationInput() {
    debounced(location)
  }
</script>

<section>
  <input type="text" bind:value={location} on:input={handleLocationInput} />

  {#if location}
    {#await promise}
      <p>...waiting</p>
    {:then locations}
      {#if locations.length}
        {#each locations as location (location.id)}
          <LocationItem {location} />
        {/each}
      {/if}
    {:catch error}
      <p style="color: red">{error.message}</p>
    {/await}
  {/if}
</section>
