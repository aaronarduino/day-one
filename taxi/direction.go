package taxi

type Direction struct {
	Facing int
}

// Turn takes the "L" or "R" direction and
// returns a cardial direction
func (d *Direction) Turn(way string) string {
	compass := []string{"N", "E", "S", "W"}
	if way == "L" {
		d.Facing = checkBounds(d.Facing - 1)
	} else {
		d.Facing = checkBounds(d.Facing + 1)
	}
	return compass[d.Facing]
}

func checkBounds(i int) int {
	if i > 3 {
		return 0
	} else if i < 0 {
		return 3
	}
	return i
}
