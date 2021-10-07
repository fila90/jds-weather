<script lang="ts">
  import type { TypeLocationItem } from '../helpers/type'

  import { geoLocation } from '../helpers/store'

  export let location: TypeLocationItem

  function onLocationSelect() {
    geoLocation.setGeoLocation({
      lat: `${location.lat}`,
      lon: `${location.lon}`
    })
  }
</script>

<div on:click={onLocationSelect}>
  {location.name}
</div>
