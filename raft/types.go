// Package raft implements the Raft consensus algorithm as a pure state machine
// no I/O, doesnt read clock, no network interaction, no disk access or goroutines
// Callers feed it events () and it returns a batch of Actions describing the side effects
// it wants performed
package raft

// Term, Index, NodeID aliases
// With distinct types instead of uint64 for all,
// we get compile time errors instead of silent
// bugs when a Term is passed instead of an Index etc.
type Term uint64
type Index uint64
type NodeID string

// NoNode is the zero value for "no vote cast" / "no known leader"
const NoNode NodeID = ""

type Role uint8

const (
	Follower Role = iota
	Candidate
	Leader
)

func (r Role) String() string {
	switch r {
	case Follower:
		return "follower"
	case Candidate:
		return "candidate"
	case Leader:
		return "leader"
	}
	return "unknown"
}

// One slot in the replicated log.
// Data is []byte and opaque to this package
// Consensus layer mustn know what a command menas:
type LogEntry struct {
	Term  Term   `json:"term"`
	Index Index  `json:"index"`
	Data  []byte `json:"data,omitempty"`
}

// --------------------
// Messages
//
// This package models the two RPCs specified in the paper
// as four one-way messages instead.
// Request/response as usual would imply a caller blocking on a reply,
// which doesnt go with no I/O core
// With one-way sends, msgs are independent event the simulator can drop,
// delay, dupe or reorder.
type PayLoad interface {
	isPayload()
}

type Message struct {
	From    NodeID  `json:"from"`
	To      NodeID  `json:"to"`
	PayLoad PayLoad `json:"payload`
}

type RequestVoteReq struct {
	Term         Term   `json:"term"`
	CandidateID  NodeID `json:"candidate_id"`
	LastLogIndex Index  `json:"last_log_index"`
}

type RequestVoteResp struct {
	Term        Term `json:"term"`
	VoteGranted bool `json:"vote_granted"`
}

type AppendEntriesReq struct {
	Term         Term       `json:"term"`
	LeaderID     NodeID     `json:"leader_id"`
	PrevLogIndex Index      `json:"prev_log_index"`
	PrevLogTerm  Term       `json:"prev_log_term"`
	Entries      []LogEntry `json:"entries,omitempty"`
	LeaderCommit Index      `json:"leader_commit"`
}

type AppendEntriesResp struct {
	Term    Term `json:"term"`
	Success bool `json:"success"`

	// MatchIndex is the highest index the follower now has that agrees with the leader
	// paper fig 2.
	MatchIndex Index `json:"match_index"`

	// Rejection hints
	// paper's recovery decrements nextIndex by one per round trip
	// After a long partition, the cluster looks hung to a client
	// which shows up in the failure matrix as an availability
	// bug instead of slow backup.
	// The hint lets leader skip an entire term per round trip. (5.3)
	ConflictIndex Index `json:"conflict_index"`
	ConflictTerm  Term  `json:"conflict_term"`
}

func (*RequestVoteReq) isPayload()    {}
func (*RequestVoteResp) isPayload()   {}
func (*AppendEntriesReq) isPayload()  {}
func (*AppendEntriesResp) isPayload() {}
