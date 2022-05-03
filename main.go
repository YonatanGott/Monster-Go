package main

import (
	"github.com/YonatanGott/Monster-Go/interaction"

	"github.com/YonatanGott/Monster-Go/actions"
)

var (
	currentRound = 0
	gameData     = []interaction.RoundData{}
)

func main() {
	startGame()

	winner := "" // "Player" || "Monster"

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowActions(isSpecialRound)

	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerDmg int
	var monsterDmg int
	var playerHeal int

	if userChoice == "ATK" {
		playerDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHeal = actions.HealPlayer()
	} else {
		playerDmg = actions.AttackMonster(true)
	}

	monsterDmg = actions.AttackPlayer()

	playerHP, monsterHP := actions.GetHP()

	data := interaction.RoundData{
		Action:     userChoice,
		PlayerHp:   playerHP,
		MonsterHp:  monsterHP,
		PlayerDmg:  playerDmg,
		PlayerHeal: playerHeal,
		MonsterDmg: monsterDmg,
	}
	interaction.PrintRoundStatus(&data)
	gameData = append(gameData, data)

	if playerHP <= 0 {
		return "Monster"
	} else if monsterHP <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameData)
}
