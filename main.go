package main

import "fmt"

// Struct da conta
type Conta struct {
	Titular string
	saldo   float64
}

// Métodos
func (c *Conta) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
	}
}

func (c *Conta) Sacar(valor float64) bool {
	if valor <= c.saldo {
		c.saldo -= valor
		return true
	}
	return false
}

func (c Conta) VerSaldo() float64 {
	return c.saldo
}

func main() {
	conta := Conta{Titular: "Bernardo"}
	conta.Depositar(1000)
	fmt.Println("Saldo:", conta.VerSaldo())

	if conta.Sacar(300) {
		fmt.Println("Saque realizado.")
	} else {
		fmt.Println("Saldo insuficiente.")
	}

	fmt.Println("Saldo final:", conta.VerSaldo())
}
