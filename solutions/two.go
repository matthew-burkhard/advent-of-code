package solutions

import (
	"main/pkg"
	"strings"

	"github.com/pkg/errors"
)

func Day_2() (interface{}, interface{}, error) {

	i := pkg.New_Input(2)
	err := i.Parse(pkg.ClearAllEmptyLines)
	if err != nil {
		return nil, nil, err
	}

	first_total_score := 0
	second_total_score := 0

	for i, line := range i.Slices() {

		moves := strings.Split(line, " ")
		if len(moves) != 2 {
			return nil, nil, errors.Errorf("unable to parse line %d: `%s`\n", i, line)
		}
		opponent_move := rps(moves[0])
		my_move := rps(moves[1])
		first_total_score += my_move.score(opponent_move)
	}

	for i, line := range i.Slices() {
		moves := strings.Split(line, " ")
		if len(moves) != 2 {
			return nil, nil, errors.Errorf("unable to parse line %d: `%s`\n", i, line)
		}
		opponent_move := rps(moves[0])
		my_outcome := wld(moves[1])
		second_total_score += my_outcome.points() + my_outcome.move(opponent_move)
	}

	return first_total_score, second_total_score, nil

}

type rps string
type wld string

const (
	their_rock     = rps("A")
	their_paper    = rps("B")
	their_scissors = rps("C")
	my_rock        = rps("X")
	my_paper       = rps("Y")
	my_scissors    = rps("Z")
	lose_points    = 0
	draw_points    = 3
	win_points     = 6
	lose           = wld("X")
	draw           = wld("Y")
	win            = wld("Z")
)

func (r rps) string() string {
	switch r {
	case my_rock:
		return "rock"
	case their_rock:
		return "rock"
	case my_paper:
		return "paper"
	case their_paper:
		return "paper"
	case my_scissors:
		return "scissors"
	case their_scissors:
		return "scissors"
	default:
		return ""
	}
}

func (r rps) points() int {
	switch r {
	case my_rock:
		return 1
	case my_paper:
		return 2
	case my_scissors:
		return 3
	default:
		return 0
	}
}

func (r rps) win(opponent rps) int {

	rockDraw := r == my_rock && opponent == their_rock
	paperDraw := r == my_paper && opponent == their_paper
	scissorsDraw := r == my_scissors && opponent == their_scissors

	rockBeatsScissors := r == my_rock && opponent == their_scissors
	scissorsBeatPaper := r == my_scissors && opponent == their_paper
	paperBeatsRock := r == my_paper && opponent == their_rock

	if rockDraw || paperDraw || scissorsDraw {
		return draw_points
	}
	if rockBeatsScissors || scissorsBeatPaper || paperBeatsRock {
		return win_points
	} else {
		return lose_points
	}
}

func (r rps) score(opponent rps) int {

	return r.win(opponent) + r.points()

}

func (w wld) points() int {
	switch w {
	case lose:
		return lose_points
	case win:
		return win_points
	case draw:
		return draw_points
	default:
		return 0
	}
}

func (w wld) move(opponent rps) int {
	switch opponent {
	case their_paper:
		switch w {
		case win:
			return my_scissors.points()
		case lose:
			return my_rock.points()
		case draw:
			return my_paper.points()
		}
	case their_rock:
		switch w {
		case win:
			return my_paper.points()
		case lose:
			return my_scissors.points()
		case draw:
			return my_rock.points()
		}
	case their_scissors:
		switch w {
		case win:
			return my_rock.points()
		case lose:
			return my_paper.points()
		case draw:
			return my_scissors.points()
		}
	}
	return 0
}
