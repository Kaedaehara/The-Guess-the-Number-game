package main

import (
	"TheGameGuessTheNumber/gamelogic"
	"TheGameGuessTheNumber/storage"
	"bufio"
	"fmt"
)

type Game struct {
	reader            *bufio.Reader
	maxNumber         int
	remainingAttempts int
	secretNumber      int
	pastAttempts      []int
}

func (g *Game) Run() {
	for {
		g.pastAttempts = nil

		g.maxNumber, g.remainingAttempts = g.chooseDifficulty()

		g.secretNumber = gamelogic.GenerateSecret(g.maxNumber)

		fmt.Printf("Игра 'Угадай число' - от 1 до %d началась!💡\n", g.maxNumber)
		fmt.Printf("Угадайте число за %s%d попыток!%s😏\n", gamelogic.ColorYellow, g.remainingAttempts, gamelogic.ColorReset)

		won := false

		for g.remainingAttempts > 0 {

			isLastTry := g.remainingAttempts == 1

			userGuess := g.readGuess(g.maxNumber)

			g.pastAttempts = append(g.pastAttempts, userGuess)

			resultCompare := gamelogic.CompareGuess(userGuess, g.secretNumber)

			if resultCompare == gamelogic.Equal {
				fmt.Println(gamelogic.ColorGreen + "Вы угадали!🙌\nИгра закончена!" + gamelogic.ColorReset)
				won = true
				break
			}

			if !isLastTry {
				switch resultCompare {
				case gamelogic.Greater:
					fmt.Println("Секретное число меньше👇")
				case gamelogic.Less:
					fmt.Println("Секретное число больше👆")
				}

				gamelogic.PrintHint(userGuess, g.secretNumber, g.remainingAttempts, g.pastAttempts)
			}

			g.remainingAttempts--

			if g.remainingAttempts != 0 {
				fmt.Printf("Осталось попыток: %s%d%s\n", gamelogic.ColorYellow, g.remainingAttempts, gamelogic.ColorReset)
			}
		}

		if !won {
			fmt.Printf(gamelogic.ColorRed+"Вы проиграли!😢\nСекретное число было: %d\n"+gamelogic.ColorReset, g.secretNumber)
		}

		result := gamelogic.MakeGameResult(won, len(g.pastAttempts))

		if err := storage.SaveGameResult("results.json", result); err != nil {
			fmt.Println("Не смог сохранить результат:", err)
		}

		if !g.askPlayAgain() {
			return
		}

	}

}
