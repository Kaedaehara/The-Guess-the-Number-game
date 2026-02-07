package main

import (
	"TheGameGuessTheNumber/gamelogic"
	"fmt"
	"strconv"
	"strings"
)

func (g *Game) readLine() string {
	s, err := g.reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(s)
}

func (g *Game) askPlayAgain() bool {
	fmt.Println("Хотите сыграть ещё раз?\nВведите \"да\" или \"нет\"")

	for {
		answer := strings.ToLower(g.readLine())

		switch answer {
		case "да", "д", "yes", "y":
			return true
		case "нет", "н", "no", "n":
			return false
		default:
			fmt.Println("Ошибка. Введите \"да\" или \"нет\"!")
		}
	}
}

func (g *Game) readGuess(max int) int {
	for {
		fmt.Print("Введите число: ")

		input, err := g.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода. Повторите.")
			continue
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Нужно ввести ОДНО целое число.")
			continue
		}

		if guess <= 0 {
			fmt.Println("Число должно быть больше нуля.")
			continue
		}

		if guess > max {
			fmt.Printf("Число должно быть не больше %d.\n", max)
			continue
		}

		return guess
	}
}

func (g *Game) chooseDifficulty() (int, int) {
	fmt.Println("Выберите уровень сложности:")
	fmt.Println(gamelogic.ColorGreen + "1 — Easy   (от 1 до 50, 15 попыток)" + gamelogic.ColorReset)
	fmt.Println(gamelogic.ColorYellow + "2 — Medium (от 1 до 100, 10 попыток)" + gamelogic.ColorReset)
	fmt.Println(gamelogic.ColorRed + "3 — Hard   (от 1 до 200, 5 попыток)" + gamelogic.ColorReset)
	fmt.Printf(
		"Введите %s1%s, %s2%s или %s3%s:\n",
		gamelogic.ColorGreen, gamelogic.ColorReset,
		gamelogic.ColorYellow, gamelogic.ColorReset,
		gamelogic.ColorRed, gamelogic.ColorReset,
	)

	for {
		diffMode := g.readLine()

		switch diffMode {
		case "1":
			return 50, 15
		case "2":
			return 100, 10
		case "3":
			return 200, 5
		default:
			fmt.Println("Некорректный ввод. Повторите ещё раз.")
		}
	}
}
