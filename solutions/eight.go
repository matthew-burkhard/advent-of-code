package solutions

import (
	"fmt"
	"main/pkg"
	"strconv"

	"github.com/pkg/errors"
)

var forest forest_holder

type forest_holder struct {
	rows    int
	columns int
	trees   [][]tree
}

func (f *forest_holder) init(rows, columns int) {
	f.rows = rows
	f.columns = columns
	f.trees = make([][]tree, rows)
	for i, _ := range f.trees {
		f.trees[i] = make([]tree, columns)
	}
}

func (f *forest_holder) visible_from_outside_count() int {
	count := 0
	for _, x := range f.trees {
		for _, y := range x {
			if y.visible_from_outside() {
				count++
			}
		}
	}
	return count
}

func (f *forest_holder) add_tree(x, y int, height int) {

	tree := tree{
		height: height,
		x:      x,
		y:      y,
	}

	f.trees[x][y] = tree
}

func (f *forest_holder) top_scenic_score() int {
	top_score := 0

	for _, x := range f.trees {
		for _, y := range x {
			new_score := y.scenic_score()
			if new_score > top_score {
				top_score = new_score
			}
		}
	}

	return top_score
}

type tree struct {
	x      int
	y      int
	height int
}

func (t tree) scenic_score() int {
	numberOfTreesSeen := func(trees []tree) int {
		treesSeen := 0
		for _, ot := range trees {
			treesSeen++
			t.debug("comparing: %d vs %d\n", ot.height, t.height)
			if ot.height >= t.height {
				t.debug("found bigger value: %d vs %d\n", ot.height, t.height)
				break
			}
		}
		return treesSeen
	}

	treesSeenAbove := numberOfTreesSeen(t.trees_above())
	treesSeenBelow := numberOfTreesSeen(t.trees_below())
	treesSeenLeft := numberOfTreesSeen(t.trees_left())
	treesSeenRight := numberOfTreesSeen(t.trees_right())

	return treesSeenAbove * treesSeenBelow * treesSeenLeft * treesSeenRight
}

func (t tree) visible_from_outside() bool {
	hasBiggerValues := func(trees []tree) bool {
		for _, ot := range trees {
			//t.debug("comparing: %d vs %d\n", ot.height, t.height)
			if ot.height >= t.height {
				t.debug("found bigger value: %d vs %d\n", ot.height, t.height)
				return true
			}
		}
		return false
	}

	hasBiggerAbove := hasBiggerValues(t.trees_above())
	hasBiggerBelow := hasBiggerValues(t.trees_below())
	hasBiggerLeft := hasBiggerValues(t.trees_left())
	hasBiggerRight := hasBiggerValues(t.trees_right())

	return !hasBiggerAbove || !hasBiggerBelow || !hasBiggerLeft || !hasBiggerRight
}

func (f tree) trees_above() []tree {
	above := []tree{}
	for i := 0; i < f.x; i++ {
		above = append(above, forest.trees[i][f.y])
	}
	// above needs to be in reverse (away from f)
	above = reverse_trees(above)
	f.debug("above", trees_to_string(above))
	return above
}

func (f tree) trees_below() []tree {
	below := []tree{}
	for i := f.x + 1; i < forest.rows; i++ {
		below = append(below, forest.trees[i][f.y])
	}
	f.debug("below", trees_to_string(below))
	return below
}

func (f tree) trees_left() []tree {
	t := forest.trees[f.x][0:f.y]
	// left needs to be in reverse (away from f)
	t = reverse_trees(t)
	f.debug("left", trees_to_string(t))
	return t
}

func (f tree) trees_right() []tree {
	t := forest.trees[f.x][f.y+1:]
	f.debug("right", trees_to_string(t))
	return t
}

func trees_to_string(trees []tree) string {
	s := ""
	for _, t := range trees {
		s += strconv.Itoa(t.height)
	}
	return s
}

func reverse_trees(trees []tree) []tree {
	r := []tree{}
	for i := len(trees) - 1; i >= 0; i-- {
		r = append(r, trees[i])
	}
	return r
}

func (f tree) debug(a ...interface{}) {
	return
	if f.x != 5 || f.y != 8 {
		return
	}
	fmt.Println(a...)
}

func Day_8() (interface{}, interface{}, error) {

	i := pkg.New_Input(8)
	err := i.Parse(pkg.ClearLastEmptyLine)
	if err != nil {
		return nil, nil, err
	}
	//i.Println()

	lines := i.Slices()
	rows := len(lines)
	columns := len(lines[0])

	forest.init(rows, columns)

	// for each row of trees
	for x, l := range lines {
		// for each tree in the row
		for y, t := range l {
			height, err := strconv.Atoi(string(t))
			if err != nil {
				return nil, nil, errors.Errorf("cant parse tree %s", string(t))
			}
			forest.add_tree(x, y, height)
		}
	}

	return forest.visible_from_outside_count(), forest.top_scenic_score(), nil
}
