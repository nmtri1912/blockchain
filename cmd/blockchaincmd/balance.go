package blockchaincmd

import (
	"fmt"
	"os"

	"github.com/nmtri1912/blockchain/infra/db"
	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("cannot process")
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	balancesCmd.AddCommand(balancesListCmd)
	return balancesCmd
}

var balancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all balances.",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := db.NewStateFromDisk()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer state.Close()

		fmt.Println("Accounts balances:")
		fmt.Println("__________________")
		fmt.Println("")
		for account, balance := range state.Balances {
			fmt.Println(fmt.Sprintf("%s: %d", account, balance))
		}
	},
}
