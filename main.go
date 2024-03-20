package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"task/floodcontrol"
	"time"
)

// Config подтягивает переменные для тестов из окружения докера
func Config() (int, int) {
	sendRPS, _ := strconv.Atoi(os.Getenv("SEND_RPS"))
	userR, _ := strconv.Atoi(os.Getenv("USER_R"))
	return 1000 / sendRPS, userR
}

func main() {
	fc := floodcontrol.NewFC(context.Background())
	sendRPS, userR := Config()
	for i := 0; i < 100; i++ {
		for j := 0; j < userR; j++ {
			res, err := fc.Check(context.Background(), int64(i))
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(time.Now(), " ID: ", i, " ", res)
			time.Sleep(time.Millisecond * time.Duration(sendRPS))
		}
		fmt.Print("\n")
	}
}
