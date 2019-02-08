package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	service = "SERVER"
)

var (
	port = flag.Int("port", 8080, "http port to listen on")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("{'message':'hello from server', 'timestamp':'%s'}\n", time.Now().String())

		fmt.Print(msg)

		w.WriteHeader(200)
		w.Write([]byte(msg))
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("%s listening on 0.0.0.0:%d", service, *port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Server closed: ", err)
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stop

	log.Printf("%s shutting down ...\n", service)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Error while shutting down the server !")
		log.Fatal(err)
	}

	log.Printf("%s down\n", service)
}
