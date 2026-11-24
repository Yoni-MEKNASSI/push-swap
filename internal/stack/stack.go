// Package stack contient les piles a et b, les 11 instructions et la lecture des arguments.
// Le sommet d'une pile est son premier élément (indice 0).
package stack

import (
	"strconv"
	"strings"
)

type Stacks struct{ A, B []int }

// Parse transforme les arguments en nombres.
// ok vaut false si un élément n'est pas un entier ou s'il y a un doublon.
func Parse(args []string) (nums []int, ok bool) {
	seen := map[int]bool{}
	for _, arg := range args {
		for _, word := range strings.Fields(arg) {
			n, err := strconv.Atoi(word)
			if err != nil || seen[n] {
				return nil, false
			}
			seen[n] = true
			nums = append(nums, n)
		}
	}
	return nums, true
}

// Apply exécute une instruction. Elle renvoie false si l'instruction n'existe pas.
func (s *Stacks) Apply(op string) bool {
	switch op {
	case "pa":
		if len(s.B) > 0 {
			s.A = append([]int{s.B[0]}, s.A...)
			s.B = s.B[1:]
		}
	case "pb":
		if len(s.A) > 0 {
			s.B = append([]int{s.A[0]}, s.B...)
			s.A = s.A[1:]
		}
	case "sa":
		swap(s.A)
	case "sb":
		swap(s.B)
	case "ss":
		swap(s.A)
		swap(s.B)
	case "ra":
		s.A = rotate(s.A)
	case "rb":
		s.B = rotate(s.B)
	case "rr":
		s.A, s.B = rotate(s.A), rotate(s.B)
	case "rra":
		s.A = reverse(s.A)
	case "rrb":
		s.B = reverse(s.B)
	case "rrr":
		s.A, s.B = reverse(s.A), reverse(s.B)
	default:
		return false
	}
	return true
}

// IsSorted renvoie true si a est triée par ordre croissant et b est vide.
func (s *Stacks) IsSorted() bool {
	for i := 1; i < len(s.A); i++ {
		if s.A[i-1] > s.A[i] {
			return false
		}
	}
	return len(s.B) == 0
}

// swap échange les deux premiers éléments.
func swap(l []int) {
	if len(l) > 1 {
		l[0], l[1] = l[1], l[0]
	}
}

// rotate : le premier élément devient le dernier.
func rotate(l []int) []int {
	if len(l) < 2 {
		return l
	}
	return append(append([]int{}, l[1:]...), l[0])
}

// reverse : le dernier élément devient le premier.
func reverse(l []int) []int {
	if len(l) < 2 {
		return l
	}
	return append([]int{l[len(l)-1]}, l[:len(l)-1]...)
}
