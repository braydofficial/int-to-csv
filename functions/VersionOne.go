package functions

import (
	"fmt"
	"os"
	"strconv"
)

func VersionOne() {
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

	// Ask for container type (CT)
	fmt.Printf("\033[33mINFO: \033[34mType in the container type:\033[0m ")
	containertype := ""
	fmt.Scanln(&containertype)

	UpCount_v1(count, count2, duns, containertype)
}
