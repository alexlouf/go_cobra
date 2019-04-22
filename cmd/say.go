package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "Une commande qui répète un paramètre",
	Long: `Fonction qui répète le paramètre reçu.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("say called")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Provide item to the say command")
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)
}
