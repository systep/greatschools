package greatschools

import "testing"
import "io/ioutil"
import "encoding/json"
import Ω "github.com/onsi/gomega"

const (
	testLat      = 33.39657
	testLon      = -112.03422
	testLevel    = "e"
	httpProxyURL = "97.77.104.22"
	filePath     = "test/test.json"
)

func TestAPIWithProxy(t *testing.T) {
	Ω.RegisterTestingT(t)

	expectedJSON, err := ioutil.ReadFile(filePath)

	if err != nil {
		t.Fatal(err.Error())
	}

	c := New()

	resp, err := c.GetSchools(&Request{
		Lat:   testLat,
		Lon:   testLon,
		Level: testLevel,
		Proxy: httpProxyURL,
	})

	if err != nil {
		t.Fatal(err.Error())
	}

	data, err := json.Marshal(resp)

	if err != nil {
		t.Fatal(err.Error())
	}

	Ω.Ω(data).Should(Ω.MatchJSON(expectedJSON), "JSON Mismatch")
}
