package main

import (
	"shell-terminal-go/core"
)

func main() {
	shell := core.NewShell() // Inicializa o shell
	shell.Run()              // Executa o shell
}
