package main

import (
	"contas"
	"fmt"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}
	status := contaDoGustavo.Transferir(200, &contaDaSilvia)

	fmt.Println(status)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	// contaDaSilvia := contaCorrente{}
	// contaDaSilvia.titular = "Silvia"
	// contaDaSilvia.saldo = 500

	// fmt.Println(contaDaSilvia.saldo)
	// status, valor := contaDaSilvia.Depositar(2000)
	// fmt.Println(status, valor)

}
