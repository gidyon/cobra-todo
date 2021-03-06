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

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a task in todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return delete(cmd)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	delCmd.Flags().Int32P("id", "i", 0, "delete task in todo with given id")
}

func delete(cmd *cobra.Command) error {
	d, err := cmd.Flags().GetInt32("id")
	if err != nil {
		return fmt.Errorf("parsing done flag failed: %v", err)
	}

	_, err = client.Delete(ctx, &todo.TaskId{TaskId: d})
	if err != nil {
		return fmt.Errorf("could not delete task in the backend: %v", err)
	}

	fmt.Printf("task with id: %v deleted successfully\n", d)
	return nil
}
