package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// printTriangle(1)
	// fmt.Println(decideFilm(2))
	fmt.Println(generatePassword("fazztrack", "stronga"))
}

func printTriangle(num int) {
	if num >= 2 {
		for i := num; i > 0; i-- {
			for space := 0; space < num-i; space++ {
				fmt.Print(" ")
			}
			for j := (i * 2) - 1; j > 0; j-- {
				fmt.Print("*")
			}
			fmt.Println("")
		}
	} else {
		fmt.Println("Masukkan angka lebih dari 1 untuk mencetak segitiga")
	}
}

func decideFilm(duration int) string {
	data := [...]int{1, 7, 3, 4, 8, 9, 3}
	var result string
	for i, firstFilm := range data {
		for j, secondFilm := range data {
			if i != j && firstFilm+secondFilm == duration {
				result = fmt.Sprintf(`Rekomendasi Film dengan durasi penerbangan %d jam
Film ke-%d dengan durasi %d jam dan film ke-%d dengan durasi %d jam`, duration, i+1, firstFilm, j+1, secondFilm)
				return result
			} else if i == len(data)-1 && secondFilm == duration {
				result = fmt.Sprintf(`Rekomendasi Film dengan durasi penerbangan %d jam
Film ke-%d dengan durasi %d jam`, duration, j+1, secondFilm)
				return result
			}
		}
	}
	result = fmt.Sprintf("Tidak ada rekomendasi film yang pas untuk penerbangan dengan durasi %d jam", duration)
	return result
}

func generatePassword(password string, level string) string {
	rand.Seed(time.Now().Unix())
	const (
		charSet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numberSet = "1234567890"
		symbolSet = "~!@#$%^&*()_=-+,>.<?/[}{]"
	)
	var result string
	salt := len(password) / 3
	addSalt := func(source string) string {
		var saltResult string
		for i := 0; i < salt; i++ {
			saltResult += string(source[rand.Intn(len(source))])
		}
		return saltResult
	}

	levelOne := func(password string) string {
		if len(password) < 6 {
			salt = 6 - len(password)
		}
		slice := make([]string, len(password))
		for i, char := range password {
			if rand.Int()%2 == 0 {
				slice[i] = strings.ToUpper(string(char))
			} else {
				slice[i] = strings.ToLower(string(char))
			}
			result = strings.Join(slice, "")
		}
		result += addSalt(charSet)
		return result
	}

	levelTwo := func(password string) string {
		result += addSalt(numberSet)
		return result
	}

	levelThree := func(password string) string {
		result += addSalt(symbolSet)
		return result
	}

	switch strings.ToLower(level) {
	case "low":
		levelOne(password)
	case "med":
		levelTwo(levelOne(password))
	case "strong":
		levelThree(levelTwo(levelOne(password)))
	default:
		return "Level tidak tersedia. (Opsi level: low, med, strong)"
	}

	// finalResult := []rune(result)
	// rand.Shuffle(len(finalResult), func(i, j int) {
	// 	finalResult[i], finalResult[j] = finalResult[j], finalResult[i]
	// })
	return string(result)
}
