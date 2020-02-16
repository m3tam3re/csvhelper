package csvhelper

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// TODO add test for GetLines method

func ExampleCreateCsvReader() {
	r, err := createCsvReader("sample_header.csv", ',')
	r1, err := createCsvReader("sample_headless.csv", ';')

	rc, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	rc1, err := r1.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rc)
	fmt.Println(rc1)

	//Output:
	//[[c1 c2 c3 c4 c5] [v1 v2 v3 v4 v5]]
	//[[v1,v2,v3,v4,v5]]
}

func TestCreateCsvReader(t *testing.T) {
	testcases := []struct {
		test     string
		file     string
		sep      rune
		expected *csv.Reader
		pass     bool
	}{
		{test: "Header missing", file: "sample_headless.csv", sep: ',', expected: csv.NewReader(strings.NewReader("v1,v2,v3,v4,v5")), pass: true},
		{test: "Header file missing", file: "sample_missing.csv", sep: ',', expected: csv.NewReader(strings.NewReader("v1,v2,v3,v4,v5")), pass: false},
		{test: "Header missing fail", file: "sample_headless.txt", sep: ',', expected: csv.NewReader(strings.NewReader("v1,v2,v3,v4,v5")), pass: false},
		{test: "Header columns", file: "sample_header.csv", sep: ',', expected: csv.NewReader(strings.NewReader("c1,c2,c3,c4,c5\nv1,v2,v3,v4,v5")), pass: true},
		{test: "Header columns fail", file: "sample_header.txt", sep: ',', expected: csv.NewReader(strings.NewReader("c1,c2,c3,c4,c5\nv1,v2,v3,v4,v5")), pass: false},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%s: %s, separator: %s", tc.test, tc.file, string(tc.sep)), func(t *testing.T) {
			r, err := createCsvReader(tc.file, tc.sep)
			if err != nil {
				if !tc.pass {
					t.Log("this test was expected to fail!")
				}
				t.Fatal(err)
			}
			tc.expected.Comma = tc.sep
			want, err := tc.expected.ReadAll()
			if err != nil {
				t.Errorf("error reading from csvreader: %s", err)
			}
			got, err := r.ReadAll()
			if err != nil {
				t.Errorf("error reading from csvreader: %s", err)
			}
			if !reflect.DeepEqual(want, got) {
				if !tc.pass {
					t.Log("this test was expected to fail!")
				}
				t.Errorf("want: %s, got: %s", want, got)
			}
		},
		)
	}
}
