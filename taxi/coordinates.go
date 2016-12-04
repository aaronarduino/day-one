package taxi

import (
	"log"
	"strconv"
)

type Coordinates struct {
	X int
	Y int
}

// Move takes the direction and len of move and
// updates Coordinates struct
func (c *Coordinates) Move(dir, length string) {
	i, err := strconv.Atoi(length)
	if err != nil {
		log.Fatal(err)
	}

	switch dir {
	case "N":
		c.Y += i
	case "S":
		c.Y -= i
	case "E":
		c.X += i
	case "W":
		c.X -= i
	}
}
