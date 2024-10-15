package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Ls: Lista arquivos no diretório atual
func Ls(currentDir string) {
	cmd := exec.Command("ls")
	cmd.Dir = currentDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
	}
}

// Echo: Exibe uma mensagem no terminal
func Echo(args []string) {
	message := strings.Join(args, " ")
	fmt.Println(message)
}

// Mkdir: Cria um diretório no caminho atual
func Mkdir(currentDir string, args []string) {
	if len(args) < 1 {
		fmt.Println("Por favor, forneça o nome do diretório.")
		return
	}
	dirPath := filepath.Join(currentDir, args[0])
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Println("Erro ao criar diretório:", err)
	} else {
		fmt.Println("Diretório criado:", dirPath)
	}
}

// Cd: Muda o diretório atual
func Cd(currentDir *string, args []string) {
	if len(args) < 1 {
		fmt.Println("Por favor, forneça o caminho do diretório.")
		return
	}
	newDir := filepath.Join(*currentDir, args[0])
	err := os.Chdir(newDir)
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
	} else {
		*currentDir = newDir
	}
}

// Rm: Remove um arquivo ou diretório
func Rm(currentDir string, args []string) {
	if len(args) < 1 {
		fmt.Println("Por favor, forneça o nome do arquivo ou diretório.")
		return
	}
	filePath := filepath.Join(currentDir, args[0])
	err := os.RemoveAll(filePath)
	if err != nil {
		fmt.Println("Erro ao remover arquivo ou diretório:", err)
	} else {
		fmt.Println("Arquivo ou diretório removido:", filePath)
	}
}

// Cat: Exibe o conteúdo de um arquivo
func Cat(currentDir string, args []string) {
	if len(args) < 1 {
		fmt.Println("Por favor, forneça o nome do arquivo.")
		return
	}
	filePath := filepath.Join(currentDir, args[0])
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	} else {
		fmt.Println(string(content))
	}
}

// Clear: Limpa a tela do terminal
func Clear() {
	fmt.Print("\033[H\033[2J")
}

// Sysinfo: Exibe informações do sistema
func Sysinfo() {
	user, _ := os.UserHomeDir()
	hostname, _ := os.Hostname()
	fmt.Println("Usuário:", user)
	fmt.Println("Hostname:", hostname)
	fmt.Println("Sistema Operacional:", runtime.GOOS)
	fmt.Println("Arquitetura:", runtime.GOARCH)
}

// Date: Exibe a data e hora atuais
func Date() {
	currentTime := time.Now()
	fmt.Println("Data e Hora:", currentTime.Format("02/01/2006 15:04:05"))
}

// History: Exibe o histórico de comandos
func History(history []string) {
	for i, cmd := range history {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}

// Help: Exibe os comandos disponíveis
func Help() {
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
}

// Touch: Cria um novo arquivo vazio
func Touch(currentDir string, args []string) {
	if len(args) < 1 {
		fmt.Println("Por favor, forneça o nome do arquivo.")
		return
	}
	filePath := filepath.Join(currentDir, args[0])
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
	} else {
		fmt.Println("Arquivo criado:", filePath)
	}
	file.Close()
}
