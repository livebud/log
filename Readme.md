# Log

[![Go Reference](https://pkg.go.dev/badge/github.com/livebud/log.svg)](https://pkg.go.dev/github.com/livebud/log)

Log utilities for the upcoming `log/slog` package in the Go v1.21.

![log](https://github.com/livebud/log/assets/170299/f520d535-99ab-4db6-915c-ef06af4fa831)

## Features

- Built on top of the new `log/slog` package
- Pretty `console` handler for terminals
- Adds a level filter handler
- Adds a multi-logger

## Install

```sh
go get github.com/livebud/log
```

**Note:** This package depends on `log/slog`, which is only available for Go v1.21+.

## Example

```go
log := log.Multi(
  log.Filter(log.LevelInfo, &log.Console{Writer: os.Stderr}),
  slog.NewJSONHandler(os.Stderr, nil),
)
log.WithGroup("hello").Debug("world", "args", 10)
log.Info("hello", "planet", "world", "args", 10)
log.Warn("hello", "planet", "world", "args", 10)
log.Error("hello world", "planet", "world", "args", 10)
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT
