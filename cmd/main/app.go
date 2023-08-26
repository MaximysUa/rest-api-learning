package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"rest-api-learning/internal/user"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("create handler")
	handler := user.NewHandler()
	handler.Register(router)

	log.Println("start")
	start(router)
}

func start(router *httprouter.Router) {
	//router.GET("/api/:name", IndexHandler)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.Serve(listener))
}
