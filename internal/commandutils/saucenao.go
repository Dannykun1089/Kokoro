//
//Functions relating to the saucenao API
//

package commandutils

import (
	"strconv"

	"github.com/dannykun1089/kokoro/v2/types"
)

//ProcessSaucenaoResponse - Convert the string type similarity data into a float type so we can actually work with it numerically
func ProcessSaucenaoResponse(results *types.SauceNaoJSONResponse) {
	for i := 0; i < len(results.Results); i++ {
		floatValue, _ := strconv.ParseFloat(results.Results[i].Header.Similarity, 64)
		results.Results[i].Header.FloatSimilarity = floatValue
	}
}
