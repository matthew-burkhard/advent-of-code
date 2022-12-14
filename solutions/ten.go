package solutions

import (
	"fmt"
	"main/pkg"
	"strconv"
	"strings"
)

func Day_10() (interface{}, interface{}, error) {

	i := pkg.New_Input(10)
	err := i.Parse(pkg.ClearLastEmptyLine)
	if err != nil {
		return nil, nil, err
	}

	cycle_number := 0
	x_register := 1

	signal_strength := func() int {
		return cycle_number * x_register
	}

	signal_map := map[int]int{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}

	check_signal_map := func() {
		_, found := signal_map[cycle_number]
		if found {
			signal_map[cycle_number] = signal_strength()
		}
	}

	debug := true

	crt_lit_pixel := byte('#')
	crt_dim_pixel := byte('.')

	// draw a 6x40 crt, hxw - default to LIT
	crt := make([][]byte, 6)
	for i, _ := range crt {
		crt[i] = make([]byte, 40)
		for j, _ := range crt[i] {
			crt[i][j] = crt_lit_pixel
		}
	}

	// helper if x-register has sprite visible
	// isSpriteVisible := func(pixel int) bool {
	// 	return pixel+1 <= x_register && x_register >= pixel-1
	// }

	drawPixel := func() {
		row := cycle_number / 40
		col := cycle_number % 40
		isSpriteVisible := col >= x_register-1 && col <= x_register+1
		pixel := crt_dim_pixel
		if isSpriteVisible {
			pixel = crt_lit_pixel
		}
		crt[row][col] = pixel
	}

	for i, l := range i.Slices() {

		// change this to see specific indexes w/ debugging
		if i == 0 {
			debug = false
		}

		if debug {
			fmt.Println("Processing:", l)
		}

		parts := strings.Split(l, " ")
		instruction := parts[0]
		switch instruction {
		case "noop":
			drawPixel()
			cycle_number++
			check_signal_map()
		case "addx":
			// check writing/calculating to signal map for each cycle
			drawPixel()
			cycle_number += 1
			check_signal_map()

			if debug {
				fmt.Println("Cycle", cycle_number, "|", "Register", x_register)
			}

			drawPixel()
			cycle_number += 1

			check_signal_map()
			// the register is added AFTER the end of the cycle
			addx, _ := strconv.Atoi(parts[1])
			x_register += addx

		}

		// now draw!

		if debug {
			fmt.Println("Cycle", cycle_number, "|", "Register", x_register)
			fmt.Println()
		}

	}

	fmt.Println(signal_map)
	sum := 0
	for _, d := range signal_map {
		sum += d
	}
	fmt.Println("Total Signal", sum)
	// 8040 - too low
	// 12,220 - wrong :(

	// print crt
	for _, i := range crt {
		for _, j := range i {
			fmt.Print(string(j))
		}
		fmt.Println()
	}

	/*
		###..####..##..###..#..#.###..####.###..
		#..#....#.#..#.#..#.#.#..#..#.#....#..#.
		#..#...#..#....#..#.##...#..#.###..###..
		###...#...#.##.###..#.#..###..#....#..#.
		#....#....#..#.#....#.#..#....#....#..#.
		#....####..###.#....#..#.#....####.###..
	*/

	return sum, "PZGPKPEB", nil
}
