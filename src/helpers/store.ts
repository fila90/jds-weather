import { writable, derived } from 'svelte/store'

import {
	searchCurrentWeather,
	searchForecastWeather,
	searchHistoryWeather
} from './api'

import type { TypeGeoLocation } from './type'

function createGeoLocationStore() {
	const defaultLat = localStorage.getItem('lat')
	const defaultLon = localStorage.getItem('lon')
	const { subscribe, set } = writable<TypeGeoLocation>({
		lat: defaultLat,
		lon: defaultLon
	})

	return {
		subscribe,
		setGeoLocation: (geoLocation: TypeGeoLocation) => {
			localStorage.setItem('lat', `${geoLocation.lat}`)
			localStorage.setItem('lon', `${geoLocation.lon}`)
			set(geoLocation)
		}
	}
}

export const geoLocation = createGeoLocationStore()

export const currentWeather = derived(
	geoLocation,
	($geoLocation, set) => {
		searchCurrentWeather($geoLocation).then(res => set(res.current))
	},
	null
)

export const forecastWeather = derived(
	geoLocation,
	($geoLocation, set) => {
		Promise.all([
			searchForecastWeather($geoLocation),
			searchHistoryWeather($geoLocation)
		])
			.then(([forecast, history]) => {
				const weather = new Map([
					['yesterday', history.forecast.forecastday[0]],
					['today', forecast.forecast.forecastday[0]],
					['tomorrow', forecast.forecast.forecastday[1]]
				])
				set(weather)
			})
			.catch(err => {
				throw Error(err)
			})
	},
	new Map([
		['yesterday', null],
		['today', null],
		['tomorrow', null]
	])
)

// export const historyWeather = derived(
//   geoLocation,
//   ($geoLocation, set) => {
//     searchHistoryWeather($geoLocation).then((res) =>
//       set(res.forecast.forecastday[0].day)
//     );
//   },
//   null
// );
