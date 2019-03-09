package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/ironarachne/regiongen"
	"github.com/kataras/iris"
	"github.com/patrickmn/go-cache"
)

func main() {
	app := iris.New()
	c := cache.New(5*time.Minute, 10*time.Minute)

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("regiongend")
	})

	app.Get("/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			ctx.Writef("error while trying to parse id parameter")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		region, found := c.Get("region_" + strconv.FormatInt(id, 10))

		if found {
			ctx.JSON(region)
		} else {
			rand.Seed(id)
			region := regiongen.GenerateRegion("random")
			c.Set("region_"+strconv.FormatInt(id, 10), region, cache.DefaultExpiration)

			ctx.JSON(region)
		}

	})

	app.Get("/{id:int64}/heraldry.svg", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			ctx.Writef("error while trying to parse id parameter")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		region, found := c.Get("region_" + strconv.FormatInt(id, 10))

		if found {
			regionData := region.(regiongen.Region)
			ctx.ContentType("image/svg+xml")
			ctx.Writef(regionData.RulerHeraldry)
		} else {
			rand.Seed(id)
			region := regiongen.GenerateRegion("random")
			c.Set("region_"+strconv.FormatInt(id, 10), region, cache.DefaultExpiration)
			ctx.ContentType("image/svg+xml")
			ctx.Writef(region.RulerHeraldry)
		}
	})

	app.Run(iris.Addr(":7970"))
}
