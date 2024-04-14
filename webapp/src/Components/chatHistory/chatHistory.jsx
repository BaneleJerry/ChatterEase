import React from "react";
import "./chatHistory.css";

export default function ChatHistory(props) {
  const messages = props.chatHistory.map((msg, index) => (
    <p key={index}>{msg.data}</p>
  ));

  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {messages}
    </div>
  );
}
