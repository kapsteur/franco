package franco

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestFranco(t *testing.T) {

	fisxtureFile, e := ioutil.ReadFile("fixture.json")
	if e != nil {
		log.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var fisxtureData map[string]interface{}
	err := json.Unmarshal(fisxtureFile, &fisxtureData)
	if e != nil {
		log.Printf("Error during languages decoding: %v\n", err)
		os.Exit(1)
	}
	for code, values := range fisxtureData {
		for _, v := range values.([]interface{}) {
			res := DetectOne(v.(string))
			if code != res.Code {
				t.Errorf("FixtureCode:%s  DetectedCode:%s", code, res.Code)
			}
		}
	}
}
