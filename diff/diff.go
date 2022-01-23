package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := generator.Int()
	p := generator.Int()
	fmt.Println("Shared numbers:", g, p, "\n")

	SecretNumberAlice := generator.Int()
	ExponentiationAlice := g ^ SecretNumberAlice%p

	fmt.Println("Secret number Alice : ", SecretNumberAlice, "exponentiation Alice:", ExponentiationAlice, "\n")

	SecretNumberBob := generator.Int()
	ExponentiationBob := g ^ SecretNumberBob%p
	fmt.Println("Secret number Bob : ", SecretNumberBob, "exponentiation Bob:", ExponentiationBob, "\n")

	akey := ExponentiationBob ^ SecretNumberAlice%p
	bkey := ExponentiationAlice ^ SecretNumberBob%p
	fmt.Println("Secret key Alice:", akey, "Secret key Bob: ", bkey)

}
