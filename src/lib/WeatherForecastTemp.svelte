<script lang="ts">
	import type { TypeForecastDay } from '../helpers/type'

	import { customRound } from '../helpers/util'

	export let weather: TypeForecastDay
	export let todayWeather: TypeForecastDay
	export let isToday: boolean

	let tempMax: string
	let tempMin: string

	$: {
		tempMax = customRound(weather.day.maxtemp_c - todayWeather.day.maxtemp_c)
		tempMin = customRound(weather.day.mintemp_c - todayWeather.day.mintemp_c)
	}
</script>

<div class="wrap">
	<h4>Temp</h4>
	{isToday ? weather.day.maxtemp_c : tempMax}°C /
	{isToday ? weather.day.mintemp_c : tempMin}°C
</div>

<style>
	.wrap {
		border: 2px solid var(--color-text-secondary);
		border-radius: 0.5rem;
		padding: 0.5rem;
		margin: 0 0.5rem;
		display: flex;
		flex-direction: column;
		align-items: center;
	}
</style>
