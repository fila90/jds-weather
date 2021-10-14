<script lang="ts">
	import { FORECAST_DAY } from './helpers/const'
	import { geoLocation } from './helpers/store'
	import LocationSearch from './lib/LocationSearch.svelte'
	import WeatherForecast from './lib/WeatherForecast.svelte'
</script>

<main>
	{#if !$geoLocation.lat || !$geoLocation.lon}
		<LocationSearch />
	{:else}
		<WeatherForecast forecastDay={FORECAST_DAY.YESTERDAY} />
		<WeatherForecast forecastDay={FORECAST_DAY.TODAY} />
		<WeatherForecast forecastDay={FORECAST_DAY.TOMORROW} />
	{/if}
</main>
