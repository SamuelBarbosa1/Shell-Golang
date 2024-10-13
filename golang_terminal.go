package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentDir, _ := os.Getwd() // Obtém o diretório atual

	fmt.Println("Terminal do Samuel usango Golang")
	fmt.Println("Digite 'sair' para fechar o terminal.")
	fmt.Println("Comandos disponíveis: ls, echo, mkdir, cd, rm")

	for {
		// Exibe o prompt com o diretório atual
		fmt.Printf("gt:%s> ", currentDir)

		// Lê a entrada do usuário
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler o comando:", err)
			continue
		}

		// Remove espaços em branco e quebras de linha
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ") // Divide a entrada por espaço para obter o comando e os argumentos

		// Processa comandos
		switch args[0] {
		case "sair":
			fmt.Println("Saindo do Golang Terminal...")
			return

		case "ls":
			cmd := exec.Command("ls")
			cmd.Dir = currentDir // Define o diretório atual para o comando
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Erro ao listar arquivos:", err)
			}

		case "echo":
			message := strings.Join(args[1:], " ") // Junta os argumentos como mensagem
			fmt.Println(message)

		case "mkdir":
			if len(args) < 2 {
				fmt.Println("Por favor, forneça o nome do diretório.")
				continue
			}
			dirPath := filepath.Join(currentDir, args[1])
			err := os.Mkdir(dirPath, 0755)
			if err != nil {
				fmt.Println("Erro ao criar diretório:", err)
			} else {
				fmt.Println("Diretório criado:", dirPath)
			}

		case "cd":
			if len(args) < 2 {
				fmt.Println("Por favor, forneça o caminho do diretório.")
				continue
			}
			newDir := filepath.Join(currentDir, args[1])
			err := os.Chdir(newDir)
			if err != nil {
				fmt.Println("Erro ao mudar de diretório:", err)
			} else {
				currentDir = newDir // Atualiza o diretório atual
			}

		case "rm":
			if len(args) < 2 {
				fmt.Println("Por favor, forneça o nome do arquivo ou diretório.")
				continue
			}
			filePath := filepath.Join(currentDir, args[1])
			err := os.RemoveAll(filePath)
			if err != nil {
				fmt.Println("Erro ao remover arquivo ou diretório:", err)
			} else {
				fmt.Println("Arquivo ou diretório removido:", filePath)
			}

		default:
			fmt.Println("Comando não reconhecido:", args[0])
		}
	}
}
