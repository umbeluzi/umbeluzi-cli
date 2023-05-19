// Copyright 2023 Edson Michaque
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/umbeluzi/umbeluzi-cli/internal"
	"github.com/umbeluzi/umbeluzi-cli/internal/config"
)

func NewCmdWhoami(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Check identity",
		RunE: func(cmd *cobra.Command, args []string) error {
			internal.SetupIO(cmd, opts)

			cfg, err := config.New()
			if err != nil {
				return err
			}

			cmd.Println(cfg)

			return nil
		},
	}

	return cmd
}
