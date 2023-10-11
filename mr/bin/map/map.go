package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func run() error {
	csvReader := csv.NewReader(bufio.NewReader(os.Stdin))

	// read header
	_, err := csvReader.Read()
	if err != nil {
		return err
	}

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Printf("%s, %d\n", rec[5], 1)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
