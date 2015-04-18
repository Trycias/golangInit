package main

import (
	"log"
	"net/http"
"os"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Meteo struct{}

func (o Meteo) GetRainInfo() int {
	//devrait retourner une valeur de 0 à 10
	return 5
} 

func main() {

//	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", StaticHandler)

/* inutile, c'est juste pour tester
	m := Meteo{}
	log.Println(m.GetRainInfo())
*/
	//init gorilla rpc
	pRpcServer :=rpc.NewServer()
	pRpcServer.RegisterCodec(json.NewCodec(), "application/json")
	pRpcServer.RegisterService(new(Meteo), "")
	http.Handle("/rpc",pRpcServer)



	log.Println("Server started")
	if e := http.ListenAndServe(":8080", nil); e != nil {
		log.Fatal(e)
	}
}


func StaticHandler(w http.ResponseWriter, pRequest *http.Request){
	url := pRequest.URL.String()
	if url == "/" {
		url = "/index.html"
	}
	/* à la place du e, qui recupére l'erreur, 
	 * on peut mettre _ pour ne pas recevoir / gérer l'erreur
	 */
	currentPath, _ := os.Getwd()
	filePath := currentPath + url
	http.ServeFile(w, pRequest, filePath)
}