package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttack bool) string {
	for {
		playerChoice, _ := getPlayerInput()
		if playerChoice == "1" {
			return "ATK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && specialAttack {
			return "SA"
		}
		fmt.Println("User input failed")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.TrimSpace(userInput)
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, nil
}
