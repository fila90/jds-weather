<script lang="ts">
	import type { TypeLocationItem } from '../helpers/type'

	import { geoLocation } from '../helpers/store'

	export let location: TypeLocationItem

	function onLocationSelect() {
		geoLocation.setGeoLocation({
			lat: `${location.lat}`,
			lon: `${location.lon}`
		})
	}
</script>

<div class="location-item" on:click={onLocationSelect}>
	{location.name}
</div>

<style>
	.location-item {
		padding: 0.5rem 1rem;
		width: 100%;
		cursor: pointer;
	}
	.location-item:not(:last-child) {
		border-bottom: 1px solid var(--color-text-secondary);
	}

	.location-item:hover {
		background-color: var(--color-active);
		color: var(--color-white);
	}
</style>
