package utils

func Pow(a, step, mod int64) int64 {
	a = a % mod

	if step == 0 {
		return 1
	} else if step == 1 {
		return a
	}

	h := Pow(a * a % mod, step / 2, mod)

	if step % 2 == 0 {
		return h
	} else {
		return h * a % mod
	}
}

func GCD(a, b int64) int64 {
	if b == 0 {
		return a
	} else {
		return GCD(b, a % b)
	}
}

func LCM(a, b int64) int64 {
	return (a / GCD(a, b)) * b
}
