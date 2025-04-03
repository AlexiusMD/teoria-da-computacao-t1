package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseDecimalToBinary(number uint64) string {
	return strconv.FormatUint(number, 2)
}

func countStates(scanner *bufio.Scanner, states *map[string]string, stateOrder *[]string) uint8 {
	var stateCount uint8

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		line = strings.TrimSpace(line)

		(*states)[line] = strconv.Itoa(int(stateCount))
		*stateOrder = append(*stateOrder, line)
		stateCount++
	}

	return stateCount
}

func convertStatesIntoBinCode(stateCount uint8, states *map[string]string) {
	if stateCount == 0 {
		return
	}

	bitsNeeded := 0
	maxValue := stateCount - 1

	for maxValue > 0 {
		bitsNeeded++
		maxValue >>= 1
	}

	if bitsNeeded == 0 && stateCount > 0 {
		bitsNeeded = 1
	}

	for key, value := range *states {
		stateNum, _ := strconv.Atoi(value)
		binCode := parseDecimalToBinary(uint64(stateNum))

		for len(binCode) < bitsNeeded {
			binCode = "0" + binCode
		}

		(*states)[key] = binCode
	}
}

func moveScannerToNextHeader(scanner *bufio.Scanner, header string) error {
	for scanner.Scan() {
		linha := scanner.Text()

		if strings.Split(linha, " ")[0] == "--" {
			continue
		}

		if linha == header {
			break
		}
	}

	if scanner.Text() != header {
		return fmt.Errorf("cabeçalho '%s' não encontrado no arquivo", header)
	}

	return nil
}

func encondeTransitionsIntoNonPureBinaryInstructions(scanner *bufio.Scanner, transitions *[]string, states *map[string]string) {
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(strings.TrimSpace(line), " ")

		if line == "" {
			break
		}

		if splitLine[0] == "--" {
			continue
		}

		var currentTransition string
		directionBit := "0"

		if splitLine[4] == ">" {
			directionBit = "1"
		}

		currentTransition = (*states)[splitLine[0]] + splitLine[1] + (*states)[splitLine[2]] + splitLine[3] + directionBit

		*transitions = append(*transitions, currentTransition)
	}
}

func createFinalEncodedTuringMachineString(transitions *[]string) string {
	encodedMachine := "$000*"

	for _, v := range *transitions {
		encodedMachine += strings.ReplaceAll(v, "*", " ") + "*"
	}

	return encodedMachine + "$"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run parser.go <caminho-do-arquivo>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %s\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	states := make(map[string]string)
	orderedStates := make([]string, 0)

	err = moveScannerToNextHeader(scanner, "states:")
	if err != nil {
		fmt.Printf("Erro: %s\n", err)
		os.Exit(1)
	}

	stateCount := countStates(scanner, &states, &orderedStates)

	convertStatesIntoBinCode(stateCount, &states)

	err = moveScannerToNextHeader(scanner, "transitions:")
	if err != nil {
		fmt.Printf("Erro: %s\n", err)
		os.Exit(1)
	}

	transitions := make([]string, 0)

	encondeTransitionsIntoNonPureBinaryInstructions(scanner, &transitions, &states)

	encodedMachine := createFinalEncodedTuringMachineString(&transitions)

	fmt.Println(encodedMachine)

	// for i, v := range transitions {
	// 	fmt.Println("Transition " + strconv.Itoa(i) + " is: " + v)
	// }

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %s\n", err)
		os.Exit(1)
	}

}
