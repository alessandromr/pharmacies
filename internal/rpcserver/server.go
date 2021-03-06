package rpcserver

import (
	"github.com/alessandromr/pharmacy/presentation/jsonrpc"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func NewJsonRpcServer() *mux.Router {
	r := mux.NewRouter()

	//Create RPC Server
	jsonRPC := rpc.NewServer()
	jsonCodec := json.NewCodec()
	jsonRPC.RegisterCodec(jsonCodec, "application/json")
	jsonRPC.RegisterCodec(jsonCodec, "application/json; charset=UTF-8")

	jsonRPC.RegisterService(new(jsonrpc.Pharmacy), "")
	r.Handle("/api", jsonRPC)
	return r
}
