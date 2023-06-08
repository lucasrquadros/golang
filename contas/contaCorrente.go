package contas

import "github.com/lucasrquadros/golang/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado não foi realizado com sucesso"
	} else {
		return "saldo infuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Saque realizado realizado com sucesso", c.saldo
	} else {
		return "saldo infuficiente", c.saldo
	}
}

func (c *ContaCorrente) Transferir(ValorDaTransferencia float64, ContaDestino *ContaCorrente) bool {
	if ValorDaTransferencia < c.saldo && ValorDaTransferencia > 0 {
		c.saldo -= ValorDaTransferencia
		ContaDestino.Depositar(ValorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo

}
