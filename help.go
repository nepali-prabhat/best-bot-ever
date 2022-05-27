package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	var msg strings.Builder
	fmt.Fprintf(&msg, "<@%v> ", m.Author.ID)
	if strings.HasSuffix(command, "greet") {
		msg.WriteString(greetHelpText())
	} else if strings.HasSuffix(command, "insult") {
		msg.WriteString(insultHelpText())
	} else if strings.HasSuffix(command, "boobs") {
		msg.WriteString(boobsHelpText())
	} else {
		msg.WriteString(helpText())
	}
	_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
	if err != nil {
		fmt.Println("Error sending help, ", err)
	}
}

func handleHelpFallback(s *discordgo.Session, m *discordgo.MessageCreate) {
	var msg strings.Builder
	// Send a text message with the list of Gophers
	authorId := m.Author.ID
	fmt.Fprintf(&msg, "Yes master <@%v> \n %v \n", authorId, selectRandom((greetings)))
	msg.WriteString(helpText())
	_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
	if err != nil {
		fmt.Println("Error greeting, ", err)
	}
}
