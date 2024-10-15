package core

import (
	"bufio"
	"fmt"
	"os"
	"shell-terminal-go/commands"
	"strings"
)

type Shell struct {
	history []string
}

func NewShell() *Shell {
	return &Shell{history: []string{}}
}

func (s *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)
	currentDir, _ := os.Getwd()

	fmt.Println("Terminal do Samuel em Golang - Versão 3.1")
	fmt.Println("Digite 'sair' para fechar o terminal.")
	fmt.Println("Digite 'help' para ver os comandos disponíveis.")

	for {
		fmt.Printf("gt:%s> ", currentDir)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		s.history = append(s.history, input)

		switch args[0] {
		case "sair":
			fmt.Println("Saindo do Golang Terminal...")
			return
		case "ls":
			commands.Ls(currentDir)
		case "echo":
			commands.Echo(args[1:])
		case "mkdir":
			commands.Mkdir(currentDir, args[1:])
		case "cd":
			commands.Cd(&currentDir, args[1:])
		case "rm":
			commands.Rm(currentDir, args[1:])
		case "cat":
			commands.Cat(currentDir, args[1:])
		case "clear":
			commands.Clear()
		case "sysinfo":
			commands.Sysinfo()
		case "date":
			commands.Date()
		case "history":
			commands.History(s.history)
		case "help":
			commands.Help()
		case "touch":
			commands.Touch(currentDir, args[1:])
		default:
			fmt.Println("Comando não reconhecido:", args[0])
		}
	}
}
