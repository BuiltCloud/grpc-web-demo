package helloworld

import (
	"context"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

// RunGrpcWebServer run grpc-web server.
func RunGrpcWebServer(ctx context.Context, s *grpc.Server, listenAddress string, allowedHeaders []string) {
	headers := []string{
		"x-grpc-web",
		"content-type",
		"content-length",
		"accept-encoding",
	}

	headers = append(headers, allowedHeaders...)

	opts := []grpcweb.Option{
		// gRPC-Web compatibility layer with CORS configured to accept on every request
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
			return true
		}),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	}
	wrappedGrpc := grpcweb.WrapServer(s, opts...)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsAcceptableGrpcCorsRequest(req) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ", "))
			return
		}
		if wrappedGrpc.IsGrpcWebSocketRequest(req) || wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(w, req)
			return
		}

		http.DefaultServeMux.ServeHTTP(w, req)
	})

	grpcweb := &http.Server{Addr: listenAddress, Handler: handler}
	fmt.Println("GRPCWeb listen address ",  listenAddress)
	go func() {
		if err := grpcweb.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("GrpcWeb Serve error.", err)
		}
	}()

	<-ctx.Done()

	if err := grpcweb.Close(); err != nil {
		fmt.Println("GrpcWeb server close error.", err)
	} else {
		fmt.Println("GrpcWeb server stopped.")
	}
}
