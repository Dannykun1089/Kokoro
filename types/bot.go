//
//The bot's data structures
//

package types

//InstanceStruct - Stores API keys and the bot's token
type InstanceStruct struct {
	DiscordToken string `json:"discord_token"`
	APIKeys      struct {
		SauceNao string `json:"saucenao"`
	} `json:"api_keys"`
}

//CommandDataStruct - Stores data for bot commands
type CommandDataStruct struct {
	Emotes struct {
		Pats []string `json:"pats"`
		Hugs []string `json:"hugs"`
	} `json:"emotes"`
	RolfQuotes []string `json:"rolf_quotes"`
}
