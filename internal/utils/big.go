package utils

import "math/big"

func IsZero(a *big.Int) bool {
	return a.Cmp(big.NewInt(0)) == 0
}

func IsOne(a *big.Int) bool {
	return a.Cmp(big.NewInt(1)) == 0
}

func BigMul(a, b, mod *big.Int) *big.Int {
	bigAB := big.NewInt(0).Mul(a, b)
	return bigAB.Mod(bigAB, mod)
}

func BigPow(a, step, mod *big.Int) *big.Int {
	a = big.NewInt(0).Mod(a, mod)

	if IsZero(step) {
		return big.NewInt(1)
	} else if IsOne(step) {
		return a
	}

	h := BigPow(BigMul(a, a, mod), big.NewInt(0).Div(step, big.NewInt(2)), mod)
	flag := big.NewInt(0).Mod(step, big.NewInt(2))

	if IsZero(flag) {
		return h
	} else {
		return BigMul(h, a, mod)
	}
}

func BigGCD(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	} else {
		return BigGCD(b, big.NewInt(0).Mod(a, b))
	}
}

func BigLCM(a, b *big.Int) *big.Int {
	g := BigGCD(a, b)
	h := big.NewInt(0).Div(a, g)
	return h.Mul(h, b)
}

func BigDec(a *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, big.NewInt(1))
}

func BigInc(a *big.Int) *big.Int {
	return big.NewInt(0).Add(a, big.NewInt(1))
}

func BigSquare(a *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, a)
}