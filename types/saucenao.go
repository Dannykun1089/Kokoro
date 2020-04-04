//
//Data structures for use with the saucenao API
//

package types

//SauceNaoJSONResponse - A JSON struct for sorting out the SauceNao responses
type SauceNaoJSONResponse struct {
	Results []struct { //Image matches
		Header struct { //Match header data
			Similarity      string  `json:"similarity"`
			FloatSimilarity float64 //This value is filled in later because saucenao dosent know how to write APIs
		} `json:"header"`
		Data struct { //Match main data
			ExtURLs []string `json:"ext_urls"` //Image URLs associated with the match, usually only 1 item long
			Source  string   `json:"source"`   //Name of the site it was retrieved from
		} `json:"data"`
	} `json:"results"`
}
