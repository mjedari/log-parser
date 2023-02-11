package app

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"time"
)

const (
	BufferSize = 16 * 1024 * 1024 // 4MB
)

type LogParser struct {
	reader *bufio.Reader
}

func NewLogParser(file *LogFile) *LogParser {
	reader := bufio.NewReader(file.file)
	return &LogParser{reader: reader}
}

func (p *LogParser) Parse() {
	startTime := time.Now()
	p.parseChunk()
	endTime := time.Now()
	fmt.Println("Duration: ", endTime.Sub(startTime).String())
}

func (p *LogParser) parseChunk() {
	nBytes, nChunks := int64(0), int64(0)
	buf := make([]byte, 0, BufferSize)
	for {
		n, err := p.reader.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))
		// process buf
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}
