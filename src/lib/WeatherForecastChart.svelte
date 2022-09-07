<script lang="ts">
	import type { TypeForecastDay } from 'src/helpers/type'

	import Chart, { type ChartItem } from 'chart.js/auto'

	import { FORECAST_DAY } from '../helpers/const'

	import { activeWeatherParam, forecastWeather } from '../helpers/store'
	import { formatDateToHour } from '../helpers/util'

	let weatherYesterday: TypeForecastDay
	let weatherToday: TypeForecastDay
	let weatherTomorrow: TypeForecastDay
	let canvasCtx: ChartItem
	let myChart: Chart<'line', number[], string>

	$: {
		weatherYesterday = $forecastWeather.get(FORECAST_DAY.YESTERDAY)
		weatherToday = $forecastWeather.get(FORECAST_DAY.TODAY)
		weatherTomorrow = $forecastWeather.get(FORECAST_DAY.TOMORROW)
	}

	$: if (canvasCtx && weatherYesterday && weatherToday && weatherTomorrow) {
		// if we already created the chart and data updated, destroy!!!
		if (myChart) {
			myChart.destroy()
		}

		myChart = new Chart(canvasCtx, {
			type: 'line',
			data: {
				labels: weatherToday.hour.map(hour =>
					formatDateToHour(new Date(hour.time))
				),
				datasets: [
					{
						backgroundColor: 'rgba(173,179,188,0.5)',
						borderColor: 'rgba(173,179,188,0.5)',
						borderDash: [1, 5],
						borderWidth: 1,
						data: weatherYesterday.hour.map(hour => hour[$activeWeatherParam]),
						label: 'Yesterday',
						tension: 0.1
					},
					{
						backgroundColor: 'rgba(250,38,160,1)',
						borderColor: 'rgba(250,38,160,1)',
						data: weatherToday.hour.map(hour => hour[$activeWeatherParam]),
						label: 'Today',
						tension: 0.1
					},
					{
						backgroundColor: 'rgba(47,243,224,0.6)',
						borderColor: 'rgba(47,243,224,0.6)',
						borderDash: [1, 3],
						borderWidth: 2,
						data: weatherTomorrow.hour.map(hour => hour[$activeWeatherParam]),
						label: 'Tomorrow',
						tension: 0.1
					}
				]
			},
			options: {
				plugins: {
					tooltip: {
						callbacks: {
							title: ([context]) =>
								`${context.dataset.label} @ ${context.label}`,
							label: context => {
								switch ($activeWeatherParam) {
									case 'temp_c':
										return `${context.formattedValue}°C`
									case 'pressure_mb':
										return `${context.formattedValue}°`
									case 'wind_kph':
										return `${context.formattedValue}/kph`
									case 'precip_mm':
										return `${context.formattedValue} mm`
									default:
										return context.raw ? 'Yes' : 'No'
								}
							}
						}
					}
				}
			}
		})
	}
</script>

<canvas
	bind:this={canvasCtx}
	aria-label="Visual representation of weather forecast"
	class="canvas"
/>

<style>
	.canvas {
		max-height: 50vh;
	}
</style>
