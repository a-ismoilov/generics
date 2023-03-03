package main

import "fmt"

func main() {
	var (
		intMap = map[string]int64{
			"first":  10,
			"second": 20,
		}

		floatMap = map[string]float64{
			"first":  10.0,
			"second": 20.0,
		}
	)
	fmt.Println("Non-Generic Code:")
	fmt.Println(`	+--INT--+
	| `, sumInt(intMap), `  |
	+-------+`)
	fmt.Println(`	+--FLOAT--+
	|  `, sumFloat(floatMap), `   |
	+---------+`)

	fmt.Println("Generic Code:")
	fmt.Println(`	+--INT--+
	| `, sumIntsOrFloats(intMap), `  |
	+-------+`)
	fmt.Println(`	+--FLOAT--+
	|  `, sumIntsOrFloats(floatMap), `   |
	+---------+`)
}
