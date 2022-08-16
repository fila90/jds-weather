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
		debounced(encodeURIComponent(location))
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
		<div class="location-items-wrap">
			{#await locationsQuery then locations}
				<div class="location-items-wrap__dropdown">
					{#if locations.length}
						{#each locations as location (location.id)}
							<LocationItem {location} resetLocation={handleResetLocation} />
						{/each}
					{:else}
						<p>No results</p>
					{/if}
				</div>
			{:catch error}
				<p style="color: red">{error.message}</p>
			{/await}
		</div>
	{/if}
</section>

<style>
	@keyframes showUp {
		from {
			display: block;
			opacity: 0;
		}

		to {
			opacity: 1;
		}
	}

	@keyframes showDown {
		from {
			opacity: 1;
		}

		to {
			display: none;
			opacity: 0;
		}
	}

	.location-search {
		align-items: center;
		display: flex;
		flex-direction: column;
		justify-content: center;
		margin: auto;
		max-width: 400px;
		padding: 2rem 0;
		width: 90%;
	}

	.location-search__input {
		width: 100%;
	}

	.location-search__input:focus + .location-items-wrap {
		animation-duration: 0.2s;
		animation-fill-mode: forwards;
		animation-name: showUp;
	}

	.location-search__input:not(:focus) + .location-items-wrap {
		animation-delay: 0.05s;
		animation-duration: 0.2s;
		animation-fill-mode: forwards;
		animation-name: showDown;
	}

	.location-items-wrap {
		position: relative;
		width: 100%;
	}

	.location-items-wrap__dropdown {
		background-color: var(--color-white);
		border-radius: var(--border-radius);
		border: 1px solid var(--color-text-secondary);
		padding: 0.5rem;
		position: absolute;
		top: 0.5rem;
		width: 100%;
		z-index: 1;
		max-height: 400px;
		overflow: auto;
	}
</style>
