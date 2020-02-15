package rpcserver

import (
	"net/http"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
	"github.com/alessandromr/pharmacy/presentation/jsonrpc"
	"github.com/gorilla/mux"
)


func NewJsonRpcServer() *mux.Router{
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


func test (w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Test"))
}
