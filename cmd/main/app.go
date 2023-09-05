package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	book "rest-api-learning/internal/book/db"
	"rest-api-learning/internal/config"
	"rest-api-learning/internal/user"
	"rest-api-learning/pkg/client/postgresql"
	"rest-api-learning/pkg/logging"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()
	newClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository := book.NewRepository(newClient, logger)
	all, err := repository.FindAll(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}
	for _, i := range all {
		logger.Debug(i.Name)
	}

	logger.Info("create handler")

	handler := user.NewHandler(logger)
	handler.Register(router)

	logger.Info("start")
	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()

	var listener net.Listener
	var listenErr error
	// запуск программы на сокете
	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)

		//запуск программы на порту
	} else {
		logger.Info("listen tcp socket")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	}
	if listenErr != nil {
		logger.Fatal(listenErr)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
