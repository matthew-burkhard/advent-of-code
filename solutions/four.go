package solutions

import (
	"fmt"
	"main/pkg"
	"strconv"
	"strings"
)

func Day_4() (interface{}, interface{}, error) {

	i := pkg.New_Input(4)
	err := i.Parse(pkg.ClearAllEmptyLines)
	if err != nil {
		return nil, nil, err
	}

	fullyContainCount := 0
	overlapCount := 0
	for _, s := range i.Slices() {
		pair := assignment_pair{}
		pair.decode(s)
		if pair.fully_contain() {
			fullyContainCount++
		}
		if pair.overlaps() {
			overlapCount++
		}
	}

	//fmt.Println("Pairs that Fully Contain:", fullyContainCount)
	//fmt.Println("Pairs that Overlap at All:", overlapCount)

	return fullyContainCount, overlapCount, nil
}

type assignment_pair struct {
	elfOne section_range
	elfTwo section_range
}

func (ap *assignment_pair) decode(input string) {
	parts := strings.Split(input, ",")
	ap.elfOne.decode(parts[0])
	ap.elfTwo.decode(parts[1])
}

func (ap assignment_pair) fully_contain() bool {

	a := ap.elfOne
	b := ap.elfTwo

	aContainsB := a.start <= b.start && a.end >= b.end
	bContainsA := b.start <= a.start && b.end >= a.end

	return aContainsB || bContainsA
}

func (ap assignment_pair) overlaps() bool {

	a := ap.elfOne
	b := ap.elfTwo

	aOverlapsB := a.end >= b.start && a.start <= b.end
	bOverlapsA := b.end >= a.start && b.start <= a.end

	return aOverlapsB || bOverlapsA
}

type section_range struct {
	start int
	end   int
}

func (sr *section_range) decode(input string) {
	parts := strings.Split(input, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("ERROR!  Can't parse start:", parts[0])
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("ERROR!  Can't parse end:", parts[1])
	}
	sr.start = start
	sr.end = end
}
