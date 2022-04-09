<script>
    import Normalize from './Normalize.svelte';
</script>

<Normalize/>

<svelte:head>
    <title>Buf Demo</title>
</svelte:head>

<style lang="scss">
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
  }

</style>

<div class="container">
    <div class="innerContainer">
        <h1>Buf Demo</h1>
        <p>Search for a &#128021; at petland to get started</p>
        <div>
            <input type="text" placeholder="Dog id"/>
            <button>Search</button>
        </div>
    </div>
</div>



