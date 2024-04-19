package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/handler"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	//http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/api/product", handler.ProductHandler)
	http.HandleFunc("/api/products", handler.ProductsHandler)
	http.HandleFunc("/api/mission", handler.MissionHandler)
	http.HandleFunc("/api/missions", handler.MissionsHandler)
	http.HandleFunc("/api/finance_logs", handler.FinanceLogsHandler)
	http.HandleFunc("/api/finance_log", handler.FinanceLogHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
