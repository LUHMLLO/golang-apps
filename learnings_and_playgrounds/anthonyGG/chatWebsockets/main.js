//const WebSocket = require('ws');

let socket = new Websocket("ws://localhost:3080/ws")

socket.onmessage = (event) => { 
    console.log("received from the server: ", event.data)
}

socket.send("hello from client")