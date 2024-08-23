package controllers

import (
	"github.com/Enricko/go-discord-bot/models"
	"github.com/Enricko/go-discord-bot/utils"
	"github.com/Enricko/go-discord-bot/views"
)

func Battle(userID string) []string {
	character, exists := models.Characters[userID]
	if !exists {
		return []string{views.RenderMessage("You don't have a character yet! Use !create to make one.")}
	}

	enemy := models.Character{
		Name:    "Goblin",
		Health:  utils.RandomInt(20, 50),
		Attack:  utils.RandomInt(1, 5),
		Defense: utils.RandomInt(1, 3),
	}

	messages := []string{views.RenderEnemyAppears(enemy)}

	for {
		// Player's turn
		damage := utils.Max(0, character.Attack-enemy.Defense)
		enemy.Health -= damage
		messages = append(messages, views.RenderAttack(character.Name, enemy.Name, damage))

		if enemy.Health <= 0 {
			expGain := utils.RandomInt(1, 10)
			character.Experience += expGain
			messages = append(messages, views.RenderVictory(enemy.Name, expGain))
			models.Characters[userID] = character
			return messages
		}

		// Enemy's turn
		damage = utils.Max(0, enemy.Attack-character.Defense)
		character.Health -= damage
		messages = append(messages, views.RenderAttack(enemy.Name, character.Name, damage))

		if character.Health <= 0 {
			messages = append(messages, views.RenderDefeat())
			delete(models.Characters, userID)
			return messages
		}

		messages = append(messages, views.RenderBattleStatus(character, enemy))
	}
}
