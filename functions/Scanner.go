package functions

import (
	"bufio"
	"os"
)

func Scanner(project string) string {
	scannerfunc := bufio.NewScanner(os.Stdin)
	project = scannerfunc.Text()
	return project
}
