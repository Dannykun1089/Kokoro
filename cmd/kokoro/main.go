//
//The main executable
//

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/dannykun1089/kokoro/v2/globals"
	"github.com/dannykun1089/kokoro/v2/internal/handlers"
	"github.com/dannykun1089/kokoro/v2/pkg/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	//Make the logger look pretty
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	//Read local JSON data into the program
	logrus.Info("Reading JSON data into program")
	utils.CritErrorCheck(utils.JSONFileToStruct("static/instance.json", &globals.Instance), "Error reading instance.json")
	utils.CritErrorCheck(utils.JSONFileToStruct("static/command data.json", &globals.CommandData), "Error reading command data.json")
	utils.CritErrorCheck(utils.JSONFileToStruct("static/help data.json", &globals.HelpData), "Error reading help data.json") //FINALLY SOMONE WHO'S COMPITENT! lord bless you bwmarrin
	logrus.Info("JSON data successfully read")

	//Create client interface with bot credentials
	logrus.Info("Creating new client interface")
	bot, err := discordgo.New("Bot " + globals.Instance.DiscordToken)
	utils.CritErrorCheck(err, "Error creating client interface, please check that the token is valid\n")
	logrus.Info("Client interface created. Now attatching handlers")

	//Add handlers
	bot.AddHandler(handlers.MessageHandler)

	//Open connection to discord
	logrus.Infof("Attempting connection to [%s]", discordgo.EndpointGatewayBot)
	err = bot.Open()
	utils.CritErrorCheck(err, "Error in establishing connection. Discord may be experiencing an outage or your internet connection may be experiencing issues\n")
	logrus.Infof("Connection opened to [%s]", discordgo.EndpointGatewayBot)

	//Wait on a system interupt to close the connection and exit the program
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	logrus.Warn("Closing connection and exiting")
	bot.Close()

	os.Exit(0)

}
