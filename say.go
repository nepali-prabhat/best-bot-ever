package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lukesampson/figlet/figletlib"
)

func handleSay(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	message := removeMentions(strings.TrimSpace(command[len("say"):]))
	authorId := "<@" + m.Author.ID + ">"
	var msg strings.Builder
	fmt.Fprintf(&msg, "Master %v says\n```\n", authorId)
	if font == nil {
		msg.WriteString(message)
	} else {
		var tempMsg strings.Builder
		figletlib.FPrintMsg(&tempMsg, message, font, 80, font.Settings(), "left")
		var figletMsg string
		if tempMsg.Len() > 1900 {
			tempMsg.Reset()
			figletlib.FPrintMsg(&tempMsg, "too long msg .|.", font, 80, font.Settings(), "left")
		} else {
			figletMsg = tempMsg.String()
		}
		msg.WriteString(figletMsg)
	}
	msg.WriteString("\n```")
	_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
	if err != nil {
		fmt.Println("Error saying, ", err)
	}
}
