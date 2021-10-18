package cmd

import (
	"fmt"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/business"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/factory"
	"github.com/alibaraneser/go-link-shortener-hexagonal/pkg/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Execute() {
	repo := factory.Repository()
	service := business.NewRedirectService(repo)
	handler := api.NewHandler(service)

	http.HandleFunc("/", handler.Get)
	http.HandleFunc("/save", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Printf("Listening on port %s \n", httpPort())
		errs <- http.ListenAndServe(httpPort(), nil)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated Process %s", <-errs)
}

func httpPort() string {
	port := "8080"
	if os.Getenv("HTTP_PORT") != "" {
		port = os.Getenv("HTTP_PORT")
	}
	return fmt.Sprintf(":%s", port)
}
