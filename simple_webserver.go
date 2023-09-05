package main
import(
	"fmt"
	"net/http"
	"log"
)
func main(){
	handler:=func(w http.ResponseWriter,r *http.Request){
		fmt.Fprint(w,"Hello world!..")
	}
	mux:=http.NewServeMux()
    mux.HandleFunc("/",handler)
	var port string
	fmt.Println("enter the port number (eg:8080)")
	fmt.Scan(&port)
	server:=&http.Server{
		Addr:":"+ port,
		Handler:mux,
	}
	fmt.Printf("server listing at port %s ....",port)
	log.Fatal(server.ListenAndServe())
}