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
	// Get variables for amount of serial numbers that should be created
	fmt.Println("\033[33mINFO: \033[34mPlease type the amount of serial numbers to be created\n\033[0mType 'delete' to delete the log files.")
	s := ""
	fmt.Scanln(&s)
	// Check if user wants to delete the logs
	if s == "delete" {
		os.Remove("logs.txt")
		fmt.Println("\033[33mINFO: \033[32mSuccess. \033[34mLogs have been deleted. Press ENTER to exit the program.")
		fmt.Scanln()
		os.Exit(0)
	}
	count, _ := strconv.Atoi(s)

	// Get variables for how often each serial number should be created
	fmt.Println("\033[33mINFO: \033[34mHow often should every serial number be added? \033[0m(For e.g. 2 would add every serial number 2 times: 1 1 2 2 3 3 etc.)")
	s2 := ""
	fmt.Scanln(&s2)
	count2, _ := strconv.Atoi(s2)
	upCount(count, count2)
}

func upCount(count int, count2 int) {
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

	// Creating the numbers for the labels and writing them to the CSV file
	for i := 1; i < count+1; i++ {
		tag := strconv.Itoa(i)
		for i2 := 0; i2 < count2; i2++ {
			_, err2 := f.WriteString(tag + "\n")
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}
	countMessage := strconv.Itoa(count)
	count2Message := strconv.Itoa(count2)
	sum := count * count2
	sumMessage := strconv.Itoa(sum)
	fmt.Println("\033[33mINFO: \033[32mDone. \033[34mCSV-File exported as Output.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
	InfoLogger.Println("Done. CSV-File exported as Output.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
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
