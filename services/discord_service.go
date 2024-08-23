package services

import (
	"strings"

	"github.com/Enricko/go-discord-bot/controllers"

	"github.com/bwmarrin/discordgo"
)

type DiscordService struct {
	session *discordgo.Session
}

func NewDiscordService(token string) (*DiscordService, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &DiscordService{session: session}, nil
}

func (d *DiscordService) Start() error {
	d.session.AddHandler(d.messageCreate)
	return d.session.Open()
}

func (d *DiscordService) Stop() {
	d.session.Close()
}

func (d *DiscordService) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Fields(m.Content)
	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0])
	switch command {
	case "!create":
		response := controllers.CreateCharacter(m.Author.ID, m.Author.Username)
		s.ChannelMessageSend(m.ChannelID, response)
	case "!stats":
		response := controllers.ShowStats(m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, response)
	case "!battle":
		responses := controllers.Battle(m.Author.ID)
		for _, response := range responses {
			s.ChannelMessageSend(m.ChannelID, response)
		}
	}
}
