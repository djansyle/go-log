package log

import (
	l "log"
	"os"
	"strings"
	"sync"
	"time"
)

// Log holds all the cloneLogger information
type Log struct {
	sync.Mutex
	tags []string
}

var logger *l.Logger

func init() {
	logger = l.New(os.Stdout, "", 0)
}


func cloneLogger(parent Log) Log {
	var cloned []string
	copy(cloned, parent.tags)
	return Log{ tags: cloned }
}

func showLog(level string, tags []string, message interface{}) {
	logger.Printf("%s %s:%s %v", time.Now().Format(time.RFC3339), level, strings.Join(tags, ","), message)
}

// NewLogger creates a new instance of the Log
func NewLogger(tags []string) *Log {
	return &Log{tags: tags}
}

func (l *Log) Tags(tags []string) Log {
	t := cloneLogger(*l)
	for _, v := range tags {
		t.tags = append(l.tags, v)
	}

	return t
}

// Info show an info message
func (l *Log) Info(message interface{}) {
	showLog("info", l.tags, message)
}

// Error show an error message
func (l *Log) Error(message interface{}) {
	showLog("error", l.tags, message)
}

// Warn show a warn message
func (l *Log) Warn(message interface{}) {
	showLog("warn", l.tags, message)
}

// Verbose show a verbose message
func (l *Log) Verbose(message interface{}) {
	showLog("verbose", l.tags, message)
}

// Silly show a silly message
func (l *Log) Silly(message interface{}) {
	showLog("silly", l.tags, message)
}

// Panic show a message log and panics
func (l *Log) Panic(message interface{}) {
	showLog("panic", l.tags, message)
	panic(message)
}