package main

import "log"

// ContaCorrente - struct
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}


func (c *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	}
	return "Saldo insuficiente."
}

func main() {	

	contaDaSilva :=  ContaCorrente{}
	contaDaSilva.titular = "Silva"
	contaDaSilva.saldo = 500

	

	response := contaDaSilva.sacar(-200.)

	log.Println(response)
	log.Println(contaDaSilva.saldo)
	
}
