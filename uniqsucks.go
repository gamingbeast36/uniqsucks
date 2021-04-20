package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := flag.String("p", "", "Prints out the contents of file. Example: uniqsucks -p foo.txt")
	showdupfile := flag.String("r", "", "Shows the duplicate lines of file. Example: uniqsucks -r loo.txt")
	deleteline := flag.String("d", "", "Use this flag to mention the file with duplicates. Example: uniqsucks -d foo.txt -o loo.txt")
	output := flag.String("o", "", "Use this flag only with -d to specify the output file of the cleaned text. Example: uniqsucks -d foo.txt -o loo.txt")
	flag.Parse()
	if *fileName != "" {
		cat(&fileName)
	} else if *showdupfile != "" {
		showdupline(&showdupfile)
	} else if *deleteline != "" {
		if *output != "" {
			deletelines(&deleteline, &output)
		} else {
			fmt.Println("Plese pass -o to specify output file for cleaned text")
		}

	} else {
		fmt.Println("Please pass a flag along with file name or pass --help for more info")
	}
}

func cat(name **string) {
	file := *name
	content, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("Error cannot read file")
	}
	body := string(content)
	fmt.Println(body)

}

func showdupline(name **string) {
	file := *name
	op, err := os.Open(*file)
	if err != nil {
		fmt.Println("failed to open file")
	}
	defer op.Close()
	scanner := bufio.NewScanner(op)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanner")
	}
	counts := make(map[string]int)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func deletelines(name **string, dupfile **string) {
	file := *name
	outie := *dupfile
	op, err := os.Open(*file)
	if err != nil {
		fmt.Println("failed to open file")
	}
	defer op.Close()
	scanner := bufio.NewScanner(op)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanner")
	}
	counts := make(map[string]int)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	newfile, err := os.Create(*outie)
	if err != nil {
		fmt.Println(err)
	}
	for line, n := range counts {
		if n > 0 {
			_ = newfile
			pp, err := os.OpenFile(*outie, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer pp.Close()
			var nl = fmt.Sprintln()
			pp.WriteString(line)
			pp.WriteString(nl)
		}
	}
}
