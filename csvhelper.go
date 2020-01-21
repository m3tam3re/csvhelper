// Package modules implements helper functions to work with csv files.
package csvhelper

import (
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var CsvReader *csv.Reader

func createCsvReader(path string, comma rune) (*csv.Reader, error) {
	if filepath.Ext(strings.ToLower(path)) != ".csv" {
		err := errors.New("Filetype must be .CSV")
		return nil, fmt.Errorf("Wrong filetype: %s", err)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	CsvReader = csv.NewReader(f)
	CsvReader.Comma = comma
	return CsvReader, nil
}

// CsvGetLines consumes a path to a *.csv file, a comma separator and header.
// If a header is present the returned map keys are named after the fields in the header row.
// Otherwise the returned map keys are named after the index of the field stating with 0.
// It returns the separated values of line one as a slice of string.
// All following lines are returned as a slice of slice of string
func CsvGetLines(path string, comma rune, header bool) ([]map[string]string, error) {
	var headerrow []string
	var lines []map[string]string

	r, err := createCsvReader(path, comma)
	if err != nil {
		return nil, err
	}
	lc := 1
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if lc == 1 {
			for i, v := range line {
				if header {
					headerrow = append(headerrow, v)
					continue
				}
				headerrow = append(headerrow, strconv.Itoa(i))
			}
			lc++
		}
		m := make(map[string]string)
		for i, v := range line {
			m[headerrow[i]] = v
		}
		lines = append(lines, m)
		lc++
	}

	return lines, nil
}
