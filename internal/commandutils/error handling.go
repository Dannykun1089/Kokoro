//
//Functions relating to error handling within commands
//

package commandutils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

//CommandErrorCheck - A oneliner function for error checking within a discordgo message handler, logs a message to the terminal and sends a message to the command's channel alerting the user of the error. Returns true if an error was caught
func CommandErrorCheck(err error, errorMessage string, session *discordgo.Session, messageCreate *discordgo.MessageCreate) bool {
	if err != nil {
		logrus.Warn(errorMessage)
		logrus.Warn(err)
		if _, err := session.ChannelMessageSend(messageCreate.ChannelID, errorMessage); err != nil {
			logrus.Warn("Error in sending error report back to the command author")
		}
		return true
	}
	return false
}
