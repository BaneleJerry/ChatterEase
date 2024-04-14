var socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
  console.log("Attempting to connect...");

  socket.onopen = () => {
    console.log("Connected!");
  };
};

socket.onmessage = (msg) => {
  console.log(msg);
};

socket.onclose = () => {
  console.log("Disconnected!");
};

socket.onerror = (err) => {
  console.log(err);
};

let sendMsg = (msg) => {
  console.log("Sending msg: " + msg);
  socket.send(msg);
};

export { connect, sendMsg };
