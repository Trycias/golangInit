package main

import (
	 "github.com/gorilla/rpc"
	 "github.com/gorilla/rpc/json"
	"log"
	"net/http"
	"os"
)

type Meteo struct{}
/*pour gorilla*/
type NoArgs struct{}
/*pour gorilla*/
type Reply struct{
	Result int
}

func (o *Meteo) GetRainInfo(r *http.Request, args *NoArgs,  reply *Reply) error {
	// devrait retourner une valeur entre 0 et 10
	reply.Result = 5
	return nil
}

func StaticHandler(w http.ResponseWriter, pRequest *http.Request) {
	url := pRequest.URL.String()

	if url == "/" {
		url = "/index.html"
	}

	currentPath, _ := os.Getwd()

	filePath := currentPath + url

	http.ServeFile(w, pRequest, filePath)
}

func main() {

		http.HandleFunc("/", StaticHandler)

	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	//	http.Handle("/", http.FileServer(http.Dir(".")))

	// Init gorilla RPC
	 pRpcServer := rpc.NewServer()
	 pRpcServer.RegisterCodec(json.NewCodec(), "application/json")
	 pRpcServer.RegisterService(new(Meteo), "")
	 http.Handle("/rpc", pRpcServer)

	 log.Println("Server started")

	 if e := http.ListenAndServe(":8080", nil); e != nil {
	 	log.Fatal(e)
	 }
}
