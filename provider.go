package log

import "github.com/livebud/di"

func Provider(in di.Injector) {
	di.Provide[Level](in, provideLevel)
	di.Provide[*Logger](in, provideLog)
}

func provideLevel(in di.Injector) (Level, error) {
	return LevelInfo, nil
}

func provideLog(in di.Injector) (*Logger, error) {
	return Default(), nil
}
