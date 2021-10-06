<script lang="ts">
  import { currentWeather } from "../helpers/store";
</script>

<div>
  Current data:

  {#if !$currentWeather}
    <p>...waiting</p>
  {:else}
    <img
      src={$currentWeather.condition.icon}
      alt={$currentWeather.condition.text}
    />
    {$currentWeather.temp_c}
  {/if}
</div>
