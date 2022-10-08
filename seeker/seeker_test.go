package seeker

import (
	"testing"
)

func TestNew(t *testing.T) {
	str := `hour
minute
second
day
week
month
year`

	lines := New(str)
	if lines.Length() != 7 {
		t.FailNow()
	}
}

func TestInit(t *testing.T) {
	str := `hour
minute
second
day
week
month
year`

	lines := Lines{}
	lines.Init(str)

	if lines.Length() != 7 {
		t.FailNow()
	}
}

func TestSeekToStringAndCurrentLine(t *testing.T) {
	str := `hour
minute
second
day
week
month
year`

	lines := Lines{}
	lines.Init(str)

	lines.RestCursor()
	lines.SeekToString(`week`)
	if lines.Array[lines.Cursor] != `week` {
		t.FailNow()
	}

	if lines.CurrentLine() != `week` {
		t.FailNow()
	}
}

func TestSeekToContainsWithBreak(t *testing.T) {
	str := `hour
minute
second
day
week
month
year`

	lines := Lines{}
	lines.Init(str)

	lines.RestCursor()
	if !lines.SeekToContainsWithBreak("cond", `mon`) {
		t.FailNow()
	}

	lines.RestCursor()
	if lines.SeekToContainsWithBreak("ear", `mon`) {
		t.FailNow()
	}
}

func TestSeekToStringWithBreak(t *testing.T) {
	str := `hour
minute
second
day
week
month
year`

	lines := Lines{}
	lines.Init(str)

	lines.RestCursor()
	if !lines.SeekToStringWithBreak(`day`, `month`) {
		t.FailNow()
	}

	lines.RestCursor()
	if lines.SeekToStringWithBreak(`year`, `month`) {
		t.FailNow()
	}
}
