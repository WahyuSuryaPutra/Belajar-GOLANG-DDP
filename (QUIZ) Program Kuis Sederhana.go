package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Question struct {
	Text    string
	Options []string
	Answer  string
}

func displayQuestion(q Question) {
	fmt.Println(q.Text)
	for i, option := range q.Options {
		fmt.Printf("%c. %s\n", 'A'+i, option)
	}
}

func checkAnswer(userAnswer string, correctAnswer string) bool {
	return userAnswer == correctAnswer
}

func main() {
	
	rand.Seed(time.Now().UnixNano())

	// Daftar pertanyaan
	questions := []Question{
		{
			Text:    "Apa ibukota Indonesia?",
			Options: []string{"Jakarta", "Surabaya", "Bandung", "Yogyakarta"},
			Answer:  "A",
		},
		{
			Text:    "Berapa jumlah planet di Tata Surya?",
			Options: []string{"7", "8", "9", "10"},
			Answer:  "B",
		},
		{
			Text:    "Siapakah presiden pertama Indonesia?",
			Options: []string{"Soekarno", "Soeharto", "Joko Widodo", "Megawati Soekarnoputri"},
			Answer:  "A",
		},
		{
			Text:    "Apa huruf pertama dalam alfabet?",
			Options: []string{"A", "B", "C", "D"},
			Answer:  "A",
		},
		{
			Text:    "Berapa hasil dari 5 + 3?",
			Options: []string{"7", "8", "9", "10"},
			Answer:  "B",
		},
		{
			Text:    "Apa warna langit pada siang hari?",
			Options: []string{"Merah", "Kuning", "Biru", "Hijau"},
			Answer:  "C",
		},
		{
			Text:    "Berapa jumlah kaki pada seekor anjing?",
			Options: []string{"2", "4", "6", "8"},
			Answer:  "B",
		},
		{
			Text:    "Apa nama benua terbesar di dunia?",
			Options: []string{"Asia", "Eropa", "Afrika", "Amerika"},
			Answer:  "A",
		},
		
	}

	// Mengacak urutan pertanyaan
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	
	for _, question := range questions {
		displayQuestion(question)

		
		var userAnswer string
		fmt.Print("Jawaban Anda (A/B/C/D): ")
		fmt.Scan(&userAnswer)

		
		if checkAnswer(userAnswer, question.Answer) {
			fmt.Println("Jawaban Anda benar!\n")
		} else {
			fmt.Printf("Jawaban Anda salah. Jawaban yang benar adalah %s.\n\n", question.Answer)
		}
	}
}