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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks in the todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return list()
	},
}

func init() {
	listCmd.Flags().BoolP("done", "d", false, "List tasks in todo that are done")
	listCmd.Flags().BoolP("pending", "p", false, "List tasks in todo that are not done")

	rootCmd.AddCommand(listCmd)
}

func list() error {
	l, err := client.List(ctx, &todo.Void{})
	if err != nil {
		return fmt.Errorf("could not fetch tasks: %v", err)
	}

	for _, task := range l.Tasks {
		d, err := rootCmd.Flags().GetBool("done")
		if err != nil {
			return fmt.Errorf("parsing done flag failed: %v", err)
		}
		p, err := rootCmd.Flags().GetBool("pending")
		if err != nil {
			return fmt.Errorf("parsing pending flag failed: %v", err)
		}
		if d {
			if task.GetDone() {
				fmt.Printf("😄")
			}
		}
		if p {
			if !task.GetDone() {
				fmt.Printf("😦")
			}
		}
		fmt.Printf(" %s\n", task.GetText())
	}

	return nil
}
