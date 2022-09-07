<script lang="ts">
	import type { TypeForecastDay } from './helpers/type'

	import { FORECAST_DAY } from './helpers/const'
	import { forecastWeather, geoLocation } from './helpers/store'

	import Loader from './lib/Loader.svelte'
	import LocationSearch from './lib/LocationSearch.svelte'
	import WeatherForecastChart from './lib/WeatherForecastChart.svelte'
	import WeatherForecastSelector from './lib/WeatherForecastSelector.svelte'

	let forecast: TypeForecastDay

	$: {
		forecast = $forecastWeather.get(FORECAST_DAY.TODAY)
	}
</script>

<main>
	<LocationSearch />
	{#if !$geoLocation.lat || !$geoLocation.lon}
		<div class="center">
			<h2>No location selected!</h2>
		</div>
	{:else if !forecast}
		<Loader />
	{:else}
		<WeatherForecastSelector />
		<WeatherForecastChart />
	{/if}
</main>
