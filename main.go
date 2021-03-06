package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "gospider/bootstrap"
	"gospider/global/config"
	"gospider/router"
)

func main() {
	/**** router ****/
	router := router.InitRouter()

	/**** pringln spider logo ****/
	content, _ := ioutil.ReadFile("./static/spider.txt")
	log.Println(string(content))

	/**** graceful shutdown ****/
	srv := &http.Server{
		Addr:    config.CreateSettingYmlFactory().GetString("config.application.host") + ":" + config.CreateSettingYmlFactory().GetString("config.application.port"),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Spider server will be shutting down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting..")
	/**** graceful shutdown ****/
}
