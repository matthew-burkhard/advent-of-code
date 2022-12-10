package solutions

import (
	"fmt"
	"main/pkg"
	"math"
	"strconv"
	"strings"
)

type rope struct {
	debug               bool
	uniqueTailPositions map[string]bool
	knots               []end
	// head                end
	// tail                end
}

func (r *rope) init(length int) {
	r.uniqueTailPositions = make(map[string]bool)
	r.knots = make([]end, length)
	for i := range r.knots {
		r.knots[i] = end{
			x: 0,
			y: 0,
		}
	}
}

func (r rope) isHeadKnot(index int) bool {
	return index == 0
}

func (r rope) isLastKnot(index int) bool {
	return index == len(r.knots)-1
}

func (r rope) lastKnotIndex() int {
	return len(r.knots) - 1
}

func (r *rope) move(direction string, count string) {

	if r.debug {
		fmt.Println()
		fmt.Println("~~~> Moving", direction, "count", count)
	}

	moves, err := strconv.Atoi(count)
	if err != nil {
		panic("can not parse count" + count)
	}
	// move the head and tail
	for i := 0; i < moves; i++ {

		// move head one step in new direction
		previous := r.knots[0].move(direction)

		// each knot needs to follow the one in front of it
		knotAhead := r.knots[0]

		// start at index 1 to skip head knot
		for i := 1; i <= r.lastKnotIndex(); i++ {

			// make the next knot follow the one in front
			previous = r.knots[i].follow(previous, knotAhead)

			// mark the current knot as the knot ahead for next iteration
			knotAhead = r.knots[i]

			if i == r.lastKnotIndex() {
				// record tail unique position
				r.uniqueTailPositions[r.knots[i].toString()] = true
			}

			// if r.debug {
			// 	fmt.Println("After move:")
			// 	r.print_grid()
			// }
		}
	}

	if r.debug {
		r.print_grid()
	}

}

func (r rope) count_unique_tail_positions() int {
	return len(r.uniqueTailPositions)
}

func (r rope) print_knots() {
	for i, k := range r.knots {
		fmt.Printf("Knot %d at position x,y: %s\n", i, k.toString())
	}
}

func (r rope) print_grid() {

	// printRowNum := func(y int) {
	// 	if y < 10 {
	// 		fmt.Print(0)
	// 	}
	// 	fmt.Print(y, "]")
	// }

	gridSize := 32 // -15 to 15
	halfGrid := gridSize / 2

	fmt.Println()
	for y := (halfGrid - 1); y >= -halfGrid; y-- {
		//printRowNum(y)
		for x := -halfGrid; x < halfGrid; x++ {
			display := false
			for i, k := range r.knots {
				display = k.x == x && k.y == y
				if display {
					if r.isHeadKnot(i) {
						fmt.Print("H")
					} else if r.isLastKnot(i) {
						fmt.Print("T")
					} else {
						fmt.Print(i)
					}
					break
				}
			}
			if !display {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type end struct {
	x int
	y int
}

func (e *end) move(direction string) end {
	before := *e
	switch direction {
	case "U":
		e.y += 1
	case "D":
		e.y -= 1
	case "L":
		e.x -= 1
	case "R":
		e.x += 1
	}
	return before
}

func (e *end) follow(before end, new end) end {

	previous := *e

	xDif := int(math.Abs(float64(new.x - e.x)))
	yDif := int(math.Abs(float64(new.y - e.y)))
	noMove := xDif == 0 && yDif == 0
	withinOne := xDif < 2 && yDif < 2

	if noMove || withinOne {
		// moved to the same spot, or within one, as the tail, nothing to do!
		return *e
	} else {
		e.x = before.x
		e.y = before.y
	}

	return previous

}

func (e end) toString() string {

	return strconv.Itoa(e.x) + "|" + strconv.Itoa(e.y)

}

func fake_moves() []string {
	return []string{
		"U 15",
		"L 15",
		"D 15",
		"R 30",
		"U 15",
		"L 15",
		"D 15",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
		"L 1",
		"D 1",
	}
}

func Day_9() (interface{}, interface{}, error) {

	i := pkg.New_Input(9)
	err := i.Parse(pkg.ClearLastEmptyLine)
	if err != nil {
		return nil, nil, err
	}

	short_rope := rope{}
	short_rope.init(2)

	long_rope := rope{}
	long_rope.init(10)

	// fmt.Println("Starting!")
	// long_rope.print_grid()

	// Swap these to enable your testing script!
	moves := i.Slices()
	//moves := fake_moves()
	for i, l := range moves {

		if i == 10 {
			long_rope.debug = true
		} else if i == 16 {
			long_rope.debug = false
		}

		parts := strings.Split(l, " ")
		direction := parts[0]
		count := parts[1]
		short_rope.move(direction, count)
		long_rope.move(direction, count)
	}
	//i.Println()
	//long_rope.print_grid()
	long_rope.print_knots()

	// 4930 - too high for long rope!
	// 3341 - still too high for long rope :/

	return short_rope.count_unique_tail_positions(), long_rope.count_unique_tail_positions(), nil
}
