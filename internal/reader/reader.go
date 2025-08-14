package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return []int{0}
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Ошибка закрытия файла: %v\n", err)
		}
	}()

	var arrayPhone []int

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("Ошибка чтения файла: %v\n", err)
			}
			break
		}
		digitPhone, err := strconv.Atoi(string(line[1:12]))
		if err != nil {
			log.Fatalf("Траблы: %v", err)
		}
		arrayPhone = append(arrayPhone, digitPhone)
	}
	return arrayPhone
}
