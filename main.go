package main

import (
	"fmt"
	functions "int2csv/functions"
	"os"
)

func main() {
	fmt.Println("\033[33mINFO: \033[34mPlease choose which type of CSV to be generated. Type 1 for old version or 2 for new version.\n\033[0m ")
	csv_version := " "
	fmt.Scanln(&csv_version)
	// Check which version has been chosen
	if csv_version == "1" {
		functions.VersionOne()
	} else if csv_version == "2" {
		functions.VersionTwo()
	} else {
		fmt.Println("\033[33mINFO: \033[34mYou have to choose between version 1 or 2. Press ENTER to exit the program.")
		fmt.Scanln()
		os.Exit(0)
	}
}
