package PriorityQueue

import (
	"testing"
)

func TestIntQueue(t *testing.T) {
	pq := New[int](func(i, j int) bool {
		return i < j
	})

	pq.Push(2)
	pq.Push(3)
	pq.Push(1)

	if pq.Len() != 3 {
		t.Error("pe.Len() != 3")
	}

	for i := 1; i <= 3; i++ {
		v, _ := pq.Pop()
		if v != i {
			t.Error("pop data not match")
		}
	}
}

func TestStringQueue(t *testing.T) {
	pq := New[string](func(i, j string) bool {
		return i < j
	})

	pq.Push("xxx")
	pq.Push("abc")
	pq.Push("aaa")

	if pq.Len() != 3 {
		t.Error("pe.Len() != 3")
	}

	v, _ := pq.Pop()
	if v != "aaa" {
		t.Error("pop data not aaa")
	}

	v, _ = pq.Pop()
	if v != "abc" {
		t.Error("pop data not abc")
	}

	v, _ = pq.Pop()
	if v != "xxx" {
		t.Error("pop data not xxx")
	}
}

func TestEmpty(t *testing.T) {

	var pq Queue[int]

	if pq.Len() != 0 {
		t.Error("pq.Len() != 0")
	}

	_, ok := pq.Pop()

	if ok {
		t.Error("queue not empty")
	}
}

type pair struct {
	x, y int
}

func TestQueuePair(t *testing.T) {
	pq := New[pair](func(i, j pair) bool {
		return i.x < j.x
	})

	pq.Push(pair{2, 1})
	pq.Push(pair{3, -1})
	pq.Push(pair{1, 10})

	if pq.Len() != 3 {
		t.FailNow()
	}

	for i := 1; i <= 3; i++ {
		v, ok := pq.Pop()
		if v.x != i || ok == false {
			t.FailNow()
		}
	}

	if pq.Len() != 0 {
		t.FailNow()
	}

	_, ok := pq.Pop()
	if ok == true {
		t.FailNow()
	}
}

type datas struct {
	x, y, z int
	s       string
}

func TestQueueDatas(t *testing.T) {
	pq := New[datas](func(i, j datas) bool {
		return i.s < j.s
	})

	pq.Push(datas{2, 1, 3, "aaa"})
	pq.Push(datas{3, -1, 3, "bbb"})
	pq.Push(datas{1, 10, 4, "ccc"})
	pq.Push(datas{2, 1, 3, "xyz"})
	pq.Push(datas{3, -1, 3, "bbb"})

	if pq.Len() != 5 {
		t.FailNow()
	}

	for _, e := range []string{"aaa", "bbb", "bbb", "ccc", "xyz"} {
		v, ok := pq.Pop()
		if ok == false {
			break
		}
		if v.s != e {
			t.FailNow()
		}
	}

	if pq.Len() != 0 {
		t.FailNow()
	}

	_, ok := pq.Pop()
	if ok == true {
		t.FailNow()
	}
}
