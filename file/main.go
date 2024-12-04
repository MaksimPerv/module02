package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func logTime(start time.Time) {

	duration := time.Since(start)

	fmt.Printf("Время выполнения программы: %v\n", duration)
}

func main() {
	start := time.Now()
	defer logTime(start)
	file, err := os.Open("06_task_in.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}

	outfile, err := os.Create("task_out.txt")
	writer := bufio.NewWriter(outfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i := 0
	cnt := 0
	defer func() {
		fmt.Printf("Строк было записано %d. Байтов было записано %d\n", i, cnt)
	}()
	defer outfile.Close()
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		i++
		count, _ := writer.WriteString(fmt.Sprintf("%d. ", i) + line)
		cnt += count
	}
	writer.Flush()
}
