# gobender [![Build Status][travis.svg]][travis]

A dummy ETL worker in Go.

The sample workflow is designed to a) list files within a bucket, b) read each file line-by-line and convert each to
JSON, and c) transform the JSON and upload it to an analytics database.

This is mainly a template showing how to build an ETL workflow in Golang, so there's no logic to actually work with
the bucket or with the analytics database, rather it demonstrates the pipeline.

## License

Licensed at your discretion under either:

 - [Apache License, Version 2.0](./LICENSE-APACHE)
 - [MIT License](./LICENSE-MIT)

 [travis]: https://travis-ci.org/naftulikay/gobender
 [travis.svg]: https://travis-ci.org/naftulikay/gobender.svg?branch=master
