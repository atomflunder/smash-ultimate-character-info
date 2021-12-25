package src

import (
	"bufio"
	"os"
	"strings"
)

//gets the user input
func GetUserInput() string {
	inpReader := bufio.NewReader(os.Stdin)
	inp, _ := inpReader.ReadString('\n')

	//since the json file is in all lowercase, we convert the input
	inp = strings.ToLower(inp)
	//we have to trim the whitespace from the user input, too
	inp = strings.TrimSpace(inp)

	return inp
}
