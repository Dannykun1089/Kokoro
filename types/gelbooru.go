package types

//GelbooruPost - Holds the data from one gelbooru post json object, the api vends these as an array so i have to define the variable as an array of these
type GelbooruPost struct {
	Tags    string `json:"tags"`
	FileURL string `json:"file_url"`
}
