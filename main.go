package main

import (
	"fmt"
	"oop-golang/clientes"
	"oop-golang/contas"
)

// pesquisar o conceito de ponteiro
// pesquisas porque ponteiro em linguagens pode ser algo útil
type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {
	clienteThiago := clientes.Titular{
		Nome:      "thiago",
		CPF:       "123456789",
		Profissao: "Empresário",
	}
	contaDoThiago := contas.ContaCorrente{
		Titular:       clienteThiago,
		NumeroAgencia: 12346,
		NumeroConta:   321123,
	}
	PagarBoleto(&contaDoThiago, 1000.0)

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60.0)

	// contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}

	// exemplo de um uso de ponteiro
	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	// contaDaCris.Titular = "Cris"
	contaDaCris.NumeroAgencia = 321
	contaDaCris.NumeroConta = 12345
	contaDaCris.Depositar(322)

	fmt.Println(contaDoThiago)
	// fmt.Println(contaDaBruna)
}
