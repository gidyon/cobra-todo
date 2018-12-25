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
		return list(cmd)
	},
}

func init() {
	listCmd.Flags().BoolP("done", "d", false, "List tasks in todo that are done")
	listCmd.Flags().BoolP("pending", "p", false, "List tasks in todo that are not done")
	listCmd.Flags().BoolP("all", "a", false, "List all details of a task in todo")

	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command) error {
	l, err := client.List(ctx, &todo.Void{})
	if err != nil {
		return fmt.Errorf("could not fetch tasks: %v", err)
	}

	d, err := cmd.Flags().GetBool("done")
	if err != nil {
		return fmt.Errorf("parsing done flag failed: %v", err)
	}
	p, err := cmd.Flags().GetBool("pending")
	if err != nil {
		return fmt.Errorf("parsing pending flag failed: %v", err)
	}
	a, err := cmd.Flags().GetBool("all")
	if err != nil {
		return fmt.Errorf("parsing all flag failed: %v", err)
	}

	for _, task := range l.Tasks {
		print := func(a bool) {
			printTask := func() {
				if a {
					fmt.Printf(" task: %q, id: %d, done: %v\n", task.Text, task.TaskId, task.Done)
					return
				}
				fmt.Printf(" %s\n", task.Text)
			}
			if d {
				if task.Done {
					fmt.Printf("😄 ")
					printTask()
				}
				return
			} else if p {
				if !task.Done {
					fmt.Printf("😦 ")
					printTask()
				}
				return
			}
			if task.Done {
				fmt.Printf("😄 ")
			} else {
				fmt.Printf("😦 ")
			}
			printTask()
		}
		if p && d {
			print(a)
			continue
		}
		if d {
			print(a)
			continue
		} else if p {
			print(a)
			continue
		}
		print(a)
	}

	return nil
}
