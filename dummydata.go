package main

import geojson "github.com/paulmach/go.geojson"

func makeDummyFeature() *geojson.Feature {
	f := geojson.NewPointFeature([]float64{125.6, 10.1})
	f.Properties["name"] = "Dinagat Islands"
	return f
}

func makeDummyFeatureCollection() *geojson.FeatureCollection {
	f := geojson.NewFeatureCollection()
	f.AddFeature(makeDummyFeature())
	return f
}

var dummyFeature *geojson.Feature = makeDummyFeature()
var dummyFeatureCollection *geojson.FeatureCollection = makeDummyFeatureCollection()
