package algorithms

import (
	"dokusolv/pkg/solver"
	"encoding/csv"
	"io"
	"os"
	"path"
	"strconv"
	"testing"
)

func getTestFilename() string {
	filename, found := os.LookupEnv("TEST_FILENAME")
	if found == false {
		cwd, err := os.Getwd()
		if err != nil {
			cwd = "./"
		}
		filename = path.Join(cwd, "data", "sudoku.csv")
	}

	return filename
}

func getBoardFromString(data string) *solver.Board {

	j := -1
	b := solver.Make()
	for i := 0; i < len(data); i++ {
		if i % 9 == 0 {
			j++
		}
		b.Cells[i % 9][j], _ = strconv.Atoi(data[i:i+1])
	}

	return b
}

func testAlgorithm(t *testing.T, fn solver.SudokuSolverFn) {

	filename := getTestFilename()
	t.Logf("Using test file: %s", filename)

	f, err := os.Open(filename)
	if err != nil {
		t.Errorf("Unable to open test file: %s", filename)
		return
	}

	cf := csv.NewReader(f)

	_, err = cf.Read()
	for {

		record, err := cf.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Errorf("Unable to read record.")
			break
		}

		b := getBoardFromString(record[0])
		fn(b)
	}

	defer f.Close()
}

func TestBasicAlgorithm(t *testing.T) {

	testAlgorithm(t, solver.Solve)
}
