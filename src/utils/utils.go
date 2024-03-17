package utils

import (
	"bufio"
	"os"
	"strings"
)

func GetStringFromCLI() string {
reader := bufio.NewReader(os.Stdin)
scanner := bufio.NewScanner(reader)
for scanner.Scan(){
	return 	strings.TrimSpace(scanner.Text())
}
return ""
}





