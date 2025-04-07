package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
)

func Run() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(ping)

	discord.Open()
	defer discord.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	switch {
	case strings.Contains(m.Content, "!ping"):
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case strings.Contains(m.Content, "!test"):
		s.ChannelMessageSend(m.ChannelID, "I don't like test!")
	}
}
