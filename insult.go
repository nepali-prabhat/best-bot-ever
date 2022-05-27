package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleInsult(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	mentions := getMentions(command)
	var msg strings.Builder
	if len(mentions) > 0 {
		for i, mention := range mentions {
			var fmtStr string
			if i == 0 {
				fmtStr = "%v %v"
			} else {
				fmtStr = "\n%v %v"
			}
			fmt.Fprintf(&msg, fmtStr, mention, selectRandom(insults))
		}
	} else {
		fmt.Fprintf(&msg, "\n%v", selectRandom(insults))
	}
	_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
	if err != nil {
		fmt.Println("Error sending insult, bitch ", err)
	}
}
