// App.js
import { useState, useEffect } from "react";
import { connect, sendMsg } from "./api";
import Header from "./Components/header/header";
import ChatHistory from "./Components/chatHistory/chatHistory";

function App() {
  const [chatHistory, setChatHistory] = useState([]);

  useEffect(() => {
    connect((msg) => {
      console.log("New Message");
      setChatHistory((prevChatHistory) => [...prevChatHistory, msg]);
    });
  }, []);

  const send = () => {
    sendMsg("HELLO");
  };

  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <button onClick={send}>Send</button>
    </div>
  );
}

export default App;
