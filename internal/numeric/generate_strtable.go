// +build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	finparts := [][2]string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "*/")

		f := strings.Fields(parts[0])
		num := f[1]
		if num == "999" {
			continue
		}

		fmtString := parts[len(parts)-1]
		fmtString = strings.TrimSpace(fmtString[:len(fmtString)-1])
		if fmtString[len(fmtString)-1] == ',' {
			fmtString = fmtString[:len(fmtString)-1]
		}

		finparts = append(finparts, [2]string{num, fmtString})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fout, err := os.Create("./replies_strtable.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	fmt.Fprintf(fout, "package numeric\n\nvar formatStrings map[Response]string\n\nfunc init(){\n\tformatStrings = map[Response]string{\n")

	for _, f := range finparts {
		fmt.Fprintf(fout, "\t\tResponse(%q): %s,\n", f[0], f[1])
	}

	fmt.Fprintln(fout, "}}")
}
