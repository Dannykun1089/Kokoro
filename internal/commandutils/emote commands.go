//
//Functions relating to the emote commands such as pat and hug
//

package commandutils

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/dannykun1089/kokoro/v2/globals"
)

//Emote - Emote command utility meant to prevent repetitive code and make the message handler less cluttered, good for the program, good for my sanity
func Emote(session *discordgo.Session, messageCreate *discordgo.MessageCreate, action string, emoteURLs []string, commandArgs []string) {
	//Check if somone is mentioned
	if len(commandArgs) != 0 {
		if len(commandArgs[0]) == 22 { //22 Should be the exact character length for a discord mention

			userID := commandArgs[0][3 : len(commandArgs[0])-1]
			user, err := session.User(userID)

			text := fmt.Sprintf("%s %ss %s!", messageCreate.Author.Username, action, user.Username)

			//If there are no errors and the person the user mentioned isn't themselves
			if err == nil && user.ID != messageCreate.Author.ID {
				embed := &discordgo.MessageEmbed{
					Title: text,
					Color: globals.EmbedColor,
					Image: &discordgo.MessageEmbedImage{
						URL: emoteURLs[rand.Intn(len(emoteURLs))],
					},
				}
				session.ChannelMessageSendEmbed(messageCreate.ChannelID, embed)
				return
			}
		}
	}
	//Default to emoting to the user
	text := fmt.Sprintf("Here you go %s! *%s*", messageCreate.Author.Username, action)

	embed := &discordgo.MessageEmbed{
		Title: text,
		Color: globals.EmbedColor,
		Image: &discordgo.MessageEmbedImage{
			URL: emoteURLs[rand.Intn(len(emoteURLs))],
		},
	}
	session.ChannelMessageSendEmbed(messageCreate.ChannelID, embed)
}
