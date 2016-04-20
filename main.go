package main

import (
	"flag"
	"net/http"
	log "github.com/Sirupsen/logrus"

	"github.com/ant0ine/go-json-rest/rest"
)

var (
	configFile = flag.String("conf", "config.toml", "path to toml config")
)

func main() {
	flag.Parse()
	err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Could not load config file: %s\n", *configFile)
	}

	log.Println("Starting Go React Boilerplate Server")

	// serve website
	go fileServer(config.WebRoot, config.FileServerPort)

	// start the api
	apiServer(config.APIPort)
}

// fileServer serves static files from the react client's dist folder
func fileServer(path string, addr string) {
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)

	log.Println("File Server Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// apiServer sets up API middleware, routes, and starts the API server.
func apiServer(addr string) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	// cross origin middleware allows us to use different ports FileServer and API
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true // origin == "http://127.0.0.1"
		},
		AllowedMethods:                []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:                []string{"Accept", "Content-Type", "Authorization", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	// route configurations
	router, err := rest.MakeRouter(
		rest.Get("/apps", getApps),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	log.Printf("Api Server Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, api.MakeHandler()))
}


// Route Handlers

// application represents a web application
type application struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

// apps is the list of applications we <3 the most
var apps = []application{
	application{Name:"app1", URL:"http://app1.com"},
	application{Name:"app2", URL:"http://app2.com"},
	application{Name:"app3", URL:"http://app3.com"},
}

// getApps handler returns the list of applications we <3
func getApps(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(apps)
}
