package main

import "github.com/ManasNagaraj/light-blob-db/lightblob"

func main() {
	lg := &lightblob.Lightblob{}
	lg.Dump()
	lg.Write([]byte("helloworld"))
}
