<script lang="ts">
	type LoaderSize = 's' | 'm' | 'l'

	export let size: LoaderSize = 'm'
</script>

<div class="loader-ring">
	<div class={`loader-ring__content loader-ring__content--${size}`}>
		<div />
		<div />
		<div />
		<div />
	</div>
</div>

<style>
	.loader-ring {
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 1rem;
	}

	.loader-ring__content {
		display: inline-block;
		position: relative;
		width: 4rem;
		height: 4rem;
	}

	.loader-ring__content div {
		box-sizing: border-box;
		display: block;
		position: absolute;
		margin: 0.5rem;
		border: 0.25rem solid var(--color-success);
		border-radius: 50%;
		animation: spinAnimatioin 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
		border-color: var(--color-success) transparent transparent transparent;
	}

	.loader-ring__content--s div {
		width: 1rem;
		height: 1rem;
	}

	.loader-ring__content--m div {
		width: 2rem;
		height: 2rem;
	}

	.loader-ring__content--l div {
		width: 3rem;
		height: 3rem;
	}

	.loader-ring__content div:nth-child(1) {
		animation-delay: -0.45s;
	}

	.loader-ring__content div:nth-child(2) {
		animation-delay: -0.3s;
	}

	.loader-ring__content div:nth-child(3) {
		animation-delay: -0.15s;
	}

	@keyframes spinAnimatioin {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}
</style>
