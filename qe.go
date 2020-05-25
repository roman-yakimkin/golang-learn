package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getRoots(a float64, b float64, c float64) (x1 float64, x2 float64, e error) {
	if a == 0 {
		if b == 0 {
			return 0, 0, fmt.Errorf("This equqtion is true for all of X")
		}
		x1, x2 = -c/b, -c/b
	} else {
		d := math.Sqrt(b*b - 4*a*c)
		if math.IsNaN(d) {
			return 0, 0, fmt.Errorf("This equation has no real roots")
		}
		x1, x2 = (-b+d)/(2*a), (-b-d)/(2*a)
	}
	return x1, x2, e
}

func main() {
	fmt.Println("The solving of quadratic equation:")
	fmt.Println("Input a, b and c with spaces")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inpValues := strings.Split(input, " ")
	a, err := strconv.ParseFloat(strings.TrimSpace(inpValues[0]), 64)
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.ParseFloat(strings.TrimSpace(inpValues[1]), 64)
	if err != nil {
		log.Fatal(err)
	}
	c, err := strconv.ParseFloat(strings.TrimSpace(inpValues[2]), 64)
	if err != nil {
		log.Fatal(err)
	}
	x1, x2, err := getRoots(a, b, c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The roots for quadratic equiaion %.3f*x*x + %.3f*x + %.3f = %.3f and %.3f\n", a, b, c, x1, x2)
	}
}
