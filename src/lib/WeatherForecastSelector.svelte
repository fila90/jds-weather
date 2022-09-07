<script lang="ts">
	import feather from 'feather-icons'

	import { onMount } from 'svelte'

	import { activeWeatherParam } from '../helpers/store'

	onMount(() => {
		feather.replace()
	})

	const handleSelectorClick = (param: string) => () => {
		activeWeatherParam.set(param)
	}
	console.log($activeWeatherParam)
</script>

<div class="forecast-selector">
	<button
		title="Temp"
		on:click={handleSelectorClick('temp_c')}
		class={`forecast-selector__button ${
			$activeWeatherParam === 'temp_c'
				? 'forecast-selector__button--active'
				: ''
		}`}
	>
		<i class="forecast-selector__icon" data-feather="thermometer" />
	</button>
	<button
		title="Pressure"
		on:click={handleSelectorClick('pressure_mb')}
		class={`forecast-selector__button ${
			$activeWeatherParam === 'pressure_mb'
				? 'forecast-selector__button--active'
				: ''
		}`}
	>
		<i class="forecast-selector__icon" data-feather="activity" />
	</button>
	<button
		title="Precipitation"
		on:click={handleSelectorClick('precip_mm')}
		class={`forecast-selector__button ${
			$activeWeatherParam === 'precip_mm'
				? 'forecast-selector__button--active'
				: ''
		}`}
	>
		<i class="forecast-selector__icon" data-feather="cloud-rain" />
	</button>
	<button
		title="Wind"
		on:click={handleSelectorClick('wind_kph')}
		class={`forecast-selector__button ${
			$activeWeatherParam === 'wind_kph'
				? 'forecast-selector__button--active'
				: ''
		}`}
	>
		<i class="forecast-selector__icon" data-feather="wind" />
	</button>
</div>

<style>
	.forecast-selector {
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 1rem 0;
		gap: 0.5rem;
	}

	.forecast-selector__button {
		border-radius: 4px;
		cursor: pointer;
		padding: 2px 4px;
		transition: background-color ease 0.4s;
	}

	.forecast-selector__button--active {
		background-color: var(--color-active);
		color: var(--color-white);
	}

	.forecast-selector__icon {
		vertical-align: middle;
	}
</style>
