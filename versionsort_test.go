package versionsort

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func ExampleVersionSort() {
	strs := []string{
		"v1.1",
		"v1.10",
		"v1.11",
		"v1.9",
		"v1.8",
	}
	Sort(strs, false)
	for _, s := range strs {
		fmt.Println(s)
	}
	// Output:
	// v1.1
	// v1.8
	// v1.9
	// v1.10
	// v1.11
}

func ExampleVersionSort_reverse() {
	strs := []string{
		"v1.1",
		"v1.9",
		"v1.8",
		"v1.10",
		"v1.11",
	}
	Sort(strs, true)
	for _, s := range strs {
		fmt.Println(s)
	}
	// Output:
	// v1.11
	// v1.10
	// v1.9
	// v1.8
	// v1.1
}

func ExampleVersionSort_haya14busa() {
	strs := []string{
		"haya2busa",
		"haya1busa",
		"haya14busa",
		"haya13busa",
	}
	Sort(strs, false)
	for _, s := range strs {
		fmt.Println(s)
	}
	// Output:
	// haya1busa
	// haya2busa
	// haya13busa
	// haya14busa
}

func TestVersionSort(t *testing.T) {
	testVersionSortFile(t, "test1")
}

func testVersionSortFile(t *testing.T, name string) {
	infile := fmt.Sprintf("testdata/%s.in", name)
	in, err := os.Open(infile)
	if err != nil {
		t.Errorf("no test for %q: %v", name, err)
		return
	}
	defer in.Close()
	b, err := ioutil.ReadAll(in)
	if err != nil {
		t.Errorf("failed to read data %q: %v", name, err)
		return
	}
	lines := strings.Split(string(b), "\n")

	Sort(lines, false)

	outfile := fmt.Sprintf("testdata/%s.out", name)
	out, err := os.Create(outfile)
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()
	for _, l := range lines {
		fmt.Fprintln(out, l)
	}
	okfile := fmt.Sprintf("testdata/%s.ok", name)
	ok, err := os.Open(okfile)
	if err != nil {
		t.Errorf("no ok test for %q: %v", name, err)
		return
	}
	defer ok.Close()
	b, err = exec.Command("diff", "-u", okfile, outfile).Output()
	if err != nil {
		t.Error(err)
	}
	if d := string(b); d != "" {
		t.Error(d)
	}
}
