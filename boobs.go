package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleBoobs(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	mentions := getMentions(command)
	authorIds := []string{
		"<@" + m.Author.ID + ">"}
	if len(mentions) > 0 {
		authorIds = mentions
	}
	var msg strings.Builder
	for i, author := range authorIds {
		var fmtStr string
		if i == 0 {
			fmtStr = "This for you %v\n%v"
		} else {
			fmtStr = "\nThis for you %v\n%v"
		}
		fmt.Fprintf(&msg, fmtStr, author, selectRandom(boobs))
	}
	_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
	if err != nil {
		fmt.Println("Error sending boobs, ", err)
	}
}
