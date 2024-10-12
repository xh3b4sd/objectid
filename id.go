package objectid

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"

	"github.com/xh3b4sd/tracer"
)

type ID string

func Random(num int64) ID {
	return ID(fmt.Sprintf("%d%06d", num, rand.Intn(999999)))
}

func System() ID {
	return ID("1")
}

func (i ID) Big() *big.Int {
	b, ok := big.NewInt(0).SetString(string(i), 10)
	if !ok {
		tracer.Panic(fmt.Errorf("cannot convert %s to *big.Int", i))
	}

	return b
}

func (i ID) Float() float64 {
	f, e := strconv.ParseFloat(string(i), 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return f
}

func (i ID) Int() int64 {
	n, e := strconv.ParseInt(string(i), 10, 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return n
}

func (i ID) String() string {
	return string(i)
}
