package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func hex(str string) string {
	values, _ := strconv.ParseInt(str, 16, 64)
	return strconv.Itoa(int(values))
}

func bin(str string) string {
	values, _ := strconv.ParseInt(str, 2, 64)
	return strconv.Itoa(int(values))
}

func up(str string) string {
	return strings.ToUpper(str)
}

func low(str string) string {
	return strings.ToLower(str)
}

func cap(str string) string {
	return strings.Title(str)
}

func Punctuations(value string) string {
	empty := []string{",", ".", "!", "?", ";", ":", "'", "\"", "(", ")", "-"}

	for _, k := range empty {
		value = strings.Replace(value, " "+k, k, -1)
		value = strings.Replace(value, k+" ", k, -1)
	}

	r := []rune(value)
	arr := []rune{}
	for i := 0; i < len(r)-1; i++ {
		if unicode.IsPunct(r[i]) && unicode.IsLetter(r[i+1]) || unicode.IsPunct(r[i]) && r[i+1] == ' ' {
			if r[i] == '\'' && unicode.IsLetter(r[i+1]) || r[i] == '-' && unicode.IsLetter(r[i+1]) || r[i] == '"' {
				arr = append(arr, r[i])
			} else {
				arr = append(arr, r[i], ' ')
			}
		} else if unicode.IsPunct(r[i]) && r[i+1] == ' ' && r[i+2] == '\'' { //
			arr = append(arr, r[i], ' ')
		} else {
			arr = append(arr, r[i])
		}
	}
	arr = append(arr, r[len(r)-1])

	return string(arr)
}

func ChangeA(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

func FixAgain(res string) string {
	actualRes := ""
	quotesCount := 1
	for i := 0; i < len(res); i++ {
		if (string(res[i]) == "." || string(res[i]) == "," || string(res[i]) == "!" || string(res[i]) == "?" || string(res[i]) == ":" || string(res[i]) == ";") && i != len(res)-1 {
			if string(res[i+1]) == "'" && quotesCount%2 == 1 {
				actualRes += string(res[i]) + " "
				quotesCount++
			} else {
				actualRes += string(res[i])
			}
		} else {
			actualRes += string(res[i])
		}
	}
	return actualRes
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Dosya açma hatası:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	resultFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Dosya oluşturma hatası:", err)
		return
	}
	defer resultFile.Close()

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		words = ChangeA(words)

		var output string

		for i := len(words) - 1; i >= 0; i-- {
			if words[i] == "(up)" {
				words[i-1] = up(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(up," {
				nStr := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = up(words[i-n+a-1])
				}
				i--

				continue
			} else if words[i] == "(low)" {
				words[i-1] = low(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(low," {
				nStr := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = low(words[i-n+a-1])
				}
				i--
				continue
			} else if words[i] == "(cap)" {
				words[i-1] = cap(words[i-1])
				i--

			} else if strings.Contains(words[i], ")") && words[i-1] == "(cap," {
				nStr := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = cap(words[i-n+a-1])
				}
				i--
				continue
			} else if words[i] == "(hex)" {
				words[i-1] = hex(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(hex," {
				nStr := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = hex(words[i-n+a-1])
				}
				i--
				continue
			} else if words[i] == "(bin)" {
				words[i-1] = bin(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(bin," {
				nStr := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(nStr)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = bin(words[i-n+a-1])
				}
				i--
				continue
			}
			output = words[i] + " " + output
		}

		line = strings.Join(words, " ")
		line = Punctuations(output)

		resultFile, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			fmt.Println("Dosya yazma hatası:", err)
			return
		}
		defer resultFile.Close()

		_, err = resultFile.WriteString(FixAgain(line))
		if err != nil {
			fmt.Println("Dosya yazma hatası:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Dosya okuma hatası:", err)
		return
	}
}
