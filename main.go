package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	finishdata := Modify(string((data)))
	ioutil.WriteFile("result.txt", finishdata, 0)
	file, err := os.Open("result.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	w.WriteString(string(finishdata))
	w.Flush()
}

func Modify(txt string) []byte {
	txt = Vowels(txt)
	txt = Capitalize(txt)
	txt = UpLow(txt)
	txt = UpLowNb(txt)
	txt = HexBin(txt)
	txt = Ponctu(txt)
	txt = Guillemets(txt)
	finish := []byte(txt)
	return finish
}

func Ponctu(str string) string {
	tab := []rune(str)
	for i := 1; i < len(tab); i++ {
		if i != len(tab)-1 {
			if tab[i-1] == ' ' && (tab[i] == '.' || tab[i] == ',' || tab[i] == '!' || tab[i] == '?' || tab[i] == ':' || tab[i] == ';') {
				tab[i-1], tab[i] = tab[i], tab[i-1]
				if tab[i] == tab[i+1] {
					tab = append(tab[:i], tab[i+1:]...)
				}
			} else if tab[i] == ' ' && (tab[i+1] == '.' || tab[i+1] == ',' || tab[i+1] == '!' || tab[i+1] == '?' || tab[i+1] == ':' || tab[i+1] == ';') {
				tab[i+1], tab[i] = tab[i], tab[i+1]
				if tab[i] == tab[i+1] {
					tab = append(tab[:i], tab[i+1:]...)
				}
			} else if tab[i] == ' ' && tab[i+1] == ' ' {
				tab = append(tab[:i], tab[i+1:]...)
			}
		}
	}
	return string(tab)
}

// func Guillemets(s string) string {
// 	res := ""
// 	var tab []rune
// 	var quote bool
// 	for i := 0; i < len(tab); i++ {
// 		cara := tab[i]
// 		if cara == '\'' && !quote {
// 			quote = true
// 			if tab[i+1] == ' ' {
// 				i++
// 			}
// 		} else if cara == '\'' && quote {
// 			quote = false
// 			i++
// 		}
// 		res += string(cara)
// 	}
// 	return res
// }

func Guillemets(s string) string {
	new := strings.Replace(s, "‘", "'", -1)
	tab := []rune(new)
	res := ""
	quote := false
	for i := 0; i < len(tab); i++ {
		cara := tab[i]
		if cara == '\'' && !quote {
			quote = true
			if tab[i+1] == ' ' {
				i++
			}
		} else if (cara == ' ' && tab[i+1] == '\'') && quote {
			quote = false
			cara = '\''
			i++
		}
		res += string(cara)
	}

	return res
}

func Capitalize(str string) string {
	res := strings.Split(str, " ")
	for i := 0; i < len(res)-1; i++ {
		if res[i] == "(cap)" {
			res[i-1] = strings.Title(res[i-1])
			res = append(res[:i], res[i+1:]...)
		} else if res[i] == "(cap," {
			nb := res[i+1][:len(res[i+1])-1]
			x, e := strconv.Atoi(string(nb))

			if e == nil {
			}
			for j := x; j > 0; j-- {
				res[j+i-(x+1)] = strings.Title(res[j+i-(x+1)])
			}
			res = append(res[:i], res[(i+2):]...)
		}
	}
	s := strings.Join(res, " ")
	return s
}

func UpLow(str string) string {
	res := strings.Split(str, " ")
	for i := 0; i < len(res)-1; i++ {
		if res[i] == "(up)" {
			res[i-1] = strings.ToUpper(res[i-1])
			res = append(res[:i], res[i+1:]...)
		} else if res[i] == "(low)" {
			res[i-1] = strings.ToLower(res[i-1])
			res = append(res[:i], res[i+1:]...)
		}
	}
	s := strings.Join(res, " ")
	return s
}

func UpLowNb(str string) string {
	res := strings.Split(str, " ")
	for i := 0; i < len(res)-1; i++ {
		if res[i] == "(low," {
			nb := res[i+1][:len(res[i+1])-1]
			x, e := strconv.Atoi(string(nb))

			if e == nil {
			}
			for j := x; j > 0; j-- {
				res[j+i-(x+1)] = strings.ToLower(res[j+i-(x+1)])
			}
			res = append(res[:i], res[(i+2):]...)
		} else if res[i] == "(up," {
			nb := res[i+1][:len(res[i+1])-1]
			x, e := strconv.Atoi(string(nb))

			if e == nil {
			}
			for j := x; j > 0; j-- {
				res[j+i-(x+1)] = strings.ToUpper(res[j+i-(x+1)])
			}
			res = append(res[:i], res[(i+2):]...)
		}
	}
	s := strings.Join(res, " ")
	return s
}

func HexBin(s string) string {
	res := strings.Split(s, " ")
	for i := 0; i < len(res)-1; i++ {
		if res[i] == "(hex)" {
			// hexa := res[i-1]
			hexa, err := strconv.ParseInt(res[i-1], 16, 32)
			if err != nil {
				fmt.Println(err)
			}
			decimal := strconv.FormatInt(hexa, 10)
			res[i-1] = decimal
			res = append(res[:i], res[i+1:]...)
		} else if res[i] == "(bin)" {
			bin, err := strconv.ParseInt(res[i-1], 2, 32)
			if err != nil {
				fmt.Println(err)
			}
			dec := strconv.FormatInt(bin, 10)
			res[i-1] = dec
			res = append(res[:i], res[i+1:]...)
		}
	}
	s = strings.Join(res, " ")
	return s
}

func Vowels(s string) string {
	// result := strings.Map(r, rune)
	res := strings.Split(s, " ")
	for i := 0; i < len(res); i++ {
		if res[i] == "A" {
			if IsOk(string(res[i+1][0])) {
				res[i] = "An"
			}
		} else if res[i] == "a" {
			if IsOk(string(res[i+1][0])) {
				res[i] = "an"
			}
		}
	}

	s = strings.Join(res, " ")
	return s
}

func IsOk(next string) bool {
	if next == "a" || next == "e" || next == "i" || next == "o" || next == "u" || next == "h" ||
		next == "A" || next == "E" || next == "I" || next == "O" || next == "U" || next == "H" {
		return true
	}
	return false
}
