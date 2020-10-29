package num2words

import (
	"errors"
	"fmt"
	"strings"
)

const maxValue = 999999999999999

type Words []string

var (
	smallNumbers = [...]string{
		"nol", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh",
		"sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas",
		"delapan belas", "sembilan belas",
	}

	tensNumber = [...]string{
		"", "", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh",
		"delapan puluh", "sembilan puluh",
	}

	scales = [...]string{
		"", "ribu", "juta", "miliar", "triliun",
	}
)

// ConvertDecimal converts decimal number into the words representation
// reference: https://www.jawapos.com/opini/28/04/2019/membaca-desimal/
func ConvertDecimal(number float64, precision int) (string, error) {
	if precision < 1 {
		return Convert(int64(number))
	}

	format := "%." + fmt.Sprintf("%d", precision) + "f"
	fractions := fmt.Sprintf(format, number)
	fractions = strings.Split(fractions, ".")[1]

	words := ""
	for precision > 0 {
		v := fractions[precision-1] - 48
		if v > 0 {
			break
		}
		precision--
	}

	fractions = fractions[:precision]

	for i := 1; i <= precision; i++ {
		v := fractions[precision-i] - 48
		words = smallNumbers[v] + " " + words
	}

	if words == "" {
		return Convert(int64(number))
	}

	words = words[:len(words)-1]

	numString, err := Convert(int64(number))
	if err != nil {
		return "", err
	}

	return numString + " koma " + words, nil
}

// Convert converts number into the words representation
func Convert(number int64) (string, error) {

	if number > maxValue {
		msg := fmt.Sprintf("input parameters exceed the maximum limit. Max = %d", maxValue)
		err := errors.New(msg)
		return "", err
	}

	var isMinus bool = false

	if number == 0 {
		return smallNumbers[number], nil
	} else if number < 0 {
		isMinus = true
		number = number * -1
	}

	digitGroups := splitIntoThreeDigitGroups(number)

	var groupText []Words
	for _, digit := range digitGroups {
		groupText = append(groupText, threeDigitGroupsToWords(digit))
	}

	combinedWords := combineWords(groupText)
	if isMinus {
		combinedWords = append(Words{"negatif"}, combinedWords...)
	}

	return strings.Join(combinedWords, " "), nil
}

func splitIntoThreeDigitGroups(number int64) []uint16 {
	var digitGroups []uint16
	for i := 0; i < len(scales); i++ {
		digitGroups = append(digitGroups, uint16(number%1000))
		number = number / 1000
	}
	return digitGroups
}

func threeDigitGroupsToWords(number uint16) Words {
	var words Words

	hundreds := number / 100
	tensOfNumbers := number % 100
	if hundreds != 0 {
		if hundreds == 1 {
			words = append(words, "seratus")
		} else {
			words = append(words, smallNumbers[hundreds], "ratus")
		}
	}

	tens := tensOfNumbers / 10
	units := tensOfNumbers % 10
	if tens >= 2 {
		words = append(words, tensNumber[tens])
		if units != 0 {
			words = append(words, smallNumbers[units])
		}

	} else if tensOfNumbers != 0 {
		words = append(words, smallNumbers[tensOfNumbers])
	}

	return words
}

func combineWords(groupText []Words) Words {
	var words Words = groupText[0]
	for i := 1; i < len(scales); i++ {
		if strings.Join(groupText[i], "") != "" {
			if strings.Join(groupText[i], "") == "satu" && scales[i] == "ribu" {
				words = append(Words{"seribu"}, words...)
			} else {
				words = append(append(groupText[i], scales[i]), words...)
			}
		}
	}

	return words
}
