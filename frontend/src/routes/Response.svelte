<script lang="ts">
    import {GetDogResponse} from '@buf/grpc_web_petland_dogapis/petland/dog/v1/dog_api_pb'
    import {JsonView} from '@zerodevx/svelte-json-view'

    export let resp: GetDogResponse.AsObject
    console.log(resp)

</script>

<div class="container">
    <JsonView {resp} json={resp} --nodeColor="#08fdf5" --leafStringColor="#08fd7f" --leafNumberColor="#fda608" --leafBooleanColor="#f660ff" />
</div>

<style lang="scss">
  .container {
    margin-top: 4rem;
    border-radius: 4px;
    padding: 2rem;
    background: #212834;

    & :global(*) {
      font-size: 2.1rem;
      font-family: Arial;
    }
  }
</style>