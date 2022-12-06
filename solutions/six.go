package solutions

import (
	"fmt"
	"main/pkg"
)

func Day_6() (interface{}, interface{}, error) {

	i := pkg.New_Input(6)
	err := i.Parse(pkg.ClearLastEmptyLine)
	if err != nil {
		return nil, nil, err
	}

	datastream := i.FirstLine()

	// feed characters one by one until a "marker" is found
	sop := make_marker("start-of-packet", 4)
	for i, c := range datastream {
		sop.append_rune(c, i)
		if sop.marker_check() {
			sop.status()
			break
		}
	}

	som := make_marker("start-of-message", 14)
	for i, c := range datastream {
		som.append_rune(c, i)
		if som.marker_check() {
			som.status()
			break
		}
	}

	return sop.answer, som.answer, nil

}

func make_marker(name string, length int) marker {
	return marker{
		name:           name,
		minimum_length: length,
		stream:         make([]string, 0),
	}
}

type marker struct {
	name            string
	minimum_length  int
	stream          []string
	chars_processed int
	answer          int
}

func (m *marker) append_rune(r rune, i int) {
	m.stream = append(m.stream, string(r))
	current_length := len(m.stream)
	// keep building until marker minimum length
	if current_length < m.minimum_length {
		return
	}
	// trim if too many
	if current_length > m.minimum_length {
		m.stream = m.stream[1:]
	}
	m.chars_processed = i
}

func (m *marker) message() string {
	s := ""
	for _, d := range m.stream {
		s += string(d)
	}
	return s
}

func (m *marker) marker_check() bool {

	// not ready to check yet
	if len(m.stream) < m.minimum_length {
		return false
	}

	s := m.message()

	// mark each byte in a map with a counter
	uniqueCharMap := map[byte]int{}
	for _, b := range s {
		uniqueCharMap[byte(b)]++
		// if any entries are over 1 then they have duplicate characters
		if uniqueCharMap[byte(b)] > 1 {
			return false
		}
	}

	// if we make it this far each char was used once
	return true
}

func (m *marker) status() {
	m.answer = m.chars_processed + 1
	fmt.Printf("Found %s marker at character %d with a value of %s\n", m.name, m.answer, m.message())

}
