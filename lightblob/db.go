package lightblob

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

type Lightblob struct {
	lock sync.RWMutex

	db    *os.File
	index *os.File
	tail  atomic.Uint64
}

func (l *Lightblob) Dump() {

	buffer := make([]byte, 10)
	var file *os.File
	file, err := os.OpenFile("./data/db.lb", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = os.OpenFile("./data/db.lb", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	file.ReadAt(buffer, 0)
	fmt.Print(string(buffer))
}

func (l *Lightblob) Write(data []byte) {
	buffer := make([]byte, 100)
	var file *os.File
	file, err := os.OpenFile("./data/db.lb", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Print(l.tail.Load())
	file.WriteAt(data, int64(l.tail.Load()))
	file.ReadAt(buffer, 0)
	fmt.Print(string(buffer))

	l.tail.Add(uint64(len(data)))
	fmt.Print(l.tail.Load())
}
