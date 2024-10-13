package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentDir, _ := os.Getwd() // Diretório inicial
	var history []string        // Histórico de comandos

	fmt.Println("Terminal do samuel em Golang - Versão 3.0")
	fmt.Println("Digite 'sair' para fechar o terminal.")
	fmt.Println("Digite 'help' para ver os comandos disponíveis.")

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
		args := strings.Split(input, " ") // Divide a entrada para obter o comando e os argumentos

		// Adiciona o comando ao histórico
		history = append(history, input)

		// Processa comandos
		switch args[0] {
		case "sair":
			fmt.Println("Saindo do Golang Terminal...")
			return

		case "ls":
			cmd := exec.Command("ls")
			cmd.Dir = currentDir
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Erro ao listar arquivos:", err)
			}

		case "echo":
			message := strings.Join(args[1:], " ")
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

		case "cat":
			if len(args) < 2 {
				fmt.Println("Por favor, forneça o nome do arquivo.")
				continue
			}
			filePath := filepath.Join(currentDir, args[1])
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Erro ao ler o arquivo:", err)
			} else {
				fmt.Println(string(content))
			}

		case "clear":
			fmt.Print("\033[H\033[2J") // Código de escape ANSI para limpar a tela

		case "sysinfo":
			user, _ := os.UserHomeDir()
			hostname, _ := os.Hostname()
			fmt.Println("Usuário:", user)
			fmt.Println("Hostname:", hostname)
			fmt.Println("Sistema Operacional:", runtime.GOOS)
			fmt.Println("Arquitetura:", runtime.GOARCH)

		case "date":
			currentTime := time.Now()
			fmt.Println("Data e Hora:", currentTime.Format("02/01/2006 15:04:05"))

		case "history":
			for i, cmd := range history {
				fmt.Printf("%d: %s\n", i+1, cmd)
			}

		case "help":
			fmt.Println("Comandos disponíveis:")
			fmt.Println("  ls           - Lista arquivos no diretório atual")
			fmt.Println("  echo [msg]   - Exibe a mensagem")
			fmt.Println("  mkdir [nome] - Cria um novo diretório")
			fmt.Println("  cd [path]    - Muda o diretório atual")
			fmt.Println("  rm [nome]    - Remove arquivo ou diretório")
			fmt.Println("  cat [nome]   - Exibe o conteúdo de um arquivo")
			fmt.Println("  clear        - Limpa a tela")
			fmt.Println("  sysinfo      - Exibe informações do sistema")
			fmt.Println("  date         - Exibe a data e hora atuais")
			fmt.Println("  history      - Exibe o histórico de comandos")
			fmt.Println("  touch [nome] - Cria um novo arquivo vazio")
			fmt.Println("  sair         - Fecha o terminal")

		case "touch":
			if len(args) < 2 {
				fmt.Println("Por favor, forneça o nome do arquivo.")
				continue
			}
			filePath := filepath.Join(currentDir, args[1])
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Erro ao criar o arquivo:", err)
			} else {
				fmt.Println("Arquivo criado:", filePath)
			}
			file.Close()

		default:
			fmt.Println("Comando não reconhecido:", args[0])
		}
	}
}
