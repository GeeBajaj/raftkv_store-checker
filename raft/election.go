package raft

// Candidate behavior and both directions of RequestVote

// Steps down to a follower status. Called on any higher term, and
// on discovering a current-term leader while a candidate
// stepping down doesnt clear the log or commitUndex - only the leader-only volatile state
// A deposed leader keeps its uncommitted entries;
// they are resolved by the next leader's consistency check, not by self-censoring
func (r *Raft) becomeFollower(nowMS int64, term Term, leader NodeID) {
	panic("todo")
}

// increments the term, votes for itself, asks for votes

// the self-vote is recorded in votesGranted And in votedFor,
// votedForm must be persisted before the RequestVote messages go out.
// (voting before persisting -> crash -> vote again)
func (r *Raft) becomeCandidate(nowMS int64) {
	panic("todo")
}

// initialize nextIndex/matchIndex and hearbeating

// nextIndex[peer] starts at log.LastIndex()+1 and matchIndex[peer] at 0.
// nextIndex is a guess to be corrected downwards by rejection, matchIndex is a
// fact to be established upward by successes. Initializing matchIndex the same way
// would let the leader commit entries a follower never received
func (r *Raft) becomeLeader(nowMS int64) {
	panic("todo")
}

// implements receiver rules
// grant the vote iff:
//   - req.Term is not stale (guaranteed by the guard in Step)
//   - votedFor is unset or already equals the candidate,
//   - candidate's log is aat least as up to date as ours
func (r *Raft) handleRequestVote(nowMS int64, from NodeID, req *RequestVoteReq) {
	panic("todo")
}

// counts a vote
// ignore reply if r.role != Candidate or resp.Term != r.currentTerm
// replies from a previous election can arrive late and counting one into the
// current election erroneously lead to a false majority
func (r *Raft) handleRequestVoteResp(nowMS int64, from NodeID, resp *RequestVoteResp) {
	panic("todo")
}
