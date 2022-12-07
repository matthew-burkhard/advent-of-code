package solutions

import (
	"fmt"
	"main/pkg"
	"sort"
	"strconv"
	"strings"
)

const (
	root = "/"
)

type file_parser struct {
	fileCounts  map[string]int32
	currentPath file_path
}

func (fp *file_parser) init() {
	fp.fileCounts = make(map[string]int32)
}

func (fp *file_parser) add_file_size(size int32) {
	btr := fp.currentPath.back_to_root()
	for _, r := range btr {
		//fmt.Println("Applying back to path", r)
		fp.add_file_to_path(r, size)
	}
}

func (fp *file_parser) add_file_to_path(p string, size int32) {
	_, found := fp.fileCounts[p]
	if !found {
		fp.fileCounts[p] = size
	} else {
		fp.fileCounts[p] += size
	}
}

func (fp *file_parser) feed_line(line string) {

	isCommand := is_command(line)
	currentPath := fp.currentPath

	//fmt.Println("Parsing `", line, "`")

	if isCommand {
		// this is a command, either cd or ls
		cmd := terminal_command{}
		cmd.message = line
		if cmd.is_cd() {
			// exploring a new path
			if cmd.is_move_out() {
				currentPath.move_out()
			} else {
				// move in command
				if cmd.is_root_move() {
					currentPath.set()
				} else {
					currentPath.move_in(cmd.path_argument())
				}
			}
			//fmt.Println(currentPath.string())
		} else {
			// ls command - feed files into map... do nothing?
			//fmt.Println("ls command found - file time :)")
		}
	} else {
		// this is either a directory or a file
		opt := terminal_output{}
		opt.message = line

		if opt.is_file() {
			fp.add_file_size(opt.file_size())
			//fmt.Println("-", opt.file_name(), size)
			//fmt.Println(myPath, fp.fileCounts[myPath])
		} else {
			// is directory... so do nothing?
			//fmt.Println("directory listing - nothing to do")
		}
	}

	fp.currentPath = currentPath

}

type pathDeets struct {
	path      string
	totalSize int32
}

func (fp file_parser) path_slice_max(maxSize int32) []pathDeets {
	paths := []pathDeets{}
	for p, s := range fp.fileCounts {
		if s > int32(maxSize) && maxSize != -1 {
			continue
		}
		paths = append(paths, pathDeets{
			path:      p,
			totalSize: s,
		})
	}
	return paths
}

func (fp file_parser) path_slice_min(minSize int32) []pathDeets {
	paths := []pathDeets{}
	for p, s := range fp.fileCounts {
		if s <= minSize && minSize != -1 {
			continue
		}
		paths = append(paths, pathDeets{
			path:      p,
			totalSize: s,
		})
	}
	return paths
}

func (fp file_parser) results(maxSize int32) {
	fmt.Println("final results!")
	// convert to slice
	paths := fp.path_slice_max(maxSize)
	sorter := func(i, j int) bool {
		return paths[i].totalSize > paths[j].totalSize
	}

	sort.Slice(paths, sorter)

	// for i, p := range paths {
	// 	fmt.Printf("%d) %s: %d\n", i, p.path, p.totalSize)
	// }

	// fmt.Println("1st Path", paths[0])
	// fmt.Println("2nd Path", paths[1])
	// fmt.Println("3rd Path", paths[2])

	fmt.Println(len(fp.fileCounts), "paths found")
	fmt.Println(len(paths), "filtered paths")
}

func (fp file_parser) total(maxSize int32) int32 {
	size := int32(0)
	for _, s := range fp.fileCounts {
		if s > maxSize {
			//fmt.Println(size, ">", maxSize, m)
			continue
		}
		//fmt.Println("Adding in", m, s)
		size += s
	}
	return size
}

func (fp file_parser) find_path_size(path string) int32 {
	return fp.fileCounts[path]
}

func (fp file_parser) closest_path_by_size(minimumSize int32) int32 {
	// first make paths into a slice
	paths := fp.path_slice_min(minimumSize)

	sorter := func(i, j int) bool {
		return paths[i].totalSize < paths[j].totalSize
	}

	sort.Slice(paths, sorter)

	winner := paths[0]
	fmt.Println("closest_path_by_size", winner.path, winner.totalSize)
	return winner.totalSize
}

func is_command(message string) bool {
	return strings.HasPrefix(message, "$")
}

type terminal_output struct {
	message string
}

func (o terminal_output) is_file() bool {
	return o.parts()[0] != "dir"
}

func (o terminal_output) file_size() int32 {
	if !o.is_file() {
		return -1
	}
	p := o.parts()[0]
	i, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Couldn't atoi!", p)
		return -1
	}
	return int32(i)
}

func (o terminal_output) file_name() string {
	return o.parts()[1]
}

func (o terminal_output) parts() []string {
	// parts[0] == number, or "dir"
	// parts[1] == directory name or file name
	return strings.Split(o.message, " ")
}

type terminal_command struct {
	message string
}

func (c terminal_command) parts() []string {
	// parts[0] == "$"
	// parts[1] == command
	// parts[2] == argument
	return strings.Split(c.message, " ")
}

func (c terminal_command) command() string {
	return c.parts()[1]
}

func (c terminal_command) path_argument() string {
	if len(c.parts()) < 2 {
		return ""
	}
	return c.parts()[2]
}

func (c terminal_command) is_root_move() bool {
	return c.path_argument() == root
}

func (c terminal_command) is_cd() bool {
	return c.command() == "cd"
}

func (c terminal_command) is_move_out() bool {
	return c.path_argument() == ".."
}

type file_path struct {
	path []string
}

func (fp file_path) string() string {
	s := ""
	for i, c := range fp.path {
		s += c
		if i > 0 { // add path markers for all but root
			s += root
		}
	}
	return s
}

func (fp *file_path) set() {
	fp.path = []string{root}
}

func (fp *file_path) move_out() {
	fp.path = fp.path[0 : len(fp.path)-1]
}

func (fp *file_path) move_in(folder string) {
	fp.path = append(fp.path, folder)
}

func (fp file_path) back_to_root() []string {
	pathsOutput := []string{}
	currentPaths := len(fp.path)
	for i := 0; i < currentPaths; i++ {
		pathsOutput = append(pathsOutput, fp.string())
		fp.move_out()
		//paths = append(paths, )
	}
	return pathsOutput
}

func Day_7() (interface{}, interface{}, error) {

	i := pkg.New_Input(7)
	err := i.Parse(pkg.ClearLastEmptyLine)
	if err != nil {
		return nil, nil, err
	}

	parser := file_parser{}
	parser.init()
	for _, l := range i.Slices() {
		parser.feed_line(l)
	}

	oneHundredK := int32(100000)
	oneHundredKTotal := parser.total(oneHundredK)
	parser.results(oneHundredK)

	fmt.Println()
	fmt.Println("Part Two")
	totalDiskAvailable := 70_000_000
	updateSpaceNeeded := 30_000_000
	rootSize := parser.find_path_size(root)
	spaceAvailable := totalDiskAvailable - int(rootSize)
	spaceNeeded := updateSpaceNeeded - spaceAvailable
	fmt.Println("System has a total memory of", totalDiskAvailable)
	fmt.Println("Root has a size of", rootSize)
	fmt.Println("Sadly...", spaceAvailable, "<", updateSpaceNeeded)
	fmt.Println("Looking for another", spaceNeeded)

	// Now to find the directory with the closet space to the needed amount!
	closet_path_size := parser.closest_path_by_size(int32(spaceNeeded))

	return oneHundredKTotal, closet_path_size, nil
}
