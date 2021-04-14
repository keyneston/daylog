# daylog

`daylog` is a theoretical CLI app for keeping track of your day, and then
submitting your standups.

# Theory:

```
# Add a task:
daylog add -m "Finished ticket #1573"

# Open it in $EDITOR:
daylog add

# output a completed version of the day
daylog compile
```

# Decisions

Storage Engine:

* sqlite
  - neg: requires cgo
  - pos: easily searchable

* flatfile:
  - pos: pure go
  - neg: more work to parse
  - pos: would more easily map to $EDITOR entries
