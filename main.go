package main

import (
	"bufio"
	"fmt"
	functions "int2csv/functions"
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
	fmt.Println("\033[33mINFO: \033[34mPlease choose which type of CSV to be generated. Type 1 for old version or 2 for new version.\n\033[0m ")
	csv_version := " "
	fmt.Scanln(&csv_version)
	// Check which version has been chosen
	if csv_version == "1" {
		version_one()
	} else if csv_version == "2" {
		version_two()
	} else {
		fmt.Println("\033[33mINFO: \033[34mYou have to choose between version 1 or 2. Exiting...")
		os.Exit(0)
	}
}

func version_one() {
	LogFile()
	// Get variables for amount of serial numbers that should be created
	fmt.Println("\033[33mINFO: \033[34mPlease type the amount of serial numbers to be created\n\033[0mType 'delete' to delete the log files.")
	s := ""
	fmt.Scanln(&s)
	// Check if user wants to delete the logs
	if s == "delete" {
		err := os.Remove("./logs.txt")
		if err != nil {
			fmt.Println(err)
		}
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

	// Ask for DUNS number
	fmt.Printf("\033[33mINFO: \033[34mType in the DUNS number you want to use (9 characters): ")
	duns := ""
	fmt.Scanln(&duns)

	// Ask for container type (CT)
	fmt.Printf("\033[33mINFO: \033[34mType in the container type: ")
	containertype := ""
	fmt.Scanln(&containertype)

	upCount_v1(count, count2, duns, containertype)
}

func version_two() {
	LogFile()
	// Get variables for amount of serial numbers that should be created
	fmt.Println("\033[33mINFO: \033[34mPlease type the amount of serial numbers to be created\n\033[0mType 'delete' to delete the log files.")
	s := ""
	fmt.Scanln(&s)
	// Check if user wants to delete the logs
	if s == "delete" {
		err := os.Remove("./logs.txt")
		if err != nil {
			fmt.Println(err)
		}
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

	// Ask for DUNS number
	fmt.Printf("\033[33mINFO: \033[34mType in the DUNS number you want to use (9 characters):\033[0m ")
	duns := ""
	fmt.Scanln(&duns)

	// Ask for container type (LT)
	fmt.Printf("\033[33mINFO: \033[34mType in the LT number (6 characters):\033[0m ")
	ltnumber := ""
	fmt.Scanln(&ltnumber)

	// Ask for project
	fmt.Printf("\033[33mINFO: \033[34mType in the project:\033[0m ")
	scannerproject := bufio.NewScanner(os.Stdin)
	var project string
	if scannerproject.Scan() {
		project = scannerproject.Text()
	}

	// Ask for "Bauteil"
	fmt.Printf("\033[33mINFO: \033[34mType in the Bauteil:\033[0m ")
	scannerbauteil := bufio.NewScanner(os.Stdin)
	var bauteil string
	if scannerbauteil.Scan() {
		bauteil = scannerbauteil.Text()
	}

	upCount_v2(count, count2, duns, ltnumber, project, bauteil)
}

func upCount_v1(count int, count2 int, duns string, containertype string) {
	InfoLogger.Println("Generating CSV-File as Output.csv ...")
	f, e := os.Create("./Output.csv")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	_, err2 := f.WriteString("DUNS;CT;SN\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	// Creating the numbers for the labels and writing them to the CSV file
	for i := 1; i < count+1; i++ {
		// Attach leading zeros and convert from int to string
		res := fmt.Sprintf("%03d", i)

		for i2 := 0; i2 < count2; i2++ {
			_, err2 := f.WriteString(duns + ";" + containertype + ";" + res + "\n")
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}
	countMessage := functions.IntToString(count)
	count2Message := functions.IntToString(count)
	sumMessage := functions.Sum(count, count2)
	fmt.Println("\033[33mINFO: \033[32mDone. \033[34mCSV-File exported as Output.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
	InfoLogger.Println("Done. CSV-File exported as Output.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
}

// Counting up with the variables created trough the user input to create the right amount of entries in the CSV file
func upCount_v2(count int, count2 int, duns string, ltnumber string, project string, bauteil string) {
	InfoLogger.Println("Generating CSV-File as Output-AHA.csv...")
	f, e := os.Create("./Output-AHA.csv")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	_, err2 := f.WriteString("DUNS;LT;SN;PROJEKT;BAUTEIL;DM;RFID\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	// Creating the numbers for the labels and writing them to the CSV file
	for i := 1; i < count+1; i++ {
		// Attach leading zeros and convert from int to string
		sn := fmt.Sprintf("%09d", i)

		// Concatinate RFID and DM variables
		dm := "26BUN" + duns + ltnumber + "+" + sn
		rfid := dm + "!"

		for i2 := 0; i2 < count2; i2++ {
			_, err2 := f.WriteString(duns + ";" + ltnumber + ";" + sn + ";" + project + ";" + bauteil + ";" + dm + ";" + rfid + "\n")
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}
	countMessage := functions.IntToString(count)
	count2Message := functions.IntToString(count2)
	sumMessage := functions.Sum(count, count2)
	fmt.Println("\033[33mINFO: \033[32mDone. \033[34mCSV-File exported as Output-AHA.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
	InfoLogger.Println("Done. CSV-File exported as Output-AHA.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
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
