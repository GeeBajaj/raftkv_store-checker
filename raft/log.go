package raft

import "fmt"

/*
Log is the replicated log
INDEXING is 1-based, with a sentinel entry at slot 0 holding {Term: 0, Index: 0}

This is a bit awkward but the paper (fig 2 and section 5.3) are written
in terms of 1-based indices, and PrevLogIndex=0 is the encoding of
"the log is empty, there is nothing before this"

If the code is 0-based, i have to translate indices in my head every time
i compare against the spec. Every conversion between a log index and a slice offset happens
here and only here.
*/

type Log struct {
	entries []LogEntry
}

// entries[0] is the sen
// entries[i].Index == Index(i)
func NewLog() *Log {
	return &Log{entries: []LogEntry{{Term: 0, Index: 0}}}
}

// NewLogFrom rebuilds a log after a restarted.
func NewLogFrom(persisted []LogEntry) *Log {
	l := NewLog()
	l.entries = append(l.entries, persisted...)
	return l
}

func (l *Log) LastIndex() Index { return l.entries[len(l.entries)-1].Index }

func (l *Log) LastTerm() Term { return l.entries[len(l.entries)-1].Term }

// TermAt returns the term at the entry at i, or (0, false) if i
// is beyond the end of the log
func (l *Log) TermAt(i Index) (Term, bool) {
	if i > l.LastIndex() {
		return 0, false
	}
	return l.entries[i].Term, true
}

func (l *Log) At(i Index) LogEntry { return l.entries[i] }

// Returns entries [from, LastIndex()], sharing no backing array with the log
func (l *Log) Slice(from Index) []LogEntry {
	if from > l.LastIndex() {
		return nil
	}
	out := make([]LogEntry, 0, l.LastIndex()-from+1)
	return append(out, l.entries[from:]...)
}

// Implements the AppendEntries consistency check:
// does this log contain an entry at prevIndex whose term is prevTerm
func (l *Log) Matches(prevIndex Index, prevTerm Term) bool {
	t, ok := l.TermAt(prevIndex)
	return ok && t == prevTerm
}

// implements: is the candidate's log at least as up to date as ours?
// compares last term, then length
func (l *Log) IsUpToDate(candLastIndex Index, candLastTerm Term) bool {
	l_term := l.LastTerm()
	if candLastTerm != l_term {
		return candLastTerm > l_term
	}
	return candLastIndex >= l.LastIndex()
}

// Adds entries to the end
// Callers must have already resolved conflicts via TruncateFrom
func (l *Log) Append(entries ...LogEntry) {
	// TODO: assign Index if zero, or verify contiguity. and assert
	panic("not implemented")
}

// Deletes all entries from entry i to after
// The caller must never truncate at or below commitIndex.
// Delete it is a linearizability violation that the checker should find
// Assert instead of trusting the calls
func (l *Log) TruncateFrom(i Index) {
	panic("not implemented")
}

// produce ConflictIndex/ConflictTerm hint for a rejected AppendEntries
// so the leader can back up by a term per round trip
// instead of per entry.
// if log is too short: return  (lastindex + 1, 0)
// if entry at prevIndex with wrong term, return the first index of that wrong term and the term
func (l *Log) FindConflict(prevIndex Index) (Index, Term) {
	panic("not implemented")
}

// return last index whose entry has the given term
// used by the leader to interpret a conflic hint.
// ret 0, false if absent
func (l *Log) LastIndexOfTerm(t Term) (Index, bool) {
	panic("not implemented")
}

// Check that position == index
// any off by one bugs should be caught herew
func (l *Log) assert() {
	for i, e := range l.entries {
		if e.Index != Index(i) {
			panic(fmt.Sprintf("log corrupt: entries[%d].Index == %d", i, e.Index))
		}
	}
}
