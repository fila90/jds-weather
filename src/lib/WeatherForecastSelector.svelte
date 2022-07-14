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
		on:click={handleSelectorClick('will_it_rain')}
		class={`forecast-selector__button ${
			$activeWeatherParam === 'will_it_rain'
				? 'forecast-selector__button--active'
				: ''
		}`}
	>
		<i class="forecast-selector__icon" data-feather="cloud-rain" />
	</button>
	<button
		on:click={handleSelectorClick('wind_kph')}
		class={`forecast-selector__button ${
			$activeWeatherParam === 'wind_kph'
				? 'forecast-selector__button--active'
				: ''
		}`}
	>
		<i class="forecast-selector__icon" data-feather="wind" />
	</button>
	<button
		on:click={handleSelectorClick('precip_mm')}
		class={`forecast-selector__button ${
			$activeWeatherParam === 'precip_mm'
				? 'forecast-selector__button--active'
				: ''
		}`}
	>
		<i class="forecast-selector__icon" data-feather="umbrella" />
	</button>
</div>

<style>
	.forecast-selector {
		display: flex;
		align-items: center;
	}

	.forecast-selector__button {
		border-radius: 4px;
		cursor: pointer;
		padding: 2px 4px;
		margin-right: 4px;
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
