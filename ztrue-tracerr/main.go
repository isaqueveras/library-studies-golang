package main

import (
	"log"

	a "github.com/isaqueveras/error/pessoa"
	"github.com/ztrue/tracerr"
)

var (
	mensagem string
	erro     error
)

func main() {
	// Falando
	if mensagem, erro = a.Falar("Olá mundo"); erro != nil {
		tracerr.PrintSourceColor(erro)
		return
	}

	log.Print(mensagem)

	// Comendo
	if mensagem, erro = a.Comer("Miojo"); erro != nil {
		tracerr.PrintSourceColor(erro)
		return
	}

	log.Println(mensagem)
}
