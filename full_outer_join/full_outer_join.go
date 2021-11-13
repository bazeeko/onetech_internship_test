package full_outer_join

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func scanLines(file *os.File) (lines map[string]bool) {
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	lines = make(map[string]bool)

	for scanner.Scan() {
		lines[scanner.Text()] = true
	}

	return
}

func outerJoin(a, b map[string]bool) (result []string) {
	for key := range a {
		if _, exists := b[key]; exists {
			delete(a, key)
			delete(b, key)
		}
	}

	result = make([]string, 0, len(a)+len(b))

	for key := range a {
		result = append(result, key)
	}

	for key := range b {
		result = append(result, key)
	}

	sort.Strings(result)

	return
}

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	file1, err := os.Open(f1Path)
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	file2, err := os.Open(f2Path)
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	file1Lines := scanLines(file1)
	file2Lines := scanLines(file2)

	result := outerJoin(file1Lines, file2Lines)

	resultFile, err := os.Create(resultPath)
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()

	for i := range result {
		if i != len(result)-1 {
			resultFile.WriteString(fmt.Sprintf("%s\n", result[i]))
		} else {
			resultFile.WriteString(result[i])
		}
	}
}
