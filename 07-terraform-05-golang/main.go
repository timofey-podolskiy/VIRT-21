package main

import (
    "fmt"
    "errors"
)

func metresToFeet(metres float64) float64 {
	const METRES_IN_FEET float64 = 0.3048
	return metres * METRES_IN_FEET
}

func intListMin(list []int) (int, error) {
    if (len(list) == 0) {
        return 0, errors.New("Empty list")
    }

    var min int = list[0]
    for _, value := range list {
        if (value < min) {
            min = value
        }
    }

    return min, nil
}

func printMultiplesOf3() []int {
    result := make([]int, 0);
    for i := 1; i <= 100; i++ {
		if (i % 3 == 0) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
    fmt.Print("### TASK 1 ###\n")
	fmt.Print("Enter a number: ")
	var input float64
	_, err := fmt.Scanf("%f", &input)

	if (err != nil) {
	    fmt.Print("Incorrect input")
	} else {
	    fmt.Printf("%f metres equals %f feet", input, feetToMetres(input))
	}

    fmt.Print("\n\n### TASK 2 ###\n")
    var list = []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
    min, err := intListMin(list);

    if (err != nil) {
        fmt.Print(err)
    } else {
        fmt.Printf("The minimum number of the set is %d", min)
    }

    fmt.Print("\n\n### TASK 3 ###\n")
    fmt.Print(printMultiplesOf3())
}
