package regiongen

import (
	"math/rand"

	"github.com/ironarachne/chargen"
)

func (region Region) generateRuler() chargen.Character {
	ruler := chargen.GenerateCharacterOfCulture(region.Culture)
	ruler.Profession = "noble"
	ruler.AgeCategory = "adult"
	ruler.Age = rand.Intn(40) + 25

	return ruler
}
