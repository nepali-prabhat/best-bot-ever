// some ideas:
// add ascii letter emojis as png
// react with ascii letters to whatever message is being replied to
// gender and pronoun support
// if being replied to a msg, bitch's reply will reply to that msg too
// send in image format
package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/lukesampson/figlet/figletlib"
)

// global variables
const botPrefix = "!bitch"
var (
	Token string
)
var font *figletlib.Font

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Cannot seed random")
	}
	rand.Seed(int64(binary.LittleEndian.Uint16(b[:])))

	cwd, _ := os.Getwd()
	fontsdir := filepath.Join(cwd, "fonts")

	f, err := figletlib.GetFontByName(fontsdir, "ghost")
	if err != nil {
		panic("Cannot find the ghost font")
	}
	font = f
}


func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating discord session, ", err)
		return
	}

	dg.AddHandler(messageCreate)

	// Only receive message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening connection, ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session once signal is received
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, commandErr := getCommand(m.Content)

	if commandErr != nil {
		return
	}

	// if m.Type == 19 {
	// 	fmt.Println("reply type", m.Type)
	// 	value, err := discordgo.Marshal(m.ReferencedMessage)
	// 	if err == nil {
	// 		fmt.Println("references msg type", string(value))
	// 	}
	// } else {
	// 	fmt.Println("not  reply type", m.Type)
	// }

	var msg strings.Builder
	if command == "" {
		handleHelpFallback(s, m)
	}
	// fmt.Println("command: ", command)

	if strings.HasPrefix(command, "greet") {
		handleGreet(s, m, command)
		return
	}
	if strings.HasPrefix(command, "insult") {
		handleInsult(s, m, command)
		return
	}
	if strings.HasPrefix(command, "boobs") {
		handleBoobs(s, m, command)
		return
	}
	if strings.HasPrefix(command, "test") {
		msg.WriteString("      <:test1:979686404139384832>\n <:test1:979686404139384832>")
		_, err := s.ChannelMessageSend(m.ChannelID, msg.String())
		if err != nil {
			fmt.Println("Error sending tests, ", err)
		}
		return
	}
	if strings.HasPrefix(command, "say") {
		handleSay(s, m, command)
	}
	if strings.HasPrefix(command, "show") {
		handleShow(s, m, command)
	}
	if strings.HasPrefix(command, "help") {
		handleHelp(s, m, command)
	}
	return
}
