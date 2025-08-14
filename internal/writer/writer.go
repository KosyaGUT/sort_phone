package writer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WriterFile(arr []int) {
	file, err := os.OpenFile("configs/sortedPhone.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Unable to create/open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	rw := bufio.NewWriter(file)
	defer rw.Flush()

	for _, value := range arr {
		_, err = rw.WriteString("+" + strconv.Itoa(value) + "\n")
		if err != nil {
			fmt.Printf("Ошибка записи: %v\n", err)
			return
		}
	}
}
