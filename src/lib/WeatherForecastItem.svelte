<script lang="ts">
	import type { TypeForecastDay } from '../helpers/type'

	import { customRound } from '../helpers/util'

	export let weather: TypeForecastDay
	export let todayWeather: TypeForecastDay
	export let isToday: boolean
	export let label: string
	export let prop: string

	let value = ''

	$: {
		value = customRound(weather.day[prop] - todayWeather.day[prop])
	}
</script>

<div class="wrap">
	<h4>{label}</h4>
	{isToday ? weather.day[prop] : value}
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
