package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chayutK/skill-kafka-api/repository/database"
	routes "github.com/chayutK/skill-kafka-api/router"
	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	DB = database.Sync()
	r := gin.Default()

	r.GET("/", HelloHandler)

	v := r.Group("/api/v1")
	routes.InitSkillRouter(v, DB)

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	cancelChanel := make(chan struct{})
	go func() {
		<-ctx.Done()
		fmt.Println("Server is shutting down.........")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Server closed :", err.Error())
			}
		}

		close(cancelChanel)
	}()
	if err := srv.ListenAndServe(); err != nil {
		log.Println("Error while starting server.", err.Error())
	}
	<-cancelChanel
	fmt.Println("bye")

}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, World",
	})
	// return
}
