package el_gamal

import (
	"fmt"
	"github.com/olegfomenko/homomorphic-encryption/internal/utils"
	"math/rand"
)

const (
	p int64 = 1000000007
	g int64 = 13
)

func GenKeypair() (int64, int64) {
	x := rand.Int63n(p - 1)
	y := utils.Pow(g, x, p)
	return x, y
}

func Enc(y int64, M int64) (int64, int64) {
	k := rand.Int63n(p - 1)
	a := utils.Pow(g, k, p)
	b := utils.Pow(y, k, p) * M % p
	return a, b
}

func Dec(x, a, b int64) int64 {
	return b * utils.Pow(a, p - 1 - x, p) % p
}

func Test()  {
	fmt.Println("\n--------ElGamal test--------")
	prv, pub := GenKeypair()
	fmt.Println("Private/ Public keys:", prv, pub)

	s0, s1 := Enc(pub, 102)
	fmt.Println("Encrypted data:", s0, s1)
	fmt.Println("Decrypted data:", Dec(prv, s0, s1))

	t0, t1 := Enc(pub, 5)
	t0, t1 = t0 * s0 % p, t1 * s1 % p
	fmt.Println("Decrypted multiply:", Dec(prv, t0, t1))
	fmt.Println("--------ElGamal test end--------\n\n\n")
}