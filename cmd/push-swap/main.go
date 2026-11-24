// push-swap affiche des instructions qui trient les nombres reçus.
package main

import (
	"fmt"
	"os"
	"sort"

	"push-swap/internal/stack"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	nums, ok := stack.Parse(os.Args[1:])
	if !ok {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	for _, op := range sortOps(nums) {
		fmt.Println(op)
	}
}

// sortOps renvoie les instructions qui trient nums (le premier nombre est au sommet de a).
//   - 2 nombres : sa
//   - 3 à 5 nombres : on envoie les plus petits dans b, on trie les 3 restants, on remet b dans a
//   - plus de 5 nombres : tri radix (un tour par bit)
func sortOps(nums []int) []string {
	// Chaque nombre est remplacé par son rang (0, 1, 2...) : ça marche aussi avec les négatifs.
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)
	ranks := make([]int, len(nums))
	for i, n := range nums {
		ranks[i] = sort.SearchInts(sorted, n)
	}

	s := &stack.Stacks{A: ranks}
	ops := []string{}
	do := func(op string) {
		s.Apply(op)
		ops = append(ops, op)
	}

	n := len(nums)
	switch {
	case s.IsSorted():
		// rien à faire
	case n == 2:
		do("sa")
	case n <= 5:
		for len(s.A) > 3 {
			// Les plus petits sont déjà dans b : le plus petit de a a pour rang len(b).
			i := 0
			for s.A[i] != len(s.B) {
				i++
			}
			if i <= len(s.A)/2 {
				for ; i > 0; i-- {
					do("ra")
				}
			} else {
				for ; i < len(s.A); i++ {
					do("rra")
				}
			}
			do("pb")
		}
		t, m, b := s.A[0], s.A[1], s.A[2]
		switch {
		case t < m && m < b: // déjà triés
		case t > m && m > b: // 3 2 1
			do("sa")
			do("rra")
		case t > m && b > t: // 2 1 3
			do("sa")
		case t > m: // 3 1 2
			do("ra")
		case b > t: // 1 3 2
			do("sa")
			do("ra")
		default: // 2 3 1
			do("rra")
		}
		for len(s.B) > 0 {
			do("pa")
		}
	default:
		// Radix : pour chaque bit, les 0 vont dans b, les 1 restent dans a, puis on remet b dans a.
		for bit := 0; (n-1)>>bit != 0; bit++ {
			for i := 0; i < n; i++ {
				if (s.A[0]>>bit)&1 == 1 {
					do("ra")
				} else {
					do("pb")
				}
			}
			for len(s.B) > 0 {
				do("pa")
			}
		}
	}
	return ops
}
