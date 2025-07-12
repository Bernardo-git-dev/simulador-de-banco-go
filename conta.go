package main

type Conta struct {
  Titular string
  saldo   float64
}

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
