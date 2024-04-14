import { connect, sendMsg } from "./api";

function App() {
  connect();

  var send = () => {
    sendMsg("HELLO");
  };

  return (
    <div>
      <button onClick={send}>Hit</button>
    </div>
  );
}

export default App;
