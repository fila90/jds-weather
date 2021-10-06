<script lang="ts">
  import type { TypeHourWeather } from "src/helpers/type";
  import { customRound, formatDateToHour } from "../helpers/util";

  export let dayHour: TypeHourWeather;
  export let todayHour: TypeHourWeather;
  export let isToday: boolean;
</script>

<div class="hour">
  <p>{formatDateToHour(new Date(dayHour.time))}</p>
  <p>
    {isToday ? dayHour.temp_c : customRound(dayHour.temp_c - todayHour.temp_c)}
  </p>
  <p>
    {isToday
      ? dayHour.pressure_mb
      : customRound(dayHour.pressure_mb - todayHour.pressure_mb)}
  </p>
  <p>
    {isToday
      ? dayHour.chance_of_rain
      : customRound(dayHour.chance_of_rain - todayHour.chance_of_rain)} %
  </p>
  <p>
    {isToday
      ? dayHour.precip_mm
      : customRound(dayHour.precip_mm - todayHour.precip_mm)} mm
  </p>
  <p>
    {isToday
      ? dayHour.wind_kph
      : customRound(dayHour.wind_kph - todayHour.wind_kph)}
  </p>
  <p>{dayHour.wind_dir}</p>
</div>

<style>
  .hour {
    flex-grow: 1;
    font-size: 0.8rem;
    font-weight: 400;
  }
</style>
