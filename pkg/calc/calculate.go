package calc

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Calculate(input, output string) ([]byte, error) {

	in, err := os.OpenFile(input, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(in)

	filerReader := bufio.NewReader(in)

	_ = os.Remove(output)
	out, err := os.OpenFile(output, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)

	for {
		line, _, err := filerReader.ReadLine()
		if err != nil {
			break
		}
		mathRegex := regexp.MustCompile("[0-9]+[-+*/][0-9]+")
		isMath := mathRegex.MatchString(string(line))

		if isMath {
			result := calc(line)

			writer := bufio.NewWriter(out)
			_, err := writer.WriteString(string(result))
			if err != nil {
				return nil, nil
			}
			_, err = writer.WriteString("\n")
			if err != nil {
				return nil, nil
			}
			err = writer.Flush()
			if err != nil {
				return nil, nil
			}
		}
	}

	// метод для чтения файлов
	content, err := ioutil.ReadFile(output)
	if err != nil {
		panic(err)
	}

	return content, err
}

// Калькулятор
func calc(expr []byte) (result []byte) {
	numRegex := regexp.MustCompile("[0-9]+")
	nums := numRegex.FindAllString(string(expr), -1)

	signRegex := regexp.MustCompile("[-+*/]")
	operations := signRegex.FindAllString(string(expr), -1)

	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	var equally int

	if operations[0] == "+" {
		equally = num1 + num2
	} else if operations[0] == "-" {
		equally = num1 - num2
	} else if operations[0] == "*" {
		equally = num1 * num2
	} else if operations[0] == "/" {
		equally = num1 / num2
	}

	finNum := strconv.Itoa(equally)

	var resString []string
	resString = append(resString, nums[0], operations[0], nums[1], "=", finNum)

	stringByte := strings.Join(resString, "\x20")

	result = []byte(stringByte)

	return result
}
