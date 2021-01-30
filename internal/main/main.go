package main

import (
	el_gamal "github.com/olegfomenko/homomorphic-encryption/internal/el-gamal"
	"github.com/olegfomenko/homomorphic-encryption/internal/paillier"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().Unix())
	paillier.Test()
	el_gamal.Test()
}
