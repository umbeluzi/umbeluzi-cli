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

package internal

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewCommandOpts() *CommandOptions {
	return &CommandOptions{
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

type CommandOptions struct {
	Stdout io.Writer
	Stdin  io.Reader
	Stderr io.Writer
}

func SetupIO(cmd *cobra.Command, opts *CommandOptions) {
	if opts.Stdout != nil {
		cmd.SetOut(opts.Stdout)
	}

	if opts.Stdin != nil {
		cmd.SetIn(opts.Stdin)
	}

	if opts.Stderr != nil {
		cmd.SetErr(opts.Stderr)
	}
}
