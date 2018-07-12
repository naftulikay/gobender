package etl

import (
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// FileNameConsumer consumes a channel of file names and produces an entry per line to the outgoing channel.
type FileNameConsumer struct {
	wg       sync.WaitGroup
	logger   *log.Entry
	incoming chan string
	outgoing chan string
}

// Start causes the FileNameConsumer to begin consuming and processing items from the incoming channel,
// asynchronously.
func (fnc *FileNameConsumer) Start() {
	go func() {
		fnc.wg.Add(1)
		defer fnc.wg.Done()

		for filename := range fnc.incoming {
			fnc.wg.Add(1)
			go fnc.consume(filename)
		}
	}()
}

func (fnc *FileNameConsumer) consume(filename string) {
	defer fnc.wg.Done()

	fnc.logger.Debug("consumed file name: ", filename)

	for i := range [32]int{} {
		fnc.outgoing <- fmt.Sprintf("%s:%d", filename, i+1)
	}
}

// Finish waits for all items to be consumed from the incoming channel.
func (fnc *FileNameConsumer) Finish() {
	fnc.wg.Wait()
}
