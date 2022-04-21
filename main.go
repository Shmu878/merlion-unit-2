package main

import (
	"fmt"
	"sort"
)

type Position struct {
	OperationN         int
	ProductN           int
	Description        string
	Quantity           int
}

func main() {
	Positions := []Position{
		{OperationN: 3618, ProductN: 39201, Description: "VC AGP Sp FX5500 128/128 TD oe", Quantity: 117},
		{OperationN: 3619, ProductN: 54441, Description: "Mon Benq FP71G+ (bndl) 7' TFT", Quantity: -2000},
		{OperationN: 3620, ProductN: 54441, Description: "Mon Benq FP71G+ (bndl) 7' TFT", Quantity: 2000},
		{OperationN: 3621, ProductN: 53710, Description: "Mon Benq FP71G (bndl) 17\" TFT", Quantity: -2000},
		{OperationN: 3622, ProductN: 50840, Description: "MB MSI 915P Combo2-V (bndl) (b", Quantity: 2000},
		{OperationN: 3623, ProductN: 51746, Description: "VC PCE Sp GF6600 (bndl) 128/12", Quantity: -2000},
		{OperationN: 3624, ProductN: 51746, Description: "VC PCE Sp GF6600 (bndl) 128/12", Quantity: 2000},
		{OperationN: 3625, ProductN: 50840, Description: "MB MSI 915P Combo2-V (bndl) (b", Quantity: -2000},
		{OperationN: 3626, ProductN: 53710, Description: "Mon Benq FP71G (bndl) 17\" TFT", Quantity: 2000},
		{OperationN: 3627, ProductN: 56437, Description: "DVD+/-RW NEC ND3550 (bndl)", Quantity: -2000},
	}

	for _, pos := range distinct(Positions) {
		fmt.Println(pos)
	}

	fmt.Println("")

	for _, pos := range sortUpDown(distinct(Positions)) {
		fmt.Println(pos)
	}

	fmt.Println("")

	for _, pos := range sortDownUp(distinct(Positions)) {
		fmt.Println(pos)
	}
}

func distinct(table []Position) []Position {
	var res []Position
	for _, p := range table {
		s := false
		for _, r := range res {
			if p.ProductN == r.ProductN {
				s = true
				break
			}
		}
		if !s {
			res = append(res, p)
		}
	}
	return res
}

func sortUpDown(table []Position) []Position {
	sort.Slice(table, func(i, j int) bool {
		return table[i].ProductN > table[j].ProductN
	})
	return table
}

func sortDownUp(table []Position) []Position {
	sort.Slice(table, func(i, j int) bool {
		return table[i].ProductN < table[j].ProductN
	})
	return table
}