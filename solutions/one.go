package solutions

import (
	"fmt"
	"main/pkg"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

func Day_1() (interface{}, interface{}, error) {

	i := pkg.New_Input(1)
	err := i.Parse()
	if err != nil {
		return nil, nil, err
	}

	// read input into parseable format
	elves := []elf{}
	e := elf{}
	for _, v := range i.Slices() {

		// new elf
		if v == "" {
			elves = append(elves, e)
			e = elf{}
			continue
		}

		calories, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil, errors.Errorf("error reading calorie: %s", v)
		}
		e.cals = append(e.cals, calories)
	}

	fmt.Println("Total Elves", len(elves))

	sort.Slice(elves, func(i int, j int) bool {
		return elves[i].totalCals() > elves[j].totalCals()
	})

	topElf := elves[0]
	secondElf := elves[1]
	thirdElf := elves[2]
	topThree := topElf.totalCals() + secondElf.totalCals() + thirdElf.totalCals()

	//fmt.Printf("The Elf carrying the most has %d calories\n", topElf.totalCals())
	//fmt.Printf("The top 3 Elves are carrying %d calories\n", topThree)

	return topElf.totalCals(), topThree, nil

}

type elf struct {
	cals []int
}

func (e elf) totalCals() int {
	totalCals := 0
	for _, cal := range e.cals {
		totalCals += cal
	}
	return totalCals
}
