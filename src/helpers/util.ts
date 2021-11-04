import { format } from 'date-fns'

export const debounce = (fn: (any) => void, delay: number): (any) => void => {
	let timeOutId: NodeJS.Timeout
	return function (...args) {
		if (timeOutId) {
			clearTimeout(timeOutId)
		}
		timeOutId = setTimeout(() => {
			fn(...args)
		}, delay)
	}
}

export const fetcher = async (url: string) => {
	const res = await fetch(url, { method: 'GET' })
	return await res.json()
}

export const customRound = (value: number, step = 0.5): string => {
	if (!Number(value)) {
		return '0'
	}
	const inverse = 1.0 / step
	const res = Math.round(value * inverse) / inverse
	return res > 0 ? `+${res}` : `${res}`
}

export const formatDateToISO = (date: Date): string =>
	format(date, 'yyyy-MM-dd')

export const formatDateToReadable = (date: Date): string =>
	format(date, 'MM/dd/yyyy')

export const formatDateToDay = (date: Date): string => format(date, 'iiii')

export const formatDateToHour = (date: Date): string => format(date, 'p')
