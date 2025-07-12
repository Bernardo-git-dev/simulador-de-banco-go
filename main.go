package main

import "fmt"

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
