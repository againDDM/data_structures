package main

import 	(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	PhoneBook := make(map[int]string)
	var qNumber int
	fmt.Scan(&qNumber)
	reader := bufio.NewReader(os.Stdin)
	for ; qNumber > 0; qNumber-- {
		command, _ := reader.ReadString(' ')
		switch {
		case strings.HasPrefix(command, "a"): // add
			pn, _ := reader.ReadString(' ')
			pNumber, _ := strconv.Atoi(string(pn[:len(pn)-1]))
			nm, _, _ := reader.ReadLine()
			PhoneBook[pNumber] = string(nm)
		case strings.HasPrefix(command, "d"): // del
			pn, _, _ := reader.ReadLine()
			pNumber, _ := strconv.Atoi(string(pn))
			delete(PhoneBook, pNumber)
		case strings.HasPrefix(command, "f"): // find
			pn, _, _ := reader.ReadLine()
			pNumber, _ := strconv.Atoi(string(pn))
			if name, ok := PhoneBook[pNumber]; ok {
				fmt.Println(name)
			} else {
				fmt.Println("not found")
			}
		}
	}
}
