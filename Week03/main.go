package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	Addr = "localhost"
	Port = 9999
)

func main() {
	eg := errgroup.Group{}

	sigChan := make(chan os.Signal)
	errorChan := make(chan error)

	s := http.Server{Addr: fmt.Sprintf("%s:%d", Addr, Port)}

	eg.Go(func() error {
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-sigChan
		return s.Shutdown(context.TODO())
	})

	eg.Go(func() error {
		go func() {
			errorChan <- s.ListenAndServe()
		}()
		select {
		case err := <-errorChan:
			close(sigChan)
			close(errorChan)
			return err
		}
	})

	if err := eg.Wait(); err != nil {
		log.Printf("%s", err)
	}
}
