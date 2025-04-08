package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func Run() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal("Failed to load BOT_TOKEN from environment variable", err)
	}

	discord.AddHandler(ping)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	discord.Open()
	if err != nil {
		log.Fatal("Failed to open discord connection", err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
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
