package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/toolkits/file"
)

const (
	VERSION = "0.1.0"
)

var debug_mode bool = false

func fetchAllMacros(in string) (map[string]string, error) {
	m := make(map[string]string)
	re := regexp.MustCompile("%%.*?%%")
	keys := re.FindAllString(in, -1)
	for i, v := range keys {
		_, ok := m[v]
		if ok {
			continue
		}

		m[v] = os.Getenv(strings.Split(v, "%%")[1])
		if debug_mode {
			fmt.Printf("[%d] map[%s]=%s\n", i, v, m[v])
		}
	}

	return m, nil
}

func replaceAllMacros(m map[string]string, in, outfile string) error {
	var tmpStr string = in

	for k, v := range m {
		tmpStr = strings.Replace(tmpStr, k, v, -1)
	}

	if debug_mode {
		fmt.Println("out:", tmpStr)
	}

	_, err := file.WriteString(outfile, tmpStr)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fin := flag.String("i", "file.in", "input file")
	fout := flag.String("o", "file.out", "output file")
	version := flag.Bool("v", false, "show version")
	debug := flag.Bool("d", false, "debug mode")

	flag.Parse()

	debug_mode = *debug

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	content, err := file.ToString(*fin)
	if err != nil {
		fmt.Println(err, "when read", *fin)
		os.Exit(1)
	}

	m, err := fetchAllMacros(string(content))
	if err != nil {
		fmt.Println(err, "when fetch all macros")
		os.Exit(1)
	}

	err = replaceAllMacros(m, string(content), *fout)
	if err != nil {
		fmt.Println(err, "when replace all macros")
		os.Exit(1)
	}

	if debug_mode {
		fmt.Println("changes wrote to", *fout)
	}

	return
}
