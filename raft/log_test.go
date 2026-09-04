package raft

func entries(termsByIndex ...Term) []LogEntry {
	out := make([]LogEntry, 0, len(termsByIndex))
	for i, t := range termsByIndex {
		out = append(out, LogEntry{Term: t, Index: Index(i+1)})
	}
	return out
}
