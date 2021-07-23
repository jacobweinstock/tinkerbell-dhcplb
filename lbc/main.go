package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"lbc/cmd"

	"github.com/joncrlsn/dque"
)

func main() {
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	defer done()

	if err := cmd.Execute(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
	}
}

type receiver struct {
	c chan string
}

type FIFO struct {
	que []string
	mu  sync.Mutex
}

type Item struct {
	Name string
	Id   int
}

func ItemBuilder() interface{} {
	return &Item{}
}

func main2() {
	fmt.Println("Don't communicate by sharing memory, share memory by communicating.")
	var f FIFO
	for i := 1; i < 11; i++ {
		i := i
		go f.enqueue(strconv.Itoa(i))
		time.Sleep(time.Millisecond)
	}

	f.enqueue("a")
	f.enqueue("b")
	f.enqueue("c")
	f.enqueue("d")
	f.enqueue("e")

	fmt.Println("full queue", f.queue())
	fmt.Println("full queue", f.que)
	fmt.Printf("dequeue: [%v]\n", f.dequeue())
	fmt.Printf("dequeue: [%v]\n", f.dequeue())
	fmt.Printf("dequeue: [%v]\n", f.dequeue())

	fmt.Println("full queue", f.queue())

	dir, err := ioutil.TempDir("", "lbc")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up
	qName := "item-queue"
	qDir := dir
	segmentSize := 5000

	q, _ := dque.NewOrOpen(qName, qDir, segmentSize, ItemBuilder)

	for i := 1; i < 11; i++ {
		i := i
		go q.Enqueue(&Item{fmt.Sprintf("Joe-%v", i), i})
		time.Sleep(time.Millisecond)
	}

	for i := 1; i < 11; i++ {
		v, err := q.Dequeue()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("dequeue: [%v]\n", v)
		time.Sleep(time.Millisecond)
	}
}

func (r receiver) asyncReceiver(msg string) {
	r.c <- msg
}

func (f *FIFO) queue() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.que
}

func (f *FIFO) enqueue(e string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.que = append(f.que, e)
}

func (f *FIFO) dequeue() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.que) > 0 {
		v := f.que[0]
		f.que = f.que[1:]
		return v
	}
	return ""
}
