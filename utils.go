package main

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func getCommand(content string) (string, error) {
	trimmedContent := strings.TrimSpace(content)
	if strings.HasPrefix(trimmedContent, botPrefix) {
		command := strings.TrimSpace(trimmedContent[len(botPrefix):])
		return command, nil
	} else {
		return "", errors.New("No prefix")
	}
}

func helpText() string {
	var msg strings.Builder
	fmt.Fprintf(&msg, "Call me %v and give me commands. I respond to the following: \n", botPrefix)
	msg.WriteString("- greet\n")
	msg.WriteString("- insult\n")
	msg.WriteString("- boobs\n")
	msg.WriteString("You can also type !bitch help followed by any command learn about that command. Try !bitch help insult (─‿‿─)♡")
	return msg.String()
}

func greetHelpText() string {
	return "***greet***:\nOrder me to greet yourself or anyone you mention. \n"
}

func insultHelpText() string {
	return "***insult***:\nOrder me to say hurtful things to your enemies. I'll insult whoever you mention. \n"
}

func boobsHelpText() string {
	return "***boobs***:\nOrder me to show you boobs. I might get shy sometimes. \n"
}

func selectRandom(arr []string) string {
	return arr[rand.Intn(len(arr))]
}

func getHelloBasedOnHour() string {
	now := time.Now()
	hour := now.Hour()
	if hour > 4 && hour < 12 {
		return "Ohayou!"
	} else if hour < 16 {
		return "Konnichiwa!"
	} else if hour < 23 {
		return "Konbanwa!"
	} else {
		return "Oyasumi!"
	}
}

func getMentions(command string) []string {
	mentionRegex := regexp.MustCompile(`<@!?(\d+)>`)
	mentions := mentionRegex.FindAllString(command, -1)
	return mentions
}
