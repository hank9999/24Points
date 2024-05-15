package solution24

func GetSolution(a, b, c, d int) (solutions []string) {
	
		if a+b+c+d == 24 { solutions = append(solutions, "a+b+c+d") }
	
		if a+b+d+c == 24 { solutions = append(solutions, "a+b+d+c") }
	
		if a+c+b+d == 24 { solutions = append(solutions, "a+c+b+d") }
	
		if a+c+d+b == 24 { solutions = append(solutions, "a+c+d+b") }
	
		if a+d+c+b == 24 { solutions = append(solutions, "a+d+c+b") }
	
		if a+d+b+c == 24 { solutions = append(solutions, "a+d+b+c") }
	
		if b+a+c+d == 24 { solutions = append(solutions, "b+a+c+d") }
	
		if b+a+d+c == 24 { solutions = append(solutions, "b+a+d+c") }
	
		if b+c+a+d == 24 { solutions = append(solutions, "b+c+a+d") }
	
		if b+c+d+a == 24 { solutions = append(solutions, "b+c+d+a") }
	
		if b+d+c+a == 24 { solutions = append(solutions, "b+d+c+a") }
	
		if b+d+a+c == 24 { solutions = append(solutions, "b+d+a+c") }
	
		if c+b+a+d == 24 { solutions = append(solutions, "c+b+a+d") }
	
		if c+b+d+a == 24 { solutions = append(solutions, "c+b+d+a") }
	
		if c+a+b+d == 24 { solutions = append(solutions, "c+a+b+d") }
	
		if c+a+d+b == 24 { solutions = append(solutions, "c+a+d+b") }
	
		if c+d+a+b == 24 { solutions = append(solutions, "c+d+a+b") }
	
		if c+d+b+a == 24 { solutions = append(solutions, "c+d+b+a") }
	
		if d+b+c+a == 24 { solutions = append(solutions, "d+b+c+a") }
	
		if d+b+a+c == 24 { solutions = append(solutions, "d+b+a+c") }
	
		if d+c+b+a == 24 { solutions = append(solutions, "d+c+b+a") }
	
		if d+c+a+b == 24 { solutions = append(solutions, "d+c+a+b") }
	
		if d+a+c+b == 24 { solutions = append(solutions, "d+a+c+b") }
	
		if d+a+b+c == 24 { solutions = append(solutions, "d+a+b+c") }
	
		if a+b+c-d == 24 { solutions = append(solutions, "a+b+c-d") }
	
		if a+b+d-c == 24 { solutions = append(solutions, "a+b+d-c") }
	
		if a+c+b-d == 24 { solutions = append(solutions, "a+c+b-d") }
	
		if a+c+d-b == 24 { solutions = append(solutions, "a+c+d-b") }
	
		if a+d+c-b == 24 { solutions = append(solutions, "a+d+c-b") }
	
		if a+d+b-c == 24 { solutions = append(solutions, "a+d+b-c") }
	
		if b+a+c-d == 24 { solutions = append(solutions, "b+a+c-d") }
	
		if b+a+d-c == 24 { solutions = append(solutions, "b+a+d-c") }
	
		if b+c+a-d == 24 { solutions = append(solutions, "b+c+a-d") }
	
		if b+c+d-a == 24 { solutions = append(solutions, "b+c+d-a") }
	
		if b+d+c-a == 24 { solutions = append(solutions, "b+d+c-a") }
	
		if b+d+a-c == 24 { solutions = append(solutions, "b+d+a-c") }
	
		if c+b+a-d == 24 { solutions = append(solutions, "c+b+a-d") }
	
		if c+b+d-a == 24 { solutions = append(solutions, "c+b+d-a") }
	
		if c+a+b-d == 24 { solutions = append(solutions, "c+a+b-d") }
	
		if c+a+d-b == 24 { solutions = append(solutions, "c+a+d-b") }
	
		if c+d+a-b == 24 { solutions = append(solutions, "c+d+a-b") }
	
		if c+d+b-a == 24 { solutions = append(solutions, "c+d+b-a") }
	
		if d+b+c-a == 24 { solutions = append(solutions, "d+b+c-a") }
	
		if d+b+a-c == 24 { solutions = append(solutions, "d+b+a-c") }
	
		if d+c+b-a == 24 { solutions = append(solutions, "d+c+b-a") }
	
		if d+c+a-b == 24 { solutions = append(solutions, "d+c+a-b") }
	
		if d+a+c-b == 24 { solutions = append(solutions, "d+a+c-b") }
	
		if d+a+b-c == 24 { solutions = append(solutions, "d+a+b-c") }
	
		if a+b+c*d == 24 { solutions = append(solutions, "a+b+c*d") }
	
		if a+b+d*c == 24 { solutions = append(solutions, "a+b+d*c") }
	
		if a+c+b*d == 24 { solutions = append(solutions, "a+c+b*d") }
	
		if a+c+d*b == 24 { solutions = append(solutions, "a+c+d*b") }
	
		if a+d+c*b == 24 { solutions = append(solutions, "a+d+c*b") }
	
		if a+d+b*c == 24 { solutions = append(solutions, "a+d+b*c") }
	
		if b+a+c*d == 24 { solutions = append(solutions, "b+a+c*d") }
	
		if b+a+d*c == 24 { solutions = append(solutions, "b+a+d*c") }
	
		if b+c+a*d == 24 { solutions = append(solutions, "b+c+a*d") }
	
		if b+c+d*a == 24 { solutions = append(solutions, "b+c+d*a") }
	
		if b+d+c*a == 24 { solutions = append(solutions, "b+d+c*a") }
	
		if b+d+a*c == 24 { solutions = append(solutions, "b+d+a*c") }
	
		if c+b+a*d == 24 { solutions = append(solutions, "c+b+a*d") }
	
		if c+b+d*a == 24 { solutions = append(solutions, "c+b+d*a") }
	
		if c+a+b*d == 24 { solutions = append(solutions, "c+a+b*d") }
	
		if c+a+d*b == 24 { solutions = append(solutions, "c+a+d*b") }
	
		if c+d+a*b == 24 { solutions = append(solutions, "c+d+a*b") }
	
		if c+d+b*a == 24 { solutions = append(solutions, "c+d+b*a") }
	
		if d+b+c*a == 24 { solutions = append(solutions, "d+b+c*a") }
	
		if d+b+a*c == 24 { solutions = append(solutions, "d+b+a*c") }
	
		if d+c+b*a == 24 { solutions = append(solutions, "d+c+b*a") }
	
		if d+c+a*b == 24 { solutions = append(solutions, "d+c+a*b") }
	
		if d+a+c*b == 24 { solutions = append(solutions, "d+a+c*b") }
	
		if d+a+b*c == 24 { solutions = append(solutions, "d+a+b*c") }
	
		if a+b+c/d == 24 { solutions = append(solutions, "a+b+c/d") }
	
		if a+b+d/c == 24 { solutions = append(solutions, "a+b+d/c") }
	
		if a+c+b/d == 24 { solutions = append(solutions, "a+c+b/d") }
	
		if a+c+d/b == 24 { solutions = append(solutions, "a+c+d/b") }
	
		if a+d+c/b == 24 { solutions = append(solutions, "a+d+c/b") }
	
		if a+d+b/c == 24 { solutions = append(solutions, "a+d+b/c") }
	
		if b+a+c/d == 24 { solutions = append(solutions, "b+a+c/d") }
	
		if b+a+d/c == 24 { solutions = append(solutions, "b+a+d/c") }
	
		if b+c+a/d == 24 { solutions = append(solutions, "b+c+a/d") }
	
		if b+c+d/a == 24 { solutions = append(solutions, "b+c+d/a") }
	
		if b+d+c/a == 24 { solutions = append(solutions, "b+d+c/a") }
	
		if b+d+a/c == 24 { solutions = append(solutions, "b+d+a/c") }
	
		if c+b+a/d == 24 { solutions = append(solutions, "c+b+a/d") }
	
		if c+b+d/a == 24 { solutions = append(solutions, "c+b+d/a") }
	
		if c+a+b/d == 24 { solutions = append(solutions, "c+a+b/d") }
	
		if c+a+d/b == 24 { solutions = append(solutions, "c+a+d/b") }
	
		if c+d+a/b == 24 { solutions = append(solutions, "c+d+a/b") }
	
		if c+d+b/a == 24 { solutions = append(solutions, "c+d+b/a") }
	
		if d+b+c/a == 24 { solutions = append(solutions, "d+b+c/a") }
	
		if d+b+a/c == 24 { solutions = append(solutions, "d+b+a/c") }
	
		if d+c+b/a == 24 { solutions = append(solutions, "d+c+b/a") }
	
		if d+c+a/b == 24 { solutions = append(solutions, "d+c+a/b") }
	
		if d+a+c/b == 24 { solutions = append(solutions, "d+a+c/b") }
	
		if d+a+b/c == 24 { solutions = append(solutions, "d+a+b/c") }
	
		if a+b-c+d == 24 { solutions = append(solutions, "a+b-c+d") }
	
		if a+b-d+c == 24 { solutions = append(solutions, "a+b-d+c") }
	
		if a+c-b+d == 24 { solutions = append(solutions, "a+c-b+d") }
	
		if a+c-d+b == 24 { solutions = append(solutions, "a+c-d+b") }
	
		if a+d-c+b == 24 { solutions = append(solutions, "a+d-c+b") }
	
		if a+d-b+c == 24 { solutions = append(solutions, "a+d-b+c") }
	
		if b+a-c+d == 24 { solutions = append(solutions, "b+a-c+d") }
	
		if b+a-d+c == 24 { solutions = append(solutions, "b+a-d+c") }
	
		if b+c-a+d == 24 { solutions = append(solutions, "b+c-a+d") }
	
		if b+c-d+a == 24 { solutions = append(solutions, "b+c-d+a") }
	
		if b+d-c+a == 24 { solutions = append(solutions, "b+d-c+a") }
	
		if b+d-a+c == 24 { solutions = append(solutions, "b+d-a+c") }
	
		if c+b-a+d == 24 { solutions = append(solutions, "c+b-a+d") }
	
		if c+b-d+a == 24 { solutions = append(solutions, "c+b-d+a") }
	
		if c+a-b+d == 24 { solutions = append(solutions, "c+a-b+d") }
	
		if c+a-d+b == 24 { solutions = append(solutions, "c+a-d+b") }
	
		if c+d-a+b == 24 { solutions = append(solutions, "c+d-a+b") }
	
		if c+d-b+a == 24 { solutions = append(solutions, "c+d-b+a") }
	
		if d+b-c+a == 24 { solutions = append(solutions, "d+b-c+a") }
	
		if d+b-a+c == 24 { solutions = append(solutions, "d+b-a+c") }
	
		if d+c-b+a == 24 { solutions = append(solutions, "d+c-b+a") }
	
		if d+c-a+b == 24 { solutions = append(solutions, "d+c-a+b") }
	
		if d+a-c+b == 24 { solutions = append(solutions, "d+a-c+b") }
	
		if d+a-b+c == 24 { solutions = append(solutions, "d+a-b+c") }
	
		if a+b-c-d == 24 { solutions = append(solutions, "a+b-c-d") }
	
		if a+b-d-c == 24 { solutions = append(solutions, "a+b-d-c") }
	
		if a+c-b-d == 24 { solutions = append(solutions, "a+c-b-d") }
	
		if a+c-d-b == 24 { solutions = append(solutions, "a+c-d-b") }
	
		if a+d-c-b == 24 { solutions = append(solutions, "a+d-c-b") }
	
		if a+d-b-c == 24 { solutions = append(solutions, "a+d-b-c") }
	
		if b+a-c-d == 24 { solutions = append(solutions, "b+a-c-d") }
	
		if b+a-d-c == 24 { solutions = append(solutions, "b+a-d-c") }
	
		if b+c-a-d == 24 { solutions = append(solutions, "b+c-a-d") }
	
		if b+c-d-a == 24 { solutions = append(solutions, "b+c-d-a") }
	
		if b+d-c-a == 24 { solutions = append(solutions, "b+d-c-a") }
	
		if b+d-a-c == 24 { solutions = append(solutions, "b+d-a-c") }
	
		if c+b-a-d == 24 { solutions = append(solutions, "c+b-a-d") }
	
		if c+b-d-a == 24 { solutions = append(solutions, "c+b-d-a") }
	
		if c+a-b-d == 24 { solutions = append(solutions, "c+a-b-d") }
	
		if c+a-d-b == 24 { solutions = append(solutions, "c+a-d-b") }
	
		if c+d-a-b == 24 { solutions = append(solutions, "c+d-a-b") }
	
		if c+d-b-a == 24 { solutions = append(solutions, "c+d-b-a") }
	
		if d+b-c-a == 24 { solutions = append(solutions, "d+b-c-a") }
	
		if d+b-a-c == 24 { solutions = append(solutions, "d+b-a-c") }
	
		if d+c-b-a == 24 { solutions = append(solutions, "d+c-b-a") }
	
		if d+c-a-b == 24 { solutions = append(solutions, "d+c-a-b") }
	
		if d+a-c-b == 24 { solutions = append(solutions, "d+a-c-b") }
	
		if d+a-b-c == 24 { solutions = append(solutions, "d+a-b-c") }
	
		if a+b-c*d == 24 { solutions = append(solutions, "a+b-c*d") }
	
		if a+b-d*c == 24 { solutions = append(solutions, "a+b-d*c") }
	
		if a+c-b*d == 24 { solutions = append(solutions, "a+c-b*d") }
	
		if a+c-d*b == 24 { solutions = append(solutions, "a+c-d*b") }
	
		if a+d-c*b == 24 { solutions = append(solutions, "a+d-c*b") }
	
		if a+d-b*c == 24 { solutions = append(solutions, "a+d-b*c") }
	
		if b+a-c*d == 24 { solutions = append(solutions, "b+a-c*d") }
	
		if b+a-d*c == 24 { solutions = append(solutions, "b+a-d*c") }
	
		if b+c-a*d == 24 { solutions = append(solutions, "b+c-a*d") }
	
		if b+c-d*a == 24 { solutions = append(solutions, "b+c-d*a") }
	
		if b+d-c*a == 24 { solutions = append(solutions, "b+d-c*a") }
	
		if b+d-a*c == 24 { solutions = append(solutions, "b+d-a*c") }
	
		if c+b-a*d == 24 { solutions = append(solutions, "c+b-a*d") }
	
		if c+b-d*a == 24 { solutions = append(solutions, "c+b-d*a") }
	
		if c+a-b*d == 24 { solutions = append(solutions, "c+a-b*d") }
	
		if c+a-d*b == 24 { solutions = append(solutions, "c+a-d*b") }
	
		if c+d-a*b == 24 { solutions = append(solutions, "c+d-a*b") }
	
		if c+d-b*a == 24 { solutions = append(solutions, "c+d-b*a") }
	
		if d+b-c*a == 24 { solutions = append(solutions, "d+b-c*a") }
	
		if d+b-a*c == 24 { solutions = append(solutions, "d+b-a*c") }
	
		if d+c-b*a == 24 { solutions = append(solutions, "d+c-b*a") }
	
		if d+c-a*b == 24 { solutions = append(solutions, "d+c-a*b") }
	
		if d+a-c*b == 24 { solutions = append(solutions, "d+a-c*b") }
	
		if d+a-b*c == 24 { solutions = append(solutions, "d+a-b*c") }
	
		if a+b-c/d == 24 { solutions = append(solutions, "a+b-c/d") }
	
		if a+b-d/c == 24 { solutions = append(solutions, "a+b-d/c") }
	
		if a+c-b/d == 24 { solutions = append(solutions, "a+c-b/d") }
	
		if a+c-d/b == 24 { solutions = append(solutions, "a+c-d/b") }
	
		if a+d-c/b == 24 { solutions = append(solutions, "a+d-c/b") }
	
		if a+d-b/c == 24 { solutions = append(solutions, "a+d-b/c") }
	
		if b+a-c/d == 24 { solutions = append(solutions, "b+a-c/d") }
	
		if b+a-d/c == 24 { solutions = append(solutions, "b+a-d/c") }
	
		if b+c-a/d == 24 { solutions = append(solutions, "b+c-a/d") }
	
		if b+c-d/a == 24 { solutions = append(solutions, "b+c-d/a") }
	
		if b+d-c/a == 24 { solutions = append(solutions, "b+d-c/a") }
	
		if b+d-a/c == 24 { solutions = append(solutions, "b+d-a/c") }
	
		if c+b-a/d == 24 { solutions = append(solutions, "c+b-a/d") }
	
		if c+b-d/a == 24 { solutions = append(solutions, "c+b-d/a") }
	
		if c+a-b/d == 24 { solutions = append(solutions, "c+a-b/d") }
	
		if c+a-d/b == 24 { solutions = append(solutions, "c+a-d/b") }
	
		if c+d-a/b == 24 { solutions = append(solutions, "c+d-a/b") }
	
		if c+d-b/a == 24 { solutions = append(solutions, "c+d-b/a") }
	
		if d+b-c/a == 24 { solutions = append(solutions, "d+b-c/a") }
	
		if d+b-a/c == 24 { solutions = append(solutions, "d+b-a/c") }
	
		if d+c-b/a == 24 { solutions = append(solutions, "d+c-b/a") }
	
		if d+c-a/b == 24 { solutions = append(solutions, "d+c-a/b") }
	
		if d+a-c/b == 24 { solutions = append(solutions, "d+a-c/b") }
	
		if d+a-b/c == 24 { solutions = append(solutions, "d+a-b/c") }
	
		if a+b*c+d == 24 { solutions = append(solutions, "a+b*c+d") }
	
		if a+b*d+c == 24 { solutions = append(solutions, "a+b*d+c") }
	
		if a+c*b+d == 24 { solutions = append(solutions, "a+c*b+d") }
	
		if a+c*d+b == 24 { solutions = append(solutions, "a+c*d+b") }
	
		if a+d*c+b == 24 { solutions = append(solutions, "a+d*c+b") }
	
		if a+d*b+c == 24 { solutions = append(solutions, "a+d*b+c") }
	
		if b+a*c+d == 24 { solutions = append(solutions, "b+a*c+d") }
	
		if b+a*d+c == 24 { solutions = append(solutions, "b+a*d+c") }
	
		if b+c*a+d == 24 { solutions = append(solutions, "b+c*a+d") }
	
		if b+c*d+a == 24 { solutions = append(solutions, "b+c*d+a") }
	
		if b+d*c+a == 24 { solutions = append(solutions, "b+d*c+a") }
	
		if b+d*a+c == 24 { solutions = append(solutions, "b+d*a+c") }
	
		if c+b*a+d == 24 { solutions = append(solutions, "c+b*a+d") }
	
		if c+b*d+a == 24 { solutions = append(solutions, "c+b*d+a") }
	
		if c+a*b+d == 24 { solutions = append(solutions, "c+a*b+d") }
	
		if c+a*d+b == 24 { solutions = append(solutions, "c+a*d+b") }
	
		if c+d*a+b == 24 { solutions = append(solutions, "c+d*a+b") }
	
		if c+d*b+a == 24 { solutions = append(solutions, "c+d*b+a") }
	
		if d+b*c+a == 24 { solutions = append(solutions, "d+b*c+a") }
	
		if d+b*a+c == 24 { solutions = append(solutions, "d+b*a+c") }
	
		if d+c*b+a == 24 { solutions = append(solutions, "d+c*b+a") }
	
		if d+c*a+b == 24 { solutions = append(solutions, "d+c*a+b") }
	
		if d+a*c+b == 24 { solutions = append(solutions, "d+a*c+b") }
	
		if d+a*b+c == 24 { solutions = append(solutions, "d+a*b+c") }
	
		if a+b*c-d == 24 { solutions = append(solutions, "a+b*c-d") }
	
		if a+b*d-c == 24 { solutions = append(solutions, "a+b*d-c") }
	
		if a+c*b-d == 24 { solutions = append(solutions, "a+c*b-d") }
	
		if a+c*d-b == 24 { solutions = append(solutions, "a+c*d-b") }
	
		if a+d*c-b == 24 { solutions = append(solutions, "a+d*c-b") }
	
		if a+d*b-c == 24 { solutions = append(solutions, "a+d*b-c") }
	
		if b+a*c-d == 24 { solutions = append(solutions, "b+a*c-d") }
	
		if b+a*d-c == 24 { solutions = append(solutions, "b+a*d-c") }
	
		if b+c*a-d == 24 { solutions = append(solutions, "b+c*a-d") }
	
		if b+c*d-a == 24 { solutions = append(solutions, "b+c*d-a") }
	
		if b+d*c-a == 24 { solutions = append(solutions, "b+d*c-a") }
	
		if b+d*a-c == 24 { solutions = append(solutions, "b+d*a-c") }
	
		if c+b*a-d == 24 { solutions = append(solutions, "c+b*a-d") }
	
		if c+b*d-a == 24 { solutions = append(solutions, "c+b*d-a") }
	
		if c+a*b-d == 24 { solutions = append(solutions, "c+a*b-d") }
	
		if c+a*d-b == 24 { solutions = append(solutions, "c+a*d-b") }
	
		if c+d*a-b == 24 { solutions = append(solutions, "c+d*a-b") }
	
		if c+d*b-a == 24 { solutions = append(solutions, "c+d*b-a") }
	
		if d+b*c-a == 24 { solutions = append(solutions, "d+b*c-a") }
	
		if d+b*a-c == 24 { solutions = append(solutions, "d+b*a-c") }
	
		if d+c*b-a == 24 { solutions = append(solutions, "d+c*b-a") }
	
		if d+c*a-b == 24 { solutions = append(solutions, "d+c*a-b") }
	
		if d+a*c-b == 24 { solutions = append(solutions, "d+a*c-b") }
	
		if d+a*b-c == 24 { solutions = append(solutions, "d+a*b-c") }
	
		if a+b*c*d == 24 { solutions = append(solutions, "a+b*c*d") }
	
		if a+b*d*c == 24 { solutions = append(solutions, "a+b*d*c") }
	
		if a+c*b*d == 24 { solutions = append(solutions, "a+c*b*d") }
	
		if a+c*d*b == 24 { solutions = append(solutions, "a+c*d*b") }
	
		if a+d*c*b == 24 { solutions = append(solutions, "a+d*c*b") }
	
		if a+d*b*c == 24 { solutions = append(solutions, "a+d*b*c") }
	
		if b+a*c*d == 24 { solutions = append(solutions, "b+a*c*d") }
	
		if b+a*d*c == 24 { solutions = append(solutions, "b+a*d*c") }
	
		if b+c*a*d == 24 { solutions = append(solutions, "b+c*a*d") }
	
		if b+c*d*a == 24 { solutions = append(solutions, "b+c*d*a") }
	
		if b+d*c*a == 24 { solutions = append(solutions, "b+d*c*a") }
	
		if b+d*a*c == 24 { solutions = append(solutions, "b+d*a*c") }
	
		if c+b*a*d == 24 { solutions = append(solutions, "c+b*a*d") }
	
		if c+b*d*a == 24 { solutions = append(solutions, "c+b*d*a") }
	
		if c+a*b*d == 24 { solutions = append(solutions, "c+a*b*d") }
	
		if c+a*d*b == 24 { solutions = append(solutions, "c+a*d*b") }
	
		if c+d*a*b == 24 { solutions = append(solutions, "c+d*a*b") }
	
		if c+d*b*a == 24 { solutions = append(solutions, "c+d*b*a") }
	
		if d+b*c*a == 24 { solutions = append(solutions, "d+b*c*a") }
	
		if d+b*a*c == 24 { solutions = append(solutions, "d+b*a*c") }
	
		if d+c*b*a == 24 { solutions = append(solutions, "d+c*b*a") }
	
		if d+c*a*b == 24 { solutions = append(solutions, "d+c*a*b") }
	
		if d+a*c*b == 24 { solutions = append(solutions, "d+a*c*b") }
	
		if d+a*b*c == 24 { solutions = append(solutions, "d+a*b*c") }
	
		if a+b*c/d == 24 { solutions = append(solutions, "a+b*c/d") }
	
		if a+b*d/c == 24 { solutions = append(solutions, "a+b*d/c") }
	
		if a+c*b/d == 24 { solutions = append(solutions, "a+c*b/d") }
	
		if a+c*d/b == 24 { solutions = append(solutions, "a+c*d/b") }
	
		if a+d*c/b == 24 { solutions = append(solutions, "a+d*c/b") }
	
		if a+d*b/c == 24 { solutions = append(solutions, "a+d*b/c") }
	
		if b+a*c/d == 24 { solutions = append(solutions, "b+a*c/d") }
	
		if b+a*d/c == 24 { solutions = append(solutions, "b+a*d/c") }
	
		if b+c*a/d == 24 { solutions = append(solutions, "b+c*a/d") }
	
		if b+c*d/a == 24 { solutions = append(solutions, "b+c*d/a") }
	
		if b+d*c/a == 24 { solutions = append(solutions, "b+d*c/a") }
	
		if b+d*a/c == 24 { solutions = append(solutions, "b+d*a/c") }
	
		if c+b*a/d == 24 { solutions = append(solutions, "c+b*a/d") }
	
		if c+b*d/a == 24 { solutions = append(solutions, "c+b*d/a") }
	
		if c+a*b/d == 24 { solutions = append(solutions, "c+a*b/d") }
	
		if c+a*d/b == 24 { solutions = append(solutions, "c+a*d/b") }
	
		if c+d*a/b == 24 { solutions = append(solutions, "c+d*a/b") }
	
		if c+d*b/a == 24 { solutions = append(solutions, "c+d*b/a") }
	
		if d+b*c/a == 24 { solutions = append(solutions, "d+b*c/a") }
	
		if d+b*a/c == 24 { solutions = append(solutions, "d+b*a/c") }
	
		if d+c*b/a == 24 { solutions = append(solutions, "d+c*b/a") }
	
		if d+c*a/b == 24 { solutions = append(solutions, "d+c*a/b") }
	
		if d+a*c/b == 24 { solutions = append(solutions, "d+a*c/b") }
	
		if d+a*b/c == 24 { solutions = append(solutions, "d+a*b/c") }
	
		if a+b/c+d == 24 { solutions = append(solutions, "a+b/c+d") }
	
		if a+b/d+c == 24 { solutions = append(solutions, "a+b/d+c") }
	
		if a+c/b+d == 24 { solutions = append(solutions, "a+c/b+d") }
	
		if a+c/d+b == 24 { solutions = append(solutions, "a+c/d+b") }
	
		if a+d/c+b == 24 { solutions = append(solutions, "a+d/c+b") }
	
		if a+d/b+c == 24 { solutions = append(solutions, "a+d/b+c") }
	
		if b+a/c+d == 24 { solutions = append(solutions, "b+a/c+d") }
	
		if b+a/d+c == 24 { solutions = append(solutions, "b+a/d+c") }
	
		if b+c/a+d == 24 { solutions = append(solutions, "b+c/a+d") }
	
		if b+c/d+a == 24 { solutions = append(solutions, "b+c/d+a") }
	
		if b+d/c+a == 24 { solutions = append(solutions, "b+d/c+a") }
	
		if b+d/a+c == 24 { solutions = append(solutions, "b+d/a+c") }
	
		if c+b/a+d == 24 { solutions = append(solutions, "c+b/a+d") }
	
		if c+b/d+a == 24 { solutions = append(solutions, "c+b/d+a") }
	
		if c+a/b+d == 24 { solutions = append(solutions, "c+a/b+d") }
	
		if c+a/d+b == 24 { solutions = append(solutions, "c+a/d+b") }
	
		if c+d/a+b == 24 { solutions = append(solutions, "c+d/a+b") }
	
		if c+d/b+a == 24 { solutions = append(solutions, "c+d/b+a") }
	
		if d+b/c+a == 24 { solutions = append(solutions, "d+b/c+a") }
	
		if d+b/a+c == 24 { solutions = append(solutions, "d+b/a+c") }
	
		if d+c/b+a == 24 { solutions = append(solutions, "d+c/b+a") }
	
		if d+c/a+b == 24 { solutions = append(solutions, "d+c/a+b") }
	
		if d+a/c+b == 24 { solutions = append(solutions, "d+a/c+b") }
	
		if d+a/b+c == 24 { solutions = append(solutions, "d+a/b+c") }
	
		if a+b/c-d == 24 { solutions = append(solutions, "a+b/c-d") }
	
		if a+b/d-c == 24 { solutions = append(solutions, "a+b/d-c") }
	
		if a+c/b-d == 24 { solutions = append(solutions, "a+c/b-d") }
	
		if a+c/d-b == 24 { solutions = append(solutions, "a+c/d-b") }
	
		if a+d/c-b == 24 { solutions = append(solutions, "a+d/c-b") }
	
		if a+d/b-c == 24 { solutions = append(solutions, "a+d/b-c") }
	
		if b+a/c-d == 24 { solutions = append(solutions, "b+a/c-d") }
	
		if b+a/d-c == 24 { solutions = append(solutions, "b+a/d-c") }
	
		if b+c/a-d == 24 { solutions = append(solutions, "b+c/a-d") }
	
		if b+c/d-a == 24 { solutions = append(solutions, "b+c/d-a") }
	
		if b+d/c-a == 24 { solutions = append(solutions, "b+d/c-a") }
	
		if b+d/a-c == 24 { solutions = append(solutions, "b+d/a-c") }
	
		if c+b/a-d == 24 { solutions = append(solutions, "c+b/a-d") }
	
		if c+b/d-a == 24 { solutions = append(solutions, "c+b/d-a") }
	
		if c+a/b-d == 24 { solutions = append(solutions, "c+a/b-d") }
	
		if c+a/d-b == 24 { solutions = append(solutions, "c+a/d-b") }
	
		if c+d/a-b == 24 { solutions = append(solutions, "c+d/a-b") }
	
		if c+d/b-a == 24 { solutions = append(solutions, "c+d/b-a") }
	
		if d+b/c-a == 24 { solutions = append(solutions, "d+b/c-a") }
	
		if d+b/a-c == 24 { solutions = append(solutions, "d+b/a-c") }
	
		if d+c/b-a == 24 { solutions = append(solutions, "d+c/b-a") }
	
		if d+c/a-b == 24 { solutions = append(solutions, "d+c/a-b") }
	
		if d+a/c-b == 24 { solutions = append(solutions, "d+a/c-b") }
	
		if d+a/b-c == 24 { solutions = append(solutions, "d+a/b-c") }
	
		if a+b/c*d == 24 { solutions = append(solutions, "a+b/c*d") }
	
		if a+b/d*c == 24 { solutions = append(solutions, "a+b/d*c") }
	
		if a+c/b*d == 24 { solutions = append(solutions, "a+c/b*d") }
	
		if a+c/d*b == 24 { solutions = append(solutions, "a+c/d*b") }
	
		if a+d/c*b == 24 { solutions = append(solutions, "a+d/c*b") }
	
		if a+d/b*c == 24 { solutions = append(solutions, "a+d/b*c") }
	
		if b+a/c*d == 24 { solutions = append(solutions, "b+a/c*d") }
	
		if b+a/d*c == 24 { solutions = append(solutions, "b+a/d*c") }
	
		if b+c/a*d == 24 { solutions = append(solutions, "b+c/a*d") }
	
		if b+c/d*a == 24 { solutions = append(solutions, "b+c/d*a") }
	
		if b+d/c*a == 24 { solutions = append(solutions, "b+d/c*a") }
	
		if b+d/a*c == 24 { solutions = append(solutions, "b+d/a*c") }
	
		if c+b/a*d == 24 { solutions = append(solutions, "c+b/a*d") }
	
		if c+b/d*a == 24 { solutions = append(solutions, "c+b/d*a") }
	
		if c+a/b*d == 24 { solutions = append(solutions, "c+a/b*d") }
	
		if c+a/d*b == 24 { solutions = append(solutions, "c+a/d*b") }
	
		if c+d/a*b == 24 { solutions = append(solutions, "c+d/a*b") }
	
		if c+d/b*a == 24 { solutions = append(solutions, "c+d/b*a") }
	
		if d+b/c*a == 24 { solutions = append(solutions, "d+b/c*a") }
	
		if d+b/a*c == 24 { solutions = append(solutions, "d+b/a*c") }
	
		if d+c/b*a == 24 { solutions = append(solutions, "d+c/b*a") }
	
		if d+c/a*b == 24 { solutions = append(solutions, "d+c/a*b") }
	
		if d+a/c*b == 24 { solutions = append(solutions, "d+a/c*b") }
	
		if d+a/b*c == 24 { solutions = append(solutions, "d+a/b*c") }
	
		if a+b/c/d == 24 { solutions = append(solutions, "a+b/c/d") }
	
		if a+b/d/c == 24 { solutions = append(solutions, "a+b/d/c") }
	
		if a+c/b/d == 24 { solutions = append(solutions, "a+c/b/d") }
	
		if a+c/d/b == 24 { solutions = append(solutions, "a+c/d/b") }
	
		if a+d/c/b == 24 { solutions = append(solutions, "a+d/c/b") }
	
		if a+d/b/c == 24 { solutions = append(solutions, "a+d/b/c") }
	
		if b+a/c/d == 24 { solutions = append(solutions, "b+a/c/d") }
	
		if b+a/d/c == 24 { solutions = append(solutions, "b+a/d/c") }
	
		if b+c/a/d == 24 { solutions = append(solutions, "b+c/a/d") }
	
		if b+c/d/a == 24 { solutions = append(solutions, "b+c/d/a") }
	
		if b+d/c/a == 24 { solutions = append(solutions, "b+d/c/a") }
	
		if b+d/a/c == 24 { solutions = append(solutions, "b+d/a/c") }
	
		if c+b/a/d == 24 { solutions = append(solutions, "c+b/a/d") }
	
		if c+b/d/a == 24 { solutions = append(solutions, "c+b/d/a") }
	
		if c+a/b/d == 24 { solutions = append(solutions, "c+a/b/d") }
	
		if c+a/d/b == 24 { solutions = append(solutions, "c+a/d/b") }
	
		if c+d/a/b == 24 { solutions = append(solutions, "c+d/a/b") }
	
		if c+d/b/a == 24 { solutions = append(solutions, "c+d/b/a") }
	
		if d+b/c/a == 24 { solutions = append(solutions, "d+b/c/a") }
	
		if d+b/a/c == 24 { solutions = append(solutions, "d+b/a/c") }
	
		if d+c/b/a == 24 { solutions = append(solutions, "d+c/b/a") }
	
		if d+c/a/b == 24 { solutions = append(solutions, "d+c/a/b") }
	
		if d+a/c/b == 24 { solutions = append(solutions, "d+a/c/b") }
	
		if d+a/b/c == 24 { solutions = append(solutions, "d+a/b/c") }
	
		if a-b+c+d == 24 { solutions = append(solutions, "a-b+c+d") }
	
		if a-b+d+c == 24 { solutions = append(solutions, "a-b+d+c") }
	
		if a-c+b+d == 24 { solutions = append(solutions, "a-c+b+d") }
	
		if a-c+d+b == 24 { solutions = append(solutions, "a-c+d+b") }
	
		if a-d+c+b == 24 { solutions = append(solutions, "a-d+c+b") }
	
		if a-d+b+c == 24 { solutions = append(solutions, "a-d+b+c") }
	
		if b-a+c+d == 24 { solutions = append(solutions, "b-a+c+d") }
	
		if b-a+d+c == 24 { solutions = append(solutions, "b-a+d+c") }
	
		if b-c+a+d == 24 { solutions = append(solutions, "b-c+a+d") }
	
		if b-c+d+a == 24 { solutions = append(solutions, "b-c+d+a") }
	
		if b-d+c+a == 24 { solutions = append(solutions, "b-d+c+a") }
	
		if b-d+a+c == 24 { solutions = append(solutions, "b-d+a+c") }
	
		if c-b+a+d == 24 { solutions = append(solutions, "c-b+a+d") }
	
		if c-b+d+a == 24 { solutions = append(solutions, "c-b+d+a") }
	
		if c-a+b+d == 24 { solutions = append(solutions, "c-a+b+d") }
	
		if c-a+d+b == 24 { solutions = append(solutions, "c-a+d+b") }
	
		if c-d+a+b == 24 { solutions = append(solutions, "c-d+a+b") }
	
		if c-d+b+a == 24 { solutions = append(solutions, "c-d+b+a") }
	
		if d-b+c+a == 24 { solutions = append(solutions, "d-b+c+a") }
	
		if d-b+a+c == 24 { solutions = append(solutions, "d-b+a+c") }
	
		if d-c+b+a == 24 { solutions = append(solutions, "d-c+b+a") }
	
		if d-c+a+b == 24 { solutions = append(solutions, "d-c+a+b") }
	
		if d-a+c+b == 24 { solutions = append(solutions, "d-a+c+b") }
	
		if d-a+b+c == 24 { solutions = append(solutions, "d-a+b+c") }
	
		if a-b+c-d == 24 { solutions = append(solutions, "a-b+c-d") }
	
		if a-b+d-c == 24 { solutions = append(solutions, "a-b+d-c") }
	
		if a-c+b-d == 24 { solutions = append(solutions, "a-c+b-d") }
	
		if a-c+d-b == 24 { solutions = append(solutions, "a-c+d-b") }
	
		if a-d+c-b == 24 { solutions = append(solutions, "a-d+c-b") }
	
		if a-d+b-c == 24 { solutions = append(solutions, "a-d+b-c") }
	
		if b-a+c-d == 24 { solutions = append(solutions, "b-a+c-d") }
	
		if b-a+d-c == 24 { solutions = append(solutions, "b-a+d-c") }
	
		if b-c+a-d == 24 { solutions = append(solutions, "b-c+a-d") }
	
		if b-c+d-a == 24 { solutions = append(solutions, "b-c+d-a") }
	
		if b-d+c-a == 24 { solutions = append(solutions, "b-d+c-a") }
	
		if b-d+a-c == 24 { solutions = append(solutions, "b-d+a-c") }
	
		if c-b+a-d == 24 { solutions = append(solutions, "c-b+a-d") }
	
		if c-b+d-a == 24 { solutions = append(solutions, "c-b+d-a") }
	
		if c-a+b-d == 24 { solutions = append(solutions, "c-a+b-d") }
	
		if c-a+d-b == 24 { solutions = append(solutions, "c-a+d-b") }
	
		if c-d+a-b == 24 { solutions = append(solutions, "c-d+a-b") }
	
		if c-d+b-a == 24 { solutions = append(solutions, "c-d+b-a") }
	
		if d-b+c-a == 24 { solutions = append(solutions, "d-b+c-a") }
	
		if d-b+a-c == 24 { solutions = append(solutions, "d-b+a-c") }
	
		if d-c+b-a == 24 { solutions = append(solutions, "d-c+b-a") }
	
		if d-c+a-b == 24 { solutions = append(solutions, "d-c+a-b") }
	
		if d-a+c-b == 24 { solutions = append(solutions, "d-a+c-b") }
	
		if d-a+b-c == 24 { solutions = append(solutions, "d-a+b-c") }
	
		if a-b+c*d == 24 { solutions = append(solutions, "a-b+c*d") }
	
		if a-b+d*c == 24 { solutions = append(solutions, "a-b+d*c") }
	
		if a-c+b*d == 24 { solutions = append(solutions, "a-c+b*d") }
	
		if a-c+d*b == 24 { solutions = append(solutions, "a-c+d*b") }
	
		if a-d+c*b == 24 { solutions = append(solutions, "a-d+c*b") }
	
		if a-d+b*c == 24 { solutions = append(solutions, "a-d+b*c") }
	
		if b-a+c*d == 24 { solutions = append(solutions, "b-a+c*d") }
	
		if b-a+d*c == 24 { solutions = append(solutions, "b-a+d*c") }
	
		if b-c+a*d == 24 { solutions = append(solutions, "b-c+a*d") }
	
		if b-c+d*a == 24 { solutions = append(solutions, "b-c+d*a") }
	
		if b-d+c*a == 24 { solutions = append(solutions, "b-d+c*a") }
	
		if b-d+a*c == 24 { solutions = append(solutions, "b-d+a*c") }
	
		if c-b+a*d == 24 { solutions = append(solutions, "c-b+a*d") }
	
		if c-b+d*a == 24 { solutions = append(solutions, "c-b+d*a") }
	
		if c-a+b*d == 24 { solutions = append(solutions, "c-a+b*d") }
	
		if c-a+d*b == 24 { solutions = append(solutions, "c-a+d*b") }
	
		if c-d+a*b == 24 { solutions = append(solutions, "c-d+a*b") }
	
		if c-d+b*a == 24 { solutions = append(solutions, "c-d+b*a") }
	
		if d-b+c*a == 24 { solutions = append(solutions, "d-b+c*a") }
	
		if d-b+a*c == 24 { solutions = append(solutions, "d-b+a*c") }
	
		if d-c+b*a == 24 { solutions = append(solutions, "d-c+b*a") }
	
		if d-c+a*b == 24 { solutions = append(solutions, "d-c+a*b") }
	
		if d-a+c*b == 24 { solutions = append(solutions, "d-a+c*b") }
	
		if d-a+b*c == 24 { solutions = append(solutions, "d-a+b*c") }
	
		if a-b+c/d == 24 { solutions = append(solutions, "a-b+c/d") }
	
		if a-b+d/c == 24 { solutions = append(solutions, "a-b+d/c") }
	
		if a-c+b/d == 24 { solutions = append(solutions, "a-c+b/d") }
	
		if a-c+d/b == 24 { solutions = append(solutions, "a-c+d/b") }
	
		if a-d+c/b == 24 { solutions = append(solutions, "a-d+c/b") }
	
		if a-d+b/c == 24 { solutions = append(solutions, "a-d+b/c") }
	
		if b-a+c/d == 24 { solutions = append(solutions, "b-a+c/d") }
	
		if b-a+d/c == 24 { solutions = append(solutions, "b-a+d/c") }
	
		if b-c+a/d == 24 { solutions = append(solutions, "b-c+a/d") }
	
		if b-c+d/a == 24 { solutions = append(solutions, "b-c+d/a") }
	
		if b-d+c/a == 24 { solutions = append(solutions, "b-d+c/a") }
	
		if b-d+a/c == 24 { solutions = append(solutions, "b-d+a/c") }
	
		if c-b+a/d == 24 { solutions = append(solutions, "c-b+a/d") }
	
		if c-b+d/a == 24 { solutions = append(solutions, "c-b+d/a") }
	
		if c-a+b/d == 24 { solutions = append(solutions, "c-a+b/d") }
	
		if c-a+d/b == 24 { solutions = append(solutions, "c-a+d/b") }
	
		if c-d+a/b == 24 { solutions = append(solutions, "c-d+a/b") }
	
		if c-d+b/a == 24 { solutions = append(solutions, "c-d+b/a") }
	
		if d-b+c/a == 24 { solutions = append(solutions, "d-b+c/a") }
	
		if d-b+a/c == 24 { solutions = append(solutions, "d-b+a/c") }
	
		if d-c+b/a == 24 { solutions = append(solutions, "d-c+b/a") }
	
		if d-c+a/b == 24 { solutions = append(solutions, "d-c+a/b") }
	
		if d-a+c/b == 24 { solutions = append(solutions, "d-a+c/b") }
	
		if d-a+b/c == 24 { solutions = append(solutions, "d-a+b/c") }
	
		if a-b-c+d == 24 { solutions = append(solutions, "a-b-c+d") }
	
		if a-b-d+c == 24 { solutions = append(solutions, "a-b-d+c") }
	
		if a-c-b+d == 24 { solutions = append(solutions, "a-c-b+d") }
	
		if a-c-d+b == 24 { solutions = append(solutions, "a-c-d+b") }
	
		if a-d-c+b == 24 { solutions = append(solutions, "a-d-c+b") }
	
		if a-d-b+c == 24 { solutions = append(solutions, "a-d-b+c") }
	
		if b-a-c+d == 24 { solutions = append(solutions, "b-a-c+d") }
	
		if b-a-d+c == 24 { solutions = append(solutions, "b-a-d+c") }
	
		if b-c-a+d == 24 { solutions = append(solutions, "b-c-a+d") }
	
		if b-c-d+a == 24 { solutions = append(solutions, "b-c-d+a") }
	
		if b-d-c+a == 24 { solutions = append(solutions, "b-d-c+a") }
	
		if b-d-a+c == 24 { solutions = append(solutions, "b-d-a+c") }
	
		if c-b-a+d == 24 { solutions = append(solutions, "c-b-a+d") }
	
		if c-b-d+a == 24 { solutions = append(solutions, "c-b-d+a") }
	
		if c-a-b+d == 24 { solutions = append(solutions, "c-a-b+d") }
	
		if c-a-d+b == 24 { solutions = append(solutions, "c-a-d+b") }
	
		if c-d-a+b == 24 { solutions = append(solutions, "c-d-a+b") }
	
		if c-d-b+a == 24 { solutions = append(solutions, "c-d-b+a") }
	
		if d-b-c+a == 24 { solutions = append(solutions, "d-b-c+a") }
	
		if d-b-a+c == 24 { solutions = append(solutions, "d-b-a+c") }
	
		if d-c-b+a == 24 { solutions = append(solutions, "d-c-b+a") }
	
		if d-c-a+b == 24 { solutions = append(solutions, "d-c-a+b") }
	
		if d-a-c+b == 24 { solutions = append(solutions, "d-a-c+b") }
	
		if d-a-b+c == 24 { solutions = append(solutions, "d-a-b+c") }
	
		if a-b-c-d == 24 { solutions = append(solutions, "a-b-c-d") }
	
		if a-b-d-c == 24 { solutions = append(solutions, "a-b-d-c") }
	
		if a-c-b-d == 24 { solutions = append(solutions, "a-c-b-d") }
	
		if a-c-d-b == 24 { solutions = append(solutions, "a-c-d-b") }
	
		if a-d-c-b == 24 { solutions = append(solutions, "a-d-c-b") }
	
		if a-d-b-c == 24 { solutions = append(solutions, "a-d-b-c") }
	
		if b-a-c-d == 24 { solutions = append(solutions, "b-a-c-d") }
	
		if b-a-d-c == 24 { solutions = append(solutions, "b-a-d-c") }
	
		if b-c-a-d == 24 { solutions = append(solutions, "b-c-a-d") }
	
		if b-c-d-a == 24 { solutions = append(solutions, "b-c-d-a") }
	
		if b-d-c-a == 24 { solutions = append(solutions, "b-d-c-a") }
	
		if b-d-a-c == 24 { solutions = append(solutions, "b-d-a-c") }
	
		if c-b-a-d == 24 { solutions = append(solutions, "c-b-a-d") }
	
		if c-b-d-a == 24 { solutions = append(solutions, "c-b-d-a") }
	
		if c-a-b-d == 24 { solutions = append(solutions, "c-a-b-d") }
	
		if c-a-d-b == 24 { solutions = append(solutions, "c-a-d-b") }
	
		if c-d-a-b == 24 { solutions = append(solutions, "c-d-a-b") }
	
		if c-d-b-a == 24 { solutions = append(solutions, "c-d-b-a") }
	
		if d-b-c-a == 24 { solutions = append(solutions, "d-b-c-a") }
	
		if d-b-a-c == 24 { solutions = append(solutions, "d-b-a-c") }
	
		if d-c-b-a == 24 { solutions = append(solutions, "d-c-b-a") }
	
		if d-c-a-b == 24 { solutions = append(solutions, "d-c-a-b") }
	
		if d-a-c-b == 24 { solutions = append(solutions, "d-a-c-b") }
	
		if d-a-b-c == 24 { solutions = append(solutions, "d-a-b-c") }
	
		if a-b-c*d == 24 { solutions = append(solutions, "a-b-c*d") }
	
		if a-b-d*c == 24 { solutions = append(solutions, "a-b-d*c") }
	
		if a-c-b*d == 24 { solutions = append(solutions, "a-c-b*d") }
	
		if a-c-d*b == 24 { solutions = append(solutions, "a-c-d*b") }
	
		if a-d-c*b == 24 { solutions = append(solutions, "a-d-c*b") }
	
		if a-d-b*c == 24 { solutions = append(solutions, "a-d-b*c") }
	
		if b-a-c*d == 24 { solutions = append(solutions, "b-a-c*d") }
	
		if b-a-d*c == 24 { solutions = append(solutions, "b-a-d*c") }
	
		if b-c-a*d == 24 { solutions = append(solutions, "b-c-a*d") }
	
		if b-c-d*a == 24 { solutions = append(solutions, "b-c-d*a") }
	
		if b-d-c*a == 24 { solutions = append(solutions, "b-d-c*a") }
	
		if b-d-a*c == 24 { solutions = append(solutions, "b-d-a*c") }
	
		if c-b-a*d == 24 { solutions = append(solutions, "c-b-a*d") }
	
		if c-b-d*a == 24 { solutions = append(solutions, "c-b-d*a") }
	
		if c-a-b*d == 24 { solutions = append(solutions, "c-a-b*d") }
	
		if c-a-d*b == 24 { solutions = append(solutions, "c-a-d*b") }
	
		if c-d-a*b == 24 { solutions = append(solutions, "c-d-a*b") }
	
		if c-d-b*a == 24 { solutions = append(solutions, "c-d-b*a") }
	
		if d-b-c*a == 24 { solutions = append(solutions, "d-b-c*a") }
	
		if d-b-a*c == 24 { solutions = append(solutions, "d-b-a*c") }
	
		if d-c-b*a == 24 { solutions = append(solutions, "d-c-b*a") }
	
		if d-c-a*b == 24 { solutions = append(solutions, "d-c-a*b") }
	
		if d-a-c*b == 24 { solutions = append(solutions, "d-a-c*b") }
	
		if d-a-b*c == 24 { solutions = append(solutions, "d-a-b*c") }
	
		if a-b-c/d == 24 { solutions = append(solutions, "a-b-c/d") }
	
		if a-b-d/c == 24 { solutions = append(solutions, "a-b-d/c") }
	
		if a-c-b/d == 24 { solutions = append(solutions, "a-c-b/d") }
	
		if a-c-d/b == 24 { solutions = append(solutions, "a-c-d/b") }
	
		if a-d-c/b == 24 { solutions = append(solutions, "a-d-c/b") }
	
		if a-d-b/c == 24 { solutions = append(solutions, "a-d-b/c") }
	
		if b-a-c/d == 24 { solutions = append(solutions, "b-a-c/d") }
	
		if b-a-d/c == 24 { solutions = append(solutions, "b-a-d/c") }
	
		if b-c-a/d == 24 { solutions = append(solutions, "b-c-a/d") }
	
		if b-c-d/a == 24 { solutions = append(solutions, "b-c-d/a") }
	
		if b-d-c/a == 24 { solutions = append(solutions, "b-d-c/a") }
	
		if b-d-a/c == 24 { solutions = append(solutions, "b-d-a/c") }
	
		if c-b-a/d == 24 { solutions = append(solutions, "c-b-a/d") }
	
		if c-b-d/a == 24 { solutions = append(solutions, "c-b-d/a") }
	
		if c-a-b/d == 24 { solutions = append(solutions, "c-a-b/d") }
	
		if c-a-d/b == 24 { solutions = append(solutions, "c-a-d/b") }
	
		if c-d-a/b == 24 { solutions = append(solutions, "c-d-a/b") }
	
		if c-d-b/a == 24 { solutions = append(solutions, "c-d-b/a") }
	
		if d-b-c/a == 24 { solutions = append(solutions, "d-b-c/a") }
	
		if d-b-a/c == 24 { solutions = append(solutions, "d-b-a/c") }
	
		if d-c-b/a == 24 { solutions = append(solutions, "d-c-b/a") }
	
		if d-c-a/b == 24 { solutions = append(solutions, "d-c-a/b") }
	
		if d-a-c/b == 24 { solutions = append(solutions, "d-a-c/b") }
	
		if d-a-b/c == 24 { solutions = append(solutions, "d-a-b/c") }
	
		if a-b*c+d == 24 { solutions = append(solutions, "a-b*c+d") }
	
		if a-b*d+c == 24 { solutions = append(solutions, "a-b*d+c") }
	
		if a-c*b+d == 24 { solutions = append(solutions, "a-c*b+d") }
	
		if a-c*d+b == 24 { solutions = append(solutions, "a-c*d+b") }
	
		if a-d*c+b == 24 { solutions = append(solutions, "a-d*c+b") }
	
		if a-d*b+c == 24 { solutions = append(solutions, "a-d*b+c") }
	
		if b-a*c+d == 24 { solutions = append(solutions, "b-a*c+d") }
	
		if b-a*d+c == 24 { solutions = append(solutions, "b-a*d+c") }
	
		if b-c*a+d == 24 { solutions = append(solutions, "b-c*a+d") }
	
		if b-c*d+a == 24 { solutions = append(solutions, "b-c*d+a") }
	
		if b-d*c+a == 24 { solutions = append(solutions, "b-d*c+a") }
	
		if b-d*a+c == 24 { solutions = append(solutions, "b-d*a+c") }
	
		if c-b*a+d == 24 { solutions = append(solutions, "c-b*a+d") }
	
		if c-b*d+a == 24 { solutions = append(solutions, "c-b*d+a") }
	
		if c-a*b+d == 24 { solutions = append(solutions, "c-a*b+d") }
	
		if c-a*d+b == 24 { solutions = append(solutions, "c-a*d+b") }
	
		if c-d*a+b == 24 { solutions = append(solutions, "c-d*a+b") }
	
		if c-d*b+a == 24 { solutions = append(solutions, "c-d*b+a") }
	
		if d-b*c+a == 24 { solutions = append(solutions, "d-b*c+a") }
	
		if d-b*a+c == 24 { solutions = append(solutions, "d-b*a+c") }
	
		if d-c*b+a == 24 { solutions = append(solutions, "d-c*b+a") }
	
		if d-c*a+b == 24 { solutions = append(solutions, "d-c*a+b") }
	
		if d-a*c+b == 24 { solutions = append(solutions, "d-a*c+b") }
	
		if d-a*b+c == 24 { solutions = append(solutions, "d-a*b+c") }
	
		if a-b*c-d == 24 { solutions = append(solutions, "a-b*c-d") }
	
		if a-b*d-c == 24 { solutions = append(solutions, "a-b*d-c") }
	
		if a-c*b-d == 24 { solutions = append(solutions, "a-c*b-d") }
	
		if a-c*d-b == 24 { solutions = append(solutions, "a-c*d-b") }
	
		if a-d*c-b == 24 { solutions = append(solutions, "a-d*c-b") }
	
		if a-d*b-c == 24 { solutions = append(solutions, "a-d*b-c") }
	
		if b-a*c-d == 24 { solutions = append(solutions, "b-a*c-d") }
	
		if b-a*d-c == 24 { solutions = append(solutions, "b-a*d-c") }
	
		if b-c*a-d == 24 { solutions = append(solutions, "b-c*a-d") }
	
		if b-c*d-a == 24 { solutions = append(solutions, "b-c*d-a") }
	
		if b-d*c-a == 24 { solutions = append(solutions, "b-d*c-a") }
	
		if b-d*a-c == 24 { solutions = append(solutions, "b-d*a-c") }
	
		if c-b*a-d == 24 { solutions = append(solutions, "c-b*a-d") }
	
		if c-b*d-a == 24 { solutions = append(solutions, "c-b*d-a") }
	
		if c-a*b-d == 24 { solutions = append(solutions, "c-a*b-d") }
	
		if c-a*d-b == 24 { solutions = append(solutions, "c-a*d-b") }
	
		if c-d*a-b == 24 { solutions = append(solutions, "c-d*a-b") }
	
		if c-d*b-a == 24 { solutions = append(solutions, "c-d*b-a") }
	
		if d-b*c-a == 24 { solutions = append(solutions, "d-b*c-a") }
	
		if d-b*a-c == 24 { solutions = append(solutions, "d-b*a-c") }
	
		if d-c*b-a == 24 { solutions = append(solutions, "d-c*b-a") }
	
		if d-c*a-b == 24 { solutions = append(solutions, "d-c*a-b") }
	
		if d-a*c-b == 24 { solutions = append(solutions, "d-a*c-b") }
	
		if d-a*b-c == 24 { solutions = append(solutions, "d-a*b-c") }
	
		if a-b*c*d == 24 { solutions = append(solutions, "a-b*c*d") }
	
		if a-b*d*c == 24 { solutions = append(solutions, "a-b*d*c") }
	
		if a-c*b*d == 24 { solutions = append(solutions, "a-c*b*d") }
	
		if a-c*d*b == 24 { solutions = append(solutions, "a-c*d*b") }
	
		if a-d*c*b == 24 { solutions = append(solutions, "a-d*c*b") }
	
		if a-d*b*c == 24 { solutions = append(solutions, "a-d*b*c") }
	
		if b-a*c*d == 24 { solutions = append(solutions, "b-a*c*d") }
	
		if b-a*d*c == 24 { solutions = append(solutions, "b-a*d*c") }
	
		if b-c*a*d == 24 { solutions = append(solutions, "b-c*a*d") }
	
		if b-c*d*a == 24 { solutions = append(solutions, "b-c*d*a") }
	
		if b-d*c*a == 24 { solutions = append(solutions, "b-d*c*a") }
	
		if b-d*a*c == 24 { solutions = append(solutions, "b-d*a*c") }
	
		if c-b*a*d == 24 { solutions = append(solutions, "c-b*a*d") }
	
		if c-b*d*a == 24 { solutions = append(solutions, "c-b*d*a") }
	
		if c-a*b*d == 24 { solutions = append(solutions, "c-a*b*d") }
	
		if c-a*d*b == 24 { solutions = append(solutions, "c-a*d*b") }
	
		if c-d*a*b == 24 { solutions = append(solutions, "c-d*a*b") }
	
		if c-d*b*a == 24 { solutions = append(solutions, "c-d*b*a") }
	
		if d-b*c*a == 24 { solutions = append(solutions, "d-b*c*a") }
	
		if d-b*a*c == 24 { solutions = append(solutions, "d-b*a*c") }
	
		if d-c*b*a == 24 { solutions = append(solutions, "d-c*b*a") }
	
		if d-c*a*b == 24 { solutions = append(solutions, "d-c*a*b") }
	
		if d-a*c*b == 24 { solutions = append(solutions, "d-a*c*b") }
	
		if d-a*b*c == 24 { solutions = append(solutions, "d-a*b*c") }
	
		if a-b*c/d == 24 { solutions = append(solutions, "a-b*c/d") }
	
		if a-b*d/c == 24 { solutions = append(solutions, "a-b*d/c") }
	
		if a-c*b/d == 24 { solutions = append(solutions, "a-c*b/d") }
	
		if a-c*d/b == 24 { solutions = append(solutions, "a-c*d/b") }
	
		if a-d*c/b == 24 { solutions = append(solutions, "a-d*c/b") }
	
		if a-d*b/c == 24 { solutions = append(solutions, "a-d*b/c") }
	
		if b-a*c/d == 24 { solutions = append(solutions, "b-a*c/d") }
	
		if b-a*d/c == 24 { solutions = append(solutions, "b-a*d/c") }
	
		if b-c*a/d == 24 { solutions = append(solutions, "b-c*a/d") }
	
		if b-c*d/a == 24 { solutions = append(solutions, "b-c*d/a") }
	
		if b-d*c/a == 24 { solutions = append(solutions, "b-d*c/a") }
	
		if b-d*a/c == 24 { solutions = append(solutions, "b-d*a/c") }
	
		if c-b*a/d == 24 { solutions = append(solutions, "c-b*a/d") }
	
		if c-b*d/a == 24 { solutions = append(solutions, "c-b*d/a") }
	
		if c-a*b/d == 24 { solutions = append(solutions, "c-a*b/d") }
	
		if c-a*d/b == 24 { solutions = append(solutions, "c-a*d/b") }
	
		if c-d*a/b == 24 { solutions = append(solutions, "c-d*a/b") }
	
		if c-d*b/a == 24 { solutions = append(solutions, "c-d*b/a") }
	
		if d-b*c/a == 24 { solutions = append(solutions, "d-b*c/a") }
	
		if d-b*a/c == 24 { solutions = append(solutions, "d-b*a/c") }
	
		if d-c*b/a == 24 { solutions = append(solutions, "d-c*b/a") }
	
		if d-c*a/b == 24 { solutions = append(solutions, "d-c*a/b") }
	
		if d-a*c/b == 24 { solutions = append(solutions, "d-a*c/b") }
	
		if d-a*b/c == 24 { solutions = append(solutions, "d-a*b/c") }
	
		if a-b/c+d == 24 { solutions = append(solutions, "a-b/c+d") }
	
		if a-b/d+c == 24 { solutions = append(solutions, "a-b/d+c") }
	
		if a-c/b+d == 24 { solutions = append(solutions, "a-c/b+d") }
	
		if a-c/d+b == 24 { solutions = append(solutions, "a-c/d+b") }
	
		if a-d/c+b == 24 { solutions = append(solutions, "a-d/c+b") }
	
		if a-d/b+c == 24 { solutions = append(solutions, "a-d/b+c") }
	
		if b-a/c+d == 24 { solutions = append(solutions, "b-a/c+d") }
	
		if b-a/d+c == 24 { solutions = append(solutions, "b-a/d+c") }
	
		if b-c/a+d == 24 { solutions = append(solutions, "b-c/a+d") }
	
		if b-c/d+a == 24 { solutions = append(solutions, "b-c/d+a") }
	
		if b-d/c+a == 24 { solutions = append(solutions, "b-d/c+a") }
	
		if b-d/a+c == 24 { solutions = append(solutions, "b-d/a+c") }
	
		if c-b/a+d == 24 { solutions = append(solutions, "c-b/a+d") }
	
		if c-b/d+a == 24 { solutions = append(solutions, "c-b/d+a") }
	
		if c-a/b+d == 24 { solutions = append(solutions, "c-a/b+d") }
	
		if c-a/d+b == 24 { solutions = append(solutions, "c-a/d+b") }
	
		if c-d/a+b == 24 { solutions = append(solutions, "c-d/a+b") }
	
		if c-d/b+a == 24 { solutions = append(solutions, "c-d/b+a") }
	
		if d-b/c+a == 24 { solutions = append(solutions, "d-b/c+a") }
	
		if d-b/a+c == 24 { solutions = append(solutions, "d-b/a+c") }
	
		if d-c/b+a == 24 { solutions = append(solutions, "d-c/b+a") }
	
		if d-c/a+b == 24 { solutions = append(solutions, "d-c/a+b") }
	
		if d-a/c+b == 24 { solutions = append(solutions, "d-a/c+b") }
	
		if d-a/b+c == 24 { solutions = append(solutions, "d-a/b+c") }
	
		if a-b/c-d == 24 { solutions = append(solutions, "a-b/c-d") }
	
		if a-b/d-c == 24 { solutions = append(solutions, "a-b/d-c") }
	
		if a-c/b-d == 24 { solutions = append(solutions, "a-c/b-d") }
	
		if a-c/d-b == 24 { solutions = append(solutions, "a-c/d-b") }
	
		if a-d/c-b == 24 { solutions = append(solutions, "a-d/c-b") }
	
		if a-d/b-c == 24 { solutions = append(solutions, "a-d/b-c") }
	
		if b-a/c-d == 24 { solutions = append(solutions, "b-a/c-d") }
	
		if b-a/d-c == 24 { solutions = append(solutions, "b-a/d-c") }
	
		if b-c/a-d == 24 { solutions = append(solutions, "b-c/a-d") }
	
		if b-c/d-a == 24 { solutions = append(solutions, "b-c/d-a") }
	
		if b-d/c-a == 24 { solutions = append(solutions, "b-d/c-a") }
	
		if b-d/a-c == 24 { solutions = append(solutions, "b-d/a-c") }
	
		if c-b/a-d == 24 { solutions = append(solutions, "c-b/a-d") }
	
		if c-b/d-a == 24 { solutions = append(solutions, "c-b/d-a") }
	
		if c-a/b-d == 24 { solutions = append(solutions, "c-a/b-d") }
	
		if c-a/d-b == 24 { solutions = append(solutions, "c-a/d-b") }
	
		if c-d/a-b == 24 { solutions = append(solutions, "c-d/a-b") }
	
		if c-d/b-a == 24 { solutions = append(solutions, "c-d/b-a") }
	
		if d-b/c-a == 24 { solutions = append(solutions, "d-b/c-a") }
	
		if d-b/a-c == 24 { solutions = append(solutions, "d-b/a-c") }
	
		if d-c/b-a == 24 { solutions = append(solutions, "d-c/b-a") }
	
		if d-c/a-b == 24 { solutions = append(solutions, "d-c/a-b") }
	
		if d-a/c-b == 24 { solutions = append(solutions, "d-a/c-b") }
	
		if d-a/b-c == 24 { solutions = append(solutions, "d-a/b-c") }
	
		if a-b/c*d == 24 { solutions = append(solutions, "a-b/c*d") }
	
		if a-b/d*c == 24 { solutions = append(solutions, "a-b/d*c") }
	
		if a-c/b*d == 24 { solutions = append(solutions, "a-c/b*d") }
	
		if a-c/d*b == 24 { solutions = append(solutions, "a-c/d*b") }
	
		if a-d/c*b == 24 { solutions = append(solutions, "a-d/c*b") }
	
		if a-d/b*c == 24 { solutions = append(solutions, "a-d/b*c") }
	
		if b-a/c*d == 24 { solutions = append(solutions, "b-a/c*d") }
	
		if b-a/d*c == 24 { solutions = append(solutions, "b-a/d*c") }
	
		if b-c/a*d == 24 { solutions = append(solutions, "b-c/a*d") }
	
		if b-c/d*a == 24 { solutions = append(solutions, "b-c/d*a") }
	
		if b-d/c*a == 24 { solutions = append(solutions, "b-d/c*a") }
	
		if b-d/a*c == 24 { solutions = append(solutions, "b-d/a*c") }
	
		if c-b/a*d == 24 { solutions = append(solutions, "c-b/a*d") }
	
		if c-b/d*a == 24 { solutions = append(solutions, "c-b/d*a") }
	
		if c-a/b*d == 24 { solutions = append(solutions, "c-a/b*d") }
	
		if c-a/d*b == 24 { solutions = append(solutions, "c-a/d*b") }
	
		if c-d/a*b == 24 { solutions = append(solutions, "c-d/a*b") }
	
		if c-d/b*a == 24 { solutions = append(solutions, "c-d/b*a") }
	
		if d-b/c*a == 24 { solutions = append(solutions, "d-b/c*a") }
	
		if d-b/a*c == 24 { solutions = append(solutions, "d-b/a*c") }
	
		if d-c/b*a == 24 { solutions = append(solutions, "d-c/b*a") }
	
		if d-c/a*b == 24 { solutions = append(solutions, "d-c/a*b") }
	
		if d-a/c*b == 24 { solutions = append(solutions, "d-a/c*b") }
	
		if d-a/b*c == 24 { solutions = append(solutions, "d-a/b*c") }
	
		if a-b/c/d == 24 { solutions = append(solutions, "a-b/c/d") }
	
		if a-b/d/c == 24 { solutions = append(solutions, "a-b/d/c") }
	
		if a-c/b/d == 24 { solutions = append(solutions, "a-c/b/d") }
	
		if a-c/d/b == 24 { solutions = append(solutions, "a-c/d/b") }
	
		if a-d/c/b == 24 { solutions = append(solutions, "a-d/c/b") }
	
		if a-d/b/c == 24 { solutions = append(solutions, "a-d/b/c") }
	
		if b-a/c/d == 24 { solutions = append(solutions, "b-a/c/d") }
	
		if b-a/d/c == 24 { solutions = append(solutions, "b-a/d/c") }
	
		if b-c/a/d == 24 { solutions = append(solutions, "b-c/a/d") }
	
		if b-c/d/a == 24 { solutions = append(solutions, "b-c/d/a") }
	
		if b-d/c/a == 24 { solutions = append(solutions, "b-d/c/a") }
	
		if b-d/a/c == 24 { solutions = append(solutions, "b-d/a/c") }
	
		if c-b/a/d == 24 { solutions = append(solutions, "c-b/a/d") }
	
		if c-b/d/a == 24 { solutions = append(solutions, "c-b/d/a") }
	
		if c-a/b/d == 24 { solutions = append(solutions, "c-a/b/d") }
	
		if c-a/d/b == 24 { solutions = append(solutions, "c-a/d/b") }
	
		if c-d/a/b == 24 { solutions = append(solutions, "c-d/a/b") }
	
		if c-d/b/a == 24 { solutions = append(solutions, "c-d/b/a") }
	
		if d-b/c/a == 24 { solutions = append(solutions, "d-b/c/a") }
	
		if d-b/a/c == 24 { solutions = append(solutions, "d-b/a/c") }
	
		if d-c/b/a == 24 { solutions = append(solutions, "d-c/b/a") }
	
		if d-c/a/b == 24 { solutions = append(solutions, "d-c/a/b") }
	
		if d-a/c/b == 24 { solutions = append(solutions, "d-a/c/b") }
	
		if d-a/b/c == 24 { solutions = append(solutions, "d-a/b/c") }
	
		if a*b+c+d == 24 { solutions = append(solutions, "a*b+c+d") }
	
		if a*b+d+c == 24 { solutions = append(solutions, "a*b+d+c") }
	
		if a*c+b+d == 24 { solutions = append(solutions, "a*c+b+d") }
	
		if a*c+d+b == 24 { solutions = append(solutions, "a*c+d+b") }
	
		if a*d+c+b == 24 { solutions = append(solutions, "a*d+c+b") }
	
		if a*d+b+c == 24 { solutions = append(solutions, "a*d+b+c") }
	
		if b*a+c+d == 24 { solutions = append(solutions, "b*a+c+d") }
	
		if b*a+d+c == 24 { solutions = append(solutions, "b*a+d+c") }
	
		if b*c+a+d == 24 { solutions = append(solutions, "b*c+a+d") }
	
		if b*c+d+a == 24 { solutions = append(solutions, "b*c+d+a") }
	
		if b*d+c+a == 24 { solutions = append(solutions, "b*d+c+a") }
	
		if b*d+a+c == 24 { solutions = append(solutions, "b*d+a+c") }
	
		if c*b+a+d == 24 { solutions = append(solutions, "c*b+a+d") }
	
		if c*b+d+a == 24 { solutions = append(solutions, "c*b+d+a") }
	
		if c*a+b+d == 24 { solutions = append(solutions, "c*a+b+d") }
	
		if c*a+d+b == 24 { solutions = append(solutions, "c*a+d+b") }
	
		if c*d+a+b == 24 { solutions = append(solutions, "c*d+a+b") }
	
		if c*d+b+a == 24 { solutions = append(solutions, "c*d+b+a") }
	
		if d*b+c+a == 24 { solutions = append(solutions, "d*b+c+a") }
	
		if d*b+a+c == 24 { solutions = append(solutions, "d*b+a+c") }
	
		if d*c+b+a == 24 { solutions = append(solutions, "d*c+b+a") }
	
		if d*c+a+b == 24 { solutions = append(solutions, "d*c+a+b") }
	
		if d*a+c+b == 24 { solutions = append(solutions, "d*a+c+b") }
	
		if d*a+b+c == 24 { solutions = append(solutions, "d*a+b+c") }
	
		if a*b+c-d == 24 { solutions = append(solutions, "a*b+c-d") }
	
		if a*b+d-c == 24 { solutions = append(solutions, "a*b+d-c") }
	
		if a*c+b-d == 24 { solutions = append(solutions, "a*c+b-d") }
	
		if a*c+d-b == 24 { solutions = append(solutions, "a*c+d-b") }
	
		if a*d+c-b == 24 { solutions = append(solutions, "a*d+c-b") }
	
		if a*d+b-c == 24 { solutions = append(solutions, "a*d+b-c") }
	
		if b*a+c-d == 24 { solutions = append(solutions, "b*a+c-d") }
	
		if b*a+d-c == 24 { solutions = append(solutions, "b*a+d-c") }
	
		if b*c+a-d == 24 { solutions = append(solutions, "b*c+a-d") }
	
		if b*c+d-a == 24 { solutions = append(solutions, "b*c+d-a") }
	
		if b*d+c-a == 24 { solutions = append(solutions, "b*d+c-a") }
	
		if b*d+a-c == 24 { solutions = append(solutions, "b*d+a-c") }
	
		if c*b+a-d == 24 { solutions = append(solutions, "c*b+a-d") }
	
		if c*b+d-a == 24 { solutions = append(solutions, "c*b+d-a") }
	
		if c*a+b-d == 24 { solutions = append(solutions, "c*a+b-d") }
	
		if c*a+d-b == 24 { solutions = append(solutions, "c*a+d-b") }
	
		if c*d+a-b == 24 { solutions = append(solutions, "c*d+a-b") }
	
		if c*d+b-a == 24 { solutions = append(solutions, "c*d+b-a") }
	
		if d*b+c-a == 24 { solutions = append(solutions, "d*b+c-a") }
	
		if d*b+a-c == 24 { solutions = append(solutions, "d*b+a-c") }
	
		if d*c+b-a == 24 { solutions = append(solutions, "d*c+b-a") }
	
		if d*c+a-b == 24 { solutions = append(solutions, "d*c+a-b") }
	
		if d*a+c-b == 24 { solutions = append(solutions, "d*a+c-b") }
	
		if d*a+b-c == 24 { solutions = append(solutions, "d*a+b-c") }
	
		if a*b+c*d == 24 { solutions = append(solutions, "a*b+c*d") }
	
		if a*b+d*c == 24 { solutions = append(solutions, "a*b+d*c") }
	
		if a*c+b*d == 24 { solutions = append(solutions, "a*c+b*d") }
	
		if a*c+d*b == 24 { solutions = append(solutions, "a*c+d*b") }
	
		if a*d+c*b == 24 { solutions = append(solutions, "a*d+c*b") }
	
		if a*d+b*c == 24 { solutions = append(solutions, "a*d+b*c") }
	
		if b*a+c*d == 24 { solutions = append(solutions, "b*a+c*d") }
	
		if b*a+d*c == 24 { solutions = append(solutions, "b*a+d*c") }
	
		if b*c+a*d == 24 { solutions = append(solutions, "b*c+a*d") }
	
		if b*c+d*a == 24 { solutions = append(solutions, "b*c+d*a") }
	
		if b*d+c*a == 24 { solutions = append(solutions, "b*d+c*a") }
	
		if b*d+a*c == 24 { solutions = append(solutions, "b*d+a*c") }
	
		if c*b+a*d == 24 { solutions = append(solutions, "c*b+a*d") }
	
		if c*b+d*a == 24 { solutions = append(solutions, "c*b+d*a") }
	
		if c*a+b*d == 24 { solutions = append(solutions, "c*a+b*d") }
	
		if c*a+d*b == 24 { solutions = append(solutions, "c*a+d*b") }
	
		if c*d+a*b == 24 { solutions = append(solutions, "c*d+a*b") }
	
		if c*d+b*a == 24 { solutions = append(solutions, "c*d+b*a") }
	
		if d*b+c*a == 24 { solutions = append(solutions, "d*b+c*a") }
	
		if d*b+a*c == 24 { solutions = append(solutions, "d*b+a*c") }
	
		if d*c+b*a == 24 { solutions = append(solutions, "d*c+b*a") }
	
		if d*c+a*b == 24 { solutions = append(solutions, "d*c+a*b") }
	
		if d*a+c*b == 24 { solutions = append(solutions, "d*a+c*b") }
	
		if d*a+b*c == 24 { solutions = append(solutions, "d*a+b*c") }
	
		if a*b+c/d == 24 { solutions = append(solutions, "a*b+c/d") }
	
		if a*b+d/c == 24 { solutions = append(solutions, "a*b+d/c") }
	
		if a*c+b/d == 24 { solutions = append(solutions, "a*c+b/d") }
	
		if a*c+d/b == 24 { solutions = append(solutions, "a*c+d/b") }
	
		if a*d+c/b == 24 { solutions = append(solutions, "a*d+c/b") }
	
		if a*d+b/c == 24 { solutions = append(solutions, "a*d+b/c") }
	
		if b*a+c/d == 24 { solutions = append(solutions, "b*a+c/d") }
	
		if b*a+d/c == 24 { solutions = append(solutions, "b*a+d/c") }
	
		if b*c+a/d == 24 { solutions = append(solutions, "b*c+a/d") }
	
		if b*c+d/a == 24 { solutions = append(solutions, "b*c+d/a") }
	
		if b*d+c/a == 24 { solutions = append(solutions, "b*d+c/a") }
	
		if b*d+a/c == 24 { solutions = append(solutions, "b*d+a/c") }
	
		if c*b+a/d == 24 { solutions = append(solutions, "c*b+a/d") }
	
		if c*b+d/a == 24 { solutions = append(solutions, "c*b+d/a") }
	
		if c*a+b/d == 24 { solutions = append(solutions, "c*a+b/d") }
	
		if c*a+d/b == 24 { solutions = append(solutions, "c*a+d/b") }
	
		if c*d+a/b == 24 { solutions = append(solutions, "c*d+a/b") }
	
		if c*d+b/a == 24 { solutions = append(solutions, "c*d+b/a") }
	
		if d*b+c/a == 24 { solutions = append(solutions, "d*b+c/a") }
	
		if d*b+a/c == 24 { solutions = append(solutions, "d*b+a/c") }
	
		if d*c+b/a == 24 { solutions = append(solutions, "d*c+b/a") }
	
		if d*c+a/b == 24 { solutions = append(solutions, "d*c+a/b") }
	
		if d*a+c/b == 24 { solutions = append(solutions, "d*a+c/b") }
	
		if d*a+b/c == 24 { solutions = append(solutions, "d*a+b/c") }
	
		if a*b-c+d == 24 { solutions = append(solutions, "a*b-c+d") }
	
		if a*b-d+c == 24 { solutions = append(solutions, "a*b-d+c") }
	
		if a*c-b+d == 24 { solutions = append(solutions, "a*c-b+d") }
	
		if a*c-d+b == 24 { solutions = append(solutions, "a*c-d+b") }
	
		if a*d-c+b == 24 { solutions = append(solutions, "a*d-c+b") }
	
		if a*d-b+c == 24 { solutions = append(solutions, "a*d-b+c") }
	
		if b*a-c+d == 24 { solutions = append(solutions, "b*a-c+d") }
	
		if b*a-d+c == 24 { solutions = append(solutions, "b*a-d+c") }
	
		if b*c-a+d == 24 { solutions = append(solutions, "b*c-a+d") }
	
		if b*c-d+a == 24 { solutions = append(solutions, "b*c-d+a") }
	
		if b*d-c+a == 24 { solutions = append(solutions, "b*d-c+a") }
	
		if b*d-a+c == 24 { solutions = append(solutions, "b*d-a+c") }
	
		if c*b-a+d == 24 { solutions = append(solutions, "c*b-a+d") }
	
		if c*b-d+a == 24 { solutions = append(solutions, "c*b-d+a") }
	
		if c*a-b+d == 24 { solutions = append(solutions, "c*a-b+d") }
	
		if c*a-d+b == 24 { solutions = append(solutions, "c*a-d+b") }
	
		if c*d-a+b == 24 { solutions = append(solutions, "c*d-a+b") }
	
		if c*d-b+a == 24 { solutions = append(solutions, "c*d-b+a") }
	
		if d*b-c+a == 24 { solutions = append(solutions, "d*b-c+a") }
	
		if d*b-a+c == 24 { solutions = append(solutions, "d*b-a+c") }
	
		if d*c-b+a == 24 { solutions = append(solutions, "d*c-b+a") }
	
		if d*c-a+b == 24 { solutions = append(solutions, "d*c-a+b") }
	
		if d*a-c+b == 24 { solutions = append(solutions, "d*a-c+b") }
	
		if d*a-b+c == 24 { solutions = append(solutions, "d*a-b+c") }
	
		if a*b-c-d == 24 { solutions = append(solutions, "a*b-c-d") }
	
		if a*b-d-c == 24 { solutions = append(solutions, "a*b-d-c") }
	
		if a*c-b-d == 24 { solutions = append(solutions, "a*c-b-d") }
	
		if a*c-d-b == 24 { solutions = append(solutions, "a*c-d-b") }
	
		if a*d-c-b == 24 { solutions = append(solutions, "a*d-c-b") }
	
		if a*d-b-c == 24 { solutions = append(solutions, "a*d-b-c") }
	
		if b*a-c-d == 24 { solutions = append(solutions, "b*a-c-d") }
	
		if b*a-d-c == 24 { solutions = append(solutions, "b*a-d-c") }
	
		if b*c-a-d == 24 { solutions = append(solutions, "b*c-a-d") }
	
		if b*c-d-a == 24 { solutions = append(solutions, "b*c-d-a") }
	
		if b*d-c-a == 24 { solutions = append(solutions, "b*d-c-a") }
	
		if b*d-a-c == 24 { solutions = append(solutions, "b*d-a-c") }
	
		if c*b-a-d == 24 { solutions = append(solutions, "c*b-a-d") }
	
		if c*b-d-a == 24 { solutions = append(solutions, "c*b-d-a") }
	
		if c*a-b-d == 24 { solutions = append(solutions, "c*a-b-d") }
	
		if c*a-d-b == 24 { solutions = append(solutions, "c*a-d-b") }
	
		if c*d-a-b == 24 { solutions = append(solutions, "c*d-a-b") }
	
		if c*d-b-a == 24 { solutions = append(solutions, "c*d-b-a") }
	
		if d*b-c-a == 24 { solutions = append(solutions, "d*b-c-a") }
	
		if d*b-a-c == 24 { solutions = append(solutions, "d*b-a-c") }
	
		if d*c-b-a == 24 { solutions = append(solutions, "d*c-b-a") }
	
		if d*c-a-b == 24 { solutions = append(solutions, "d*c-a-b") }
	
		if d*a-c-b == 24 { solutions = append(solutions, "d*a-c-b") }
	
		if d*a-b-c == 24 { solutions = append(solutions, "d*a-b-c") }
	
		if a*b-c*d == 24 { solutions = append(solutions, "a*b-c*d") }
	
		if a*b-d*c == 24 { solutions = append(solutions, "a*b-d*c") }
	
		if a*c-b*d == 24 { solutions = append(solutions, "a*c-b*d") }
	
		if a*c-d*b == 24 { solutions = append(solutions, "a*c-d*b") }
	
		if a*d-c*b == 24 { solutions = append(solutions, "a*d-c*b") }
	
		if a*d-b*c == 24 { solutions = append(solutions, "a*d-b*c") }
	
		if b*a-c*d == 24 { solutions = append(solutions, "b*a-c*d") }
	
		if b*a-d*c == 24 { solutions = append(solutions, "b*a-d*c") }
	
		if b*c-a*d == 24 { solutions = append(solutions, "b*c-a*d") }
	
		if b*c-d*a == 24 { solutions = append(solutions, "b*c-d*a") }
	
		if b*d-c*a == 24 { solutions = append(solutions, "b*d-c*a") }
	
		if b*d-a*c == 24 { solutions = append(solutions, "b*d-a*c") }
	
		if c*b-a*d == 24 { solutions = append(solutions, "c*b-a*d") }
	
		if c*b-d*a == 24 { solutions = append(solutions, "c*b-d*a") }
	
		if c*a-b*d == 24 { solutions = append(solutions, "c*a-b*d") }
	
		if c*a-d*b == 24 { solutions = append(solutions, "c*a-d*b") }
	
		if c*d-a*b == 24 { solutions = append(solutions, "c*d-a*b") }
	
		if c*d-b*a == 24 { solutions = append(solutions, "c*d-b*a") }
	
		if d*b-c*a == 24 { solutions = append(solutions, "d*b-c*a") }
	
		if d*b-a*c == 24 { solutions = append(solutions, "d*b-a*c") }
	
		if d*c-b*a == 24 { solutions = append(solutions, "d*c-b*a") }
	
		if d*c-a*b == 24 { solutions = append(solutions, "d*c-a*b") }
	
		if d*a-c*b == 24 { solutions = append(solutions, "d*a-c*b") }
	
		if d*a-b*c == 24 { solutions = append(solutions, "d*a-b*c") }
	
		if a*b-c/d == 24 { solutions = append(solutions, "a*b-c/d") }
	
		if a*b-d/c == 24 { solutions = append(solutions, "a*b-d/c") }
	
		if a*c-b/d == 24 { solutions = append(solutions, "a*c-b/d") }
	
		if a*c-d/b == 24 { solutions = append(solutions, "a*c-d/b") }
	
		if a*d-c/b == 24 { solutions = append(solutions, "a*d-c/b") }
	
		if a*d-b/c == 24 { solutions = append(solutions, "a*d-b/c") }
	
		if b*a-c/d == 24 { solutions = append(solutions, "b*a-c/d") }
	
		if b*a-d/c == 24 { solutions = append(solutions, "b*a-d/c") }
	
		if b*c-a/d == 24 { solutions = append(solutions, "b*c-a/d") }
	
		if b*c-d/a == 24 { solutions = append(solutions, "b*c-d/a") }
	
		if b*d-c/a == 24 { solutions = append(solutions, "b*d-c/a") }
	
		if b*d-a/c == 24 { solutions = append(solutions, "b*d-a/c") }
	
		if c*b-a/d == 24 { solutions = append(solutions, "c*b-a/d") }
	
		if c*b-d/a == 24 { solutions = append(solutions, "c*b-d/a") }
	
		if c*a-b/d == 24 { solutions = append(solutions, "c*a-b/d") }
	
		if c*a-d/b == 24 { solutions = append(solutions, "c*a-d/b") }
	
		if c*d-a/b == 24 { solutions = append(solutions, "c*d-a/b") }
	
		if c*d-b/a == 24 { solutions = append(solutions, "c*d-b/a") }
	
		if d*b-c/a == 24 { solutions = append(solutions, "d*b-c/a") }
	
		if d*b-a/c == 24 { solutions = append(solutions, "d*b-a/c") }
	
		if d*c-b/a == 24 { solutions = append(solutions, "d*c-b/a") }
	
		if d*c-a/b == 24 { solutions = append(solutions, "d*c-a/b") }
	
		if d*a-c/b == 24 { solutions = append(solutions, "d*a-c/b") }
	
		if d*a-b/c == 24 { solutions = append(solutions, "d*a-b/c") }
	
		if a*b*c+d == 24 { solutions = append(solutions, "a*b*c+d") }
	
		if a*b*d+c == 24 { solutions = append(solutions, "a*b*d+c") }
	
		if a*c*b+d == 24 { solutions = append(solutions, "a*c*b+d") }
	
		if a*c*d+b == 24 { solutions = append(solutions, "a*c*d+b") }
	
		if a*d*c+b == 24 { solutions = append(solutions, "a*d*c+b") }
	
		if a*d*b+c == 24 { solutions = append(solutions, "a*d*b+c") }
	
		if b*a*c+d == 24 { solutions = append(solutions, "b*a*c+d") }
	
		if b*a*d+c == 24 { solutions = append(solutions, "b*a*d+c") }
	
		if b*c*a+d == 24 { solutions = append(solutions, "b*c*a+d") }
	
		if b*c*d+a == 24 { solutions = append(solutions, "b*c*d+a") }
	
		if b*d*c+a == 24 { solutions = append(solutions, "b*d*c+a") }
	
		if b*d*a+c == 24 { solutions = append(solutions, "b*d*a+c") }
	
		if c*b*a+d == 24 { solutions = append(solutions, "c*b*a+d") }
	
		if c*b*d+a == 24 { solutions = append(solutions, "c*b*d+a") }
	
		if c*a*b+d == 24 { solutions = append(solutions, "c*a*b+d") }
	
		if c*a*d+b == 24 { solutions = append(solutions, "c*a*d+b") }
	
		if c*d*a+b == 24 { solutions = append(solutions, "c*d*a+b") }
	
		if c*d*b+a == 24 { solutions = append(solutions, "c*d*b+a") }
	
		if d*b*c+a == 24 { solutions = append(solutions, "d*b*c+a") }
	
		if d*b*a+c == 24 { solutions = append(solutions, "d*b*a+c") }
	
		if d*c*b+a == 24 { solutions = append(solutions, "d*c*b+a") }
	
		if d*c*a+b == 24 { solutions = append(solutions, "d*c*a+b") }
	
		if d*a*c+b == 24 { solutions = append(solutions, "d*a*c+b") }
	
		if d*a*b+c == 24 { solutions = append(solutions, "d*a*b+c") }
	
		if a*b*c-d == 24 { solutions = append(solutions, "a*b*c-d") }
	
		if a*b*d-c == 24 { solutions = append(solutions, "a*b*d-c") }
	
		if a*c*b-d == 24 { solutions = append(solutions, "a*c*b-d") }
	
		if a*c*d-b == 24 { solutions = append(solutions, "a*c*d-b") }
	
		if a*d*c-b == 24 { solutions = append(solutions, "a*d*c-b") }
	
		if a*d*b-c == 24 { solutions = append(solutions, "a*d*b-c") }
	
		if b*a*c-d == 24 { solutions = append(solutions, "b*a*c-d") }
	
		if b*a*d-c == 24 { solutions = append(solutions, "b*a*d-c") }
	
		if b*c*a-d == 24 { solutions = append(solutions, "b*c*a-d") }
	
		if b*c*d-a == 24 { solutions = append(solutions, "b*c*d-a") }
	
		if b*d*c-a == 24 { solutions = append(solutions, "b*d*c-a") }
	
		if b*d*a-c == 24 { solutions = append(solutions, "b*d*a-c") }
	
		if c*b*a-d == 24 { solutions = append(solutions, "c*b*a-d") }
	
		if c*b*d-a == 24 { solutions = append(solutions, "c*b*d-a") }
	
		if c*a*b-d == 24 { solutions = append(solutions, "c*a*b-d") }
	
		if c*a*d-b == 24 { solutions = append(solutions, "c*a*d-b") }
	
		if c*d*a-b == 24 { solutions = append(solutions, "c*d*a-b") }
	
		if c*d*b-a == 24 { solutions = append(solutions, "c*d*b-a") }
	
		if d*b*c-a == 24 { solutions = append(solutions, "d*b*c-a") }
	
		if d*b*a-c == 24 { solutions = append(solutions, "d*b*a-c") }
	
		if d*c*b-a == 24 { solutions = append(solutions, "d*c*b-a") }
	
		if d*c*a-b == 24 { solutions = append(solutions, "d*c*a-b") }
	
		if d*a*c-b == 24 { solutions = append(solutions, "d*a*c-b") }
	
		if d*a*b-c == 24 { solutions = append(solutions, "d*a*b-c") }
	
		if a*b*c*d == 24 { solutions = append(solutions, "a*b*c*d") }
	
		if a*b*d*c == 24 { solutions = append(solutions, "a*b*d*c") }
	
		if a*c*b*d == 24 { solutions = append(solutions, "a*c*b*d") }
	
		if a*c*d*b == 24 { solutions = append(solutions, "a*c*d*b") }
	
		if a*d*c*b == 24 { solutions = append(solutions, "a*d*c*b") }
	
		if a*d*b*c == 24 { solutions = append(solutions, "a*d*b*c") }
	
		if b*a*c*d == 24 { solutions = append(solutions, "b*a*c*d") }
	
		if b*a*d*c == 24 { solutions = append(solutions, "b*a*d*c") }
	
		if b*c*a*d == 24 { solutions = append(solutions, "b*c*a*d") }
	
		if b*c*d*a == 24 { solutions = append(solutions, "b*c*d*a") }
	
		if b*d*c*a == 24 { solutions = append(solutions, "b*d*c*a") }
	
		if b*d*a*c == 24 { solutions = append(solutions, "b*d*a*c") }
	
		if c*b*a*d == 24 { solutions = append(solutions, "c*b*a*d") }
	
		if c*b*d*a == 24 { solutions = append(solutions, "c*b*d*a") }
	
		if c*a*b*d == 24 { solutions = append(solutions, "c*a*b*d") }
	
		if c*a*d*b == 24 { solutions = append(solutions, "c*a*d*b") }
	
		if c*d*a*b == 24 { solutions = append(solutions, "c*d*a*b") }
	
		if c*d*b*a == 24 { solutions = append(solutions, "c*d*b*a") }
	
		if d*b*c*a == 24 { solutions = append(solutions, "d*b*c*a") }
	
		if d*b*a*c == 24 { solutions = append(solutions, "d*b*a*c") }
	
		if d*c*b*a == 24 { solutions = append(solutions, "d*c*b*a") }
	
		if d*c*a*b == 24 { solutions = append(solutions, "d*c*a*b") }
	
		if d*a*c*b == 24 { solutions = append(solutions, "d*a*c*b") }
	
		if d*a*b*c == 24 { solutions = append(solutions, "d*a*b*c") }
	
		if a*b*c/d == 24 { solutions = append(solutions, "a*b*c/d") }
	
		if a*b*d/c == 24 { solutions = append(solutions, "a*b*d/c") }
	
		if a*c*b/d == 24 { solutions = append(solutions, "a*c*b/d") }
	
		if a*c*d/b == 24 { solutions = append(solutions, "a*c*d/b") }
	
		if a*d*c/b == 24 { solutions = append(solutions, "a*d*c/b") }
	
		if a*d*b/c == 24 { solutions = append(solutions, "a*d*b/c") }
	
		if b*a*c/d == 24 { solutions = append(solutions, "b*a*c/d") }
	
		if b*a*d/c == 24 { solutions = append(solutions, "b*a*d/c") }
	
		if b*c*a/d == 24 { solutions = append(solutions, "b*c*a/d") }
	
		if b*c*d/a == 24 { solutions = append(solutions, "b*c*d/a") }
	
		if b*d*c/a == 24 { solutions = append(solutions, "b*d*c/a") }
	
		if b*d*a/c == 24 { solutions = append(solutions, "b*d*a/c") }
	
		if c*b*a/d == 24 { solutions = append(solutions, "c*b*a/d") }
	
		if c*b*d/a == 24 { solutions = append(solutions, "c*b*d/a") }
	
		if c*a*b/d == 24 { solutions = append(solutions, "c*a*b/d") }
	
		if c*a*d/b == 24 { solutions = append(solutions, "c*a*d/b") }
	
		if c*d*a/b == 24 { solutions = append(solutions, "c*d*a/b") }
	
		if c*d*b/a == 24 { solutions = append(solutions, "c*d*b/a") }
	
		if d*b*c/a == 24 { solutions = append(solutions, "d*b*c/a") }
	
		if d*b*a/c == 24 { solutions = append(solutions, "d*b*a/c") }
	
		if d*c*b/a == 24 { solutions = append(solutions, "d*c*b/a") }
	
		if d*c*a/b == 24 { solutions = append(solutions, "d*c*a/b") }
	
		if d*a*c/b == 24 { solutions = append(solutions, "d*a*c/b") }
	
		if d*a*b/c == 24 { solutions = append(solutions, "d*a*b/c") }
	
		if a*b/c+d == 24 { solutions = append(solutions, "a*b/c+d") }
	
		if a*b/d+c == 24 { solutions = append(solutions, "a*b/d+c") }
	
		if a*c/b+d == 24 { solutions = append(solutions, "a*c/b+d") }
	
		if a*c/d+b == 24 { solutions = append(solutions, "a*c/d+b") }
	
		if a*d/c+b == 24 { solutions = append(solutions, "a*d/c+b") }
	
		if a*d/b+c == 24 { solutions = append(solutions, "a*d/b+c") }
	
		if b*a/c+d == 24 { solutions = append(solutions, "b*a/c+d") }
	
		if b*a/d+c == 24 { solutions = append(solutions, "b*a/d+c") }
	
		if b*c/a+d == 24 { solutions = append(solutions, "b*c/a+d") }
	
		if b*c/d+a == 24 { solutions = append(solutions, "b*c/d+a") }
	
		if b*d/c+a == 24 { solutions = append(solutions, "b*d/c+a") }
	
		if b*d/a+c == 24 { solutions = append(solutions, "b*d/a+c") }
	
		if c*b/a+d == 24 { solutions = append(solutions, "c*b/a+d") }
	
		if c*b/d+a == 24 { solutions = append(solutions, "c*b/d+a") }
	
		if c*a/b+d == 24 { solutions = append(solutions, "c*a/b+d") }
	
		if c*a/d+b == 24 { solutions = append(solutions, "c*a/d+b") }
	
		if c*d/a+b == 24 { solutions = append(solutions, "c*d/a+b") }
	
		if c*d/b+a == 24 { solutions = append(solutions, "c*d/b+a") }
	
		if d*b/c+a == 24 { solutions = append(solutions, "d*b/c+a") }
	
		if d*b/a+c == 24 { solutions = append(solutions, "d*b/a+c") }
	
		if d*c/b+a == 24 { solutions = append(solutions, "d*c/b+a") }
	
		if d*c/a+b == 24 { solutions = append(solutions, "d*c/a+b") }
	
		if d*a/c+b == 24 { solutions = append(solutions, "d*a/c+b") }
	
		if d*a/b+c == 24 { solutions = append(solutions, "d*a/b+c") }
	
		if a*b/c-d == 24 { solutions = append(solutions, "a*b/c-d") }
	
		if a*b/d-c == 24 { solutions = append(solutions, "a*b/d-c") }
	
		if a*c/b-d == 24 { solutions = append(solutions, "a*c/b-d") }
	
		if a*c/d-b == 24 { solutions = append(solutions, "a*c/d-b") }
	
		if a*d/c-b == 24 { solutions = append(solutions, "a*d/c-b") }
	
		if a*d/b-c == 24 { solutions = append(solutions, "a*d/b-c") }
	
		if b*a/c-d == 24 { solutions = append(solutions, "b*a/c-d") }
	
		if b*a/d-c == 24 { solutions = append(solutions, "b*a/d-c") }
	
		if b*c/a-d == 24 { solutions = append(solutions, "b*c/a-d") }
	
		if b*c/d-a == 24 { solutions = append(solutions, "b*c/d-a") }
	
		if b*d/c-a == 24 { solutions = append(solutions, "b*d/c-a") }
	
		if b*d/a-c == 24 { solutions = append(solutions, "b*d/a-c") }
	
		if c*b/a-d == 24 { solutions = append(solutions, "c*b/a-d") }
	
		if c*b/d-a == 24 { solutions = append(solutions, "c*b/d-a") }
	
		if c*a/b-d == 24 { solutions = append(solutions, "c*a/b-d") }
	
		if c*a/d-b == 24 { solutions = append(solutions, "c*a/d-b") }
	
		if c*d/a-b == 24 { solutions = append(solutions, "c*d/a-b") }
	
		if c*d/b-a == 24 { solutions = append(solutions, "c*d/b-a") }
	
		if d*b/c-a == 24 { solutions = append(solutions, "d*b/c-a") }
	
		if d*b/a-c == 24 { solutions = append(solutions, "d*b/a-c") }
	
		if d*c/b-a == 24 { solutions = append(solutions, "d*c/b-a") }
	
		if d*c/a-b == 24 { solutions = append(solutions, "d*c/a-b") }
	
		if d*a/c-b == 24 { solutions = append(solutions, "d*a/c-b") }
	
		if d*a/b-c == 24 { solutions = append(solutions, "d*a/b-c") }
	
		if a*b/c*d == 24 { solutions = append(solutions, "a*b/c*d") }
	
		if a*b/d*c == 24 { solutions = append(solutions, "a*b/d*c") }
	
		if a*c/b*d == 24 { solutions = append(solutions, "a*c/b*d") }
	
		if a*c/d*b == 24 { solutions = append(solutions, "a*c/d*b") }
	
		if a*d/c*b == 24 { solutions = append(solutions, "a*d/c*b") }
	
		if a*d/b*c == 24 { solutions = append(solutions, "a*d/b*c") }
	
		if b*a/c*d == 24 { solutions = append(solutions, "b*a/c*d") }
	
		if b*a/d*c == 24 { solutions = append(solutions, "b*a/d*c") }
	
		if b*c/a*d == 24 { solutions = append(solutions, "b*c/a*d") }
	
		if b*c/d*a == 24 { solutions = append(solutions, "b*c/d*a") }
	
		if b*d/c*a == 24 { solutions = append(solutions, "b*d/c*a") }
	
		if b*d/a*c == 24 { solutions = append(solutions, "b*d/a*c") }
	
		if c*b/a*d == 24 { solutions = append(solutions, "c*b/a*d") }
	
		if c*b/d*a == 24 { solutions = append(solutions, "c*b/d*a") }
	
		if c*a/b*d == 24 { solutions = append(solutions, "c*a/b*d") }
	
		if c*a/d*b == 24 { solutions = append(solutions, "c*a/d*b") }
	
		if c*d/a*b == 24 { solutions = append(solutions, "c*d/a*b") }
	
		if c*d/b*a == 24 { solutions = append(solutions, "c*d/b*a") }
	
		if d*b/c*a == 24 { solutions = append(solutions, "d*b/c*a") }
	
		if d*b/a*c == 24 { solutions = append(solutions, "d*b/a*c") }
	
		if d*c/b*a == 24 { solutions = append(solutions, "d*c/b*a") }
	
		if d*c/a*b == 24 { solutions = append(solutions, "d*c/a*b") }
	
		if d*a/c*b == 24 { solutions = append(solutions, "d*a/c*b") }
	
		if d*a/b*c == 24 { solutions = append(solutions, "d*a/b*c") }
	
		if a*b/c/d == 24 { solutions = append(solutions, "a*b/c/d") }
	
		if a*b/d/c == 24 { solutions = append(solutions, "a*b/d/c") }
	
		if a*c/b/d == 24 { solutions = append(solutions, "a*c/b/d") }
	
		if a*c/d/b == 24 { solutions = append(solutions, "a*c/d/b") }
	
		if a*d/c/b == 24 { solutions = append(solutions, "a*d/c/b") }
	
		if a*d/b/c == 24 { solutions = append(solutions, "a*d/b/c") }
	
		if b*a/c/d == 24 { solutions = append(solutions, "b*a/c/d") }
	
		if b*a/d/c == 24 { solutions = append(solutions, "b*a/d/c") }
	
		if b*c/a/d == 24 { solutions = append(solutions, "b*c/a/d") }
	
		if b*c/d/a == 24 { solutions = append(solutions, "b*c/d/a") }
	
		if b*d/c/a == 24 { solutions = append(solutions, "b*d/c/a") }
	
		if b*d/a/c == 24 { solutions = append(solutions, "b*d/a/c") }
	
		if c*b/a/d == 24 { solutions = append(solutions, "c*b/a/d") }
	
		if c*b/d/a == 24 { solutions = append(solutions, "c*b/d/a") }
	
		if c*a/b/d == 24 { solutions = append(solutions, "c*a/b/d") }
	
		if c*a/d/b == 24 { solutions = append(solutions, "c*a/d/b") }
	
		if c*d/a/b == 24 { solutions = append(solutions, "c*d/a/b") }
	
		if c*d/b/a == 24 { solutions = append(solutions, "c*d/b/a") }
	
		if d*b/c/a == 24 { solutions = append(solutions, "d*b/c/a") }
	
		if d*b/a/c == 24 { solutions = append(solutions, "d*b/a/c") }
	
		if d*c/b/a == 24 { solutions = append(solutions, "d*c/b/a") }
	
		if d*c/a/b == 24 { solutions = append(solutions, "d*c/a/b") }
	
		if d*a/c/b == 24 { solutions = append(solutions, "d*a/c/b") }
	
		if d*a/b/c == 24 { solutions = append(solutions, "d*a/b/c") }
	
		if a/b+c+d == 24 { solutions = append(solutions, "a/b+c+d") }
	
		if a/b+d+c == 24 { solutions = append(solutions, "a/b+d+c") }
	
		if a/c+b+d == 24 { solutions = append(solutions, "a/c+b+d") }
	
		if a/c+d+b == 24 { solutions = append(solutions, "a/c+d+b") }
	
		if a/d+c+b == 24 { solutions = append(solutions, "a/d+c+b") }
	
		if a/d+b+c == 24 { solutions = append(solutions, "a/d+b+c") }
	
		if b/a+c+d == 24 { solutions = append(solutions, "b/a+c+d") }
	
		if b/a+d+c == 24 { solutions = append(solutions, "b/a+d+c") }
	
		if b/c+a+d == 24 { solutions = append(solutions, "b/c+a+d") }
	
		if b/c+d+a == 24 { solutions = append(solutions, "b/c+d+a") }
	
		if b/d+c+a == 24 { solutions = append(solutions, "b/d+c+a") }
	
		if b/d+a+c == 24 { solutions = append(solutions, "b/d+a+c") }
	
		if c/b+a+d == 24 { solutions = append(solutions, "c/b+a+d") }
	
		if c/b+d+a == 24 { solutions = append(solutions, "c/b+d+a") }
	
		if c/a+b+d == 24 { solutions = append(solutions, "c/a+b+d") }
	
		if c/a+d+b == 24 { solutions = append(solutions, "c/a+d+b") }
	
		if c/d+a+b == 24 { solutions = append(solutions, "c/d+a+b") }
	
		if c/d+b+a == 24 { solutions = append(solutions, "c/d+b+a") }
	
		if d/b+c+a == 24 { solutions = append(solutions, "d/b+c+a") }
	
		if d/b+a+c == 24 { solutions = append(solutions, "d/b+a+c") }
	
		if d/c+b+a == 24 { solutions = append(solutions, "d/c+b+a") }
	
		if d/c+a+b == 24 { solutions = append(solutions, "d/c+a+b") }
	
		if d/a+c+b == 24 { solutions = append(solutions, "d/a+c+b") }
	
		if d/a+b+c == 24 { solutions = append(solutions, "d/a+b+c") }
	
		if a/b+c-d == 24 { solutions = append(solutions, "a/b+c-d") }
	
		if a/b+d-c == 24 { solutions = append(solutions, "a/b+d-c") }
	
		if a/c+b-d == 24 { solutions = append(solutions, "a/c+b-d") }
	
		if a/c+d-b == 24 { solutions = append(solutions, "a/c+d-b") }
	
		if a/d+c-b == 24 { solutions = append(solutions, "a/d+c-b") }
	
		if a/d+b-c == 24 { solutions = append(solutions, "a/d+b-c") }
	
		if b/a+c-d == 24 { solutions = append(solutions, "b/a+c-d") }
	
		if b/a+d-c == 24 { solutions = append(solutions, "b/a+d-c") }
	
		if b/c+a-d == 24 { solutions = append(solutions, "b/c+a-d") }
	
		if b/c+d-a == 24 { solutions = append(solutions, "b/c+d-a") }
	
		if b/d+c-a == 24 { solutions = append(solutions, "b/d+c-a") }
	
		if b/d+a-c == 24 { solutions = append(solutions, "b/d+a-c") }
	
		if c/b+a-d == 24 { solutions = append(solutions, "c/b+a-d") }
	
		if c/b+d-a == 24 { solutions = append(solutions, "c/b+d-a") }
	
		if c/a+b-d == 24 { solutions = append(solutions, "c/a+b-d") }
	
		if c/a+d-b == 24 { solutions = append(solutions, "c/a+d-b") }
	
		if c/d+a-b == 24 { solutions = append(solutions, "c/d+a-b") }
	
		if c/d+b-a == 24 { solutions = append(solutions, "c/d+b-a") }
	
		if d/b+c-a == 24 { solutions = append(solutions, "d/b+c-a") }
	
		if d/b+a-c == 24 { solutions = append(solutions, "d/b+a-c") }
	
		if d/c+b-a == 24 { solutions = append(solutions, "d/c+b-a") }
	
		if d/c+a-b == 24 { solutions = append(solutions, "d/c+a-b") }
	
		if d/a+c-b == 24 { solutions = append(solutions, "d/a+c-b") }
	
		if d/a+b-c == 24 { solutions = append(solutions, "d/a+b-c") }
	
		if a/b+c*d == 24 { solutions = append(solutions, "a/b+c*d") }
	
		if a/b+d*c == 24 { solutions = append(solutions, "a/b+d*c") }
	
		if a/c+b*d == 24 { solutions = append(solutions, "a/c+b*d") }
	
		if a/c+d*b == 24 { solutions = append(solutions, "a/c+d*b") }
	
		if a/d+c*b == 24 { solutions = append(solutions, "a/d+c*b") }
	
		if a/d+b*c == 24 { solutions = append(solutions, "a/d+b*c") }
	
		if b/a+c*d == 24 { solutions = append(solutions, "b/a+c*d") }
	
		if b/a+d*c == 24 { solutions = append(solutions, "b/a+d*c") }
	
		if b/c+a*d == 24 { solutions = append(solutions, "b/c+a*d") }
	
		if b/c+d*a == 24 { solutions = append(solutions, "b/c+d*a") }
	
		if b/d+c*a == 24 { solutions = append(solutions, "b/d+c*a") }
	
		if b/d+a*c == 24 { solutions = append(solutions, "b/d+a*c") }
	
		if c/b+a*d == 24 { solutions = append(solutions, "c/b+a*d") }
	
		if c/b+d*a == 24 { solutions = append(solutions, "c/b+d*a") }
	
		if c/a+b*d == 24 { solutions = append(solutions, "c/a+b*d") }
	
		if c/a+d*b == 24 { solutions = append(solutions, "c/a+d*b") }
	
		if c/d+a*b == 24 { solutions = append(solutions, "c/d+a*b") }
	
		if c/d+b*a == 24 { solutions = append(solutions, "c/d+b*a") }
	
		if d/b+c*a == 24 { solutions = append(solutions, "d/b+c*a") }
	
		if d/b+a*c == 24 { solutions = append(solutions, "d/b+a*c") }
	
		if d/c+b*a == 24 { solutions = append(solutions, "d/c+b*a") }
	
		if d/c+a*b == 24 { solutions = append(solutions, "d/c+a*b") }
	
		if d/a+c*b == 24 { solutions = append(solutions, "d/a+c*b") }
	
		if d/a+b*c == 24 { solutions = append(solutions, "d/a+b*c") }
	
		if a/b+c/d == 24 { solutions = append(solutions, "a/b+c/d") }
	
		if a/b+d/c == 24 { solutions = append(solutions, "a/b+d/c") }
	
		if a/c+b/d == 24 { solutions = append(solutions, "a/c+b/d") }
	
		if a/c+d/b == 24 { solutions = append(solutions, "a/c+d/b") }
	
		if a/d+c/b == 24 { solutions = append(solutions, "a/d+c/b") }
	
		if a/d+b/c == 24 { solutions = append(solutions, "a/d+b/c") }
	
		if b/a+c/d == 24 { solutions = append(solutions, "b/a+c/d") }
	
		if b/a+d/c == 24 { solutions = append(solutions, "b/a+d/c") }
	
		if b/c+a/d == 24 { solutions = append(solutions, "b/c+a/d") }
	
		if b/c+d/a == 24 { solutions = append(solutions, "b/c+d/a") }
	
		if b/d+c/a == 24 { solutions = append(solutions, "b/d+c/a") }
	
		if b/d+a/c == 24 { solutions = append(solutions, "b/d+a/c") }
	
		if c/b+a/d == 24 { solutions = append(solutions, "c/b+a/d") }
	
		if c/b+d/a == 24 { solutions = append(solutions, "c/b+d/a") }
	
		if c/a+b/d == 24 { solutions = append(solutions, "c/a+b/d") }
	
		if c/a+d/b == 24 { solutions = append(solutions, "c/a+d/b") }
	
		if c/d+a/b == 24 { solutions = append(solutions, "c/d+a/b") }
	
		if c/d+b/a == 24 { solutions = append(solutions, "c/d+b/a") }
	
		if d/b+c/a == 24 { solutions = append(solutions, "d/b+c/a") }
	
		if d/b+a/c == 24 { solutions = append(solutions, "d/b+a/c") }
	
		if d/c+b/a == 24 { solutions = append(solutions, "d/c+b/a") }
	
		if d/c+a/b == 24 { solutions = append(solutions, "d/c+a/b") }
	
		if d/a+c/b == 24 { solutions = append(solutions, "d/a+c/b") }
	
		if d/a+b/c == 24 { solutions = append(solutions, "d/a+b/c") }
	
		if a/b-c+d == 24 { solutions = append(solutions, "a/b-c+d") }
	
		if a/b-d+c == 24 { solutions = append(solutions, "a/b-d+c") }
	
		if a/c-b+d == 24 { solutions = append(solutions, "a/c-b+d") }
	
		if a/c-d+b == 24 { solutions = append(solutions, "a/c-d+b") }
	
		if a/d-c+b == 24 { solutions = append(solutions, "a/d-c+b") }
	
		if a/d-b+c == 24 { solutions = append(solutions, "a/d-b+c") }
	
		if b/a-c+d == 24 { solutions = append(solutions, "b/a-c+d") }
	
		if b/a-d+c == 24 { solutions = append(solutions, "b/a-d+c") }
	
		if b/c-a+d == 24 { solutions = append(solutions, "b/c-a+d") }
	
		if b/c-d+a == 24 { solutions = append(solutions, "b/c-d+a") }
	
		if b/d-c+a == 24 { solutions = append(solutions, "b/d-c+a") }
	
		if b/d-a+c == 24 { solutions = append(solutions, "b/d-a+c") }
	
		if c/b-a+d == 24 { solutions = append(solutions, "c/b-a+d") }
	
		if c/b-d+a == 24 { solutions = append(solutions, "c/b-d+a") }
	
		if c/a-b+d == 24 { solutions = append(solutions, "c/a-b+d") }
	
		if c/a-d+b == 24 { solutions = append(solutions, "c/a-d+b") }
	
		if c/d-a+b == 24 { solutions = append(solutions, "c/d-a+b") }
	
		if c/d-b+a == 24 { solutions = append(solutions, "c/d-b+a") }
	
		if d/b-c+a == 24 { solutions = append(solutions, "d/b-c+a") }
	
		if d/b-a+c == 24 { solutions = append(solutions, "d/b-a+c") }
	
		if d/c-b+a == 24 { solutions = append(solutions, "d/c-b+a") }
	
		if d/c-a+b == 24 { solutions = append(solutions, "d/c-a+b") }
	
		if d/a-c+b == 24 { solutions = append(solutions, "d/a-c+b") }
	
		if d/a-b+c == 24 { solutions = append(solutions, "d/a-b+c") }
	
		if a/b-c-d == 24 { solutions = append(solutions, "a/b-c-d") }
	
		if a/b-d-c == 24 { solutions = append(solutions, "a/b-d-c") }
	
		if a/c-b-d == 24 { solutions = append(solutions, "a/c-b-d") }
	
		if a/c-d-b == 24 { solutions = append(solutions, "a/c-d-b") }
	
		if a/d-c-b == 24 { solutions = append(solutions, "a/d-c-b") }
	
		if a/d-b-c == 24 { solutions = append(solutions, "a/d-b-c") }
	
		if b/a-c-d == 24 { solutions = append(solutions, "b/a-c-d") }
	
		if b/a-d-c == 24 { solutions = append(solutions, "b/a-d-c") }
	
		if b/c-a-d == 24 { solutions = append(solutions, "b/c-a-d") }
	
		if b/c-d-a == 24 { solutions = append(solutions, "b/c-d-a") }
	
		if b/d-c-a == 24 { solutions = append(solutions, "b/d-c-a") }
	
		if b/d-a-c == 24 { solutions = append(solutions, "b/d-a-c") }
	
		if c/b-a-d == 24 { solutions = append(solutions, "c/b-a-d") }
	
		if c/b-d-a == 24 { solutions = append(solutions, "c/b-d-a") }
	
		if c/a-b-d == 24 { solutions = append(solutions, "c/a-b-d") }
	
		if c/a-d-b == 24 { solutions = append(solutions, "c/a-d-b") }
	
		if c/d-a-b == 24 { solutions = append(solutions, "c/d-a-b") }
	
		if c/d-b-a == 24 { solutions = append(solutions, "c/d-b-a") }
	
		if d/b-c-a == 24 { solutions = append(solutions, "d/b-c-a") }
	
		if d/b-a-c == 24 { solutions = append(solutions, "d/b-a-c") }
	
		if d/c-b-a == 24 { solutions = append(solutions, "d/c-b-a") }
	
		if d/c-a-b == 24 { solutions = append(solutions, "d/c-a-b") }
	
		if d/a-c-b == 24 { solutions = append(solutions, "d/a-c-b") }
	
		if d/a-b-c == 24 { solutions = append(solutions, "d/a-b-c") }
	
		if a/b-c*d == 24 { solutions = append(solutions, "a/b-c*d") }
	
		if a/b-d*c == 24 { solutions = append(solutions, "a/b-d*c") }
	
		if a/c-b*d == 24 { solutions = append(solutions, "a/c-b*d") }
	
		if a/c-d*b == 24 { solutions = append(solutions, "a/c-d*b") }
	
		if a/d-c*b == 24 { solutions = append(solutions, "a/d-c*b") }
	
		if a/d-b*c == 24 { solutions = append(solutions, "a/d-b*c") }
	
		if b/a-c*d == 24 { solutions = append(solutions, "b/a-c*d") }
	
		if b/a-d*c == 24 { solutions = append(solutions, "b/a-d*c") }
	
		if b/c-a*d == 24 { solutions = append(solutions, "b/c-a*d") }
	
		if b/c-d*a == 24 { solutions = append(solutions, "b/c-d*a") }
	
		if b/d-c*a == 24 { solutions = append(solutions, "b/d-c*a") }
	
		if b/d-a*c == 24 { solutions = append(solutions, "b/d-a*c") }
	
		if c/b-a*d == 24 { solutions = append(solutions, "c/b-a*d") }
	
		if c/b-d*a == 24 { solutions = append(solutions, "c/b-d*a") }
	
		if c/a-b*d == 24 { solutions = append(solutions, "c/a-b*d") }
	
		if c/a-d*b == 24 { solutions = append(solutions, "c/a-d*b") }
	
		if c/d-a*b == 24 { solutions = append(solutions, "c/d-a*b") }
	
		if c/d-b*a == 24 { solutions = append(solutions, "c/d-b*a") }
	
		if d/b-c*a == 24 { solutions = append(solutions, "d/b-c*a") }
	
		if d/b-a*c == 24 { solutions = append(solutions, "d/b-a*c") }
	
		if d/c-b*a == 24 { solutions = append(solutions, "d/c-b*a") }
	
		if d/c-a*b == 24 { solutions = append(solutions, "d/c-a*b") }
	
		if d/a-c*b == 24 { solutions = append(solutions, "d/a-c*b") }
	
		if d/a-b*c == 24 { solutions = append(solutions, "d/a-b*c") }
	
		if a/b-c/d == 24 { solutions = append(solutions, "a/b-c/d") }
	
		if a/b-d/c == 24 { solutions = append(solutions, "a/b-d/c") }
	
		if a/c-b/d == 24 { solutions = append(solutions, "a/c-b/d") }
	
		if a/c-d/b == 24 { solutions = append(solutions, "a/c-d/b") }
	
		if a/d-c/b == 24 { solutions = append(solutions, "a/d-c/b") }
	
		if a/d-b/c == 24 { solutions = append(solutions, "a/d-b/c") }
	
		if b/a-c/d == 24 { solutions = append(solutions, "b/a-c/d") }
	
		if b/a-d/c == 24 { solutions = append(solutions, "b/a-d/c") }
	
		if b/c-a/d == 24 { solutions = append(solutions, "b/c-a/d") }
	
		if b/c-d/a == 24 { solutions = append(solutions, "b/c-d/a") }
	
		if b/d-c/a == 24 { solutions = append(solutions, "b/d-c/a") }
	
		if b/d-a/c == 24 { solutions = append(solutions, "b/d-a/c") }
	
		if c/b-a/d == 24 { solutions = append(solutions, "c/b-a/d") }
	
		if c/b-d/a == 24 { solutions = append(solutions, "c/b-d/a") }
	
		if c/a-b/d == 24 { solutions = append(solutions, "c/a-b/d") }
	
		if c/a-d/b == 24 { solutions = append(solutions, "c/a-d/b") }
	
		if c/d-a/b == 24 { solutions = append(solutions, "c/d-a/b") }
	
		if c/d-b/a == 24 { solutions = append(solutions, "c/d-b/a") }
	
		if d/b-c/a == 24 { solutions = append(solutions, "d/b-c/a") }
	
		if d/b-a/c == 24 { solutions = append(solutions, "d/b-a/c") }
	
		if d/c-b/a == 24 { solutions = append(solutions, "d/c-b/a") }
	
		if d/c-a/b == 24 { solutions = append(solutions, "d/c-a/b") }
	
		if d/a-c/b == 24 { solutions = append(solutions, "d/a-c/b") }
	
		if d/a-b/c == 24 { solutions = append(solutions, "d/a-b/c") }
	
		if a/b*c+d == 24 { solutions = append(solutions, "a/b*c+d") }
	
		if a/b*d+c == 24 { solutions = append(solutions, "a/b*d+c") }
	
		if a/c*b+d == 24 { solutions = append(solutions, "a/c*b+d") }
	
		if a/c*d+b == 24 { solutions = append(solutions, "a/c*d+b") }
	
		if a/d*c+b == 24 { solutions = append(solutions, "a/d*c+b") }
	
		if a/d*b+c == 24 { solutions = append(solutions, "a/d*b+c") }
	
		if b/a*c+d == 24 { solutions = append(solutions, "b/a*c+d") }
	
		if b/a*d+c == 24 { solutions = append(solutions, "b/a*d+c") }
	
		if b/c*a+d == 24 { solutions = append(solutions, "b/c*a+d") }
	
		if b/c*d+a == 24 { solutions = append(solutions, "b/c*d+a") }
	
		if b/d*c+a == 24 { solutions = append(solutions, "b/d*c+a") }
	
		if b/d*a+c == 24 { solutions = append(solutions, "b/d*a+c") }
	
		if c/b*a+d == 24 { solutions = append(solutions, "c/b*a+d") }
	
		if c/b*d+a == 24 { solutions = append(solutions, "c/b*d+a") }
	
		if c/a*b+d == 24 { solutions = append(solutions, "c/a*b+d") }
	
		if c/a*d+b == 24 { solutions = append(solutions, "c/a*d+b") }
	
		if c/d*a+b == 24 { solutions = append(solutions, "c/d*a+b") }
	
		if c/d*b+a == 24 { solutions = append(solutions, "c/d*b+a") }
	
		if d/b*c+a == 24 { solutions = append(solutions, "d/b*c+a") }
	
		if d/b*a+c == 24 { solutions = append(solutions, "d/b*a+c") }
	
		if d/c*b+a == 24 { solutions = append(solutions, "d/c*b+a") }
	
		if d/c*a+b == 24 { solutions = append(solutions, "d/c*a+b") }
	
		if d/a*c+b == 24 { solutions = append(solutions, "d/a*c+b") }
	
		if d/a*b+c == 24 { solutions = append(solutions, "d/a*b+c") }
	
		if a/b*c-d == 24 { solutions = append(solutions, "a/b*c-d") }
	
		if a/b*d-c == 24 { solutions = append(solutions, "a/b*d-c") }
	
		if a/c*b-d == 24 { solutions = append(solutions, "a/c*b-d") }
	
		if a/c*d-b == 24 { solutions = append(solutions, "a/c*d-b") }
	
		if a/d*c-b == 24 { solutions = append(solutions, "a/d*c-b") }
	
		if a/d*b-c == 24 { solutions = append(solutions, "a/d*b-c") }
	
		if b/a*c-d == 24 { solutions = append(solutions, "b/a*c-d") }
	
		if b/a*d-c == 24 { solutions = append(solutions, "b/a*d-c") }
	
		if b/c*a-d == 24 { solutions = append(solutions, "b/c*a-d") }
	
		if b/c*d-a == 24 { solutions = append(solutions, "b/c*d-a") }
	
		if b/d*c-a == 24 { solutions = append(solutions, "b/d*c-a") }
	
		if b/d*a-c == 24 { solutions = append(solutions, "b/d*a-c") }
	
		if c/b*a-d == 24 { solutions = append(solutions, "c/b*a-d") }
	
		if c/b*d-a == 24 { solutions = append(solutions, "c/b*d-a") }
	
		if c/a*b-d == 24 { solutions = append(solutions, "c/a*b-d") }
	
		if c/a*d-b == 24 { solutions = append(solutions, "c/a*d-b") }
	
		if c/d*a-b == 24 { solutions = append(solutions, "c/d*a-b") }
	
		if c/d*b-a == 24 { solutions = append(solutions, "c/d*b-a") }
	
		if d/b*c-a == 24 { solutions = append(solutions, "d/b*c-a") }
	
		if d/b*a-c == 24 { solutions = append(solutions, "d/b*a-c") }
	
		if d/c*b-a == 24 { solutions = append(solutions, "d/c*b-a") }
	
		if d/c*a-b == 24 { solutions = append(solutions, "d/c*a-b") }
	
		if d/a*c-b == 24 { solutions = append(solutions, "d/a*c-b") }
	
		if d/a*b-c == 24 { solutions = append(solutions, "d/a*b-c") }
	
		if a/b*c*d == 24 { solutions = append(solutions, "a/b*c*d") }
	
		if a/b*d*c == 24 { solutions = append(solutions, "a/b*d*c") }
	
		if a/c*b*d == 24 { solutions = append(solutions, "a/c*b*d") }
	
		if a/c*d*b == 24 { solutions = append(solutions, "a/c*d*b") }
	
		if a/d*c*b == 24 { solutions = append(solutions, "a/d*c*b") }
	
		if a/d*b*c == 24 { solutions = append(solutions, "a/d*b*c") }
	
		if b/a*c*d == 24 { solutions = append(solutions, "b/a*c*d") }
	
		if b/a*d*c == 24 { solutions = append(solutions, "b/a*d*c") }
	
		if b/c*a*d == 24 { solutions = append(solutions, "b/c*a*d") }
	
		if b/c*d*a == 24 { solutions = append(solutions, "b/c*d*a") }
	
		if b/d*c*a == 24 { solutions = append(solutions, "b/d*c*a") }
	
		if b/d*a*c == 24 { solutions = append(solutions, "b/d*a*c") }
	
		if c/b*a*d == 24 { solutions = append(solutions, "c/b*a*d") }
	
		if c/b*d*a == 24 { solutions = append(solutions, "c/b*d*a") }
	
		if c/a*b*d == 24 { solutions = append(solutions, "c/a*b*d") }
	
		if c/a*d*b == 24 { solutions = append(solutions, "c/a*d*b") }
	
		if c/d*a*b == 24 { solutions = append(solutions, "c/d*a*b") }
	
		if c/d*b*a == 24 { solutions = append(solutions, "c/d*b*a") }
	
		if d/b*c*a == 24 { solutions = append(solutions, "d/b*c*a") }
	
		if d/b*a*c == 24 { solutions = append(solutions, "d/b*a*c") }
	
		if d/c*b*a == 24 { solutions = append(solutions, "d/c*b*a") }
	
		if d/c*a*b == 24 { solutions = append(solutions, "d/c*a*b") }
	
		if d/a*c*b == 24 { solutions = append(solutions, "d/a*c*b") }
	
		if d/a*b*c == 24 { solutions = append(solutions, "d/a*b*c") }
	
		if a/b*c/d == 24 { solutions = append(solutions, "a/b*c/d") }
	
		if a/b*d/c == 24 { solutions = append(solutions, "a/b*d/c") }
	
		if a/c*b/d == 24 { solutions = append(solutions, "a/c*b/d") }
	
		if a/c*d/b == 24 { solutions = append(solutions, "a/c*d/b") }
	
		if a/d*c/b == 24 { solutions = append(solutions, "a/d*c/b") }
	
		if a/d*b/c == 24 { solutions = append(solutions, "a/d*b/c") }
	
		if b/a*c/d == 24 { solutions = append(solutions, "b/a*c/d") }
	
		if b/a*d/c == 24 { solutions = append(solutions, "b/a*d/c") }
	
		if b/c*a/d == 24 { solutions = append(solutions, "b/c*a/d") }
	
		if b/c*d/a == 24 { solutions = append(solutions, "b/c*d/a") }
	
		if b/d*c/a == 24 { solutions = append(solutions, "b/d*c/a") }
	
		if b/d*a/c == 24 { solutions = append(solutions, "b/d*a/c") }
	
		if c/b*a/d == 24 { solutions = append(solutions, "c/b*a/d") }
	
		if c/b*d/a == 24 { solutions = append(solutions, "c/b*d/a") }
	
		if c/a*b/d == 24 { solutions = append(solutions, "c/a*b/d") }
	
		if c/a*d/b == 24 { solutions = append(solutions, "c/a*d/b") }
	
		if c/d*a/b == 24 { solutions = append(solutions, "c/d*a/b") }
	
		if c/d*b/a == 24 { solutions = append(solutions, "c/d*b/a") }
	
		if d/b*c/a == 24 { solutions = append(solutions, "d/b*c/a") }
	
		if d/b*a/c == 24 { solutions = append(solutions, "d/b*a/c") }
	
		if d/c*b/a == 24 { solutions = append(solutions, "d/c*b/a") }
	
		if d/c*a/b == 24 { solutions = append(solutions, "d/c*a/b") }
	
		if d/a*c/b == 24 { solutions = append(solutions, "d/a*c/b") }
	
		if d/a*b/c == 24 { solutions = append(solutions, "d/a*b/c") }
	
		if a/b/c+d == 24 { solutions = append(solutions, "a/b/c+d") }
	
		if a/b/d+c == 24 { solutions = append(solutions, "a/b/d+c") }
	
		if a/c/b+d == 24 { solutions = append(solutions, "a/c/b+d") }
	
		if a/c/d+b == 24 { solutions = append(solutions, "a/c/d+b") }
	
		if a/d/c+b == 24 { solutions = append(solutions, "a/d/c+b") }
	
		if a/d/b+c == 24 { solutions = append(solutions, "a/d/b+c") }
	
		if b/a/c+d == 24 { solutions = append(solutions, "b/a/c+d") }
	
		if b/a/d+c == 24 { solutions = append(solutions, "b/a/d+c") }
	
		if b/c/a+d == 24 { solutions = append(solutions, "b/c/a+d") }
	
		if b/c/d+a == 24 { solutions = append(solutions, "b/c/d+a") }
	
		if b/d/c+a == 24 { solutions = append(solutions, "b/d/c+a") }
	
		if b/d/a+c == 24 { solutions = append(solutions, "b/d/a+c") }
	
		if c/b/a+d == 24 { solutions = append(solutions, "c/b/a+d") }
	
		if c/b/d+a == 24 { solutions = append(solutions, "c/b/d+a") }
	
		if c/a/b+d == 24 { solutions = append(solutions, "c/a/b+d") }
	
		if c/a/d+b == 24 { solutions = append(solutions, "c/a/d+b") }
	
		if c/d/a+b == 24 { solutions = append(solutions, "c/d/a+b") }
	
		if c/d/b+a == 24 { solutions = append(solutions, "c/d/b+a") }
	
		if d/b/c+a == 24 { solutions = append(solutions, "d/b/c+a") }
	
		if d/b/a+c == 24 { solutions = append(solutions, "d/b/a+c") }
	
		if d/c/b+a == 24 { solutions = append(solutions, "d/c/b+a") }
	
		if d/c/a+b == 24 { solutions = append(solutions, "d/c/a+b") }
	
		if d/a/c+b == 24 { solutions = append(solutions, "d/a/c+b") }
	
		if d/a/b+c == 24 { solutions = append(solutions, "d/a/b+c") }
	
		if a/b/c-d == 24 { solutions = append(solutions, "a/b/c-d") }
	
		if a/b/d-c == 24 { solutions = append(solutions, "a/b/d-c") }
	
		if a/c/b-d == 24 { solutions = append(solutions, "a/c/b-d") }
	
		if a/c/d-b == 24 { solutions = append(solutions, "a/c/d-b") }
	
		if a/d/c-b == 24 { solutions = append(solutions, "a/d/c-b") }
	
		if a/d/b-c == 24 { solutions = append(solutions, "a/d/b-c") }
	
		if b/a/c-d == 24 { solutions = append(solutions, "b/a/c-d") }
	
		if b/a/d-c == 24 { solutions = append(solutions, "b/a/d-c") }
	
		if b/c/a-d == 24 { solutions = append(solutions, "b/c/a-d") }
	
		if b/c/d-a == 24 { solutions = append(solutions, "b/c/d-a") }
	
		if b/d/c-a == 24 { solutions = append(solutions, "b/d/c-a") }
	
		if b/d/a-c == 24 { solutions = append(solutions, "b/d/a-c") }
	
		if c/b/a-d == 24 { solutions = append(solutions, "c/b/a-d") }
	
		if c/b/d-a == 24 { solutions = append(solutions, "c/b/d-a") }
	
		if c/a/b-d == 24 { solutions = append(solutions, "c/a/b-d") }
	
		if c/a/d-b == 24 { solutions = append(solutions, "c/a/d-b") }
	
		if c/d/a-b == 24 { solutions = append(solutions, "c/d/a-b") }
	
		if c/d/b-a == 24 { solutions = append(solutions, "c/d/b-a") }
	
		if d/b/c-a == 24 { solutions = append(solutions, "d/b/c-a") }
	
		if d/b/a-c == 24 { solutions = append(solutions, "d/b/a-c") }
	
		if d/c/b-a == 24 { solutions = append(solutions, "d/c/b-a") }
	
		if d/c/a-b == 24 { solutions = append(solutions, "d/c/a-b") }
	
		if d/a/c-b == 24 { solutions = append(solutions, "d/a/c-b") }
	
		if d/a/b-c == 24 { solutions = append(solutions, "d/a/b-c") }
	
		if a/b/c*d == 24 { solutions = append(solutions, "a/b/c*d") }
	
		if a/b/d*c == 24 { solutions = append(solutions, "a/b/d*c") }
	
		if a/c/b*d == 24 { solutions = append(solutions, "a/c/b*d") }
	
		if a/c/d*b == 24 { solutions = append(solutions, "a/c/d*b") }
	
		if a/d/c*b == 24 { solutions = append(solutions, "a/d/c*b") }
	
		if a/d/b*c == 24 { solutions = append(solutions, "a/d/b*c") }
	
		if b/a/c*d == 24 { solutions = append(solutions, "b/a/c*d") }
	
		if b/a/d*c == 24 { solutions = append(solutions, "b/a/d*c") }
	
		if b/c/a*d == 24 { solutions = append(solutions, "b/c/a*d") }
	
		if b/c/d*a == 24 { solutions = append(solutions, "b/c/d*a") }
	
		if b/d/c*a == 24 { solutions = append(solutions, "b/d/c*a") }
	
		if b/d/a*c == 24 { solutions = append(solutions, "b/d/a*c") }
	
		if c/b/a*d == 24 { solutions = append(solutions, "c/b/a*d") }
	
		if c/b/d*a == 24 { solutions = append(solutions, "c/b/d*a") }
	
		if c/a/b*d == 24 { solutions = append(solutions, "c/a/b*d") }
	
		if c/a/d*b == 24 { solutions = append(solutions, "c/a/d*b") }
	
		if c/d/a*b == 24 { solutions = append(solutions, "c/d/a*b") }
	
		if c/d/b*a == 24 { solutions = append(solutions, "c/d/b*a") }
	
		if d/b/c*a == 24 { solutions = append(solutions, "d/b/c*a") }
	
		if d/b/a*c == 24 { solutions = append(solutions, "d/b/a*c") }
	
		if d/c/b*a == 24 { solutions = append(solutions, "d/c/b*a") }
	
		if d/c/a*b == 24 { solutions = append(solutions, "d/c/a*b") }
	
		if d/a/c*b == 24 { solutions = append(solutions, "d/a/c*b") }
	
		if d/a/b*c == 24 { solutions = append(solutions, "d/a/b*c") }
	
		if a/b/c/d == 24 { solutions = append(solutions, "a/b/c/d") }
	
		if a/b/d/c == 24 { solutions = append(solutions, "a/b/d/c") }
	
		if a/c/b/d == 24 { solutions = append(solutions, "a/c/b/d") }
	
		if a/c/d/b == 24 { solutions = append(solutions, "a/c/d/b") }
	
		if a/d/c/b == 24 { solutions = append(solutions, "a/d/c/b") }
	
		if a/d/b/c == 24 { solutions = append(solutions, "a/d/b/c") }
	
		if b/a/c/d == 24 { solutions = append(solutions, "b/a/c/d") }
	
		if b/a/d/c == 24 { solutions = append(solutions, "b/a/d/c") }
	
		if b/c/a/d == 24 { solutions = append(solutions, "b/c/a/d") }
	
		if b/c/d/a == 24 { solutions = append(solutions, "b/c/d/a") }
	
		if b/d/c/a == 24 { solutions = append(solutions, "b/d/c/a") }
	
		if b/d/a/c == 24 { solutions = append(solutions, "b/d/a/c") }
	
		if c/b/a/d == 24 { solutions = append(solutions, "c/b/a/d") }
	
		if c/b/d/a == 24 { solutions = append(solutions, "c/b/d/a") }
	
		if c/a/b/d == 24 { solutions = append(solutions, "c/a/b/d") }
	
		if c/a/d/b == 24 { solutions = append(solutions, "c/a/d/b") }
	
		if c/d/a/b == 24 { solutions = append(solutions, "c/d/a/b") }
	
		if c/d/b/a == 24 { solutions = append(solutions, "c/d/b/a") }
	
		if d/b/c/a == 24 { solutions = append(solutions, "d/b/c/a") }
	
		if d/b/a/c == 24 { solutions = append(solutions, "d/b/a/c") }
	
		if d/c/b/a == 24 { solutions = append(solutions, "d/c/b/a") }
	
		if d/c/a/b == 24 { solutions = append(solutions, "d/c/a/b") }
	
		if d/a/c/b == 24 { solutions = append(solutions, "d/a/c/b") }
	
		if d/a/b/c == 24 { solutions = append(solutions, "d/a/b/c") }
	
	return solutions
}
