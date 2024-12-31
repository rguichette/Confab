import {useEffect, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';



function App() {

  const [connected, setConnected] = useState(false);
  const [serverMsg, setServerMsg] = useState("")

  useEffect(() => {
    async function checkConnection() {
        try {
            const res = await fetch("http://localhost:9999/ping");

            if (res.ok) {
                const data = await res.json();
                console.log(data.message); // Logs: "Hello from the Wails server!"
                setConnected(true);
                setServerMsg(data.message)
            }
        } catch (error) {
            console.error("Connection failed:", error);
            setConnected(false);
        }
    }

    checkConnection();
}, []);




    

    return (
        <div id="App">

         {connected ? "connected to backend": "Not connected"}
         <br/>
         {connected && `Message from server: ${serverMsg}`}
        </div>
    )
}

export default App
