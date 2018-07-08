package etl

import (
	"sync"

	log "github.com/Sirupsen/logrus"
)

// ExportConsumer consumes a list of transformed lines
type ExportConsumer struct {
	wg       sync.WaitGroup
	logger   *log.Entry
	incoming chan string
}

// Start causes the ExportConsumer to begin consuming payloads from the incoming channel asynchronously.
func (ec *ExportConsumer) Start() {
	go func() {
		ec.wg.Add(1)
		defer ec.wg.Done()

		for payload := range ec.incoming {
			ec.wg.Add(1)
			go ec.consume(payload)
		}
	}()
}

func (ec *ExportConsumer) consume(payload string) {
	defer ec.wg.Done()

	ec.logger.Debug("consumed payload: ", payload)
}

// Finish causes the ExportConsumer to wait for all items to be consumed and processed from the incoming channel.
func (ec *ExportConsumer) Finish() {
	ec.wg.Wait()
}
