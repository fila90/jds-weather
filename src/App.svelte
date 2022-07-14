<script lang="ts">
	import { FORECAST_DAY } from './helpers/const'
	import { geoLocation } from './helpers/store'
	import LocationSearch from './lib/LocationSearch.svelte'
	import WeatherForecast from './lib/WeatherForecast.svelte'
	import WeatherForecastChart from './lib/WeatherForecastChart.svelte'
	import WeatherForecastSelector from './lib/WeatherForecastSelector.svelte'
</script>

<main>
	<LocationSearch />
	{#if $geoLocation.lat && $geoLocation.lon}
		<WeatherForecastSelector />
		<WeatherForecastChart />
		<WeatherForecast forecastDay={FORECAST_DAY.YESTERDAY} />
		<WeatherForecast forecastDay={FORECAST_DAY.TODAY} />
		<WeatherForecast forecastDay={FORECAST_DAY.TOMORROW} />
	{/if}
</main>
