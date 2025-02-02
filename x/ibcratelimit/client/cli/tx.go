package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"

	"github.com/provenance-io/provenance/internal/provcli"
	"github.com/provenance-io/provenance/x/ibcratelimit"
)

// NewTxCmd is the top-level command for oracle CLI transactions.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        ibcratelimit.ModuleName,
		Aliases:                    []string{"rl"},
		Short:                      "Transaction commands for the ibcratelimit module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdParamsUpdate(),
	)

	return txCmd
}

// GetCmdParamsUpdate is a command to update the params of the module's rate limiter.
func GetCmdParamsUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-params <address>",
		Short:   "Update the module's params",
		Long:    "Submit an update params via governance proposal along with an initial deposit.",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"u"},
		Example: fmt.Sprintf(`%[1]s tx ratelimitedibc update-params pb1skjwj5whet0lpe65qaq4rpq03hjxlwd9nf39lk --deposit 50000nhash`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			flagSet := cmd.Flags()
			authority := provcli.GetAuthority(flagSet)
			msg := ibcratelimit.NewMsgGovUpdateParamsRequest(authority, args[0])
			return provcli.GenerateOrBroadcastTxCLIAsGovProp(clientCtx, flagSet, msg)
		},
	}

	govcli.AddGovPropFlagsToCmd(cmd)
	provcli.AddAuthorityFlagToCmd(cmd)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
