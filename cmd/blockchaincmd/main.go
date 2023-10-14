package blockchaincmd

import (
	"github.com/spf13/cobra"
)

func StartCmd() *cobra.Command {
	var nmtri1912 = &cobra.Command{
		Use:   "nmtri1912",
		Short: "The BlockChain CLI",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	nmtri1912.AddCommand(versionCmd)
	nmtri1912.AddCommand(balancesCmd())
	nmtri1912.AddCommand(txCmd())

	return nmtri1912
}
