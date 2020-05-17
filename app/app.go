package app

import (
	"vengeful_replacer/app/config"
)

type Application struct {
	config config.Config
}

func (app *Application) Run() {
	algorithm := app.config.Algorithm
	dictionary := app.config.Dictionary
	scanner := app.config.Scanner

	algorithm.SetDictionary(dictionary)
	algorithm.SetData(scanner.GetData())
	algorithm.Run()
}

func (app Application) GetResult() string {
	return app.config.Algorithm.GetResult()
}

func New(config config.Config) Application {
	return Application{config: config}
}
