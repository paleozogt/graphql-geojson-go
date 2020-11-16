package main

import (
	"encoding/json"
	"log"

	geojson "github.com/paulmach/go.geojson"
)

type Feature struct {
	geojson.Feature
}

func (Feature) ImplementsGraphQLType(name string) bool {
	return name == "Feature"
}

func (f *Feature) UnmarshalGraphQL(input interface{}) error {
	log.Printf("Feature UnmarshalGraphQL")

	str, err := json.Marshal(input)
	if err != nil {
		return nil
	}

	gf, err := geojson.UnmarshalFeature(str)
	if err != nil {
		return nil
	}

	f.Feature = *gf
	return err
}

/////

type FeatureCollection struct {
	geojson.FeatureCollection
}

func (FeatureCollection) ImplementsGraphQLType(name string) bool {
	return name == "FeatureCollection"
}

func (f *FeatureCollection) UnmarshalGraphQL(input interface{}) error {
	log.Printf("FeatureCollection UnmarshalGraphQL")

	str, err := json.Marshal(input)
	if err != nil {
		return nil
	}

	gf, err := geojson.UnmarshalFeatureCollection(str)
	if err != nil {
		return nil
	}

	f.FeatureCollection = *gf
	return err
}
