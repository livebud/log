package log

import (
	"fmt"
	"sync"
)

func Buffer() *Buffered {
	return &Buffered{}
}

type Buffered struct {
	mu      sync.Mutex
	Entries []*Entry
}

var _ Handler = (*Buffered)(nil)
var _ Log = (*Buffered)(nil)

func (h *Buffered) Log(entry *Entry) error {
	h.Entries = append(h.Entries, entry)
	return nil
}

func (b *Buffered) Field(key string, value interface{}) Log {
	return &sublogger{b, Fields{key: value}}
}

func (b *Buffered) Fields(fields map[string]interface{}) Log {
	return &sublogger{b, fields}
}

func (b *Buffered) log(lvl Level, args []interface{}, fields Fields) error {
	b.mu.Lock()
	b.Entries = append(b.Entries, createEntry(lvl, sprint(args...), fields))
	b.mu.Unlock()
	return nil
}

func (b *Buffered) logf(lvl Level, msg string, args []interface{}, fields Fields) error {
	b.mu.Lock()
	b.Entries = append(b.Entries, createEntry(lvl, fmt.Sprintf(msg, args...), fields))
	b.mu.Unlock()
	return nil
}

func (b *Buffered) Debug(args ...interface{}) error {
	return b.log(LevelDebug, args, nil)
}

func (b *Buffered) Debugf(msg string, args ...interface{}) error {
	return b.logf(LevelDebug, msg, args, nil)
}

func (b *Buffered) Info(args ...interface{}) error {
	return b.log(LevelInfo, args, nil)
}

func (b *Buffered) Infof(msg string, args ...interface{}) error {
	return b.logf(LevelInfo, msg, args, nil)
}

func (b *Buffered) Notice(args ...interface{}) error {
	return b.log(LevelNotice, args, nil)
}

func (b *Buffered) Noticef(msg string, args ...interface{}) error {
	return b.logf(LevelNotice, msg, args, nil)
}

func (b *Buffered) Warn(args ...interface{}) error {
	return b.log(LevelWarn, args, nil)
}

func (b *Buffered) Warnf(msg string, args ...interface{}) error {
	return b.logf(LevelWarn, msg, args, nil)
}

func (b *Buffered) Error(args ...interface{}) error {
	return b.log(LevelError, args, nil)
}

func (b *Buffered) Errorf(msg string, args ...interface{}) error {
	return b.logf(LevelError, msg, args, nil)
}
