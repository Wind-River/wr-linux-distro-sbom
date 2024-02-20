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

type RootOptions struct {
	Config  string
	Verbose int
}

var _ FlagsAdder = (*RootOptions)(nil)

func DefaultRootOptions() *RootOptions {
	return &RootOptions{}
}

func (o *RootOptions) AddFlags(cmd *cobra.Command, v *viper.Viper) error {
	cmd.PersistentFlags().CountVarP(&o.Verbose, "verbose", "v", "increase verbosity (-v = info, -vv = debug)")

	return bindRootConfigOptions(cmd.PersistentFlags(), v)
}

func bindRootConfigOptions(flags *pflag.FlagSet, v *viper.Viper) error {
	if err := v.BindPFlag("verbosity", flags.Lookup("verbose")); err != nil {
		return err
	}
	return nil
}
