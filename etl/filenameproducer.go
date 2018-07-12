package etl

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

// FileNameProducer sends a list of filenames into an outgoing channel.
type FileNameProducer struct {
	logger   *log.Entry
	outgoing chan string
}

// Produce causes the FileNameProducer to produce lines into the outgoing channel, blocking the current thread.
func (fnp *FileNameProducer) Produce() {
	for i := range [32]int{} {
		fnp.outgoing <- fmt.Sprintf("file-%d", i)
	}
}
