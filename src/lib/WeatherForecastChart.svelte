<script lang="ts">
	import type { TypeForecastDay } from 'src/helpers/type'

	import Chart, { type ChartItem } from 'chart.js/auto'

	import { FORECAST_DAY } from '../helpers/const'

	import { activeWeatherParam, forecastWeather } from '../helpers/store'
	import { formatDateToHour } from '../helpers/util'

	import Loader from './Loader.svelte'

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
						label: 'Yesterday',
						borderColor: 'rgba(251,86,7,0.2)',
						backgroundColor: 'rgba(251,86,7,0.2)',
						data: weatherYesterday.hour.map(hour => hour[$activeWeatherParam])
					},
					{
						label: 'Today',
						borderColor: 'rgba(255,0,110,0.6)',
						backgroundColor: 'rgba(255,0,110,0.6)',
						data: weatherToday.hour.map(hour => hour[$activeWeatherParam])
					},
					{
						label: 'Tomorrow',
						borderColor: 'rgba(58,134,255,0.4)',
						backgroundColor: 'rgba(58,134,255,0.4)',
						data: weatherTomorrow.hour.map(hour => hour[$activeWeatherParam])
					}
				]
			},
			options: {
				elements: {
					line: {
						borderWidth: 2
					}
				}
			}
		})
	}
</script>

<canvas
	bind:this={canvasCtx}
	aria-label="Visual representation of weather forecast"
	role="img"
	height="50"
/>
