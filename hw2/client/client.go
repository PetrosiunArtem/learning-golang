package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const defaultHTTPTimeout = 30 * time.Second
const hardOpTimeout = 15 * time.Second

type DecodeResponse struct {
	OutputString string `json:"outputString"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: defaultHTTPTimeout,
		},
	}
}

func (client *Client) GetVersion() error {
	fmt.Println("GET /version")

	resp, err := client.httpClient.Get(client.baseURL + "/version")
	if err != nil {
		return fmt.Errorf("ошибка при запросе version: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка чтения ответа version: %v", err)
	}

	fmt.Printf("Статус: %d\n", resp.StatusCode)
	fmt.Printf("Тело ответа: %s\n", string(body))
	return nil
}

func (client *Client) Decode(inputString string) error {
	fmt.Println("POST /decode")

	requestData := map[string]string{
		"inputString": inputString,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга JSON: %v", err)
	}

	response, err := client.httpClient.Post(client.baseURL+"/decode", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ошибка при запросе decode: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("ошибка чтения ответа decode: %v", err)
	}

	fmt.Printf("Статус: %d\n", response.StatusCode)

	var decodeResp DecodeResponse
	if err := json.Unmarshal(body, &decodeResp); err != nil {
		log.Printf("Ошибка парсинга JSON ответа: %v", err)
		fmt.Printf("Тело ответа: %s\n", string(body))
	} else {
		fmt.Printf("Декодированная строка: %s\n", decodeResp.OutputString)
	}

	return nil
}

func (client *Client) HardOp() error {
	fmt.Println("GET /hard-op")

	ctx, cancel := context.WithTimeout(context.Background(), hardOpTimeout)
	defer cancel()

	requestWithContext, err := http.NewRequestWithContext(ctx, "GET", client.baseURL+"/hard-op", nil)
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %v", err)
	}

	startTime := time.Now()
	fmt.Printf("Запрос отправлен в: %s\n", startTime.Format("15:04:05"))

	response, err := client.httpClient.Do(requestWithContext)

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		endTime := time.Now()
		duration := endTime.Sub(startTime)

		fmt.Printf("Запрос отменен в: %s\n", endTime.Format("15:04:05"))
		fmt.Printf("Время выполнения до отмены: %.2f секунд\n", duration.Seconds())
		fmt.Println("Результат: Запрос отменен - превышено время ожидания (15 секунд)")
		return fmt.Errorf("таймаут запроса: операция заняла более %.0f секунд", hardOpTimeout.Seconds())
	}

	if err != nil {
		return fmt.Errorf("ошибка при запросе hard-op: %v", err)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("Запрос завершен в: %s\n", endTime.Format("15:04:05"))
	fmt.Printf("Время выполнения: %.2f секунд\n", duration.Seconds())
	fmt.Printf("Статус: %d\n", response.StatusCode)
	return nil
}

func (client *Client) RunAllMethods() error {
	fmt.Println("Запуск всех методов сервера...")
	if err := client.GetVersion(); err != nil {
		return fmt.Errorf("ошибка в GetVersion: %v", err)
	}
	if err := client.Decode("SGVsbG8sIFdvcmxkIQ=="); err != nil { // "Hello, World!" в base64
		return fmt.Errorf("ошибка в Decode: %v", err)
	}
	if err := client.HardOp(); err != nil {
		return fmt.Errorf("ошибка в HardOp: %v", err)
	}
	return nil
}
