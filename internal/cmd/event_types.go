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

func NewCmdEventType(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "event-type",
		Short:   "Manage event types",
		Aliases: []string{"event-types"},
	}

	cmd.AddCommand(NewCmdEventTypeList(opts))
	cmd.AddCommand(NewCmdEventTypeDelete(opts))
	cmd.AddCommand(NewCmdEventTypeCreate(opts))
	cmd.AddCommand(NewCmdEventTypeGet(opts))

	return cmd
}

func NewCmdEventTypeList(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List event types",
		Example: heredoc.Doc(`
			xibugo event-type list
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

func NewCmdEventTypeDelete(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an event type type",
		Args:  cobra.NoArgs,
		Example: heredoc.Doc(`
			xibugo event-type delete 123
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

func NewCmdEventTypeCreate(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an event type type",
		Example: heredoc.Doc(`
			xibugo event-type create --webhook 123
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

func NewCmdEventTypeGet(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve an event type type",
		Example: heredoc.Doc(`
			xibugo event-type get 123
			xibugo event-type get 123
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
