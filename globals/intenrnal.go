//
//Variables which hold the bot's internal data structures
//

package globals

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dannykun1089/kokoro/v2/types"
)

var (
	//Instance - Holds API keys and the bots token
	Instance types.InstanceStruct
	//CommandData - Holds data for use in bot commands
	CommandData types.CommandDataStruct
	//HelpData - Holds the data for the help message
	HelpData discordgo.MessageEmbed
)
