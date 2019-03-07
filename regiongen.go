package regiongen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/climategen"
	"github.com/ironarachne/heraldry"
	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/orggen"
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
	Organizations []orggen.Organization
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
	climate := climategen.Climate{}

	if regionType == "random" {
		climate = climategen.Generate()
	} else {
		climate = climategen.GetClimate(regionType)
	}

	regionType = climate.Name

	region.Biome = climate.Name

	region.Class = randomClass()

	newTown := towngen.GenerateTown("city", regionType)
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown = towngen.GenerateTown("random", regionType)
		region.Towns = append(region.Towns, newTown)
	}

	numberOfOrgs := rand.Intn(3) + 1
	newOrg := orggen.Organization{}

	for i := 0; i < numberOfOrgs; i++ {
		newOrg = orggen.Generate()
		region.Organizations = append(region.Organizations, newOrg)
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
