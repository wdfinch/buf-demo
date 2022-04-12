<script lang="ts">
	import { GetDogResponse } from '@buf/grpc_web_petland_dogapis/petland/dog/v1/dog_api_pb.js'
	import { JsonView } from '@zerodevx/svelte-json-view'

	export let resp: GetDogResponse.AsObject
</script>

<div class="container">
	<img src={resp.dog.profileImageUrl} alt="dog" />
	<JsonView
		{resp}
		json={resp}
		--nodeColor="#08fdf5"
		--leafStringColor="#08fd7f"
		--leafNumberColor="#fda608"
		--leafBooleanColor="#f660ff"
	/>
</div>

<style lang="scss">
	img {
		width: 100%;
		border-radius: 4px;
		margin-bottom: 2rem;
	}

	.container {
		width: 70%;
		margin-top: 4rem;
		border-radius: 4px;
		padding: 2.2rem;
		background: #212834;
		margin-bottom: 4rem;

		& :global(*) {
			font-size: 2.1rem;
			font-family: Arial;
		}
	}
</style>
