package main

import (
	"math/rand"
	"testing"

	"push-swap/internal/stack"
)

// check applique les instructions renvoyées par sortOps et vérifie que a est triée.
func check(t *testing.T, nums []int) int {
	s := &stack.Stacks{A: append([]int{}, nums...)}
	ops := sortOps(nums)
	for _, op := range ops {
		s.Apply(op)
	}
	if !s.IsSorted() {
		t.Fatalf("%v : pas trié (a=%v b=%v)", nums, s.A, s.B)
	}
	return len(ops)
}

func TestSmallCases(t *testing.T) {
	if check(t, []int{}) != 0 || check(t, []int{42}) != 0 || check(t, []int{1, 2, 3, 4}) != 0 {
		t.Error("aucune instruction attendue")
	}
}

func TestAllPermutationsUpToFive(t *testing.T) {
	var permute func(nums []int, k int)
	permute = func(nums []int, k int) {
		if k == len(nums) {
			if count := check(t, nums); len(nums) == 5 && count > 12 {
				t.Errorf("%v : %d instructions (max attendu 12)", nums, count)
			}
			return
		}
		for i := k; i < len(nums); i++ {
			nums[k], nums[i] = nums[i], nums[k]
			permute(nums, k+1)
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
	for n := 2; n <= 5; n++ {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = i*10 - 15 // avec des négatifs
		}
		permute(nums, 0)
	}
}

func TestRandomLists(t *testing.T) {
	for _, n := range []int{6, 7, 10, 100, 500} {
		nums := rand.Perm(n)
		for i := range nums {
			nums[i] -= n / 2
		}
		check(t, nums)
	}
}
