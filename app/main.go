package main


import (
    "fmt"
    "log"
    "net/http"
    "os"
    "io"
    )


func main() {

    http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
        fmt.Println("APP:", os.Getenv("APP"))
        appName := os.Getenv("APP")
        w.WriteHeader(http.Statusok)
        io.WriteString(w, "Inicio"+appName)
    })

    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("APP:", os.Getenv("APP"))
        appName := os.Getenv("APP")
        w.WriteHeader(http.Statusok)
        userName := r.URL.Query().Get("name")
        if userName != "" {
            io.WriteString(w, appName+":Username is " + userName)
        }
    })

    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatalf("Could not start the server: %s\n", err.Error())
    }
    log.Println("Server listening on port 8081")


}

