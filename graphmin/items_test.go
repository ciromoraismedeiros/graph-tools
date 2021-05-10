package graphmin

import (
	"fmt"
	// "github.com/ciromdrs/graph-tools/ccfpq"
	ds "github.com/ciromdrs/graph-tools/datastructures"
	. "github.com/ciromdrs/graph-tools/util"
	"testing"
)

func TestAugItem(t *testing.T) {
	S := ds.NewSimpleVertex("S")
	a := ds.NewSimpleVertex("a")
	b := ds.NewSimpleVertex("b")
	c := ds.NewSimpleVertex("c")
	s := ds.NewSimpleVertex("s")
	o := ds.NewSimpleVertex("o")
	e1 := newEdge(s, a, o)

	rule := []ds.Vertex{S, a, b, c}
	item := newAugItem(rule)
	if item.rule[0] != S || item.rule[1] != a || item.rule[2] != b ||
		item.rule[3] != c {
		t.Fatalf("Wrong rule. Expected %v, got %v", rule, item.rule)
	}
	if len(item.edges) < 3 {
		t.Fatalf("Expected edges of length 3, got %v", len(item.edges))
	}

	AssertPanic(t, func() { item.addEdge(e1, 0) },
		fmt.Sprintf("Should not add inexistent edge %v.", e1))
	e1.exists = true
	item.addEdge(e1, 1)
	if item.edges[1][0] != e1 {
		t.Fatalf("Eror adding edge. Expected %v, got %v", e1, item.edges[1][0])
	}
	{
		want := itemPos{item: item, pos: 0}
		if e1.dependencies[0].item != item || e1.dependencies[0].pos != 1 {
			t.Fatalf("Eror adding dependency. Expected %v, got %v",
				want, e1.dependencies[0])
		}
	}
	AssertPanic(t, func() { item.addEdge(e1, 0) },
		fmt.Sprintf("Should not add duplicated dependency %v %d.", e1, 0))
	e2 := newEdge(s, b, o)
	e2.exists = true
	AssertPanic(t, func() { item.addEdge(e2, 0) },
		"Should not add edge with wrong predicate b.")
}

func TestAugItemSet(t *testing.T) {
	f := NewHashFactory()
	f.NewAugItemSet(0)
}