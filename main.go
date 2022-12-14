package main

import (
	"fmt"
	"main/solutions"
	"time"
)

const forceDay = 11

type PuzzleResponse interface {
}

type Puzzle struct {
	Solution  func() (interface{}, interface{}, error)
	AnswerOne interface{}
	AnswerTwo interface{}
}

func (p Puzzle) Solve() bool {

	pass := true
	a1Str := "=="
	a2Str := "=="

	a1, a2, err := p.Solution()
	if err != nil {
		fmt.Printf("error running solution: %v\n", err)
	}

	if a1 != p.AnswerOne {
		pass = false
		a1Str = "!="
	} else if a2 != p.AnswerTwo {
		pass = false
		a2Str = "!="
	}

	fmt.Println("~~~~~~~~~~~")
	fmt.Printf("- Answer 1: %#v %s %#v\n", a1, a1Str, p.AnswerOne)
	fmt.Printf("- Answer 2: %#v %s %#v\n", a2, a2Str, p.AnswerTwo)
	fmt.Printf("Result: ")
	if pass {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL!!!")
	}
	fmt.Println("~~~~~~~~~~~")

	return pass

}

func main() {

	codeDay := map[int]Puzzle{
		1: {
			Solution:  solutions.Day_1,
			AnswerOne: 68923,
			AnswerTwo: 200044,
		},
		2: {
			Solution:  solutions.Day_2,
			AnswerOne: 14827,
			AnswerTwo: 13889,
		},
		3: {
			Solution:  solutions.Day_3,
			AnswerOne: 8401,
			AnswerTwo: 2641,
		},
		4: {
			Solution:  solutions.Day_4,
			AnswerOne: 518,
			AnswerTwo: 909,
		},
		5: {
			Solution:  solutions.Day_5,
			AnswerOne: "CWMTGHBDW",
			AnswerTwo: "SSCGWJCRB",
		},
		6: {
			Solution:  solutions.Day_6,
			AnswerOne: 1262,
			AnswerTwo: 3444,
		},
		7: {
			Solution:  solutions.Day_7,
			AnswerOne: int32(1297683),
			AnswerTwo: int32(5756764),
		},
		8: {
			Solution:  solutions.Day_8,
			AnswerOne: 1776,
			AnswerTwo: 234416,
		},
		// 9: {
		// 	Solution:  solutions.Day_9,
		// 	AnswerOne: 6337,
		// 	AnswerTwo: "?",
		// },
		10: {
			Solution:  solutions.Day_10,
			AnswerOne: 13680,
			AnswerTwo: "PZGPKPEB",
		},
	}

	allPassed := true
	current_day := time.Now().Day()
	start_day := 1
	if forceDay != 0 {
		start_day = forceDay
	}
	for i := start_day; i <= current_day; i++ {
		day, found := codeDay[i]
		if !found || day.Solution == nil {
			fmt.Printf("day %d not implemented...!\n", i)
			allPassed = false
			break
		}
		fmt.Printf("~~~~~~~~~~~Day %d~~~~~~~~~~~\n", i)
		pass := day.Solve()

		if !pass {
			fmt.Printf("Unable to pass day %d!  Can't move on...\n", i)
			allPassed = false
			break
		} else {
			fmt.Println("")
		}

	}

	if allPassed {
		fmt.Println("Congrats, all caught up!")
	} else {
		fmt.Println("Try again...")
	}

}
