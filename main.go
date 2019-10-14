package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func generateMatrix(m, n int) *[][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(100)
		}
	}
	return &matrix
}

func blankMatrix(m, n int) *[][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	return &matrix
}

func sortRow(m *[][]int) {
	var wg sync.WaitGroup
	wg.Add(len(*m))
	for i := range *m {
		go func(i int) {
			sort.Ints((*m)[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func transpose(m, t *[][]int) {
	k := 0
	h := 0
	for i := range *m {
		for j := range (*m)[i] {
			if k == len(*t) {
				k = 0
				h++
			}
			(*t)[k][h] = (*m)[i][j]
			k++
		}
	}
}

func untranspose(m, t *[][]int) {
	k := 0
	h := 0
	for i := range *m {
		for j := range (*m)[i] {
			if k == len(*t) {
				k = 0
				h++
			}
			(*m)[i][j] = (*t)[k][h]
			k++
		}
	}
}

func slide(m *[][]int) *[][]int {
	s := make([][]int, len(*m)+1)
	for i := range s {
		s[i] = make([]int, len((*m)[0]))
	}

	for i := 0; i < len((*m)[0]); i++ {
		s[0][i] = MinInt
		s[len(s)-1][i] = MaxInt
	}

	k := 0
	h := len((*m)[0]) / 2

	for i := range *m {
		for j := range (*m)[i] {
			if h == len((*m)[i]) {
				h = 0
				k++
			}
			s[k][h] = (*m)[i][j]
			h++
		}
	}
	return &s
}

func unslide(s, m *[][]int) {
	k := 0
	h := len((*m)[0]) / 2
	for i := range *m {
		for j := range (*m)[i] {
			if h == len((*s)[k]) {
				h = 0
				k++
			}
			(*m)[i][j] = (*s)[k][h]
			h++
		}
	}
}

func print(m *[][]int) {
	for i := range *m {
		fmt.Println((*m)[i])
	}
}

func getSize() (int, int, error) {
	var m, n int
	fmt.Println("Enter the number of rows: ")
	_, e := fmt.Scan(&m)
	if e != nil {
		return 0, 0, e
	}
	fmt.Println("Enter the number of columns: ")
	_, e = fmt.Scan(&n)
	if e != nil {
		return 0, 0, e
	}
	return m, n, nil
}

func showAll() bool {
	var s string
	fmt.Println("Would you like to show all of the steps? (y/n)")
	for {
		_, _ = fmt.Scan(&s)
		if s == "y" || s == "Y" {
			return true
		} else if s == "n" || s == "N" {
			return false
		}
		fmt.Println("Please enter y or n")
	}
}

func invalid(m, n int) bool {
	if n%m != 0 || 2*(m-1)*(m-1) > n {
		return true
	}
	return false
}

func yesPrint(r, c int) {
	m := generateMatrix(c, r)
	fmt.Println("\nInitial Matrix")
	sortRow(m)
	t := blankMatrix(c, r)
	transpose(m, t)
	fmt.Println("\nSorted initial matrix")
	print(m)
	fmt.Println("\nTransposed matrix")
	print(t)
	fmt.Println("\nSorted transpose")
	sortRow(t)
	print(t)
	fmt.Println("\nUntransposed matrix")
	untranspose(m, t)
	t = nil
	print(m)
	fmt.Println("\nSorted untransposed")
	sortRow(m)
	print(m)
	fmt.Println("\nSlide")
	s := slide(m)
	print(s)
	fmt.Println("\nSorted slide")
	sortRow(s)
	print(s)
	fmt.Println("\nUnslide")
	unslide(s, m)
	print(m)
	s = nil
	fmt.Println("\nFinal sort")
	sortRow(m)
	print(m)
}

func noPrint(r, c int) {
	m := generateMatrix(c, r)
	start := time.Now()
	sortRow(m)
	t := blankMatrix(c, r)
	transpose(m, t)
	sortRow(t)
	untranspose(m, t)
	t = nil
	sortRow(m)
	s := slide(m)
	sortRow(s)
	unslide(s, m)
	s = nil
	sortRow(m)
	elapsed := time.Since(start)
	fmt.Printf("\nThe sort took %s\n", elapsed)
}

func main() {
	fmt.Println("This is a program designed to show Grouping Sort as described on the test with an additional sort applied at the end")
	fmt.Println("Also note, the matrix will be be row-major rather than column-major")
	for {
		m, n, e := getSize()
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		if invalid(m, n) {
			fmt.Println("Invalid matrix size")
		} else if !showAll() {
			noPrint(m, n)
		} else {
			yesPrint(m, n)
		}
	}
}
