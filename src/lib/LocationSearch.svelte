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

<section class="location-search">
	<input
		type="text"
		bind:value={location}
		on:input={handleLocationInput}
		placeholder="Search your location"
		class="location-search__input"
	/>

	{#if location}
		{#await promise}
			<p>...waiting</p>
		{:then locations}
			{#if locations.length}
				<div class="location-items-wrap">
					{#each locations as location (location.id)}
						<LocationItem {location} />
					{/each}
				</div>
			{/if}
		{:catch error}
			<p style="color: red">{error.message}</p>
		{/await}
	{/if}
</section>

<style>
	.location-search {
		display: flex;
		margin: auto;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 20rem 0;
		width: 90%;
		max-width: 400px;
	}

	.location-search__input {
		width: 100%;
	}

	.location-items-wrap {
		margin-top: 1rem;
	}
</style>
