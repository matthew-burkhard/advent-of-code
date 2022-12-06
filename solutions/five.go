package solutions

import (
	"fmt"
	"main/pkg"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Day_5() (interface{}, interface{}, error) {

	i := pkg.New_Input(5)
	err := i.Parse(pkg.ClearLastEmptyLine)
	if err != nil {
		return nil, nil, err
	}

	d, err := i.SplitByNewline()
	if err != nil {
		return nil, nil, err
	}

	// fmt.Println("Original input")
	// for _, l := range d[0] {
	// 	fmt.Println(l)
	// }

	// Parse first data set
	stacks, err := parse_stacks(d[0])
	if err != nil {
		return nil, nil, err
	}

	//fmt.Printf("Found %d stacks\n", len(stacks))
	// for i, s := range stacks {
	// 	fmt.Println(i+1, s)
	// }

	moves, err := parse_instructions(d[1])
	if err != nil {
		return nil, nil, err
	}
	//fmt.Printf("Found %d moves\n", len(moves))

	fmt.Println("Applying moves in normal order...")
	for _, m := range moves {
		//fmt.Println("Applying move", m.raw)
		from := stacks[m.from-1]
		to := stacks[m.to-1]
		for i := 0; i < m.quantity; i++ {
			p := from.pop()
			to.push(p)
		}
		stacks[m.from-1] = from
		stacks[m.to-1] = to
	}

	combo_one := combo(stacks)

	// reparse original stacks
	stacks, err = parse_stacks(d[0])
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Applying moves in reverse order...")
	for _, m := range moves {
		//fmt.Println("Applying move", m.raw)
		from := stacks[m.from-1]
		to := stacks[m.to-1]

		// push items into a reverse stack
		rstack := []string{}
		for i := 0; i < m.quantity; i++ {
			rstack = append(rstack, from.pop())
		}

		// apply reverse stack
		for i := len(rstack) - 1; i >= 0; i-- {
			to.push(rstack[i])
		}

		stacks[m.from-1] = from
		stacks[m.to-1] = to
		//fmt.Println("=================")
		//fmt.Println("after the move...")
		// for i, s := range stacks {
		// 	fmt.Println(i+1, s)
		// }
	}

	// fmt.Println("here we are after!")
	// for i, s := range stacks {
	// 	fmt.Println(i+1, s)
	// }
	combo_two := combo(stacks)

	return combo_one, combo_two, nil
}

func combo(stacks []crate_stack) string {
	combo := ""
	for _, r := range stacks {
		combo += r.top_crate()
	}
	combo = strings.ReplaceAll(combo, "[", "")
	combo = strings.ReplaceAll(combo, "]", "")
	return combo
}

type crate_stack struct {
	stack []string
}

func (s *crate_stack) pop() string {
	// grab the top item
	popped := s.stack[len(s.stack)-1]
	// trip top item
	s.stack = s.stack[:len(s.stack)-1]
	// return popped item
	return popped
}

func (s *crate_stack) push(crate string) {
	s.stack = append(s.stack, crate)
}

func (s crate_stack) top_crate() string {
	if len(s.stack) == 0 {
		return "[]"
	}
	return s.stack[len(s.stack)-1]
}

func parse_stacks(input []string) ([]crate_stack, error) {
	stacks := []crate_stack{}

	// Let's utilize the space/tab offset to determine which stack it needs to be in
	stackNumbers := strings.TrimSpace(input[len(input)-1])
	numbers := strings.Split(stackNumbers, " ")
	lastStack := numbers[len(numbers)-1]
	stackCount, err := strconv.Atoi(strings.TrimSpace(lastStack))
	if err != nil {
		return stacks, err
	}
	stacks = make([]crate_stack, stackCount)
	//fmt.Println("Stack Count", stackCount)

	input = input[:len(input)-1]

	// Print original output
	// fmt.Println("Original input")
	// for _, l := range input {
	// 	fmt.Println(l)
	// }

	// we need to traverse the input reversed to get the stack in the right order
	for i := len(input) - 1; i >= 0; i-- {
		l := input[i]
		for j, c := range l {
			if c != '[' {
				continue
			}
			// We've found the start to a new crate
			crate := l[j : j+3]
			//fmt.Println("New crate at", j)
			stack_num := j / 4
			//fmt.Println("Stack?", stack_num)
			stacks[stack_num].push(crate)
		}
	}

	// fmt.Println("Output!")
	// for _, s := range stacks {
	// 	fmt.Println(s.stack)
	// }

	return stacks, nil
}

type move struct {
	raw      string
	quantity int
	from     int
	to       int
}

func (m *move) parse(input string) error {
	m.raw = input
	q := strings.Index(input, "move")
	f := strings.Index(input, "from")
	t := strings.Index(input, "to")

	//fmt.Println("input", input)
	//fmt.Println(q, f, t)

	qt, err := strconv.Atoi(input[q+5 : f-1])
	if err != nil {
		return err
	}
	fr, err := strconv.Atoi(input[f+5 : t-1])
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(input[t+3:])
	if err != nil {
		return err
	}

	//fmt.Println(qt, fr, to)

	m.quantity = qt
	m.from = fr
	m.to = to
	return nil
}

func parse_instructions(input []string) ([]move, error) {
	m := make([]move, 0)
	for i, l := range input {
		mv := move{}
		err := mv.parse(l)
		if err != nil {
			return m, errors.Wrapf(err, "error parsing line %d", i)
		}
		m = append(m, mv)
	}
	return m, nil
}
