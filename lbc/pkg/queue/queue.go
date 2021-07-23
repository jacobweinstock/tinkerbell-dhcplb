package queue

import "sync"

type Data struct {
	Backend string
	MAC     string
}

type FIFO struct {
	que []Data
	mu  sync.Mutex
}

func (f *FIFO) Queue() []Data {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.que
}

func (f *FIFO) Enqueue(e Data) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.que = append(f.que, e)
}

func (f *FIFO) Dequeue() Data {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.que) > 0 {
		v := f.que[0]
		f.que = f.que[1:]
		return v
	}
	return Data{}
}

func (f *FIFO) Len() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.que)
}
