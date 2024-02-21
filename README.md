# Go-gRPC-Demo
<h3>Why?</h3>
<p>
In a REST or GraphQL architecture we have<br/>
    Client sending request<br/>
    Server giving response<br/>
  <br/>
But this process is
    <li>Slow</li>
    <li>Synchronous</li>
    <li>Not Scalable</li>
  <br/>

<h3>How?</h3>
How to make the client server intercation Scalable?<br/>

For that we use Remote Procedure Calls (RPC)<br/>

RPC directly calls functions from client to server
Instead of JSON, we use protobuf files
Here, payload size is small and accelerates communication
<br/>
<br/>
Client Server communication can happen via Streaming (sequences of mesages)<br/>
Stream is continous flow of data<br/>
Its asynchronous<br/>
Extremely scalable<br/>
Usecase: MICROSERVICES!, Blockchains<br/>
<hr/>
</p>
There are mainly 4 types of communications in GRPC

1. Unary API
    <p>Regular request-response</p>

2. Server Streaming
     <p>Client sends a request to server</p>
     <p>Server sends a strea of data to client </p>

3. Client Streaming
     <p>Client sends a stream of data to server</p>
     <p>Server provides a siple response to client</p>

4. Bi-Directional Streaming
     <p>Both client and server can communicate via streaming.</p>
     <p>And they can do this parallely, </p>
     <p>Not like request response wehre client sneds first and then server gives response - not like this.</p>
     <p>Its like a two way-traffic</p>
     <p>Even though its a stream, not a queue, sequence of messages is preserved</p>
<hr/>

<h3>To Run the project</h3>
Install the dependencies

    go mod tidy

Run the server

    go run server/main.go

Run the client

    go run client/main.go


