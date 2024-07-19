package raft_test

import (
	"io"
	"net/http"

	"github.com/alekseyvit/raft"
	"github.com/alekseyvit/raft/transport"
	"github.com/alekseyvit/raft/transport/raftgrpc"
	"github.com/alekseyvit/raft/transport/rafthttp"
	"google.golang.org/grpc"
)

type stateMachine struct{}

func (stateMachine) Apply([]byte)                           {}
func (stateMachine) Snapshot() (r io.ReadCloser, err error) { return }
func (stateMachine) Restore(io.ReadCloser) (err error)      { return }

func Example_gRPC() {
	srv := grpc.NewServer()
	node := raft.NewNode(stateMachine{}, transport.GRPC)
	raftgrpc.RegisterHandler(srv, node.Handler())
}

func Example_http() {
	node := raft.NewNode(stateMachine{}, transport.HTTP)
	handler := rafthttp.Handler(node.Handler())
	_ = http.Server{
		Handler: handler,
	}
}
