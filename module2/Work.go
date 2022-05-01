package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func init() {
	err := os.Setenv("VERSION", "V1.0")
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", rootHandler)
	mux.HandleFunc("/healthz", health)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	t := time.Now()
	remoteAddr := req.RemoteAddr
	fmt.Printf("访问时间：%4d-%02d-%02d %02d:%02d:%02d  访问IP：%s\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), remoteAddr)

	username := req.FormValue("username")
	password := req.FormValue("password")

	resHeader := res.Header()
	version := os.Getenv("VERSION")
	resHeader.Add("username", username)
	resHeader.Add("password", password)
	resHeader.Add("version", version)
	fmt.Fprintf(res, "Hello World!")
}

func health(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	fmt.Fprintf(res, "200")
}
