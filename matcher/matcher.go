package matcher

import (
	"github.com/cset"
)

var males = cset.New()
var females = cset.New()

var paidMales = cset.New()
var paidFemales = cset.New()

func AddMale(name string) {
	males.Add(name)
}

func RemoveMale(name string) {
	males.Remove(name)
}

func AddFemale(name string) {
	females.Add(name)
}

func RemoveFemale(name string) {
	females.Remove(name)
}

func HandlePaidMale(name string) {
	if males.Has(name) && !paidMales.Has(name) {
		paidMales.Add(name)
		males.Remove(name)
	}
}
