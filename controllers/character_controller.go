package controllers

import (
	"github.com/Enricko/go-discord-bot/models"
	"github.com/Enricko/go-discord-bot/utils"
	"github.com/Enricko/go-discord-bot/views"
)

func CreateCharacter(userID, username string) string {
	if _, exists := models.Characters[userID]; exists {
		return views.RenderMessage("You already have a character!")
	}

	character := models.Character{
		Name:       username,
		Health:     utils.RandomInt(50, 100),
		Attack:     utils.RandomInt(1, 10),
		Defense:    utils.RandomInt(1, 5),
		Experience: 0,
	}

	models.Characters[userID] = character
	return views.RenderCharacterStats(character, "Character created!")
}

func ShowStats(userID string) string {
	character, exists := models.Characters[userID]
	if !exists {
		return views.RenderMessage("You don't have a character yet! Use !create to make one.")
	}

	return views.RenderCharacterStats(character, "")
}
