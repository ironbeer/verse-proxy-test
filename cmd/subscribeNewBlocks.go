package cmd

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var subscribeNewBlocksCmd = &cobra.Command{
	Use: "subscribe-new-blocks",
	Run: func(cmd *cobra.Command, args []string) {
		rpc, err := cmd.Flags().GetString("rpc")
		if err != nil {
			panic(err)
		}

		client, err := ethclient.Dial(rpc)
		if err != nil {
			panic(err)
		}

		headers := make(chan *types.Header)
		sub, err := client.SubscribeNewHead(context.Background(), headers)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()

		for h := range headers {
			fmt.Printf("number=%s hash=%s\n", h.Number, h.Hash())
		}
	},
}

func init() {
	rootCmd.AddCommand(subscribeNewBlocksCmd)
}
