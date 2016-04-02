// Copyright 2016 The JoeFriday authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package joefriday gets facts.
package joefriday

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// TODO: make current/better implementation
type Error struct {
	Type string
	Op   string
	Err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %q: %s", e.Type, e.Op, e.Err)
}

// A Proc holds everything related to a proc file and
type Proc struct {
	sync.Mutex // protects the following fields
	*os.File
	Buf  *bufio.Reader
	Line []byte // current line
	Val  []byte
}

// Creats a Proc using the file handle.
func NewProc(f *os.File) Proc {
	return Proc{File: f, Buf: bufio.NewReader(f), Val: make([]byte, 0, 32)}
}

// ProfileSerializer is implemented by any profiler that has a Get method
// that returns the data as serialized bytes.
type ProfileSerializer interface {
	Get() ([]byte, error)
}

// ProfileSerializerTicker is implemented by any profiler that defines a
// Ticker method that uses a ticker to get the current data.  It is used
// for ongoing monitoring of something.  The current data is sent to the
// out channel as serialized bytes.
type ProfileSerializerTicker interface {
	Ticker(tick time.Duration, out chan []byte, done chan struct{}, errs chan error)
}

func (p *Proc) Reset() error {
	p.Lock()
	defer p.Unlock()
	_, err := p.File.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}
	p.Buf.Reset(p.File)
	p.Val = p.Val[:0]
	return nil
}

// Column returns a right justified string of width w.
// TODO: replace with text/tabwriter
func Column(w int, s string) string {
	pad := w - len(s)
	padding := make([]byte, pad)
	for i := 0; i < pad; i++ {
		padding[i] = 0x20
	}
	return fmt.Sprintf("%s%s", string(padding), s)
}

// Int64Column takes an int64 and returns a right justified string of width w.
func Int64Column(w int, v int64) string {
	s := strconv.FormatInt(v, 10)
	return Column(w, s)
}
