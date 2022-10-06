import React from 'react';
import logo from './logo.svg';
import './App.css';
import Go from "./wasm_exec"

function App() {

  const go = new Go();

  const wasmBrowserInstantiate = async (wasmModuleUrl: string, importObject) => {
    let response = undefined;

    // Check if the browser supports streaming instantiation
    if (WebAssembly.instantiateStreaming) {
      // Fetch the module, and instantiate it as it is downloading
      response = await WebAssembly.instantiateStreaming(
          fetch(wasmModuleUrl),
          importObject
      );
    } else {
      // Fallback to using fetch to download the entire module
      // And then instantiate the module
      const fetchAndInstantiateTask = async () => {
        const wasmArrayBuffer = await fetch(wasmModuleUrl).then(response =>
            response.arrayBuffer()
        );
        return WebAssembly.instantiate(wasmArrayBuffer, importObject);
      };

      response = await fetchAndInstantiateTask();
    }

    return response;
  };

  const runWasmAdd = async (x: number, y: number) => {
    // Get the importObject from the go instance.
    const importObject = go.importObject;

    // Instantiate our wasm module
    const wasmModule = await wasmBrowserInstantiate("../main.wasm", importObject);

    // Allow the wasm_exec go instance, bootstrap and execute our wasm module
    go.run(wasmModule.instance);

    // Call the Add function export from wasm, save the result
    const addResult = wasmModule.instance.exports.add(x, y);

    // Set the result onto the body
    document.body.textContent = `Hello World! addResult: ${addResult}`;
  };

  runWasmAdd(24, 53);
  runWasmAdd(24, 3);
  runWasmAdd(4, 3);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
