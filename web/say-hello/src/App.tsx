import React, { useState } from "react";
import logo from "./logo.svg";
import "./App.css";

interface Message {
  msg: string;
}

function App() {
  let [msg, setMsg] = useState("");

  const sayHello = function () {
    fetch("/api/hello")
      .then((res) => {
        if (!res.ok) {
          throw Error("hello request failed");
        }
        return res.json();
      })
      .then((data: Message) => {
        setMsg(data.msg);
      })
      .catch((err) => {
        console.log(`fetch failed ${err}`);
      });
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Welcome to say hello, you latest message is: {msg}</p>
        <button onClick={sayHello}>Fetch</button>
      </header>
    </div>
  );
}

export default App;
