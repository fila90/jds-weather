import type {
	TypeCurrentWeather,
	TypeForecastDay,
	TypeGeoLocation,
	TypeLocation,
	TypeLocationItem
} from './type'

import { startOfYesterday } from 'date-fns'

import { fetcher, formatDateToISO } from './util'


const getSearchApiUrl = (): string => {
	const base = `${import.meta.env.VITE_API_URL}/search.json`
	return addApiSearchParams(base, { key: `${import.meta.env.VITE_API_KEY}` })
}

const getCurrentApiUrl = (): string => {
	const base = `${import.meta.env.VITE_API_URL}/current.json`
	return addApiSearchParams(base, { key: `${import.meta.env.VITE_API_KEY}` })
}

const getForecastApiUrl = (): string => {
	const base = `${import.meta.env.VITE_API_URL}/forecast.json`
	return addApiSearchParams(base, { key: `${import.meta.env.VITE_API_KEY}` })
}

const getHistoryApiUrl = (): string => {
	const base = `${import.meta.env.VITE_API_URL}/history.json`
	return addApiSearchParams(base, { key: `${import.meta.env.VITE_API_KEY}` })
}

const addApiSearchParams = (
	baseUrl: string,
	querySearch: { [key: string]: string }
): string => {
	const url = new URL(baseUrl)
	const { searchParams } = url

	for (const key in querySearch) {
		searchParams.append(key, querySearch[key])
	}

	return url.href
}

export const searchLocation = async (
	location: string
): Promise<Array<TypeLocationItem>> => {
	const url = addApiSearchParams(getSearchApiUrl(), { q: location })

	return await fetcher(url)
}

export const searchCurrentWeather = (
	location: TypeGeoLocation
): Promise<{
	current: TypeCurrentWeather
	location: TypeLocation
}> => {
	const url = addApiSearchParams(getCurrentApiUrl(), {
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
	const url = addApiSearchParams(getForecastApiUrl(), {
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
	const url = addApiSearchParams(getHistoryApiUrl(), {
		q: `${location.lat},${location.lon}`,
		dt: formatDateToISO(startOfYesterday())
	})

	return fetcher(url)
}
