# daylog

`daylog` is a small highly specialised CLI app for keeping track of your day,
and then submitting your standups.

# Use:

```
# Add a task:
daylog add "Finished ticket #1573"
daylog blocked Got stuck working on ticket \#1983
daylog next "Vacation for the" "next week"

# Open it in $EDITOR:
daylog edit

# Show what is on disk:
daylog show

# output a completed version of the day
daylog compile
```

# Installation

```shell
brew tap keyneston/tap
brew install daylog
```

With go installed:

```shell
go install github.com/keyneston/daylog@latest
```

Or download the [latest release](https://github.com/keyneston/daylog/releases/) and add to your path.
