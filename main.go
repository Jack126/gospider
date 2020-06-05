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

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/**** pringln spider ****/
	content, _ := ioutil.ReadFile("./static/spider.txt")
	log.Println(string(content))

	/**** route start ****/
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	/**** route end ****/

	/**** graceful shutdown ****/
	srv := &http.Server{
		Addr:    ":8082",
		Handler: r,
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
