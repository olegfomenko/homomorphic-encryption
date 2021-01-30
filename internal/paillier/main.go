package paillier

import (
	"fmt"
	"github.com/olegfomenko/homomorphic-encryption/internal/utils"
	"math/big"
	"math/rand"
)

const (
	p int64 = 1_000_000_007
	q int64 = 41
)

var P = big.NewInt(p)
var Q = big.NewInt(q)

var n = big.NewInt(0).Mul(P, Q)
var nn = utils.BigSquare(n)

func l_func(u, n *big.Int) *big.Int {
	h1 := utils.BigDec(u)
	return h1.Div(h1, n)
}

func rev_func(a *big.Int) *big.Int {
	h1 := big.NewInt(0).Sub(n, P)
	h2 := big.NewInt(0).Sub(n, Q)
	h1.Mul(h1, h2)
	h1.Sub(h1, big.NewInt(1))
	return utils.BigPow(a, h1, n)
}

func GenKeypair() (*big.Int, *big.Int, *big.Int, *big.Int) {
	g := big.NewInt(rand.Int63n(p))

	l := utils.BigLCM(utils.BigDec(P), utils.BigDec(Q))

	gL := utils.BigPow(g, l, nn)
	u := rev_func(l_func(gL, n))

	return n, g, l, u
}

func Enc(n, g, M *big.Int) *big.Int {
	r := big.NewInt(rand.Int63n(p))
	s1 := utils.BigPow(g, M, nn)
	s2 := utils.BigPow(r, n, nn)
	return utils.BigMul(s1, s2, nn)
}

func Dec(n, l, u, c *big.Int) *big.Int {
	h := utils.BigPow(c, l, nn)
	hL := l_func(h, n)
	return utils.BigMul(hL, u, n)
}

func Test()  {
	fmt.Println("\n--------Paillier test--------")

	n, g, l, u := GenKeypair()
	fmt.Println("Private / Public keys:", n, g, l, u)

	s1 := Enc(n, g, big.NewInt(102))
	fmt.Println("Encrypted data 102:", s1)
	fmt.Println("Decrypted data 102:", Dec(n, l, u, s1))


	s2 := Enc(n, g, big.NewInt(200))
	fmt.Println("Encrypted data 200:", s2)
	fmt.Println("Decrypted data 200:", Dec(n, l, u, s2))

	s := utils.BigMul(s1, s2, nn)
	fmt.Println("Decrypted sum:", Dec(n, l, u, s))

	fmt.Println("--------Paillier test end--------\n\n\n")
}