package PriorityQueue

// Queue ... Priority Queue
type Queue[T any] struct {
	data []T
	less func(i, j T) bool
}

// New ... Create new priority queue
func New[T any](less func(i, j T) bool) Queue[T] {
	var ret Queue[T]
	ret.data = make([]T, 0)
	ret.less = less
	return ret
}

func (q Queue[T]) Len() int {
	return len(q.data)
}

func (q Queue[T]) Swap(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *Queue[T]) Push(x T) {
	q.data = append(q.data, x)
	cur := q.Len()
	parent := cur / 2
	for cur != 1 {
		if q.less(q.data[cur-1], q.data[parent-1]) {
			q.Swap(cur-1, parent-1)
		} else {
			break
		}
		cur = parent
		parent = cur / 2
	}
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.Len() == 0 {
		var item T
		return item, false
	}
	old := *q
	n := len(old.data)
	item := old.data[0]

	old.data[0] = old.data[n-1]
	old.data = old.data[:n-1]
	cur := 1
	for {
		nxt0 := cur * 2
		nxt1 := cur*2 + 1
		if nxt0 > len(old.data) {
			break
		}
		nxt := nxt0
		if nxt1 <= len(old.data) && old.less(q.data[nxt1-1], q.data[nxt0-1]) {
			nxt = nxt1
		}
		if old.less(q.data[nxt-1], q.data[cur-1]) {
			old.Swap(nxt-1, cur-1)
		} else {
			break
		}

		cur = nxt
	}

	*q = old
	return item, true
}
