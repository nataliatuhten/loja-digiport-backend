package pessoa

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func ListaPessoa() []Pessoa {
	pessoas := []Pessoa{
		{Nome: "Nati", Idade: 21},
		{Nome: "Maria", Idade: 25},
	}
	return pessoas
}
