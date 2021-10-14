<script lang="ts">
	import { searchLocation } from '../helpers/api'
	import { debounce } from '../helpers/util'
	import LocationItem from './LocationItem.svelte'

	let location = ''
	let locationsQuery = Promise.resolve([])

	const debounced = debounce((l: string) => {
		locationsQuery = searchLocation(l)
	}, 300)

	function onLocationInput() {
		debounced(location)
	}

	function handleResetLocation() {
		location = ''
	}
</script>

<section class="location-search">
	<input
		type="text"
		bind:value={location}
		on:input={onLocationInput}
		placeholder="Find your location"
		class="location-search__input"
	/>

	{#if location}
		{#await locationsQuery}
			<h3>...waiting</h3>
		{:then locations}
			{#if locations.length}
				<div class="location-items-wrap">
					{#each locations as location (location.id)}
						<LocationItem {location} resetLocation={handleResetLocation} />
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
		padding: 2rem 0;
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
