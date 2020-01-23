package main

import (
	"Effective/models"
	"Effective/pkg/logging"
	"Effective/pkg/setting"
	"Effective/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       setting.ServerSetting.ReadTimeOut,
		ReadHeaderTimeout: setting.ServerSetting.ReadHeaderTimeOut,
		WriteTimeout:      setting.ServerSetting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown:%s", err)
	}

	log.Println("Server exiting")

	s.ListenAndServe()
}
