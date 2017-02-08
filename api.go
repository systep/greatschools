package greatschools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36"
	baseURL   = "http://www.greatschools.org/geo/boundary/ajax/getAssignedSchoolByLocation.json"
)

// Client is the GreatSchools client. It contains all the different resources available.
type Client struct {
	Debug bool
}

// New creates a new GreatSchools client
func New() *Client {
	return &Client{}
}

// Request contains information to sent to the api endpoint
type Request struct {
	Lat   float64
	Lon   float64
	Level string // school level. e.g. "e"
	Proxy string
}

// Response contains Results from the API request
type Response struct {
	Results Results `json:"results"`
}

// Results is a slice of Result
type Results []Result

// Result directly corresponds to the JSON returned by the API
type Result struct {
	Level   string  `json:"level"`
	Schools Schools `json:"schools"`
}

// Schools is a slice of school
type Schools []School

// School contains school data
type School struct {
	Lon           float64 `json:"lon"`
	GradeRange    string  `json:"gradeRange"`
	State         string  `json:"state"`
	Type          string  `json:"type"`
	SchoolType    string  `json:"schoolType"`
	URL           string  `json:"url"`
	Distance      float64 `json:"distance"`
	DistrictID    int     `json:"districtId"`
	Address       Address `json:"address"`
	NumReviews    int     `json:"numReviews"`
	IsNewGSRating bool    `json:"isNewGSRating"`
	Name          string  `json:"name"`
	Rating        int     `json:"rating"`
	ParentRating  int     `json:"parentRating"`
	Grades        string  `json:"grades"`
	Lat           float64 `json:"lat"`
	ID            int     `json:"id"`
}

// Address contains address data
type Address struct {
	Street2      string `json:"street2"`
	Zip          string `json:"zip"`
	Street1      string `json:"street1"`
	CityStateZip string `json:"cityStateZip"`
}

// GetSchools fetches schoold from greatschools.org API
func (c *Client) GetSchools(r *Request) (*Response, error) {
	if r.Lat == 0 || r.Lon == 0 {
		return nil, fmt.Errorf("must provide lat and lon")
	}

	sURL := fmt.Sprintf("%s?lat=%f&lon=%f&level=%s", baseURL, r.Lat, r.Lon, r.Level)

	if c.Debug {
		log.Printf("sURL: %s", sURL)
	}

	client := &http.Client{}

	if len(r.Proxy) > 0 {
		_, err := url.Parse(r.Proxy)
		if err == nil {
			os.Setenv("HTTP_PROXY", r.Proxy)
		}
	}

	req, err := http.NewRequest("GET", sURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var response *Response
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
