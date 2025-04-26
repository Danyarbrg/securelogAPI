package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	Text   string
	Answer string
}

func main() {
	questions := []Question{
		{"Самый большой океан?", "тихий"},
		{"Сколько планет в Солнечной системе?", "8"},
		{"Столица Франции?", "париж"},
	}

	score := 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🎉 Викторина! Ответь на вопросы (все ответы в нижнем регистре):")
	
	for _, q := range questions {
		fmt.Printf("\nВопрос: %s\n", q.Text)
		fmt.Print("Твой ответ: ")
		
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(strings.ToLower(userAnswer))
		
		if userAnswer == q.Answer {
			fmt.Println("✅ Верно!")
			score++
		} else {
			fmt.Printf("❌ Неверно! Правильный ответ: %s\n", q.Answer)
		}
	}

	fmt.Printf("\n🏁 Твой результат: %d/%d\n", score, len(questions))
	
	switch {
	case score == len(questions):
		fmt.Println("Ты гений! 🧠")
	case score >= len(questions)/2:
		fmt.Println("Неплохо! 👍")
	default:
		fmt.Println("Попробуй ещё раз! 😊")
	}
}
