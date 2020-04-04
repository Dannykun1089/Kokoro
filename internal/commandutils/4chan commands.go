//
//Functions used in 4Chan related commands such as cute and weebm
//

package commandutils

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/moshee/go-4chan-api/api"
)

//RandomFourchanMedia - Gets a random peice of media off of a specified board on 4Chan
func RandomFourchanMedia(board string, session *discordgo.Session, messageCreate *discordgo.MessageCreate) string {
	APIThreadIDs, err := api.GetThreads(board)
	if CommandErrorCheck(err, "Error contacting 4Chan's API", session, messageCreate) {
		return ""
	}

	threadIDs := recompileThreadIDs(APIThreadIDs)

	thread, err := api.GetThread(board, threadIDs[rand.Intn(len(threadIDs))])
	if CommandErrorCheck(err, "Error contacting 4Chan's API", session, messageCreate) {
		return ""
	}

	var post *api.Post

	for {
		post = thread.Posts[rand.Intn(len(thread.Posts))]
		if post.ImageURL() != "" {
			return post.ImageURL()
		}
	}
}

//Turns the 2D array returned from the API wrapper into a regular one
func recompileThreadIDs(APIThreadIDs [][]int64) []int64 {
	var retval []int64

	for _, element := range APIThreadIDs {
		for _, nestedElement := range element {
			retval = append(retval, nestedElement)
		}
	}
	return retval
}
