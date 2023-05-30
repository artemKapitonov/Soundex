package service

import (
	"strconv"
	"strings"

	"github.com/artemKapitonov/soundex/internal/models"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

var sound map[int][]rune = map[int][]rune{
	1: {'b', 'f', 'p', 'v'},
	2: {'c', 'g', 'j', 'k', 'q', 's', 'x', 'z'},
	3: {'d', 't'},
	4: {'l'},
	5: {'m', 'n'},
	6: {'r'},
}

var ignoreSymbols []string = []string{"a", "e", "i", "o", "u", "y"}

func runeInSlice(slice []rune, sy rune) bool {
	for _, elem := range slice {
		if elem == sy {
			return true
		}
	}
	return false
}

func SoundexName(name string) string {
	var resultName string

RowLoops:
	for indx, i := range name {
		if indx == 0 {
			resultName += strings.Title(string(i))
			continue
		}

		if i == 'h' || i == 'w' {
			continue
		}

		for key, val := range sound {
			if runeInSlice(val, i) {
				resultName += strconv.Itoa(key)
				continue RowLoops
			}
		}

		resultName += string(i)
	}

	resultName = DeleteDublicate([]byte(resultName))
	resultName = DeleteIgnoreSymbols(resultName)

	return resultName
}

func DeleteDublicate(name []byte) string {

	for indx := 1; indx < len(name)-1; indx++ {
		if name[indx] == name[indx+1] {
			name = append(name[:indx], name[indx+1:]...)
		}
	}

	return string(name)
}

func DeleteIgnoreSymbols(name string) string {
	var result = name

	for _, i := range ignoreSymbols {
		result = strings.Replace(result, i, "", 100)
	}

	return result
}

func (s *Service) Soundex(names models.Names) []string {

	var result []string

	text := strings.Split(names.Names, ", ")

	for _, name := range text {

		soundex := SoundexName(name)

		if len(soundex) < 4 {
			soundex += "0000"
		}

		result = append(result, soundex[:4])
	}

	return result
}
