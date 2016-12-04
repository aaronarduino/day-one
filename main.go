package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/aaronarduino/day-one/taxi"
)

func main() {
	pos := taxi.Coordinates{}
	orient := taxi.Direction{}

	r := csv.NewReader(bufio.NewReader(os.Stdin))

	dir, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, inst := range dir[0] {
		tmp := strings.Replace(inst, " ", "", -1) // rm spaces
		way := tmp[:1]
		dur := tmp[1:]
		pos.Move(orient.Turn(way), dur)
	}
	fmt.Println("Moves:", math.Abs(float64(pos.X+pos.Y)))
}
