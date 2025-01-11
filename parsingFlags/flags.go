package parsingflags

import (
	"flag"
	"fmt"
)

func ParseFlags() {
	minusA := flag.Bool("a", false, "minus a, false default")
	minusB := flag.Bool("b", true, "minus b, true default")
	minusC := flag.Int("c", int(0), "minus c, int default zero")
	minusD := flag.String("d", "", "minus d, default empty string")

	flag.Parse()

	fmt.Println("-a: ", *minusA)
	fmt.Println("-b: ", *minusB)
	fmt.Println("-c: ", *minusC)
	fmt.Println("-d: ", *minusD)

	for i, val := range flag.Args() {
		fmt.Printf("index: %d, value: %s \n", i, val)
	}
}
