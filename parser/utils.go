package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createFileBuffer(file *os.File) *bufio.Scanner {
	return bufio.NewScanner(file)
}
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
