package main

import (
	"fmt"
	"github.com/YAWAL/AdrianTheProphet/parser"
)

var path = "C:/Users/vyavo/go/src/github.com/YAWAL/AdrianTheProphet/dataset/prepared/Results_797-1750.csv"

func main()  {
	results, err := parser.ParseDataset(path)
	if err != nil{
		fmt.Println(err)
	}
	for num, result := range results {
		fmt.Println(num, " : ",result)
	}
}
