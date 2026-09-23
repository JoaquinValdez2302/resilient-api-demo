package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time" 
)

func main() {
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "Servidor_Local"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" { return }
		mensaje := fmt.Sprintf("✅ Respuesta generada por: %s\n", serverName)
		fmt.Print(mensaje)
		fmt.Fprint(w, mensaje)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("💥 ¡ALERTA! %s simulando un fallo crítico y apagándose...\n", serverName)
		os.Exit(1)
	})

	port := "8080"
	
	// SIMULAMOS UN ARRANQUE LENTO (5 segundos) PARA QUE HAPROXY DETECTE LA CAÍDA
	fmt.Printf("⏳ %s reiniciando... simulando carga de base de datos (5s)...\n", serverName)
	time.Sleep(5 * time.Second)

	fmt.Printf("🚀 %s levantado y escuchando en el puerto %s...\n", serverName, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s\n", err)
	}
}