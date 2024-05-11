### gRPC App

This gRPC app demonstrates various request methods. To set it up, follow these steps:

1. Download dependencies:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
   export PATH="$PATH:$(go env GOPATH)/bin"
   sudo apt install -y protobuf-compiler
   ```

2. Uncomment the method you want to test in the file `client/main.go`:
   - `callSayHello(client)` for a single request from client to server.
   - `callSayHelloServerStream(client, names)` for a set of requests from client to server.
   - `callSayHelloClientStream(client, names)` for a set of requests from server to client.
   - `callHelloBidirectionalStream(client, names)` for bidirectional requests.

3. Run the app:
   - Open one terminal and run the server:
     ```bash
     cd server
     go run main.go
     ```
   - Open another terminal and run the client:
     ```bash
     cd client
     go run main.go
     ```

### Notes
- Make sure to have the `protoc` compiler and gRPC tools installed.
