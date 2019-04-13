package regiongen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/climategen"
	"github.com/ironarachne/culturegen"
	"github.com/ironarachne/heraldry"
	"github.com/ironarachne/orggen"
	"github.com/ironarachne/random"
	"github.com/ironarachne/towngen"
)

// Region is a map region
type Region struct {
	Biome         string
	Culture       culturegen.Culture
	Climate       climategen.Climate
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
	class := random.ItemFromThresholdMap(classes)
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
	region.Climate = climate
	region.Culture = culturegen.GenerateCulture()
	region.Culture = region.Culture.SetClimate(region.Biome)

	region.Class = randomClass()

	newTown := towngen.GenerateTown("city", regionType)
	newTown = towngen.SetCulture(region.Culture, newTown)
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown = towngen.GenerateTown("random", regionType)
		newTown = towngen.SetCulture(region.Culture, newTown)
		region.Towns = append(region.Towns, newTown)
	}

	numberOfOrgs := rand.Intn(3) + 1
	newOrg := orggen.Organization{}

	for i := 0; i < numberOfOrgs; i++ {
		newOrg = orggen.Generate()
		region.Organizations = append(region.Organizations, newOrg)
	}

	region.Ruler = region.generateRuler()

	region.RulerTitle = region.Class.RulerTitleFemale
	if region.Ruler.Gender == "male" {
		region.RulerTitle = region.Class.RulerTitleMale
	}

	device := heraldry.Generate()
	region.RulerHeraldry = device.RenderToSVG(320, 420)
	region.RulerBlazon = device.RenderToBlazon()

	regionName := region.Culture.Language.RandomName()
	region.Name = strings.Title(regionName)

	return region
}
