# GeoJSON with GraphQL

[GeoJSON](https://geojson.org/) is a really popular format for exchanging GIS data.

[GraphQL](https://graphql.org/) is a really popular protocol for exchanging data.

Can't we get these two together?

It turns out that GraphQL can't represent data the way GeoJSON wants to.  In particular, properties (maps) [can't be easily represented](https://github.com/graphql/graphql-spec/issues/101).  Coordinates are also [problematic](https://medium.com/@brygrill/creating-a-geojson-featurecollection-type-for-graphql-352591451b4a).

Since we can't encode the schema in GraphQL's SDL syntax, we can fallback to using [custom scalars](https://www.apollographql.com/docs/apollo-server/schema/scalars-enums/).  In particular, we push the entire parsing/serialization onto the custom scalar, leaving the resulting SDL very minimal:

```
scalar Feature
scalar FeatureCollection
```

This project uses [graphql-go](https://github.com/graph-gophers/graphql-go) for the backend and adds GeoJSON support via [go.geojson](https://github.com/paulmach/go.geojson).

To see this in action, simply check out this project and:
```
go build
./graphql-geojson-go
```

Then open your browser to `http://localhost:8080/graphiql.html` to see the GraphiQL IDE.

Here's some example queries:

<table>
    <tr>
        <th>Query</th>
        <th>Response</th>
    </tr>
    <tr>
        <td style="vertical-align:top">
            <pre>
query {
  feature
}
            </pre>
        </td>
        <td style="vertical-align:top">
            <pre>
{
  "data": {
    "feature": {
      "type": "Feature",
      "geometry": {
        "type": "Point",
        "coordinates": [
          [
            125.6,
            10.1
          ]
        ]
      },
      "properties": {
        "name": "Dinagat Islands"
      }
    }
  }
}
            </pre>
        </td>
    </tr>
</table>

And some example mutations:

<table>
    <tr>
        <th>Query</th>
        <th>Response</th>
    </tr>
    <tr>
        <td style="vertical-align:top">
            <pre>
mutation setFeatureCollection($featureCollection: FeatureCollection!) {
  setFeatureCollection(featureCollection:$featureCollection)
}
            </pre>
        </td>
        <td style="vertical-align:top" rowspan="2">
            <pre>
{
  "data": {
    "setFeatureCollection": {
      "type": "FeatureCollection",
      "features": [
        {
          "type": "Feature",
          "geometry": {
            "type": "Point",
            "coordinates": [
              125.6,
              10.1
            ]
          },
          "properties": {
            "name": "Dinagat Islands"
          }
        }
      ]
    }
  }
}
            </pre>
        </td>
    </tr>
    <tr>
        <td style="vertical-align:top">
            <pre>
{
	"featureCollection": {
		"type": "FeatureCollection",
		"features": [{
			"type": "Feature",
			"geometry": {
				"type": "Point",
				"coordinates": [125.6, 10.1]
			},
			"properties": {
				"name": "Dinagat Islands"
			}
		}]
	}
}
            </pre>
        </td>
    </tr>
</table>
