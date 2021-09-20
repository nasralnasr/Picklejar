package main

import (
  "context"

  "picklejar/router"
  "picklejar/database"
  "net/http"
  "time"
  "os"
  "os/signal"
  "log"
)

const (
	AccessKeyId	= "AKIAVOZFKTVU75M4G3YB"
	SecretAccessKey	= "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"
	Region		= "us-east-2"
)

func main(){

    routerInit := router.InitRouter()
    database.ConnectDatabase()

    server := &http.Server{
      Addr:           ":8080",
      Handler:        routerInit,
      ReadTimeout:    10 * time.Second,
      WriteTimeout:   10 * time.Second,
      MaxHeaderBytes: 1 << 20,
    }

    go func() {
      server.ListenAndServe()
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
      log.Fatal("Server Shutdown:", err)
    }
      log.Println("Server exiting")
}
