package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func run() error {
	regionMap := map[string]int64{}
	r := bufio.NewReader(os.Stdin)

	for {
		record, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		kv := strings.Split(record, ", ")
		region := kv[0]

		valueStr := strings.TrimSpace(kv[1])
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			return err
		}

		regionMap[region] += value
	}

	for region, count := range regionMap {
		fmt.Printf("%s, %d\n", region, count)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
