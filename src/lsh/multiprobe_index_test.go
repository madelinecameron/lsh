package lsh

import "testing"

func Test_NewMultiprobeLsh(t *testing.T) {
	lsh := NewMultiprobeLsh(100, 5, 5, 5.0, 64)
	if len(lsh.SimpleIndex.tables) != 5 {
		t.Error("Lsh init fail")
	}
	t.Logf("Scores %v", lsh.scores)
	t.Logf("Perturbation sets: %v", lsh.perturbSets)
	for i, v := range lsh.perturbSets {
		t.Logf("Set: %d, Score: %f, Set contents: %v", i, lsh.getScore(&v), v)
	}
	for i, perSet := range lsh.perturbVecs {
		for j, perTable := range perSet {
			t.Logf("Set: %d, Table: %d, Vec: %v", i, j, perTable)
		}
	}

}

func Test_MultiprobeLshQueryK(t *testing.T) {
	lsh := NewMultiprobeLsh(100, 5, 5, 5.0, 10)
	points := randomPoints(10, 100, 32.0)
	insertedKeys := make([]int, 10)
	for i, p := range points {
		lsh.Insert(p, i)
		insertedKeys[i] = i
	}
	// Use the inserted points as queries, and
	// verify that we can get back each query itself
	for i, key := range insertedKeys {
		result := make(chan int)
		go func() {
			lsh.QueryK(points[i], 10, result)
			close(result)
		}()
		found := false
		for foundKey := range result {
			if foundKey == key {
				found = true
			}
		}
		if !found {
			t.Error("Query fail")
		}
	}
}

/*
func Test_MultiprobeLshQueryK(t *testing.T) {
	lsh := NewMultiprobeLsh(100, 5, 5, 5.0, 64)
	points := randomPoints(10, 100, 32.0)
	insertedKeys := make([]int, 10)
	for i, p := range points {
		lsh.Insert(p, i)
		insertedKeys[i] = i
	}

	// Query a single point, should obtain back the entire index.
	result := make(chan int)
	go func() {
		lsh.QueryK(points[0], 10, result)
		close(result)
	}()
	actual := make([]int, 0)
	for key := range result {
		actual = append(actual, key)
	}

	// Use the inserted points as queries, and
	// verify that we can get back each query itself
	for _, key := range insertedKeys {
		found := false
		for foundKey := range actual {
			if foundKey == key {
				found = true
			}
		}
		if !found {
			t.Error("Query fail")
		}
	}
}
*/
