package actions

import (
	"math/rand"
	"time"
)

var (
	randSrc = rand.NewSource(time.Now().UnixNano())
	randGen = rand.New(randSrc)
)

var (
	currentMonsterHP = MONSTER_HP
	currentPlayerHP  = PLAYER_HP
)

func AttackMonster(specialAttack bool) int {
	minAttack := PLAYER_MIN_ATK
	maxAttack := PLAYER_MAX_ATK

	if specialAttack {
		minAttack = PLAYER_MIN_ATK_SPECAIL
		maxAttack = PLAYER_MAX_ATK_SPECAIL
	}

	dmgValue := random(minAttack, maxAttack)
	currentMonsterHP -= dmgValue

	return dmgValue
}

func HealPlayer() int {
	healValue := random(PLAYER_MIN_HEAL, PLAYER_MAX_HEAL)

	hpDiff := PLAYER_HP - currentPlayerHP
	if hpDiff >= healValue {
		currentPlayerHP += healValue
		return healValue
	} else {
		currentPlayerHP = PLAYER_HP
		return hpDiff
	}
}

func AttackPlayer() int {
	minAttack := MONSTER_MIN_ATK
	maxAttack := MONSTER_MAX_ATK

	dmgValue := random(minAttack, maxAttack)
	currentPlayerHP -= dmgValue

	return dmgValue
}

func GetHP() (int, int) {
	return currentPlayerHP, currentMonsterHP
}

func random(min int, max int) int {
	return randGen.Intn(max-min) + min
}
