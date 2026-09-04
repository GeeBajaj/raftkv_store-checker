package raft

import "math/rand"

// Hold the tunables
type Config struct {
	// bounds for the randomized election timeoout
	// broadtcastTime << electionTimeout << Mean time b/w failures
	ElectionTimeoutMin int64
	ElectionTimeoutMax int64
	HeartbeatInterval  int64

	// caps how uch a header ships in one msg.
	// without a cap,a follower that was partitioned for a long time triggers
	// one giant msg that stalls the event loop
	MaxEntriesPerAppend int
}

// consensus core. Not safe for concurrent use: the driver owns it from a single goroutine
// so there is no lock to hold across an RPC and no way to commit deadlock
type Raft struct {
	id    NodeID
	peers []NodeID
	cfg   Config

	// no package-level source so that a seed fully determines a run
	rng *rand.Rand

	// ----- persistent state --------------------
	// every fiend in this group must be covered by a PersistState or PersistEntries
	// action becfore any message that depends on it is sent
	currentTerm Term
	votedFor    NodeID
	log         *Log

	// ----- volatile state, leaders only --------
	role        Role
	commitIndex Index
	lastApplied Index
	leaderID    NodeID

	// ----- timers ------------------------------
	electionDeadline  int64
	heartbeatDeadline int64

	// tracks the current election
	// a set so that a duplicated RequestVoteResp (which the simulator will test)
	// cannot be double counted into a fake majority
	votesGranted map[NodeID]bool

	// accumulates actions during one event
	pending []Action
}

func New(id NodeID, peers []NodeID, log *Log, hardTerm Term, hardVote NodeID, cfg Config, rng *rand.Rand) *Raft {
	panic("not implemented")
}

// --------- three entry points, each returns the actions the driver must perform -------------

// advances logical time. nowMS from the driver (lets sim run a ten-min scenario in 50ms wall clock)
func (r *Raft) Tick(nowMS int64) []Action {
	panic("not implemented")
}

// process one inbound message
func (r *Raft) Step(nowMS int64, m Message) []Action {
	// guard 1 - stale term : drop anything from an older term
	// guard 2 - higher termm: step down to follower and adrop the term before
	// looking at what kind of msg this is
	panic("not implemented")
}

// append a command to the leader's log.
// return the index the caller should watch for, and false if the node is not the leader
// Because it returns an index, not a result, the core doesnt know whether the entry will commit:
// a leader can be deposted with uncommited entries that are discarded. m
func (r *Raft) Propose(nowMS int64, data []byte) (Index, bool) {
	panic("not implemented")
}

// starts a linearizable read barrier: records commitIndex and pings a quorum
// to confirm this node is still leader.
// Emits ReadIndexReady when confirmed
func (r *Raft) ReadIndex(nowMS int64, ReadID uint64) bool {
	panic("not implemented")
}

// to do, add commit rules, helpers
