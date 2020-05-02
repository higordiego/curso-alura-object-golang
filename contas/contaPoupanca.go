package contas

import "github.com/higordiego/curso-alura-object-golang/clientes"

// ContaPoupanca - struct
type ContaPoupanca struct {
	TItular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}


// ObterSaldo - handler
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
// Depositar - handler
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!"
	}
	return "valor do deposito menor que 0."
}

// Sacar - handler
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	}
	return "Saldo insuficiente."
}

// Transferir - handler
func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) bool {
	if valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}