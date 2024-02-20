/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package cli

import (
	"os"
	"strings"

	"github.com/Wind-River/lds/cmd/lds/cli/options"
	"github.com/Wind-River/lds/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// New constructs the root command.
func New() (*cobra.Command, error) {
	app := &config.Application{}

	v := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")))

	ro := options.DefaultRootOptions()
	po := options.DefaultPackagesOptions()

	packagesCmd := Packages(v, app, ro, po)

	rootCmd := &cobra.Command{
		Use:   "lds [command]",
		Short: "Container Image Scanner",
	}

	cmds := []*cobra.Command{
		packagesCmd,
	}

	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd, err := New()
	if err != nil {
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
