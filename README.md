# Go-gRPC-Demo
<hr/>
<p>
In a REST or GraphQL architecture we have<br/>
    Client sending request<br/>
    Server giving response<br/>
  <br/>
But this process is <br/>
    Slow<br/>
    Synchronous<br/>
    Not Scalable<br/>
  <br/>
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
    Regular request-response

2. Server Streaming
    Client sends a request to server
    Server sends a strea of data to client 

3. Client Streaming
    Client sends a stream of data to server
    Server provides a siple response to client

4. Bi-Directional Streaming
    Both client and server can communicate via streaming.
    And they can do this parallely, 
    Not like request response wehre client sneds first and then server gives response - not like this.
    Its like a two way-traffic
    Even though its a stream, not a queue, sequence of messages is preserved
<hr/>
