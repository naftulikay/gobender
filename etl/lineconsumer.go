package etl

import (
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// LineConsumer consumes a list of lines from a channel, transforms them, and sends them to the outgoing channel.
type LineConsumer struct {
	wg       sync.WaitGroup
	logger   *log.Entry
	incoming chan string
	outgoing chan string
}

// Start causes the LineConsumer to begin consuming lines from the incoming channel asynchronously.
func (lc *LineConsumer) Start() {
	go func() {
		lc.wg.Add(1)
		defer lc.wg.Done()

		for line := range lc.incoming {
			lc.wg.Add(1)
			go lc.consume(line)
		}
	}()
}

func (lc *LineConsumer) consume(line string) {
	lc.logger.Debug("consumed line: ", line)

	lc.outgoing <- lc.transform(line)

	defer lc.wg.Done()
}

func (lc *LineConsumer) transform(line string) string {
	return fmt.Sprintf("payload{%s}", line)
}

// Finish causes LineConsumer to wait for all lines to be consumed from the incoming channel.
func (lc *LineConsumer) Finish() {
	lc.wg.Wait()
}
