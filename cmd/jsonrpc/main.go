package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alessandromr/pharmacy/datalayer/memory"
	"github.com/alessandromr/pharmacy/internal/rpcserver"
	"github.com/alessandromr/pharmacy/script"
)

func main() {
	pharmaciesDL := memory.PharmaciesMemory{}

	go script.SyncData(&pharmaciesDL)
	time.Sleep(time.Second * 2)

	r := rpcserver.NewJsonRpcServer()
	log.Fatal(http.ListenAndServe(":8080", r))
}
