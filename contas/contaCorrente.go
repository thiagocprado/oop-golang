package contas

import "oop-golang/clientes"

// struct - mesmo propósito de uma classe porém não temos getters e setters
type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!", c.saldo
	} else {
		return "Saldo insuficiente!", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0

	if podeDepositar {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Saldo insuficiente!", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia <= c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)

		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObTerSaldo() float64 {
	return c.saldo
}
