package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleGreet(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	mentions := getMentions(command)
	authorIds := []string{"<@" + m.Author.ID + ">"}
	if len(mentions) > 0 {
		authorIds = mentions
	}
	var msg strings.Builder
	fmt.Fprintf(&msg, "%v Master %v \n %v \n", getHelloBasedOnHour(), strings.Join(authorIds, " "), selectRandom((greetings)))
	_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
	if err != nil {
		fmt.Println("Error greeting, ", err)
	}
}
