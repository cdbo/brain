package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cdbo/brain/x/membership/types"
)

var _ = strconv.Itoa(0)

func CmdAddMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-member [address] [nickname] [status]",
		Short: "Broadcast message add-member",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argAddress := args[0]
             argNickname := args[1]
             argStatus := args[2]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMember(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argNickname,
				argStatus,
				
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}