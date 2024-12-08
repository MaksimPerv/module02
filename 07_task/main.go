package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("07_task_in.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	i := 0
	outfile, _ := os.Create("07_task_out.txt")
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)

	scanner := bufio.NewScanner(file)

	defer func() (err error) {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case fmt.Sprintf("parse error: empty field on string %d", i):
				writer.Flush()
				content, _ := os.ReadFile("07_task_out.txt")
				fmt.Println(string(content))
				err = fmt.Errorf("parse error: empty field on string %d", i)
			default:
				panic("critical")
			}

		}
		return
	}()

	for scanner.Scan() {
		i++
		fields := strings.Split(scanner.Text(), "|")
		if len(fields) != 3 {
			panic(fmt.Sprintf("parse error: empty field on string %d", i))
		}
		for _, field := range fields {
			if strings.TrimSpace(field) == "" {
				panic(fmt.Sprintf("parse error: empty field on string %d", i))
			}
		}
		writer.WriteString(fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nIn City: %s\n\n", i, fields[0], fields[1], fields[2]))

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	writer.Flush()
}
