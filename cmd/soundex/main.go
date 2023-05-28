package main

import (
	"github.com/artemKapitonov/soundex/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	app, err := app.New()
	if err != nil {
		logrus.Fatalf("Starting fail")
	}

	if err := app.Run(); err != nil {
		logrus.Fatalf("Can't run app: %s", err.Error())
	}
}
