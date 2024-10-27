package ascii

import (

	// "piscine/handlers"

	"os"
	"strings"
)

// function converter to map the ascii characters with the file ascii arts
func Converter(file string) map[rune][]string {
	res := make(map[rune][]string)
	asciiChar := ' '
	data, err := os.ReadFile(file)
	if err != nil {
		// handler.NotFoundHandler
		return res
	}
	split := []string{}

	if file == "thinkertoy.txt" {
		split = strings.Split(string(data), "\r\n")
	} else {
		split = strings.Split(string(data), "\n")
	}

	for i := 1; i < len(split); {
		arrays := []string{}
		for j := 0; j < 8; j++ {
			arrays = append(arrays, split[i])
			i++
		}
		i++
		res[asciiChar] = arrays
		asciiChar++
	}
	return res
}

// function to print the ascii art
func PrintAsci(file, input string) string {
	res := ""
	input = strings.ReplaceAll(input, "\n", "\\n")
	lines := strings.Split(input, "\\n")
	ascii := Converter(file)
	if len(ascii) == 0 {
		return "nil"
	}
	count := 0
	for _, line := range lines {
		if line == "" {
			count++
			if count < len(lines) {
				res += "\n"
			}
		} else {
			for c := 0; c < 8; c++ {
				for _, char := range line {
					if asciiArt, ok := ascii[char]; ok {
						res += asciiArt[c]
					} else {
						res += " "
					}
				}
				res += "\n"
			}
		}
	}
	return res
}
