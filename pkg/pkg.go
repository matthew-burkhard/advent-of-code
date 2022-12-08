package pkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Pkg is my collection of helpers
const (
	input_file = "input.txt"
)

func New_Input(day int) Input {
	return Input{
		day: day,
	}
}

type Input struct {
	day       int
	raw       string
	rawSliced []string
}

func (i Input) filePath() string {
	return fmt.Sprintf("input/%s/%s", strconv.Itoa(i.day), input_file)
}

type InputParseOption func(*InputConfig)

type InputConfig struct {
	ClearAllEmptyLines bool
	ClearLastEmptyLine bool
}

func ClearAllEmptyLines(i *InputConfig) {
	i.ClearAllEmptyLines = true
}

func ClearLastEmptyLine(i *InputConfig) {
	i.ClearLastEmptyLine = true
}

func (i *Input) Parse(options ...InputParseOption) error {

	config := InputConfig{}
	for _, o := range options {
		o(&config)
	}

	// read day's input
	path := i.filePath()
	//fmt.Println("Reading file at " + path)
	d, err := os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "error reading input file")
	}

	// copy raw
	i.raw = string(d)

	// slice by newlines
	i.rawSliced = strings.Split(string(d), "\n")

	if config.ClearAllEmptyLines {
		filteredLines := []string{}
		for _, i := range i.rawSliced {
			if i != "" {
				filteredLines = append(filteredLines, i)
			}
		}
		i.rawSliced = filteredLines
	} else if config.ClearLastEmptyLine {
		lastIdx := len(i.rawSliced) - 1
		lastLine := i.rawSliced[lastIdx]
		if lastLine == "" {
			fmt.Println("empty last line detected, removing")
			i.rawSliced = i.rawSliced[:lastIdx]
		}
	}

	//fmt.Println("Successfully parsed input file!")
	fmt.Println(path, len(i.raw), "characters", len(i.rawSliced), "lines")

	return nil
}

func (i Input) SplitByNewline() ([][]string, error) {
	// split data by empty line
	for idx, l := range i.rawSliced {
		if l == "" {
			fParts := [][]string{
				i.rawSliced[:idx],
				i.rawSliced[idx+1:],
			}
			fmt.Println(len(fParts[0]), "lines on part 1")
			fmt.Println(len(fParts[1]), "lines on part 2")
			return fParts, nil
		}
	}
	return [][]string{}, errors.New("newline not found to split input on")
}

func (i Input) FirstLine() string {
	return i.rawSliced[0]
}

func (i Input) Slices() []string {
	return i.rawSliced
}

func (i Input) Println() {
	for i, l := range i.Slices() {
		idx := fmt.Sprintf("%d", i)
		if i < 10 {
			idx = "00" + idx
		} else if i < 100 {
			idx = "0" + idx
		}
		fmt.Printf("[%s] %s\n", idx, l)
	}
}
