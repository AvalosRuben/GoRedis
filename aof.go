package main

import (
	"bufio"
	"os"
	"sync"
)

type AOF struct {
	file *os.File
	rd *bufio.Reader
	mu sync.Mutex
}