package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func main() {
	LogFile()
	fmt.Println("Bitte hier angeben wie viele Etiketten erstellt werden sollen. (250 erstellt 500 Etiketten. Jeweils 2 Stück bis zur Zahl 250.")
	s := ""
	fmt.Scanln(&s)
	count, _ := strconv.Atoi(s)
	upCount(count)
}

func upCount(count int) {
	InfoLogger.Println("Generating CSV-File as Output.csv ...")
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
	InfoLogger.Println("Done. CSV-File exported as Output.csv")
}

func LogFile() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.Println("Starting the application...")
}
