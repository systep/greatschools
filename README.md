# GreatSchools [![CircleCI](https://circleci.com/gh/skyline-ai/greatschools.svg?style=svg&circle-token=03da7c8ea3bead0f24e2beb905dba35ca74e7dcf)](https://circleci.com/gh/skyline-ai/greatschools)

The following is an unoffical go client for the GreatSchools API. GreatSchools is a website that provides stats about great schools in geos.
This unoffical software is not affiliated with GreatSchools.org in any way. Visit GreatSchools [here](http://www.greatschools.org/).

## Usage Example

Initialize a client, and use the `GetSchools` function, passing lat/lon coordinates.

```go
package main

import (
	"flag"
	"log"

	"encoding/json"

	"github.com/skyline-ai/greatschools"
)

var (
	lat   = flag.Float64("lat", 33.39657, "latitude")
	lon   = flag.Float64("lon", -112.03422, "longitude")
	level = flag.String("level", "", "school level")
)

func init() {
	flag.Parse()
}

func main() {
	c := greatschools.New()

	resp, err := c.GetSchools(&greatschools.Request{
		Lat:   *lat,
		Lon:   *lon,
		Level: *level,
	})

	if err != nil {
		log.Fatalln(err)
	}

	j, err := json.MarshalIndent(resp.Results[0], " ", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("response: %s", j)

	// this will print for example:
	//     2017/02/08 12:07:13 response: {
	//  	"level": "e",
	//  	"schools": [
	//  		{
	//  			"lon": -112.037,
	//  			"gradeRange": "K-8",
	//  			"state": "AZ",
	//  			"type": "school",
	//  			"schoolType": "public",
	//  			"url": "/arizona/phoenix/1252-T-G-Barr-School/",
	//  			"distance": 0.774568217624737,
	//  			"districtId": 1142,
	//  			"address": {
	//  				"street2": "",
	//  				"zip": "85042",
	//  				"street1": "2041 East Vineyard",
	//  				"cityStateZip": "Phoenix, AZ  85042"
	//  			},
	//  			"numReviews": 6,
	//  			"isNewGSRating": true,
	//  			"name": "T G Barr School",
	//  			"rating": 2,
	//  			"parentRating": 4,
	//  			"grades": "KG,1,2,3,4,5,6,7,8",
	//  			"lat": 33.3856,
	//  			"id": 1252
	//  		}
	//  	]
	//  }
}

```