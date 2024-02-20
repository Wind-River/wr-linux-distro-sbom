/*
 * Copyright (C) 2023 Wind River Systems, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package options

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type FlagsAdder interface {
	AddFlags(cmd *cobra.Command, v *viper.Viper) error
}
