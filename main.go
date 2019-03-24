package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/GnuYtuce/ytuyemekhane-api/routers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := flag.String("port", os.Getenv("PORT"), "http service port")
	flag.Parse()
	if *port == "" {
		log.Fatalf("Please define the port number with the command line flags or the environment variable")
	}

	log.Printf("Go server listening on port %s", *port)

	r := httprouter.New()
	r.GET("/", routers.FoodListByCurrentTime)
	r.GET("/:year", routers.FoodListByCertainYear)
	r.GET("/:year/:month", routers.FoodListByCertainYearAndMonth)
	r.GET("/:year/:month/:day", routers.FoodListByCertainTime)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
