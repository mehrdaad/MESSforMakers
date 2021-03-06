package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/makeict/MESSforMakers/util"
)

//Main reads the configuration immediately and dies if it can't be read.
//There is no default configuration and no command line flags.
//All options are contained in the config.json file
func main() {
	//Read the configuration and die if it can't be read.
	//This does NOT guarantee that sensible options have been set, only that the file can be read.
	config, err := util.InitConfig("config.json")
	if err != nil {
		//No log files have been set up yet, so just dump the error to stdout
		fmt.Print("Cannot parse the configuration file")
		panic(1)
	}

	app, err := newApplication(config)
	if err != nil {
		fmt.Printf("Could not create the application: %v", err)
	}
	defer app.Logger.Close()

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Addr:         fmt.Sprintf("%s:%d", config.App.Host, config.App.Port),
		ErrorLog:     app.Logger.Logger,
		Handler:      app.Router,
	}

	app.Logger.Println("Starting Application on :" + strconv.Itoa(app.port))
	app.Logger.Fatal(srv.ListenAndServe())

}
