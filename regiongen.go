package regiongen

import (
	"strings"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/heraldry"
	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/towngen"
	"github.com/ironarachne/utility"
)

// Region is a map region
type Region struct {
	Biome         string
	Capital       string
	Class         RegionClass
	Name          string
	Ruler         chargen.Character
	RulerBlazon   string
	RulerHeraldry string
	RulerTitle    string
	Towns         []towngen.Town
}

// RegionClass is a class of region
type RegionClass struct {
	MaxNumberOfTowns int
	MinNumberOfTowns int
	Name             string
	RulerTitleFemale string
	RulerTitleMale   string
}

func randomClass() RegionClass {
	class := utility.RandomItemFromThresholdMap(classes)
	regionClass := classData[class]

	return regionClass
}

// GenerateRegion generates a random region
func GenerateRegion(regionType string) Region {
	region := Region{}

	if regionType == "random" {
		regionType = utility.RandomItem(availableRegions)
	}

	region.Biome = regionType

	region.Class = randomClass()

	newTown := towngen.GenerateTown("city", regionType)
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown = towngen.GenerateTown("random", regionType)
		region.Towns = append(region.Towns, newTown)
	}

	region.Ruler = chargen.GenerateCharacter()
	region.Ruler.Profession = "noble"
	region.RulerTitle = region.Class.RulerTitleFemale
	if region.Ruler.Gender == "male" {
		region.RulerTitle = region.Class.RulerTitleMale
	}

	device := heraldry.Generate()
	region.RulerHeraldry = device.RenderToSVG(320, 420)
	region.RulerBlazon = device.RenderToBlazon()

	regionName := naminglanguage.GeneratePlace()
	region.Name = strings.Title(regionName.Name)

	return region
}
