// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"

	"github.com/gidyon/ctodo/todo"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marks task i todo with given id as done",
	RunE: func(cmd *cobra.Command, args []string) error {
		return done()
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	doneCmd.Flags().Int32P("id", "i", 0, "marks done task with given id")
}

func done() error {
	d, err := rootCmd.Flags().GetInt32("id")
	if err != nil {
		return fmt.Errorf("parsing done flag failed: %v", err)
	}

	_, err = client.Done(ctx, &todo.TaskId{TaskId: d})
	if err != nil {
		return fmt.Errorf("could not mark task as done: %v", err)
	}

	fmt.Println("task marked as done successfully")
	return nil
}
