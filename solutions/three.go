package solutions

import (
	"main/pkg"
)

func Day_3() (interface{}, interface{}, error) {

	i := pkg.New_Input(3)
	err := i.Parse(pkg.ClearAllEmptyLines)
	if err != nil {
		return nil, nil, err
	}

	data := i.Slices()

	first_priority_total := 0
	second_priority_total := 0

	for _, d := range data {
		r := rucksack{}
		r.identify(d)
		r.duplicate_item()
		r.priority_calculation()

		first_priority_total += r.priority
	}

	groups := len(data) / 3
	rucksackGroups := make([]rucksackGroup, groups)

	for i, g := range rucksackGroups {
		bi := i * 3
		g.memberOne.identify(data[bi])
		g.memberTwo.identify(data[bi+1])
		g.memberThree.identify(data[bi+2])

		rucksackGroups[i] = g
		g.find_badge()
		g.priority_calculation()

		second_priority_total += g.priority
	}

	return first_priority_total, second_priority_total, nil
}

type rucksack struct {
	raw                string
	first_compartment  string
	second_compartment string
	duplicate          rune
	priority           int
}

func (r *rucksack) identify(items string) {
	r.raw = items
	half := len(items) / 2
	r.first_compartment = items[0:half]
	r.second_compartment = items[half:]
}

func (r *rucksack) duplicate_item() rune {
	charMap := map[rune]int{}

	for _, c := range r.first_compartment {
		charMap[c] += 1
	}

	for _, c := range r.second_compartment {
		if charMap[c] > 0 {
			r.duplicate = c
			break
		}
	}

	return r.duplicate
}

func (r *rucksack) priority_calculation() int {
	a := int(r.duplicate)
	if a > 96 {
		a -= 96
	} else {
		a -= 38
	}
	r.priority = a
	return a
}

type rucksackGroup struct {
	created     bool
	memberOne   rucksack
	memberTwo   rucksack
	memberThree rucksack
	badge       rune
	priority    int
}

func (rg *rucksackGroup) find_badge() rune {

	common := []rune{}

	for _, r1 := range rg.memberOne.raw {
		for _, r2 := range rg.memberTwo.raw {
			if r1 == r2 {
				common = append(common, r1)
			}
		}
	}

	for _, c := range common {
		for _, r3 := range rg.memberThree.raw {
			if c == r3 {
				rg.badge = c
				break
			}
		}
	}

	return rg.badge

}

func (rg *rucksackGroup) priority_calculation() int {
	a := int(rg.badge)
	if a > 96 {
		a -= 96
	} else {
		a -= 38
	}
	rg.priority = a
	return a
}
