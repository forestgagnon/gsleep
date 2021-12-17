# gsleep

sleep command accepting Go duration syntax.

## Installation

```bash
go install github.com/forestgagnon/gsleep@latest
```

## Usage

```bash
# Sleep 5 minutes
gsleep 5m

# Sleep 3 hours and 20 minutes
gsleep 3h20m && say check on that thing
```
## Duration syntax

https://pkg.go.dev/time#ParseDuration

> ParseDuration parses a duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
