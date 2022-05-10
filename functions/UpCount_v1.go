package functions

import (
	"fmt"
	"log"
	"os"
)

// Counting up with the variables created trough the user input to create the right amount of entries in the CSV file
func UpCount_v1(count int, count2 int, duns string, containertype string) {
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
	countMessage := IntToString(count)
	count2Message := IntToString(count2)
	sumMessage := Sum(count, count2)
	fmt.Println("\033[33mINFO: \033[32mDone. \033[34mCSV-File exported as Output.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
	InfoLogger.Println("Done. CSV-File exported as Output.csv. Every serial number has been created " + count2Message + " times. All together " + countMessage + " serial numbers have been created. Summed up this will create data for " + sumMessage + " labels.")
}
