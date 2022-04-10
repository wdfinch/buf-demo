<script lang="ts">
    import Normalize from './Normalize.svelte';

    let id: number | null = 0;
    let isInputValid: boolean | null = null
    const max = 10000

    function handleSearch() {
        if (!Number.isInteger(id) || id > max) {
            isInputValid = false
        }
    }

</script>

<Normalize/>

<svelte:head>
    <title>Buf Demo</title>
</svelte:head>

<div class="container">
    <div class="innerContainer">
        <h1>Buf Demo</h1>
        <p>Search for a &#128021; at petland to get started</p>
        <div>
            <input type=number min=0 max=max bind:value={id} on:input={isInputValid = null} placeholder="Dog id"/>
            <button on:click={handleSearch}>Search</button>

            {#if isInputValid === false}
                <div class="error">ID must be an integer less than {max}</div>
            {/if}
        </div>
    </div>
</div>

<style lang="scss">
  .error {
    color: red;
    margin-top: 1rem;
  }

  .container {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
  }

  h1 {
    margin-bottom: 2rem;
  }

  .innerContainer {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  p {
    margin-bottom: 4rem;
  }

  button {
    font-size: 2rem;
    padding: 1.2rem 2.2rem;
    background-color: azure;
    transition: background-color 1ms linear;
    cursor: pointer;
    border: 2px solid black;

    &:hover {
      background-color: #e5ffff;
    }
  }

  button, input {
    border: 2px solid black;
    border-radius: 4px;
  }

  input {
    border: 2px solid black;
    font-size: 2rem;
    padding: 1.2rem;
    margin-right: 2rem;
    width: 20rem;
  }
</style>

