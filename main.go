package main

import (
	"test-task-intech/app"
	config "test-task-intech/config"
)

func main() {
	println("Запуск сервера...")
	config := config.LoadConfigDB()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}