package main

import (
	"context"
	"log"
)

// Resolver graphql resolver
type Resolver struct {
}

// Feature method
func (r *Resolver) Feature(ctx context.Context) (Feature, error) {
	log.Printf("Feature")

	feature := &Feature{}
	feature.Feature = *dummyFeature

	return *feature, nil
}

// SetFeature method
func (r *Resolver) SetFeature(ctx context.Context, args struct{ Feature Feature }) (Feature, error) {
	log.Printf("SetFeature")
	return args.Feature, nil
}

// Feature method
func (r *Resolver) FeatureCollection(ctx context.Context) (FeatureCollection, error) {
	log.Printf("FeatureCollection")

	featureCollection := &FeatureCollection{}
	featureCollection.FeatureCollection = *dummyFeatureCollection

	return *featureCollection, nil
}

// SetFeature method
func (r *Resolver) SetFeatureCollection(ctx context.Context, args struct{ FeatureCollection FeatureCollection }) (FeatureCollection, error) {
	log.Printf("SetFeatureCollection")
	return args.FeatureCollection, nil
}
