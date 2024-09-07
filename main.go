package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

// Handler que maneja las solicitudes y responde con el nombre del servidor
func serverNameHandler(w http.ResponseWriter, r *http.Request) {
    serverName, err := os.Hostname()
    if err != nil {
        http.Error(w, "No se pudo obtener el nombre del servidor", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "v2.3 - Nombre del Servidor: %s \n", serverName)
}

func main() {
    http.HandleFunc("/", serverNameHandler)

    // Especifica el puerto en el que el servidor escuchará
    port := "1683"
    fmt.Printf("Servidor corriendo en el puerto %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
