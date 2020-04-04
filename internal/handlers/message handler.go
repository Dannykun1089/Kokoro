//
//This is fired whenever a message is sent. Mainly used to implement commands
//

package handlers

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dannykun1089/kokoro/v2/globals"
	"github.com/dannykun1089/kokoro/v2/internal/commandutils"
	"github.com/dannykun1089/kokoro/v2/pkg/utils"
	"github.com/dannykun1089/kokoro/v2/types"
)

//MessageHandler - Is fired whenever a message is sent in any channel the bot has access to
func MessageHandler(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {

	//Return if the message sender is the bot itself
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	//In order to prevent complications with people using upercase and lower case in the commands, the bot defaults the content to lowercase
	messageCreate.Content = strings.ToLower(messageCreate.Content)

	//Split the message up into words separated by the space character so we can easily access them for later use, mainly for determining if the message is a valid command
	splitContent := strings.Split(messageCreate.Content, " ")

	//Constructs a command from the split up message content, if unable to, the variable "command" will be an empty string, which will checked in order to verify if the message contains a command
	var (
		command     string
		commandArgs []string
	)
	if len(messageCreate.Content) > len(globals.CommandPrefix) { //Checks if the length of the message's content is larger than the length of the command prefix to prevent an out of bounds error
		if messageCreate.Content[:len(globals.CommandPrefix)] == globals.CommandPrefix && len(splitContent) > 1 { //Checks if the message content starts with the command prefix and has text after that, that could be a command
			command = splitContent[1]      //The text after the command prefix
			commandArgs = splitContent[2:] //Like the original splitContent but with the command prefix and command removed from it, so we can treat it like traditional arguments
		}
	}

	//If there was text after the command prefix its ran through this to see if its a command that works
	if command != "" {
		switch command {
		case "help":
			session.ChannelMessageSendEmbed(messageCreate.ChannelID, &globals.HelpData)
		case "hug":
			commandutils.Emote(session, messageCreate, "hug", globals.CommandData.Emotes.Hugs, commandArgs)
		case "pat":
			commandutils.Emote(session, messageCreate, "pat", globals.CommandData.Emotes.Pats, commandArgs)
		case "cute":
			mediaURL := commandutils.RandomFourchanMedia("c", session, messageCreate)
			session.ChannelMessageSend(messageCreate.ChannelID, mediaURL)
		case "weebm":
			mediaURL := commandutils.RandomFourchanMedia("wsg", session, messageCreate)
			session.ChannelMessageSend(messageCreate.ChannelID, mediaURL)
		case "sauce":

			//Check if there is actually an argument passed into the command
			if len(commandArgs) == 0 {
				session.ChannelMessageSend(messageCreate.ChannelID, fmt.Sprintf("Please provide me with a link to an image, e.g: ```%s sauce https://i.imgur.com/JzhkHsy.jpg```", globals.CommandPrefix))
				return
			}

			//Check if the link they've given is actually correct and usable
			if valid, _ := regexp.MatchString(`http://|https://.*\.png|\.jpg|\.gif`, commandArgs[0]); !valid {
				session.ChannelMessageSend(messageCreate.ChannelID, "The link dosen't lead to a png, jpg or gif file")
				return
			}

			//Construct API request from the arguments
			apiURL := "https://saucenao.com/search.php?api_key=" + globals.Instance.APIKeys.SauceNao + "&output_type=2&url=" + commandArgs[0]

			//Hit the API for a response
			var sauceResults types.SauceNaoJSONResponse
			if err := utils.JSONLinkToStruct(apiURL, &sauceResults); err != nil {
				if !commandutils.CommandErrorCheck(err, "Error contacting saucenao API, either the bot has fucked up majorly, or, saucenao has fucked up majorly", session, messageCreate) {
					return
				}
			}

			//Fix it up
			commandutils.ProcessSaucenaoResponse(&sauceResults)

			matchFound := false
			for i := 0; i < len(sauceResults.Results); i++ {
				if sauceResults.Results[i].Header.FloatSimilarity > globals.MatchThreshold {
					session.ChannelMessageSend(messageCreate.ChannelID, sauceResults.Results[i].Data.ExtURLs[0])
					matchFound = true
				}
			}

			if !matchFound {
				session.ChannelMessageSend(messageCreate.ChannelID, "The saucenao API couldn't find any matches")
			}

		case "avatar":
			//This is kinda similar to the emote command, but its different in a few ways, and i dont wanna convolute anything with function names and parameters, so yea
			if len(commandArgs) != 0 {
				if len(commandArgs[0]) == 22 {

					userID := commandArgs[0][3 : len(commandArgs[0])-1]
					user, err := session.User(userID)

					text := fmt.Sprintf("[Avatar of %s](%s)", user.Username, user.AvatarURL("512"))

					if err == nil {
						embed := &discordgo.MessageEmbed{
							Description: text,
							Color:       globals.EmbedColor,
							Image: &discordgo.MessageEmbedImage{
								URL: user.AvatarURL("256"),
							},
						}
						session.ChannelMessageSendEmbed(messageCreate.ChannelID, embed)
						return
					}
				}
			}
			//Default to sending the user's avatar
			text := fmt.Sprintf("[Avatar of %s](%s)", messageCreate.Author.Username, messageCreate.Author.AvatarURL("512"))

			embed := &discordgo.MessageEmbed{
				Description: text,
				Color:       globals.EmbedColor,
				Image: &discordgo.MessageEmbedImage{
					URL: messageCreate.Author.AvatarURL("256"),
				},
			}
			session.ChannelMessageSendEmbed(messageCreate.ChannelID, embed)

		case "rolf":
			session.ChannelMessageSend(messageCreate.ChannelID, globals.CommandData.RolfQuotes[rand.Intn(len(globals.CommandData.RolfQuotes))])
		case "catch":
			session.ChannelMessageSend(messageCreate.ChannelID, "https://youtu.be/dlAbFOkqxno")
		default:
			session.ChannelMessageSend(messageCreate.ChannelID, "This command aint implemented yet mate")
		}
	}
}
