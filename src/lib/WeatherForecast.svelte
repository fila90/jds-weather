<script lang="ts">
	import type { TypeForecastDay } from '../helpers/type'

	import { FORECAST_DAY } from '../helpers/const'
	import { forecastWeather } from '../helpers/store'
	import { formatDateToDay, formatDateToReadable } from '../helpers/util'

	import WeatherForecastHour from './WeatherForecastHour.svelte'
	import WeatherForecastItem from './WeatherForecastItem.svelte'
	import WeatherForecastTemp from './WeatherForecastTemp.svelte'

	export let forecastDay: string
	const isToday = forecastDay === FORECAST_DAY.TODAY

	let weather: TypeForecastDay
	let todayWeather: TypeForecastDay

	$: {
		weather = $forecastWeather.get(forecastDay)
		todayWeather = $forecastWeather.get(FORECAST_DAY.TODAY)
	}
</script>

{#if !weather || !todayWeather}
	<p>...loading</p>
{:else}
	<div class="forecast">
		<div class="forecast__day">
			<div>
				<span class="forecast__day-day">
					{formatDateToDay(new Date(weather.date))}
				</span>
				<span class="forecast__day-date">
					{formatDateToReadable(new Date(weather.date))}
				</span>
			</div>

			<div class="forecast__day-content">
				<WeatherForecastTemp {weather} {todayWeather} {isToday} />
				<WeatherForecastItem
					{weather}
					{todayWeather}
					{isToday}
					label="Wind"
					prop="maxwind_kph"
				/>
				<WeatherForecastItem
					{weather}
					{todayWeather}
					{isToday}
					label="Humidity"
					prop="avghumidity"
				/>
				<WeatherForecastItem
					{weather}
					{todayWeather}
					{isToday}
					label="Precip"
					prop="totalprecip_mm"
				/>
				<WeatherForecastItem
					{weather}
					{todayWeather}
					{isToday}
					label="UV"
					prop="uv"
				/>
				<WeatherForecastItem
					{weather}
					{todayWeather}
					{isToday}
					label="Visibility"
					prop="avgvis_km"
				/>
			</div>

			<img src={weather.day.condition.icon} alt={weather.day.condition.text} />
		</div>

		<hr />

		<div class="forecast__hour-wrap">
			<div class="forecast__hour-legend">
				<p>hour</p>
				<p>temp</p>
				<p>pressure</p>
				<p>rain</p>
				<p>precip</p>
				<p>wind</p>
				<p>direction</p>
			</div>

			<div class="forecast__hour-content">
				{#each weather.hour as hour, index (hour.time)}
					<WeatherForecastHour
						dayHour={hour}
						todayHour={todayWeather.hour[index]}
						{isToday}
					/>
				{/each}
			</div>
		</div>
	</div>
{/if}

<style>
	.forecast {
		border-radius: 0.25rem;
		margin: 0.5rem 0;
		padding: 0.5rem;
		border: 1px solid var(--color-text-secondary);
	}

	.forecast__day {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.forecast__day-day {
		font-weight: bold;
		font-size: 1.5rem;
	}

	.forecast__day-date {
		color: var(--color-text-secondary);
		font-size: 0.8rem;
	}

	.forecast__day-content {
		display: flex;
		align-items: center;
	}

	hr {
		border: 1px solid var(--color-text-secondary);
		max-width: 80%;
		margin: 2rem auto;
	}

	.forecast__hour-wrap {
		display: flex;
	}

	.forecast__hour-legend {
		font-size: 0.8rem;
		font-weight: bold;
	}

	.forecast__hour-content {
		display: flex;
		flex-grow: 1;
		padding: 0 1rem;
		justify-content: stretch;
	}
</style>
