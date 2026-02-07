package gamelogic

import (
	"TheGameGuessTheNumber/storage"
	"fmt"
	"math/rand"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

type CompareResult int

const (
	Equal CompareResult = iota
	Less
	Greater
)

func CompareGuess(guess, secret int) CompareResult {
	if guess == secret {
		return Equal
	}
	if guess > secret {
		return Greater
	}
	return Less
}

func MakeGameResult(won bool, attempts int) storage.GameResult {
	outcome := "Проигрыш"
	if won {
		outcome = "Выигрыш"
	}
	return storage.GameResult{
		Date:     time.Now().Format("02.01.2006 15:04:05"),
		Outcome:  outcome,
		Attempts: attempts,
	}
}

func GenerateSecret(max int) int {
	return rand.Intn(max) + 1
}

func PrintHint(guess, secret, remaining int, past []int) {
	diff := guess - secret
	if diff < 0 {
		diff = diff * -1
	}

	if remaining > 1 {
		switch {
		case diff <= 5:
			fmt.Println("🔥  Горячо - ты почти угадал! 🔥")

		case diff <= 15:
			fmt.Println("🙂  Тепло - ты движешься в верном направлении! 🙂")

		default:
			fmt.Println("❄️  Холодно - совсем далеко ❄️")
		}
	}

	fmt.Printf("Твои предыдущие попытки:%s%v%s\n", ColorYellow, past, ColorReset)
}
