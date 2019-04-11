package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/random"
	"github.com/ironarachne/regiongen"

	"github.com/patrickmn/go-cache"
)

func main() {
	r := chi.NewRouter()
	centralCache := cache.New(5*time.Minute, 10*time.Minute)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())

		region := regiongen.GenerateRegion("random")

		json.NewEncoder(w).Encode(region)
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/heraldry", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")

			heraldry, found := centralCache.Get("heraldry_" + id)

			if found {
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(heraldry.(string)))
			} else {
				random.SeedFromString(id)
				region := regiongen.GenerateRegion("random")
				centralCache.Set("region_"+id, region, cache.DefaultExpiration)
				centralCache.Set("heraldry_"+id, region.RulerHeraldry, cache.DefaultExpiration)

				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(region.RulerHeraldry))
			}
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")

			region, found := centralCache.Get("region_" + id)

			if found {
				json.NewEncoder(w).Encode(region)
			} else {
				random.SeedFromString(id)

				region := regiongen.GenerateRegion("random")
				centralCache.Set("region_"+id, region, cache.DefaultExpiration)
				centralCache.Set("heraldry_"+id, region.RulerHeraldry, cache.DefaultExpiration)

				json.NewEncoder(w).Encode(region)
			}
		})
	})

	fmt.Println("Region Generator API is online.")
	log.Fatal(http.ListenAndServe(":7970", r))
}
