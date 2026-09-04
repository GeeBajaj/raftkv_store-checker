package raft

// Action is a side effect the core wants the driver to perform
// Within one b atch, every PersistState and persistEntries ction must be
// durably on disk BEFORE any Send action in the same batch leaves the process

// raft requires currentTerm, votedFor and the log to be persisted before responding
// to any RPC.

type Action interface {
	isAction()
}

type Send struct {
	Msg Message
}

type PersistState struct {
	Term     Term
	VotedFor NodeID
}

type PersistEntries struct {
	TruncateFrom Index
	Entries      []LogEntry
}

type Apply struct {
	Entries []LogEntry
}

// reports that a linearizable read barrier has been confirmed by a quorum
// the driver should wait until lastApplied >= Index,
// then serve the reads tagged with ReadID from local state
type ReadIndexReady struct {
	ReadID uint64
	Index  Index
}

func (Send) isAction()           {}
func (PersistState) isAction()   {}
func (PersistEntries) isAction() {}
func (Apply) isAction()          {}
func (ReadIndexReady) isAction() {}
