# raftkv

Raft-replicated key-value store, build to be brozen by a linearizaibility checker and to have every failure
replay exactly from a seed.

- `raft/` pure state machine, no goroutines or clocks or disk or network interactions.
Takes in events, returns actions

- driver to perform actions

- simulator to run the whole cluster in-process