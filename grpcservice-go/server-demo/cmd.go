package main

import (
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	helloworld "grpcwebgo"
	"net/http"
)
// protoc --go_out=plugins=grpc:. greet.proto
func main()  {
	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &helloworld.TestServiceImpl{})
	wrappedGrpc := grpcweb.WrapServer(grpcServer)


	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedGrpc.ServeHTTP(resp, req)
		/*
			https://github.com/improbable-eng/grpc-web/tree/master/go/grpcweb
					wrappedServer.ServeHTTP(resp, req)
				}
				// Fall back to other servers.
				http.DefaultServeMux.ServeHTTP(resp, req)
		*/
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", 5001),
		Handler: http.HandlerFunc(handler),
	}
	http2.ConfigureServer(&httpServer, nil)
	//if err := httpServer.ListenAndServe(); err != nil {
	//	grpclog.Fatalf("failed starting http server: %v", err)
	//}
	httpServer.ListenAndServeTLS("E:/gRpc/examples/grpcservice-go/certs/localhost.cert", "E:/gRpc/examples/grpcservice-go/certs/localhost.key")
	/*	reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}*/

	//headers := http.Header{}
	//headers.Add("Access-Control-Request-Method", "POST")
	//headers.Add("Access-Control-Request-Headers", "origin, x-something-custom, x-grpc-web, accept")
	//req := httptest.NewRequest("OPTIONS", "http://host/grpc/improbable.grpcweb.test.TestService/Echo", nil)
	//req.Header = headers
	//resp := httptest.NewRecorder()
	//wrappedServer.ServeHTTP(resp, req)
	//
	//
	//tlsHttpServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
	//	if wrappedGrpc.IsGrpcWebRequest(req) {
	//		wrappedGrpc.ServeHTTP(resp, req)
	//	}
	//	// Fall back to other servers.
	//	http.DefaultServeMux.ServeHTTP(resp, req)
	//})
}
