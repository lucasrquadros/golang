package contas

import "github.com/lucasrquadros/golang/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado não foi realizado com sucesso"
	} else {
		return "saldo infuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Saque realizado realizado com sucesso", c.Saldo
	} else {
		return "saldo infuficiente", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(ValorDaTransferencia float64, ContaDestino *ContaCorrente) bool {
	if ValorDaTransferencia < c.Saldo && ValorDaTransferencia > 0 {
		c.Saldo -= ValorDaTransferencia
		ContaDestino.Depositar(ValorDaTransferencia)
		return true
	} else {
		return false
	}
}
