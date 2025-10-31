package ringbuffer

type Ring[T any] struct {
	data []T

	// next position to write
	head int

	// oldest element
	tail int

	// current number of elements
	size int

	// possible element number
	capacity int
}

// Creates a ring buffer with the given capacity.
// Panics if capacity is less or equal to 0.
func New[T any](capacity int) *Ring[T] {
	if capacity <= 0 {
		panic("ringbuffer: capacity must be greater than 0")
	}
	return &Ring[T]{
		data:     make([]T, capacity),
		capacity: capacity,
	}
}

// Put a new value into the buffer.
// If the buffer is full, the oldest element is overwritten.
// We use modulo operations `%` to wrap around.
func (r *Ring[T]) Write(value T) {
	r.data[r.head] = value
	r.head = (r.head + 1) % r.capacity
	if r.size == r.capacity {
		r.tail = (r.tail + 1) % r.capacity
	} else {
		r.size++
	}
}

// Get the oldest element, and remove it.
// The second return value is `false` if the buffer is empty.
func (r *Ring[T]) Read() (T, bool) {
	var zero T
	if r.size == 0 {
		return zero, false
	}
	v := r.data[r.tail]
	r.tail = (r.tail + 1) % r.capacity
	r.size--
	return v, true
}

// Get the oldest element.
// The second return value is `false` if the buffer is empty.
func (r *Ring[T]) Peek() (T, bool) {
	var zero T
	if r.size == 0 {
		return zero, false
	}
	return r.data[r.tail], true
}

// Get the current contents in logical order (oldest to newest).
func (r *Ring[T]) Snapshot() []T {
	if r.size == 0 {
		return nil
	}
	out := make([]T, r.size)
	for i := range r.size {
		j := (r.tail + i) % r.capacity
		out[i] = r.data[j]
	}
	return out
}

// Get the number of elements currently stored.
func (r *Ring[T]) Len() int {
	return r.size
}

// Get the total capacity.
func (r *Ring[T]) Cap() int {
	return r.capacity
}

// Clear the buffer but keep the capacity.
func (r *Ring[T]) Reset() {
	r.head = 0
	r.tail = 0
	r.size = 0
}
