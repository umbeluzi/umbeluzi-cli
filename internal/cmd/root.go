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
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/umbeluzi/umbeluzi-cli/internal"
)

var (
	configFile string
	profile    string
)

const (
	flagAccount             = "account"
	flagAccessToken         = "access-token"
	flagBaseURL             = "base-url"
	flagSandbox             = "sandbox"
	flagProfile             = "profile"
	flagConfig              = "config-file"
	envPrefix               = "UMBELUZI"
	defaultProfile          = "default"
	defaultConfigFileFormat = "yaml"
)

func NewCmdRoot(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "umbeluzi",
		SilenceUsage: true,
	}

	cmd.AddCommand(NewCmdConfig(opts))
	cmd.AddCommand(NewCmdWhoami(opts))
	cmd.AddCommand(NewCmdWebhook(opts))
	cmd.AddCommand(NewCmdSubscription(opts))
	cmd.AddCommand(NewCmdEvent(opts))
	cmd.AddCommand(NewCmdEventType(opts))
	cmd.AddCommand(NewCmdVersion(opts))

	cobra.OnInitialize(lookupConfigFiles)

	cmd.PersistentFlags().String(flagBaseURL, "", "Base URL")
	cmd.PersistentFlags().String(flagAccessToken, "", "Access token")
	cmd.PersistentFlags().StringVarP(&configFile, flagConfig, "c", "", "Configuration file")
	cmd.PersistentFlags().StringVar(&profile, flagProfile, "default", "Profile")

	viper.SetEnvPrefix(envPrefix)
	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		panic(err)
	}

	return cmd
}

func lookupConfigFiles() {
	var err error

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else if configFile := os.Getenv("UMBELUZI_CONFIG_FILE"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		configHome := os.Getenv("XDG_CONFIG_HOME")
		if configHome == "" {
			configHome, err = os.UserConfigDir()
			cobra.CheckErr(err)
		}

		viper.AddConfigPath(filepath.Join(configHome, "xibugo"))
		viper.AddConfigPath("/etc/umbeluzi")
		viper.SetConfigType(defaultConfigFileFormat)

		if profile != "" {
			viper.SetConfigName(profile)
		} else if configProfile := os.Getenv("UMBELUZI_PROFILE"); configProfile != "" {
			viper.SetConfigName(configProfile)
		} else {
			viper.SetConfigName(defaultProfile)
		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Found error: ", err.Error())
		}
	}
}
