package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	mux := app.routes()

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", mux)
}
