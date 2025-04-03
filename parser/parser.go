package parser

import (
	"fmt"
	"os"
)

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

	scanner := createFileBuffer(file)

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

	fmt.Printf(encodedMachine)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %s\n", err)
		os.Exit(1)
	}

}
