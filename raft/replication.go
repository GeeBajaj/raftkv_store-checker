package raft

// leader behaviour and both directions of AppendEntries

// sends an AppendEntries to every peer, carrying entries from
// nextIndex[peer] onward (capped by MaxEntriesPerAppend).
// empty entry slice is a hearbeat
func (r *Raft) broadcastAppend(nowMS int64) {
	panic("todo")
}

// implements the follower side

/*
Order of ops:
 1. reset election timer, on any valid current-term appendEntries including a rejected one.
    The msg proves a live leader. Only resetting on success means a follower's log can diverge, time out
    and start unnecessary elections
 2. consistency check PrevLogIndex/PrevLogTerm, on failure reply with conflict hint and stop.
 3. reconcile entries. Truncate only at the first index where term differs
 4. commitIndex = min(LeaderCommit, index of last new entry). Taking leader commit directly
    can cause commits past the end of the follower's log
*/
func (r *Raft) handleAppendEntries(nowMS int64, from NodeID, req *AppendEntriesReq) {
	panic("todo")
}

// update nextIndex, matchIndex

/*
on success: matchIndex[peer] = resp.MatchIndex (taken from the reply, so out of order replices cannot drop it backwards)
nextIndex = matchIndex + 1, then maybeAdvanceCommit

on failure: use ConflictTerm/ConflicEntries to back nextIndex up by a term.
If the leader has ConflictTerm in its own log, jump to the last index of that term; otherwise jump to ConflictIndex

Guard the whole thing on r.role == Leader and resp.Term == r.currentTerm.

nextIndex can not be below 1
*/
func (r *Raft) handleAppendEntriesResp(nowMS int64, from NodeID, resp *AppendEntriesResp) {
	panic("todo")
}
