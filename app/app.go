package app

import (
	"fmt"
	"path"
	"runtime"
	"sync"
)

type Application struct {
	ProcParams *ProcParams
	Config     *Config
}

var (
	appOnce sync.Once
	app     *Application
)

func GetApp() *Application {
	appOnce.Do(func() {
		app = new(Application)
	})

	return app
}

func (a *Application) Prepare() error {
	runtime.GOMAXPROCS(runtime.NumCPU()) //todo

	if err := a.initParams(); err != nil {
		return err
	}

	if err := a.initConfig(); err != nil {
		return err
	}

	return nil
}

func (a *Application) initConfig() error {

	config, err := LoadConfig(path.Join(a.ProcParams.ConfigPath, fmt.Sprintf("config_%s.json", a.ProcParams.Environment)))
	if err != nil {
		return err
	}
	a.Config = config

	return nil
}

func (a *Application) initParams() error {
	params, err := LoadParams()
	if err != nil {
		return err
	}

	a.ProcParams = params
	return nil
}
