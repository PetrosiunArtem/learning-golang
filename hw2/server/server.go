package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const apiVersion string = "v1.2.3"
const serverShutdownTimeout = 10 * time.Second
const port = ":8080"

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /version", versionHandler)
	mux.HandleFunc("POST /decode", decodeHandler)
	mux.HandleFunc("GET /hard-op", hardOpHandler)

	server := &http.Server{Addr: port, Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	<-osSignals

	log.Println("Graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited properly")
}

func GetPort() string {
	return port
}

func versionHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(apiVersion))
}

func hardOpHandler(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(time.Duration(rand.Intn(11)+10) * time.Second)
	switch rand.Intn(2) {
	case 0:
		writer.WriteHeader(200)
	case 1:
		writer.WriteHeader(500)
	}
}

func decodeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var data map[string]string
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}

	output, err := decodeToString(data["inputString"])
	if err != nil {
		http.Error(writer, "Decode failed", http.StatusBadRequest)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{"outputString": output})
}

func decodeToString(inputString string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(inputString)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
