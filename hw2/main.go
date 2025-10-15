package main

import (
	"fmt"
	"log"
	"time"

	"github.com/PetrosiunArtem/learning-golang/hw2/client"
	"github.com/PetrosiunArtem/learning-golang/hw2/server"
)

const serverBootTime = 3 * time.Second

func main() {
	go server.Run()

	fmt.Println("Ожидание запуска сервера...")
	time.Sleep(serverBootTime)

	port := server.GetPort()
	cli := client.New("http://localhost" + port)
	if err := cli.RunAllMethods(); err != nil {
		log.Fatalf("Ошибка при выполнении запросов: %v", err)
	}

	fmt.Println("\nКлиент завершил работу.")
}
