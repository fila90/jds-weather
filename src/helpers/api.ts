import type {
	TypeCurrentWeather,
	TypeForecastDay,
	TypeGeoLocation,
	TypeLocation,
	TypeLocationItem
} from './type'

import { startOfYesterday } from 'date-fns'

import { API } from './const'
import { fetcher, formatDateToISO } from './util'

const addApiSearchParams = (
	baseUrl: string,
	querySearch: { [key: string]: string }
): string => {
	const url = new URL(window.location.origin + baseUrl)
	const { searchParams } = url

	for (const key in querySearch) {
		searchParams.append(key, querySearch[key])
	}

	return url.href
}

export const searchLocation = async (
	location: string
): Promise<Array<TypeLocationItem>> => {
	const url = addApiSearchParams(API.SEARCH, { q: location })

	return await fetcher(url)
}

export const searchCurrentWeather = (
	location: TypeGeoLocation
): Promise<{
	current: TypeCurrentWeather
	location: TypeLocation
}> => {
	const url = addApiSearchParams(API.CURRENT, {
		q: `${location.lat},${location.lon}`
	})

	return fetcher(url)
}

export const searchForecastWeather = (
	location: TypeGeoLocation
): Promise<{
	location: TypeLocation
	current: TypeCurrentWeather
	forecast: { forecastday: TypeForecastDay[] }
}> => {
	const url = addApiSearchParams(API.FORECAST, {
		q: `${location.lat},${location.lon}`,
		days: '2'
	})

	return fetcher(url)
}

export const searchHistoryWeather = (
	location: TypeGeoLocation
): Promise<{
	location: TypeLocation
	forecast: { forecastday: TypeForecastDay[] }
}> => {
	const url = addApiSearchParams(API.HISTORY, {
		q: `${location.lat},${location.lon}`,
		dt: formatDateToISO(startOfYesterday())
	})

	return fetcher(url)
}
