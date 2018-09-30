# Go Simple Logger

GoLang library for leveled logging.

Supports colors (shell colors) for each level.

## Levels
- Info
- Debug
- Trace
- Warning
- Error
- Critical

# Installation
```sh
$ go get github.com/rowdyroad/go-simple-logger
```

# Example
```go
package main

import log "github.com/rowdyroad/go-simple-logger"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortlevel | log.Lcolor)
	log.Info("info")
	log.Debug("info")
	log.Trace("info")
	log.Warn("info")
	log.Error("info")
	log.Crit("info")
}

```

