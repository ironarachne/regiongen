package regiongen

import (
	"github.com/ironarachne/chargen"
	"github.com/ironarachne/towngen"
)

// Region is a map region
type Region struct {
	Towns []towngen.Town
	Name  string
	Class string
	Ruler chargen.Character
}

// GenerateRegion generates a random region
func GenerateRegion() {
	region := Region{}
	newTown := towngen.Generate()
	region.Towns = append(region.Towns, newTown)

	for i := 0; i < 2; i++ {
		newTown = towngen.Generate()
		region.Towns = append(region.Towns, newTown)
	}

	region.Ruler = chargen.Generate()

	region.Class = "fiefdom"
	region.Name = "region"

	return region
}
