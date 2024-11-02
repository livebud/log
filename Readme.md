# Log

[![Go Reference](https://pkg.go.dev/badge/github.com/livebud/log.svg)](https://pkg.go.dev/github.com/livebud/log)

Simple logger for Go.

⚠️ Deprecated in favor of: https://github.com/matthewmueller/logs.

![log](https://user-images.githubusercontent.com/170299/272141127-b7357640-0418-4739-9b4f-199662da4a1e.png)

## Features

- Pretty `console` handler for terminals
- Adds a level filter handler
- json and logfmt handlers
- Adds a multi-logger

## Install

```sh
go get github.com/livebud/log
```

## Example

```go
log := log.Multi(
  log.Filter(log.LevelInfo, log.Console(color.Ignore(), os.Stderr)),
  log.Json(os.Stderr),
)
log.Debug("world", "args", 10)
log.Info("hello", "planet", "world", "args", 10)
log.Warn("hello", "planet", "world", "args", 10)
log.Error("hello world", "planet", "world", "args", 10)
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT
