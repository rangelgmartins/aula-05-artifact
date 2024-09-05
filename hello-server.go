package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    
    // Adicionando verificação de erro na inicialização do servidor
    err := http.ListenAndServe(":13000", nil)
    if err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    // Adicionando uma verificação simples para um path vazio
    if r.URL.Path == "/" {
        fmt.Fprintf(w, "Hello, World!")
    } else {
        fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    }
}
