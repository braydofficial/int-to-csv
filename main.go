package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Bitte hier angeben wie viele Etiketten erstellt werden sollen. (250 erstellt 500 Etiketten. Jeweils 2 Stück bis zur Zahl 250.")
	s := ""
	fmt.Scanln(&s)
	count, _ := strconv.Atoi(s)
	upCount(count)
}

func upCount(count int) {
	f, e := os.Create("./Output.csv")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	_, err2 := f.WriteString("SN\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	for i := 1; i < count+1; i++ {
		tag := strconv.Itoa(i)
		_, err2 := f.WriteString(tag + "\n" + tag + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
