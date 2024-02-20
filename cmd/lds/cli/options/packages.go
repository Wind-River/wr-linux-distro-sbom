/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package options

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type PackagesOptions struct {
	OutputFormat string
	OutputFile   string
}

var _ FlagsAdder = (*PackagesOptions)(nil)

func DefaultPackagesOptions() *PackagesOptions {
	return &PackagesOptions{
		OutputFormat: "spdx-json",
		OutputFile:   "",
	}
}

func (o *PackagesOptions) AddFlags(cmd *cobra.Command, v *viper.Viper) error {
	cmd.Flags().StringVar(&o.OutputFile, "file", "", "file to write the report")

	return bindPackagesConfigOptions(cmd.Flags(), v)
}

func bindPackagesConfigOptions(flags *pflag.FlagSet, v *viper.Viper) error {
	if err := v.BindPFlag("outputfile", flags.Lookup("file")); err != nil {
		return err
	}
	return nil
}
