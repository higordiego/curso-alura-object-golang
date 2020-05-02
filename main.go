package main

import (
	"log"

	"github.com/higordiego/curso-alura-object-golang/clientes"
	"github.com/higordiego/curso-alura-object-golang/contas"
)


// PagarBoleto - handler
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}


func main() {

	clientBruno := clientes.Titular{Nome: "Bruno", CPF:"03506838326", Profissao:"Analista"}
	contaDoBruno := contas.ContaCorrente{Titular: clientBruno, NumeroAgencia: 123, NumeroConta: 20}

	PagarBoleto(&contaDoBruno, 60)

	contaLuisa := contas.ContaPoupanca{}
	contaLuisa.Depositar(500)
	PagarBoleto(&contaLuisa, 400)
	log.Println(contaLuisa.ObterSaldo())
}
