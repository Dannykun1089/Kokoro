//
//Functions relating to anything to do with http in general
//

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//JSONLinkToStruct - Hits the given URL for a JSON reponse and unmarshals it into the struct
func JSONLinkToStruct(URL string, struc interface{}) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, struc); err != nil {
		return err
	}
	return nil
}
