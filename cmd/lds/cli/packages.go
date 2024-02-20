/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package cli

import (
	"fmt"
	"log"

	"github.com/Wind-River/lds/cmd/lds/cli/options"
	"github.com/Wind-River/lds/cmd/lds/cli/packages"
	"github.com/Wind-River/lds/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Packages(v *viper.Viper, app *config.Application, ro *options.RootOptions, po *options.PackagesOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "packages SOURCE",
		Short: "Generate a package SBOM",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := app.LoadAllValues(v, ro.Config); err != nil {
				return fmt.Errorf("invalid application config: %w", err)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := packages.Run(app, args); err != nil {
				// FIXME: handle error properly
				panic(err)
			}
		},
	}

	if err := po.AddFlags(cmd, v); err != nil {
		log.Fatal(err)
	}

	return cmd
}
