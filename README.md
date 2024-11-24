# Chat Backend

This is a simple WebSocket-based chat server written in Go. The server allows multiple users to connect, send messages, and broadcast messages to all connected clients. It supports handling username uniqueness and user connection management.

## Features

- **WebSocket support**: Real-time messaging over WebSocket connections.
- **Username registration**: Users must choose a unique username to connect to the server.
- **Broadcast messaging**: Messages are broadcasted to all connected clients except the sender.
- **Graceful disconnects**: Handles user disconnections and cleans up resources accordingly.
- **Error handling**: Provides appropriate error messages when usernames are already taken or other issues arise.

## Technologies

- **Go**: The backend server is built with Go.
- **Gorilla WebSocket**: The server uses the [Gorilla WebSocket](https://github.com/gorilla/websocket) package to handle WebSocket connections.
- **HTTP**: The server uses the built-in HTTP package to handle WebSocket upgrade requests.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) installed on your system.
- A terminal with basic command-line knowledge.

### Steps to Install and Run

1. **Clone the repository**:
   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/kingmaj0r/chat-backend.git
   cd chat-backend
   ```

2. **Install dependencies**:
   Run the following command to install the required Go modules (if not already done):

   ```bash
   go mod tidy
   ```

3. **Build the server**:
   Compile the Go application into an executable:

   ```bash
   go build -o chat-backend
   ```

4. **Run the server**:
   Start the WebSocket chat server:

   ```bash
   ./chat-backend
   ```

   The server will start on port `8080`. You can change the port by modifying the code if needed.

5. **Access the server**:
   The WebSocket server will be accessible at `ws://localhost:8080/ws`. You can connect to this WebSocket endpoint from a frontend client (e.g., a web or mobile app).

## Usage

1. **Connect to the server**:
   Clients must first connect to the WebSocket server by providing a unique username. If the username is already taken, the server will respond with an error message.

2. **Send and receive messages**:
   Once connected, clients can send and receive messages in real time. Messages will be broadcasted to all clients, including the sender, but the sender's name is only included for other clients, not for themselves.

3. **Disconnecting**:
   When a client disconnects, the server cleans up the resources, and all connected clients will be notified of the disconnection.

## Code Structure

The project is organized into the following structure:

```
chat-backend/
│
├── main.go                 # Main entry point for the application. Sets up HTTP server and WebSocket handler.
└── src/                    
    ├── message.go          # Defines the Message struct used for communication between server and clients.
    ├── websocket.go        # Contains WebSocket handling logic (connection management, broadcasting, etc.).
```

### Description of Key Files:

- **`main.go`**: 
  - This is the entry point for the Go application. It initializes the HTTP server and sets up the WebSocket handler (`/ws` endpoint) using the `src.HandleConnections` function.
  - The server listens on port `8080` by default.

- **`src/message.go`**: 
  - This file contains the `Message` struct used to serialize and deserialize messages exchanged between clients and the server.
  - It defines two fields:
    - `Sender`: The username of the person sending the message.
    - `Text`: The content of the message.

- **`src/websocket.go`**: 
  - This file handles all WebSocket-related functionality.
  - It manages client connections, message broadcasting, and username registration. The `HandleConnections` function is responsible for upgrading the HTTP connection to a WebSocket connection and managing the client's session.
  - The `broadcast` function is used to send messages to all connected clients except the sender.

## Example Flow

1. A user connects to the WebSocket server with the message `{"sender": "Alice", "text": ""}`.
2. The server checks if the username "Alice" is already taken. If it's available, the user is connected, and a "welcome" message is broadcasted to all other clients.
3. Alice sends a message to the server, e.g., `{"sender": "Alice", "text": "Hello, everyone!"}`.
4. The server broadcasts the message to all other connected clients, including Alice.
5. When a user disconnects, the server broadcasts a "User has left" message.

## Example Client (JavaScript)

Here's an example of how you might connect to this WebSocket server from a frontend application:

```javascript
const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = () => {
  console.log("Connected to the server");

  // Send username
  socket.send(JSON.stringify({ sender: "Alice", text: "" }));
};

socket.onmessage = (event) => {
  const message = JSON.parse(event.data);
  console.log(`${message.sender}: ${message.text}`);
};

socket.onclose = () => {
  console.log("Disconnected from the server");
};
```

## License

This project is licensed under the MIT License.
