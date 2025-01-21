/*
Copyright © 2025 raghavyuva <raghavyuva@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fromStringCmd represents the fromString command
var fromStringCmd = &cobra.Command{
	Use:   "fromString",
	Short: "converts string to json",
	Long: `fromString converts string to json and prints out the result. 
Example: json-parser fromString '{"name": "Tony Stark", "age": 22}'
		will print {"name": "Tony Stark", "age": 22}
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fromString called with args:", args)
	},
}

func init() {
	rootCmd.AddCommand(fromStringCmd)
}
