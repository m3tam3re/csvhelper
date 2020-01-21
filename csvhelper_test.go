package csvhelper

import (
	"fmt"
	"testing"
)

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
	r, err := createCsvReader("sample_header.csv", ',')
	r1, err := createCsvReader("sample_headless.csv", ';')
	if err != nil {
		t.Errorf("Error testing 'createCsvReader()' %s", err)
	}
	rec, err := r.ReadAll()
	rec1, err := r1.ReadAll()
	t.Log(rec)
	t.Log(rec1)
}
