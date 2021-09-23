# gofbp 

This repo holds the beginning of an FBP implementation in Go

# Warning: Latest version crashes near the end of a run!

Features include:

- delayed start of goroutines (FBP processes), unless `MustRun` attribute is specified or the process has no non-IIP inputs (same as JavaFBP delayed start feature) 
- optional output ports - see https://github.com/jpaulm/gofbp/blob/master/components/testrtn/writetoconsole.go


Test cases as follows:

- 2 Senders, one Receiver - merging first come, first served

- 2 Senders, with outputs concatenated using ConcatStr

- stream of IPs being distributed among several Receivers using RoundRobinSender 

- file being written to console

To run them, position to your `GitHub\gofbp` directory, and do any of the following:

- `go test -run Merge -count=1`
- `go test -run Concat -count=1`
- `go test -run RRDist -count=1`
- `go test -run ShowFile -count=1`

`go test` runs them all, in sequence


Note: way too much logging - have to make that optional - use a JSON file...?  Issue raised for this...
