package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sayhelloCmd = &cobra.Command{
	Use:   "sayhello",
	Short: "return hello world",
	Long: `Return hello world ou remplace world par le nom passé dans le flag name.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _:= cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}
		fmt.Println("Hello "+name)
	},
}

func init() {
	sayCmd.AddCommand(sayhelloCmd)
	sayhelloCmd.Flags().StringP("name", "n", "", "Set your name")
}
