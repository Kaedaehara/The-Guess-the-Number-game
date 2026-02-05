package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type GameResult struct {
	Date     string `json:"Дата"`
	Outcome  string `json:"Исход"`
	Attempts int    `json:"Количество затраченных попыток"`
}

type CompareResult int

const (
	Equal CompareResult = iota
	Less
	Greater
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

func main() {

gameLoop:
	for {
		pastAttempts := []int{}

		maxNumber, remainingAttempts := chooseDifficulty()

		secretNumber := generateSecret(maxNumber)

		fmt.Printf("Игра 'Угадай число' - от 1 до %d началась!💡\n", maxNumber)
		fmt.Printf("Угадайте число за %s%d попыток!%s😏\n", ColorYellow, remainingAttempts, ColorReset)

		won := false

		for remainingAttempts > 0 { // Цикл угадывания

			userGuess := readGuess(maxNumber) // Ввод и валидация числа

			pastAttempts = append(pastAttempts, userGuess) // Добавляем предыдущую попытку пользователя

			resultCompare := compareGuess(userGuess, secretNumber) // Сравнение числа пользователя и загаданного

			switch resultCompare {
			case Equal:
				fmt.Println(ColorGreen + "Вы угадали!🙌\nИгра закончена!" + ColorReset)
				won = true
			case Greater:
				fmt.Println("Секретное число меньше👇")
			case Less:
				fmt.Println("Секретное число больше👆")
			}

			if won {
				break
			}

			printHint(userGuess, secretNumber, remainingAttempts, pastAttempts)

			remainingAttempts--

			fmt.Printf("Осталось попыток: %s%d%s\n", ColorYellow, remainingAttempts, ColorReset)
		}

		if !won {
			fmt.Printf(ColorRed+"Вы проиграли!😢\nСекретное число было: %d\n"+ColorReset, secretNumber)
		}

		result := makeGameResult(won, len(pastAttempts))

		if err := saveGameResult("results.json", result); err != nil {
			fmt.Println("Не смог сохранить результат:", err)
		}

		if askPlayAgain() {
			continue gameLoop
		}
		return
	}
}

func saveGameResult(filename string, result GameResult) error {
	// 1) читаем файл (если нет — считаем что результатов пока нет) json в слайс байтов
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			data = []byte("[]") // пустой список в JSON
		} else {
			return err
		}
	}

	if strings.TrimSpace(string(data)) == "" {
		data = []byte("[]")
	}

	// 2) превращаем JSON-текст в слайс структур
	var results []GameResult
	if err := json.Unmarshal(data, &results); err != nil {
		return err
	}

	// 3) добавляем новый результат
	results = append(results, result)

	// 4) превращаем обратно в JSON-текст
	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	// 5) пишем файл обратно
	return os.WriteFile(filename, out, 0644)
}

func readGuess(max int) int {
	for {
		var input string
		fmt.Scanln(&input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Нужно ввести число. Повторите попытку!")
			continue
		}
		if guess <= 0 {
			fmt.Println("Число должно быть больше нуля. Повторите попытку!")
			continue
		}
		if guess > max {
			fmt.Printf("Число должно быть не больше %d. Повторите попытку!\n", max)
			continue
		}
		return guess
	}
}

func printHint(guess, secret, remaining int, past []int) {
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

func compareGuess(guess, secret int) CompareResult {
	if guess == secret {
		return Equal
	}
	if guess > secret {
		return Greater
	}
	return Less
}

func generateSecret(max int) int {
	return rand.Intn(max) + 1
}

func chooseDifficulty() (int, int) {
	fmt.Println("Выберите уровень сложности:")
	fmt.Println(ColorGreen + "1 — Easy   (от 1 до 50, 15 попыток)" + ColorReset)
	fmt.Println(ColorYellow + "2 — Medium (от 1 до 100, 10 попыток)" + ColorReset)
	fmt.Println(ColorRed + "3 — Hard   (от 1 до 200, 5 попыток)" + ColorReset)
	fmt.Printf(
		"Введите %s1%s, %s2%s или %s3%s:\n",
		ColorGreen, ColorReset,
		ColorYellow, ColorReset,
		ColorRed, ColorReset,
	)

	for {
		var diffMode string
		fmt.Scanln(&diffMode)

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

func askPlayAgain() bool {
	fmt.Println("Хотите сыграть ещё раз?\nВведите \"да\" или \"нет\"")

	for {
		var answer string
		fmt.Scanln(&answer)

		switch strings.ToLower(strings.TrimSpace(answer)) {
		case "да", "д", "yes", "y":
			return true
		case "нет", "н", "no", "n":
			return false
		default:
			fmt.Println("Ошибка. Введите \"да\" или \"нет\"!")
		}
	}
}

func makeGameResult(won bool, attempts int) GameResult {
	outcome := "Проигрыш"
	if won {
		outcome = "Выигрыш"
	}
	return GameResult{
		Date:     time.Now().Format("02.01.2006 15:04:05"),
		Outcome:  outcome,
		Attempts: attempts,
	}
}
