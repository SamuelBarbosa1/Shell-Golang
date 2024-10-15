package utils

import (
	"fmt"
	"strings"
)

// CheckError: Função para lidar com erros de forma centralizada
func CheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

// JoinArgs: Junta os argumentos a partir de uma posição específica
func JoinArgs(args []string, start int) string {
	return strings.Join(args[start:], " ")
}
