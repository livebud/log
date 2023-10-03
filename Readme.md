# Log

[![Go Reference](https://pkg.go.dev/badge/github.com/livebud/log.svg)](https://pkg.go.dev/github.com/livebud/log)

Simple logger for Go.

![log](https://github.com/livebud/log/assets/170299/f520d535-99ab-4db6-915c-ef06af4fa831)

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
