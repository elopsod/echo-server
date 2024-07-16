package main

import (
	"context"
	"crypto/tls"
	"fmt"
	env "github.com/caitlinelfring/go-env-default"
	pb "github.com/elopsod/echo-server/echoServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {

	httpPort := env.GetDefault("HTTP_PORT", "8080")
	httpsPort := env.GetDefault("HTTPS_PORT", "8443")
	grpcPort := env.GetDefault("GRPC_PORT", "50051")
	grpcsPort := env.GetDefault("GRPCS_PORT", "50053")

	certFile := "certs/server.crt"
	keyFile := "certs/server.key"

	certificate, _ := tls.LoadX509KeyPair(certFile, keyFile)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}
	tlsCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
	})

	var wg sync.WaitGroup
	wg.Add(1)

	// HTTP server setup
	go func() {
		defer wg.Done()

		httpHandler := http.NewServeMux()
		httpHandler.HandleFunc("/", HttpPing)

		httpServer := &http.Server{
			Addr:    fmt.Sprintf(":%s", httpPort),
			Handler: httpHandler,
		}
		log.Printf("HTTP Server is listening on port %s", httpPort)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to serve HTTP: %v", err)
		}
	}()

	// HTTPS server setup
	go func() {
		defer wg.Done()

		httpsHandler := http.NewServeMux()
		httpsHandler.HandleFunc("/", HttpPing)

		httpsServer := &http.Server{
			Addr:      fmt.Sprintf(":%s", httpsPort),
			Handler:   httpsHandler,
			TLSConfig: tlsConfig,
		}

		http.HandleFunc("/", HttpPing)
		log.Printf("HTTPS Server is listening on port %s", httpsPort)
		if err := httpsServer.ListenAndServeTLS(certFile, keyFile); err != nil {
			log.Fatalf("failed to serve HTTPS: %v", err)
		}
	}()
	// GRPC server setup
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {

			log.Fatalf("Failed to listen: %v", err)
		}

		srv := grpc.NewServer()
		pb.RegisterEchoServerServer(srv, &server{})

		reflection.Register(srv)

		log.Printf("GRPC Server is listening on port %s", grpcPort)
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// GRPCS server setup
	go func() {
		defer wg.Done()
		srv := grpc.NewServer(grpc.Creds(tlsCreds))

		pb.RegisterEchoServerServer(srv, &server{})
		reflection.Register(srv)

		listener, _ := net.Listen("tcp", fmt.Sprintf(":%s", grpcsPort))

		log.Printf("GRPCS Server is listening on port %s", grpcsPort)
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	wg.Wait()
}

func HttpPing(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	host := r.Host
	method := r.Method
	protocol := r.Proto
	hostname, _ := os.Hostname()
	var scheme string = "http"
	if r.TLS != nil {
		scheme = "https"
	}
	body, _ := io.ReadAll(r.Body)

	headers := []string{}
	for key, values := range r.Header {
		headers = append(headers, fmt.Sprintf("%s: %s", key, strings.Join(values, " ")))
	}

	params := []string{}
	for key, values := range r.URL.Query() {
		params = append(params, fmt.Sprintf("%s: %s", key, strings.Join(values, " ")))
	}
	// Send headers, path, and body in the response
	fmt.Fprintln(w, "Host:", host)
	fmt.Fprintln(w, "Path:", path)
	fmt.Fprintln(w, "Method:", method)
	fmt.Fprintln(w, "Protocol:", protocol)
	fmt.Fprintln(w, "Scheme:", scheme)
	fmt.Fprintln(w, "Hostname:", hostname)
	fmt.Fprintln(w, "Headers:\n\t", strings.Join(headers[:], "\n\t "))
	fmt.Fprintln(w, "Params:\n\t", strings.Join(params[:], "\n\t "))
	fmt.Fprintln(w, "Body:\n\t", string(body))
}

type server struct {
	pb.EchoServerServer
}

func (s *server) GrpcPing(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	name := req.GetName()
	md, _ := metadata.FromIncomingContext(ctx)

	response := &pb.Response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Headers: fmt.Sprintf("%s", md),
	}
	return response, nil
}
