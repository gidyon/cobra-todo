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
	"strings"

	"github.com/gidyon/ctodo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := cmd.Flags().GetString("task")
		if err != nil {
			return fmt.Errorf("getting text flag value failed: %v", err)
		}
		if text == "" {
			text = strings.Join(args, " ")
		}
		return add(cmd, text)
	},
}

func init() {
	addCmd.Flags().BoolP("done", "d", false, "Mark task as done")
	addCmd.Flags().StringP("task", "t", "", "Add the given text as task to todo")

	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, text string) error {
	d, err := cmd.Flags().GetBool("done")
	if err != nil {
		return fmt.Errorf("getting done flag value failed: %v", err)
	}

	_, err = client.Add(ctx, &todo.Task{Text: text, Done: d})
	if err != nil {
		return fmt.Errorf("could not add task in the backend: %v", err)
	}

	fmt.Println("task added successfully")
	return nil
}
