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
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/umbeluzi/umbeluzi-cli/internal"
	"github.com/umbeluzi/umbeluzi-cli/internal/config"
)

const (
	formatJSON  = "json"
	formatYAML  = "yaml"
	formatTable = "table"
	formatText  = "text"
)

func NewCmdWebhook(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "webhook",
		Short:   "Manage webhooks",
		Aliases: []string{"webhooks"},
	}

	cmd.AddCommand(NewCmdWebhookList(opts))
	cmd.AddCommand(NewCmdWebhookDelete(opts))
	cmd.AddCommand(NewCmdWebhookCreate(opts))
	cmd.AddCommand(NewCmdWebhookGet(opts))

	return cmd
}

func NewCmdWebhookList(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List webhooks",
		Example: heredoc.Doc(`
			xibugo webhook list
		`),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
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

func NewCmdWebhookDelete(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a webhook",
		Args:  cobra.NoArgs,
		Example: heredoc.Doc(`
			xibugo webhook delete 123
		`),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
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

func NewCmdWebhookCreate(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a webhook",
		Example: heredoc.Doc(`
			xibugo webhook create
		`),
		Args: cobra.MaximumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
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

func NewCmdWebhookGet(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve a webhook",
		Example: heredoc.Doc(`
			xibugo webhook get 123
		`),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
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
