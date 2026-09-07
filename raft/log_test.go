package raft

import "testing"

func entries(termByIndex ...Term) []LogEntry {
	out := make([]LogEntry, 0, len(termByIndex))
	for i, t := range termByIndex {
		out = append(out, LogEntry{Term: t, Index: Index(i + 1)})
	}
	return out
}

func TestIsUpToDate(t *testing.T) {
	l := NewLogFrom(entries(1, 1, 2))

	cases := []struct {
		name         string
		candIdx      Index
		candTerm     Term
		wantUpToDate bool
	}{
		{"Identical logs", 3, 2, true},
		{"higher last term wins even if shorter", 1, 3, true},
		{"lower last term loses even if longer", 9, 1, false},
		{"same term, longer wins", 4, 2, true},
		{"same term, shorter loses", 2, 2, false},
		{"empty candidate log", 0, 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := l.IsUpToDate(c.candIdx, c.candTerm); got != c.wantUpToDate {
				t.Fatalf("IsUpToDate(%d, %d) = %v, want %v", c.candIdx, c.candTerm, got, c.wantUpToDate)
			}
		})
	}
}

func TestTermAtDistinguishesAbsentFromZero(t *testing.T) {
	l := NewLogFrom(entries(1, 1))

	if term, ok := l.TermAt(0); !ok || term != 0 {
		t.Fatalf("sentinel: got (%d, %v), want (0, true)", term, ok)
	}
	if _, ok := l.TermAt(3); ok {
		t.Fatal("index past end reported as present")
	}

	if l.Matches(3, 0) {
		t.Fatal("matches accepted as index past end of the log")
	}
}

func TestSliceDoesNotAliasLog(t *testing.T) {
	l := NewLogFrom(entries(1, 1, 1))
	sent := l.Slice(1)
	if len(sent) != 3 {
		t.Fatalf("Slice(1) returned %d entries, want 3", len(sent))
	}
	sent[0].Term = 99
	if got, _ := l.TermAt(1); got == 99 {
		t.Fatal("Slice aliases the log's backing array")
	}
}

func TestTruncateAndAppend(t *testing.T) {
	t.Skip(" -_- ")
	l := NewLogFrom(entries(1, 1, 2))
	l.TruncateFrom(3)
	l.Append(LogEntry{Term: 3, Index: 3})
	if l.LastTerm() != 3 || l.LastIndex() != 3 {
		t.Fatalf("after truncate+append: last = (%d, %d), want (3,3)", l.LastIndex(), l.LastTerm())
	}
}

func TestTruncateBelowCommitPanics(t *testing.T) {
	t.Skip(" -_- ")
}

func TestFindConflict(t *testing.T) {
	t.Skip(" -_- ")
}
