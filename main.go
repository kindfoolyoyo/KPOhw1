package main

import (
	"GOhw1/app"
	"GOhw1/console"
	"os"
)

func main() {
	container := app.NewContainer()                  // DI: зависимости приложения собираются в одном контейнере.
	console.Run(container.Park, os.Stdin, os.Stdout) // DI: готовый Park передаётся консольному слою извне.
}
