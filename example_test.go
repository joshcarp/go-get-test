package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	pingv1 "github.com/joshcarp/go-get-test/gen/connect/ping/v1"
	pingv1connect "github.com/joshcarp/go-get-test/gen/connect/ping/v1/pingv1connect"
	"net/http"
)

// ExamplePingServer implements some trivial business logic. The Protobuf
// definition for this API is in proto/connect/ping/v1/ping.proto.
type ExamplePingServer struct {
	pingv1connect.UnimplementedPingServiceHandler
}

// Ping implements pingv1connect.PingServiceHandler.
func (*ExamplePingServer) Ping(
	_ context.Context,
	request *connect.Request[pingv1.PingRequest],
) (*connect.Response[pingv1.PingResponse], error) {
	return connect.NewResponse(
		&pingv1.PingResponse{
			Number: request.Msg.Number,
			Text:   request.Msg.Text,
		},
	), nil
}

func Example_handler() {
	// protoc-gen-connect-go generates constructors that return plain net/http
	// Handlers, so they're compatible with most Go HTTP routers and middleware
	// (for example, net/http's StripPrefix). Each handler automatically supports
	// the Connect, gRPC, and gRPC-Web protocols.
	mux := http.NewServeMux()
	mux.Handle(
		pingv1connect.NewPingServiceHandler(
			&ExamplePingServer{}, // our business logic
		),
	)
	// You can serve gRPC's health and server reflection APIs using
	// github.com/bufbuild/connect-grpchealth-go and
	// github.com/bufbuild/connect-grpcreflect-go.
	_ = http.ListenAndServeTLS(
		"localhost:8080",
		"internal/testdata/server.crt",
		"internal/testdata/server.key",
		mux,
	)
	// To serve HTTP/2 requests without TLS (as many gRPC clients expect), import
	// golang.org/x/net/http2/h2c and golang.org/x/net/http2 and change to:
	// _ = http.ListenAndServe(
	// 	"localhost:8080",
	// 	h2c.NewHandler(mux, &http2.Server{}),
	// )
}
