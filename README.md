## Sample gRPC Requests

* Reflection
  ```shell
  $ grpcurl -v -plaintext  -proto ./pkg/greeter/greeter.proto localhost:80 list
  ```
  ```text
  EchoServer
  grpc.reflection.v1.ServerReflection
  grpc.reflection.v1alpha.ServerReflection
  ```

* GRPC request
  ```shell
  grpcurl -d '{"name": "El"}' --plaintext 127.0.0.1:50051  EchoServer/GrpcPing
  ```
  ```text
  {
  "message": "Hello, El!",
  "headers": "map[:authority:[127.0.0.1:50053] content-type:[application/grpc] grpc-accept-encoding:[gzip] user-agent:[grpcurl/v1.8.9 grpc-go/1.57.0]]"
  }
  ```
* GRPCs request
  ```shell
  grpcurl -d '{"name": "El"}' --insecure 127.0.0.1:50053   EchoServer/GrpcPing
  ```
  ```text
  {
  "message": "Hello, El!",
  "headers": "map[:authority:[127.0.0.1:50053] content-type:[application/grpc] grpc-accept-encoding:[gzip] user-agent:[grpcurl/v1.8.9 grpc-go/1.57.0]]"
  }
  ```

## Sample REST Requests

* Using HTTP/1.1 

  ```shell
  $ curl -v --raw -d '{"Message":"Hello, El!"}' --http1.1 http://localhost:8080
  ```
  
  ```text
  > POST / HTTP/1.1
  > Host: localhost:8080
  > User-Agent: curl/7.81.0
  > Accept: */*
  > Content-Length: 24
  > Content-Type: application/x-www-form-urlencoded
  >
  * Mark bundle as not supporting multiuse
    < HTTP/1.1 200 OK
    < Date: Fri, 15 Dec 2023 09:49:42 GMT
    < Content-Length: 240
    < Content-Type: text/plain; charset=utf-8
    <
    Host: localhost:8080
    Path: /
    Method: POST
    Protocol: HTTP/1.1
    Scheme: http
    Hostname: pc
    Headers:
      Content-Length: 24
      Content-Type: application/x-www-form-urlencoded
      User-Agent: curl/7.81.0
      Accept: */*
    Body:
    {"Message":"Hello, El!"}
  ```

* Using HTTP/2

  ```shell
  $ curl -vk --raw -d '{"Message":"Hello, El!"}' --http2 https://localhost:8443
  ```
    
  ```text
  > POST / HTTP/2
  > Host: localhost:8443
  > user-agent: curl/7.81.0
  > accept: */*
  > content-length: 24
  > content-type: application/x-www-form-urlencoded
  >
  * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * Connection state changed (MAX_CONCURRENT_STREAMS == 250)!
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    * TLSv1.2 (OUT), TLS header, Supplemental data (23):
    * We are completely uploaded and fine
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
    * TLSv1.2 (IN), TLS header, Supplemental data (23):
      < HTTP/2 200
      < content-type: text/plain; charset=utf-8
      < content-length: 241
      < date: Fri, 15 Dec 2023 09:50:56 GMT
      <
      Host: localhost:8443
      Path: /
      Method: POST
      Protocol: HTTP/2.0
      Scheme: https
      Hostname: pc
      Headers:
        Content-Length: 24
        Content-Type: application/x-www-form-urlencoded
        User-Agent: curl/7.81.0
        Accept: */*
      Body:
      {"Message":"Hello, El!"}
  ```
  
## Proto
* Generate proto
  ```shell
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  protoc *.proto --go_out=./ --go-grpc_out=./
  ```
