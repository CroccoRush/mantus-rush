package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"user_server/api"
	conf "user_server/configs"
)

var (
	STAGE        string
	Gwg          sync.WaitGroup
	onStartWg    sync.WaitGroup
	globsig      chan os.Signal
	interruption chan bool
	serverConf   conf.ServerConfig
)

func init() {
	STAGE = "development"
	Gwg.Add(30)
	if err := os.Setenv("SERV", ":8000"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("NAME_VER", "militaris"); err != nil {
		log.Fatal(err)
	}
	globsig = make(chan os.Signal, 1)
	interruption = make(chan bool)
	signal.Notify(globsig, os.Interrupt)
	serverConf.Read("SERV")
}

func main() {
	log.Println("server starting...")
	onStartWg.Add(1)
	go api.Init(serverConf, interruption, &Gwg, &onStartWg)
	onStartWg.Wait()

watchdog:
	for {
		select {
		case <-globsig:
			log.Println("...interrupting")
			close(globsig)
			close(interruption)
			onStartWg.Wait()
			Gwg.Wait()
			log.Println("server stopped...")
			break watchdog
		}
	}
}
