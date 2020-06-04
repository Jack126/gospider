package destroy

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	go func() {
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
	}()

}
