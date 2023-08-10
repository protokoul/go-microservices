package main

import "log"
import "fmt"
import "net/http"

const webPort = "80"

type Config struct {

}

func main(){
	app := Config{}

	log.Printf("Starting broker service on %s\n", webPort)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}