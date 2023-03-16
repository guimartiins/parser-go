package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting...")
	parser()
}

type Item struct {
	name     string
	quantity int
	price    int
}

func parser() {
	list := []Item{}
	data, err := os.Open("test.csv")

	if err != nil {
		panic(err)
	}

	defer data.Close()

	csvReader := csv.NewReader(data)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		values := strings.Split(rec[0], ";")

		qtd, _ := strconv.Atoi(values[1])

		prc, _ := strconv.Atoi(values[2])

		list = append(list, Item{name: values[0], quantity: qtd, price: prc})
	}

	fmt.Println(list)
}
