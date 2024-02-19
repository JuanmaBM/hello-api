package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/hello", helloHandler).Methods("GET")
    r.HandleFunc("/hello/{name}", helloNameHandler).Methods("GET")
    r.HandleFunc("/health", healthHandler).Methods("GET")

    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Hello World!")
}

func helloNameHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    message := fmt.Sprintf("Hello %s!", name)

    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintln(w, message)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"status": "ok"}

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    w.Write(jsonResponse)
}

