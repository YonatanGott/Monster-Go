package interaction

import (
	"fmt"
	"os"
)

type RoundData struct {
	Action     string
	PlayerDmg  int
	PlayerHeal int
	MonsterDmg int
	PlayerHp   int
	MonsterHp  int
}

func PrintGreeting() {
	fmt.Println("MONSTER HUNTER")
	fmt.Println("Starting a new Game...")
}

func ShowActions(specialAttack bool) {
	fmt.Println("Choose Your Action")
	fmt.Println("------------------")
	fmt.Println("(1) Attack")
	fmt.Println("(2) Heal")

	if specialAttack {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatus(data *RoundData) {
	if data.Action == "ATK" {
		fmt.Printf("Player Attacked for %v damage. \n", data.PlayerDmg)
	} else if data.Action == "SA" {
		fmt.Printf("Player Attacked for %v special damage. \n", data.PlayerDmg)
	} else {
		fmt.Printf("Player Healed for %v damage. \n", data.PlayerHeal)
	}

	fmt.Printf("Monster Attacked for %v damage. \n", data.MonsterDmg)
	fmt.Printf("Player Health: %v. \n", data.PlayerHp)
	fmt.Printf("Monster Health: %v. \n", data.MonsterHp)
}

func DeclareWinner(winner string) {
	fmt.Println("------------------")
	fmt.Println("Game Over")
	fmt.Println("------------------")
	fmt.Printf("%v Won! \n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file, err := os.Create("game-log.txt")
	if err != nil {
		fmt.Println("Saving log file failed")
		return
	}

	for i, v := range *rounds {
		logEntry := map[string]string{
			"Round":          fmt.Sprint(i + 1),
			"Action":         v.Action,
			"Player Damage":  fmt.Sprint(v.PlayerDmg),
			"Monster Damage": fmt.Sprint(v.MonsterDmg),
			"Player Heal":    fmt.Sprint(v.PlayerHeal),
			"Player Health":  fmt.Sprint(v.PlayerHp),
			"Monster Health": fmt.Sprint(v.MonsterHp),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)
		if err != nil {
			fmt.Println("Writing log file failed")
			continue
		}
	}
	file.Close()
	fmt.Println("File Saved")
}
