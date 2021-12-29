package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Sanjungliu/golang-startup/config"
	"github.com/Sanjungliu/golang-startup/internal/httpserver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var appsrvr *http.Server

	config := config.Init()

	log.Println("Environment:", config.Environment())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	db, err := gorm.Open(postgres.Open(config.DBConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	app := buildApp(ctx, config, db)

	// Application server
	go func(ctx context.Context) {
		appsrvr = httpserver.New(app)
		appsrvr.Addr = ":8080"
		log.Printf("About to listen on %s. Go to http://127.0.0.1%s", appsrvr.Addr, appsrvr.Addr)
		if e := appsrvr.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			log.Fatal("Application Server:", e)
		}
	}(ctx)
	<-quit
	shutdownServers(ctx, 60*time.Second, appsrvr)
	cancel()

}

func shutdownServers(ctx context.Context, timeout time.Duration, servers ...*http.Server) {
	wg := new(sync.WaitGroup)
	for _, srvr := range servers {
		if srvr == nil {
			continue
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, s *http.Server) {
			defer wg.Done()
			cto, cancel := context.WithTimeout(ctx, timeout)
			if e := s.Shutdown(cto); e != nil && e != http.ErrServerClosed {
				log.Printf("Shutdown failed for server in address: %s, %v", s.Addr, e)
			}
			cancel()
		}(wg, srvr)
	}
	wg.Wait()
}
