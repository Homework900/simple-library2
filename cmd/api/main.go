package main

import (
    "encoding/json"
    "net/http"
    "time"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    resp := struct {
        Status      string `json:"status"`
        ProjectName string `json:"project_name"`
    }{Status: "available", ProjectName: "Library API"}

    data, err := json.Marshal(resp)
    if err != nil {
        http.Error(w, "Internal Server Error", 500)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
    resp := struct {
        Version    string `json:"version"`
        ServerTime string `json:"server_time"`
    }{Version: "1.0.0", ServerTime: time.Now().Format(time.RFC3339)}

    data, _ := json.Marshal(resp)
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("This page does not exist"))
}

func main() {
    http.HandleFunc("/healthcheck", healthCheckHandler)
    http.HandleFunc("/api/v1/info", infoHandler)
    http.HandleFunc("/notfound", notFoundHandler)
    http.ListenAndServe(":8080", nil)
}