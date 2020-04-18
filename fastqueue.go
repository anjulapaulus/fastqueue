package fastqueue

//Queue ...
type Queue struct {
	length int
	ids    []int
	values []float64
}

//Push funtion is used to insert ids and items to queue
func (q *Queue) Push(id int, value float64) {
	q.ids = append(q.ids, id)
	q.values = append(q.values, value)
	q.length++
	pos := q.length - 1

	for pos > 0 {
		parent := (pos - 1) >> 1
		parentValue := q.values[parent]
		if value >= parentValue {
			break
		}
		q.ids[pos] = q.ids[parent]
		q.values[pos] = parentValue
		pos = parent
	}
	q.ids[pos] = id
	q.values[pos] = value
}

//Pop funtion is used to remove and return top item in queue
func (q *Queue) Pop() (int, float64) {
	if q.length == 0 {
		return 0, 0
	}

	top := q.ids[0]
	topValue := q.values[0]
	q.length--

	if q.length > 0 {
		q.ids[0] = q.ids[q.length]
		q.values[0] = q.values[q.length]

		id := q.ids[0]
		value := q.values[0]
		halfLength := q.length >> 1
		pos := 0

		for pos < halfLength {
			left := (pos << 1) + 1
			right := left + 1
			bestIndex := q.ids[left]
			bestValue := q.values[left]
			rightValue := q.values[right]

			if right < q.length && rightValue < bestValue {
				left = right
				bestIndex = q.ids[right]
				bestValue = rightValue
			}
			if bestValue >= value {
				break
			}

			q.ids[pos] = bestIndex
			q.values[pos] = bestValue
			pos = left
		}
		q.ids[pos] = id
		q.values[pos] = value
	}
	return top, topValue
}

//Peek returns the top item id from queue without removal
func (q *Queue) Peek() int {
	if q.length == 0 {
		return 0
	}
	return q.ids[0]
}

//PeekValue returns the top item value from queue without removal
func (q *Queue) PeekValue() float64 {
	if q.length == 0 {
		return 0
	}
	return q.values[0]
}

//Len returns the length of the queue
func (q *Queue) Len() int {
	return q.length
}
