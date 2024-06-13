package api

import (
	"log"
	"net/http"
	"sync"
	conf "user_server/configs"
	"user_server/pkg"
)

func Init(conf conf.ServerConfig, globsig chan bool, wg *sync.WaitGroup, onStartWg *sync.WaitGroup) {
	defer wg.Done()
	innerch := make(chan bool)

	log.Printf("http server started")
	router := NewRouter()
	server := &http.Server{
		Addr:    conf.Addr,
		Handler: router,
	}

	go pkg.Dog(globsig, innerch, server, "server")
	onStartWg.Done()

	log.Println("http server started on :8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	close(innerch)
}
