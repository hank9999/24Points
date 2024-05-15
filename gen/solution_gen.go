package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/mowshon/iterium"
)

const tmpl = `package solution24

func GetSolution(a, b, c, d int) (solutions []string) {
	{{ range $_, $s := .}}
		if {{$s}} == 24 { solutions = append(solutions, "{{$s}}") }
	{{ end }}
	return solutions
}
`

func main() {
	operators := []string{"+", "-", "*", "/"}
	numbers := []string{"a", "b", "c", "d"}

	operator_combo_iter := iterium.Product(operators, 3)
	number_combo_iter := iterium.Permutations(numbers, 4)

	wr := &bytes.Buffer{}

	operator_combos, err := operator_combo_iter.Slice()
	if err != nil {
		panic(err)
	}
	number_combos, err := number_combo_iter.Slice()
	if err != nil {
		panic(err)
	}

	solutions := []string{}
	for _, s := range operator_combos {
		fmt.Println(s)
		for _, n := range number_combos {
			solutions = append(solutions, n[0]+s[0]+n[1]+s[1]+n[2]+s[2]+n[3])
		}
	}

	err = template.Must(template.New("24solution").Parse(tmpl)).Execute(wr, solutions)

	if err != nil {
		panic(err)
	}

	os.WriteFile("solution24.go", wr.Bytes(), os.ModePerm)
}
