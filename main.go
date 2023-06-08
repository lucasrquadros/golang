package main

import (
	"fmt"

	"github.com/lucasrquadros/golang/clientes"
	"github.com/lucasrquadros/golang/contas"
)

func main() {

	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "093.140.459-27",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	fmt.Println(contaDoBruno)

}
