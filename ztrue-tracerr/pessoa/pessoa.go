package pessoa

// Falar uma pessoa começa a falar
func Falar(palavra string) (string, error) {
	// return "A pessoa está falando " + palavra, tracerr.New("Aconteceu um erro enquanto a pessoa estava falando")
	return "A pessoa está falando: " + palavra, nil
}

// Comer uma pessoa começa a comer
func Comer(comida string) (string, error) {
	// return "A pessoa está comendo " + comida, tracerr.New("Aconteceu um erro enquanto a pessoa estava comendo")
	return "A pessoa está comendo " + comida, nil
}
