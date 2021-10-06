<script lang="ts">
  import { geoLocation } from "../helpers/store";

  import type { TypeLocationItem } from "src/helpers/type";

  export let location: TypeLocationItem;

  function onLocationSelect() {
    geoLocation.setGeoLocation({
      lat: `${location.lat}`,
      lon: `${location.lon}`,
    });
  }
</script>

<div on:click={onLocationSelect}>
  {location.name}
</div>
