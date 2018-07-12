package etl

import (
	log "github.com/Sirupsen/logrus"
)

// Run essentiall starts the execution.
func Run() {
	shoop()
}

// Shoop will generate a list of "files" and put them into the channel.
func shoop() {
	// imma firin mah lazor
	fileNameChan := make(chan string)
	lineChan := make(chan string)
	exportChan := make(chan string)

	log.SetLevel(log.DebugLevel)
	log.Info("shoop de whoop")

	producer := FileNameProducer{
		logger:   log.WithFields(log.Fields{"logger": "FileNameProducer"}),
		outgoing: fileNameChan,
	}

	fileNameConsumer := FileNameConsumer{
		logger:   log.WithFields(log.Fields{"logger": "FileNameConsumer"}),
		incoming: fileNameChan,
		outgoing: lineChan,
	}

	lineConsumer := LineConsumer{
		logger:   log.WithFields(log.Fields{"logger": "LineConsumer"}),
		incoming: lineChan,
		outgoing: exportChan,
	}

	exportConsumer := ExportConsumer{
		logger:   log.WithFields(log.Fields{"logger": "ExportConsumer"}),
		incoming: exportChan,
	}

	fileNameConsumer.Start()
	lineConsumer.Start()
	exportConsumer.Start()

	producer.Produce()
	close(fileNameChan)

	fileNameConsumer.Finish()
	close(lineChan)

	lineConsumer.Finish()
	close(exportChan)

	exportConsumer.Finish()
}
