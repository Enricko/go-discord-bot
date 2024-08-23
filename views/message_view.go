package views

import (
	"fmt"

	"github.com/Enricko/go-discord-bot/models"
)

func RenderMessage(message string) string {
	return message
}

func RenderCharacterStats(character models.Character, prefix string) string {
	return fmt.Sprintf("%s%s's stats: Health: %d, Attack: %d, Defense: %d, Experience: %d",
		prefix, character.Name, character.Health, character.Attack, character.Defense, character.Experience)
}

func RenderEnemyAppears(enemy models.Character) string {
	return fmt.Sprintf("A wild %s appears! Health: %d, Attack: %d, Defense: %d",
		enemy.Name, enemy.Health, enemy.Attack, enemy.Defense)
}

func RenderAttack(attacker, defender string, damage int) string {
	return fmt.Sprintf("%s deals %d damage to %s", attacker, damage, defender)
}

func RenderVictory(enemyName string, expGain int) string {
	return fmt.Sprintf("You defeated the %s! Gained %d experience.", enemyName, expGain)
}

func RenderDefeat() string {
	return "You were defeated! Your character has been reset."
}

func RenderBattleStatus(player, enemy models.Character) string {
	return fmt.Sprintf("%s Health: %d | %s Health: %d", player.Name, player.Health, enemy.Name, enemy.Health)
}
