# WebSocket Chat Application

This is a simple WebSocket-based chat application written in Go. It allows multiple users to join a chat room and communicate in real-time.

## Features

- Real-time chat using WebSocket
- Simple, minimalistic design
- Easy to deploy and run

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Git](https://git-scm.com/) for version control

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/chat-app.git
    cd chat-app
    ```

2. Build and run the application:

    ```bash
    go run main.go
    ```

3. Open your browser and go to `http://localhost:8080` to start using the chat.

## Usage

- Enter a username and start chatting with other users connected to the application.
- Messages are sent in real-time to all users in the chat room.

## Folder Structure

- `main.go` - Entry point of the application, handles WebSocket connections and message broadcasting.
- `templates/` - Contains HTML templates for the chat interface.
- `static/` - Contains any static files like CSS or JavaScript (if needed).

## CI/CD Pipeline

This project includes a CI/CD pipeline using GitHub Actions:

- **Builds** the application on each push or pull request.
- **Runs tests** to ensure functionality.
- **Deploys** to the specified server (or can be customized for cloud deployment).

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---

Feel free to customize this `README.md` as you see fit!
