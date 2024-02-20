/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package config

import (
	"github.com/spf13/viper"
)

type Application struct {
	Verbosity  int
	OutputFile string
}

func (app *Application) LoadAllValues(v *viper.Viper, configPath string) error {
	if err := v.Unmarshal(app); err != nil {
		return err
	}

	return nil
}

func (app *Application) GetOutputFormat() string {
	return "spdx-json@2.3"
}

func (app *Application) GetOutputFile() string {
	return app.OutputFile
}
