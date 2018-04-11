package main

import (
	"fmt"

	cs "github.com/tenglun/comscape"
)

func main() {
	scape, err := cs.MakeDAO(false)
	if err != nil {
		fmt.Println(err)
	}

	res, err := scape.GetMatrixData("nuttin")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", res)
}
