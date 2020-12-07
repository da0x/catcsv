//
//   main.go
//   catcsv
//
//   Copyright 2020 Daher Alfawares
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License
//   You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2
//
//   Unless required by applicable law or agreed to in writing,
//   distributed under the License is distributed on an "AS IS" BASIS
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
//   See the License for the specific language governing permissions
//   limitations under the License

package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/da0x/olog"
)

func readFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func readStdin() [][]string {
	csvReader := csv.NewReader(os.Stdin)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse input as CSV", err)
	}
	return records
}

func main() {
	if len(os.Args) != 2 {
		println("usage: $ csv <filename>")
	}
	filename := os.Args[1]
	if filename == "-" {
		olog.Print(readStdin())
	} else {
		olog.Print(filename, readFile(filename))
	}
}
