package seeker

import "strings"

type Lines struct {
	Cursor int
	Array  []string
}

func New(str string) *Lines {
	ls := &Lines{}
	ls.InitWithSep(str, "\n")
	return ls
}

// Init from string, split by "\n"
func (ls *Lines) Init(str string) {
	ls.InitWithSep(str, "\n")
}

// InitWithSep from string, split by specified sep string
// empty lines will be ignored
// TODO config ignore empty lines or not
func (ls *Lines) InitWithSep(str, sep string) {
	ls.Cursor = 0
	for _, str := range strings.Split(str, sep) {
		str = strings.TrimSpace(str)
		if len(str) > 0 {
			ls.Array = append(ls.Array, str)
		}
	}
}

func (ls *Lines) RestCursor() {
	ls.Cursor = 0
}

func (ls *Lines) SeekToString(str string) bool {
	for {
		ls.Cursor++
		if ls.Cursor >= len(ls.Array) {
			ls.Cursor = 0
			return false
		}
		if ls.Array[ls.Cursor] == str {
			return true
		}
	}
}

func (ls *Lines) SeekToContains(contains string) bool {
	for {
		ls.Cursor++
		if ls.Cursor >= len(ls.Array) {
			ls.Cursor = 0
			return false
		}
		if strings.Contains(ls.Array[ls.Cursor], contains) {
			return true
		}
	}
}

func (ls *Lines) SeekToStringWithBreak(str, brk string) bool {
	for {
		ls.Cursor++
		if ls.Cursor >= len(ls.Array) {
			ls.Cursor = 0
			return false
		}
		if ls.Array[ls.Cursor] == str {
			return true
		}
		if len(brk) > 0 && ls.Array[ls.Cursor] == brk {
			return false
		}
	}
}

func (ls *Lines) SeekToContainsWithBreak(contains, brk string) bool {
	for {
		ls.Cursor++
		if ls.Cursor >= len(ls.Array) {
			ls.Cursor = 0
			return false
		}
		if strings.Contains(ls.Array[ls.Cursor], contains) {
			return true
		}
		if len(brk) > 0 && strings.Contains(ls.Array[ls.Cursor], brk) {
			return false
		}
	}
}

func (ls *Lines) CurrentLine() string {
	return ls.Array[ls.Cursor]
}

func (ls *Lines) Length() int {
	return len(ls.Array)
}
