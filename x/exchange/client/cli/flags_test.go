package cli_test

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/provenance-io/provenance/testutil/assertions"
	"github.com/provenance-io/provenance/x/exchange"
	"github.com/provenance-io/provenance/x/exchange/client/cli"
)

const (
	flagBool        = "bool"
	flagInt         = "int"
	flagString      = "string"
	flagStringSlice = "string-slice"
	flagUintSlice   = "uint-slice"
	flagUint32      = "uint32"
)

func TestMarkFlagsRequired(t *testing.T) {
	flagOne := "one"
	flagTwo := "two"
	flagThree := "three"
	expAnnotations := map[string][]string{
		cobra.BashCompOneRequiredFlag: {"true"},
	}

	tests := []struct {
		name     string
		names    []string
		expPanic string
	}{
		{
			name:     "no names",
			names:    []string{},
			expPanic: "",
		},
		{
			name:     "one name, exists",
			names:    []string{flagOne},
			expPanic: "",
		},
		{
			name:     "one name, not found",
			names:    []string{"nope"},
			expPanic: "error marking --nope flag required on testing: no such flag -nope",
		},
		{
			name:     "three names, first not found",
			names:    []string{"gold", flagThree, flagThree},
			expPanic: "error marking --gold flag required on testing: no such flag -gold",
		},
		{
			name:     "three names, second not found",
			names:    []string{flagOne, "missing", flagThree},
			expPanic: "error marking --missing flag required on testing: no such flag -missing",
		},
		{
			name:     "three names, third not found",
			names:    []string{flagOne, flagThree, "derp"},
			expPanic: "error marking --derp flag required on testing: no such flag -derp",
		},
		{
			name:     "three names, all exist",
			names:    []string{flagOne, flagThree, flagThree},
			expPanic: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := &cobra.Command{
				Use: "testing",
				RunE: func(cmd *cobra.Command, args []string) error {
					return errors.New("the command should not have been run")
				},
			}
			cmd.Flags().String(flagOne, "", "The one")
			cmd.Flags().Bool(flagTwo, false, "The next best")
			cmd.Flags().Int(flagThree, 0, "Bronze")

			testFunc := func() {
				cli.MarkFlagsRequired(cmd, tc.names...)
			}
			assertions.RequirePanicEquals(t, testFunc, tc.expPanic, "MarkFlagsRequired(%q)", tc.names)
			if len(tc.expPanic) > 0 {
				return
			}

			cmdFlags := cmd.Flags()

			for _, name := range tc.names {
				flag := cmdFlags.Lookup(name)
				if assert.NotNil(t, flag, "The --%s flag", name) {
					actAnnotations := flag.Annotations
					assert.Equal(t, expAnnotations, actAnnotations, "The --%s flag annotations", name)
				}
			}
		})
	}
}

func TestAddFlagsAdmin(t *testing.T) {
	expAnnotations := map[string][]string{
		mutExc: {cli.FlagAdmin + " " + cli.FlagAuthority},
		oneReq: {flags.FlagFrom + " " + cli.FlagAdmin + " " + cli.FlagAuthority},
	}

	cmd := &cobra.Command{
		Use: "testing",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("the command should not have been run")
		},
	}

	cmd.Flags().String(flags.FlagFrom, "", "The from flag")
	cli.AddFlagsAdmin(cmd)

	adminFlag := cmd.Flags().Lookup(cli.FlagAdmin)
	if assert.NotNil(t, adminFlag, "The --%s flag", cli.FlagAdmin) {
		expUsage := "The admin (defaults to --from account)"
		actUsage := adminFlag.Usage
		assert.Equal(t, expUsage, actUsage, "The --%s flag usage", cli.FlagAdmin)
		actAnnotations := adminFlag.Annotations
		assert.Equal(t, expAnnotations, actAnnotations, "The --%s flag annotations", cli.FlagAdmin)
	}

	authorityFlag := cmd.Flags().Lookup(cli.FlagAuthority)
	if assert.NotNil(t, authorityFlag, "The --%s flag", cli.FlagAuthority) {
		expUsage := "Use the governance module account for the admin"
		actUsage := authorityFlag.Usage
		assert.Equal(t, expUsage, actUsage, "The --%s flag usage", cli.FlagAuthority)
		actAnnotations := authorityFlag.Annotations
		assert.Equal(t, expAnnotations, actAnnotations, "The --%s flag annotations", cli.FlagAuthority)
	}

	flagFrom := cmd.Flags().Lookup(flags.FlagFrom)
	if assert.NotNil(t, flagFrom, "The --%s flag", flags.FlagFrom) {
		fromExpAnnotations := map[string][]string{oneReq: expAnnotations[oneReq]}
		actAnnotations := flagFrom.Annotations
		assert.Equal(t, fromExpAnnotations, actAnnotations, "The --%s flag annotations", flags.FlagFrom)
	}
}

func TestReadFlagsAdminOrFrom(t *testing.T) {
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.String(cli.FlagAdmin, "", "The admin")
		flagSet.Bool(cli.FlagAuthority, false, "Use authority")
		return flagSet
	}

	tests := []struct {
		name      string
		flagSet   func() *pflag.FlagSet
		flags     []string
		clientCtx client.Context
		expAddr   string
		expErr    string
	}{
		{
			name: "wrong admin flag type",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.Int(cli.FlagAdmin, 0, "The admin")
				flagSet.Bool(cli.FlagAuthority, false, "Use authority")
				return flagSet
			},
			expErr: "trying to get string value of flag of type int",
		},
		{
			name: "wrong authority flag type",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.String(cli.FlagAdmin, "", "The admin")
				flagSet.Int(cli.FlagAuthority, 0, "Use authority")
				return flagSet

			},
			expErr: "trying to get bool value of flag of type int",
		},
		{
			name:    "admin flag given",
			flags:   []string{"--" + cli.FlagAdmin, "theadmin"},
			expAddr: "theadmin",
		},
		{
			name:    "authority flag given",
			flags:   []string{"--" + cli.FlagAuthority},
			expAddr: cli.AuthorityAddr.String(),
		},
		{
			name:      "from address given",
			clientCtx: client.Context{FromAddress: sdk.AccAddress("FromAddress_________")},
			expAddr:   sdk.AccAddress("FromAddress_________").String(),
		},
		{
			name:   "nothing given",
			expErr: "no <admin> provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.flagSet == nil {
				tc.flagSet = goodFlagSet
			}
			flagSet := tc.flagSet()
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var addr string
			testFunc := func() {
				addr, err = cli.ReadFlagsAdminOrFrom(tc.clientCtx, flagSet)
			}
			require.NotPanics(t, testFunc, "ReadFlagsAdminOrFrom")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagsAdminOrFrom error")
			assert.Equal(t, tc.expAddr, addr, "ReadFlagsAdminOrFrom address")
		})
	}
}

func TestReadFlagsAdminOrFromOrDefault(t *testing.T) {
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.String(cli.FlagAdmin, "", "The admin")
		flagSet.Bool(cli.FlagAuthority, false, "Use authority")
		return flagSet
	}

	tests := []struct {
		name      string
		flagSet   func() *pflag.FlagSet
		flags     []string
		clientCtx client.Context
		def       string
		expAddr   string
		expErr    string
	}{
		{
			name: "wrong admin flag type",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.Int(cli.FlagAdmin, 0, "The admin")
				flagSet.Bool(cli.FlagAuthority, false, "Use authority")
				return flagSet
			},
			def:     "robin",
			expAddr: "robin",
			expErr:  "trying to get string value of flag of type int",
		},
		{
			name: "wrong authority flag type",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.String(cli.FlagAdmin, "", "The admin")
				flagSet.Int(cli.FlagAuthority, 0, "Use authority")
				return flagSet

			},
			def:     "scout",
			expAddr: "scout",
			expErr:  "trying to get bool value of flag of type int",
		},
		{
			name:    "admin flag given",
			flags:   []string{"--" + cli.FlagAdmin, "theadmin"},
			def:     "kerry",
			expAddr: "theadmin",
		},
		{
			name:    "authority flag given",
			flags:   []string{"--" + cli.FlagAuthority},
			def:     "logan",
			expAddr: cli.AuthorityAddr.String(),
		},
		{
			name:      "from address given",
			clientCtx: client.Context{FromAddress: sdk.AccAddress("FromAddress_________")},
			def:       "ezra",
			expAddr:   sdk.AccAddress("FromAddress_________").String(),
		},
		{
			name:    "only a default",
			def:     "drew",
			expAddr: "drew",
		},
		{
			name:   "nothing given",
			expErr: "no <admin> provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.flagSet == nil {
				tc.flagSet = goodFlagSet
			}
			flagSet := tc.flagSet()
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var addr string
			testFunc := func() {
				addr, err = cli.ReadFlagsAdminOrFromOrDefault(tc.clientCtx, flagSet, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagsAdminOrFromOrDefault")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagsAdminOrFromOrDefault error")
			assert.Equal(t, tc.expAddr, addr, "ReadFlagsAdminOrFromOrDefault address")
		})
	}
}

func TestReadFlagAuthority(t *testing.T) {
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.String(cli.FlagAuthority, "", "The authority")
		return flagSet
	}

	tests := []struct {
		name    string
		flagSet func() *pflag.FlagSet
		flags   []string
		expAddr string
		expErr  string
	}{
		{
			name: "wrong flag type",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.Int(cli.FlagAuthority, 0, "The authority")
				return flagSet

			},
			expAddr: cli.AuthorityAddr.String(),
			expErr:  "trying to get string value of flag of type int",
		},
		{
			name:    "provided",
			flagSet: goodFlagSet,
			flags:   []string{"--" + cli.FlagAuthority, "usemeinstead"},
			expAddr: "usemeinstead",
		},
		{
			name:    "not provided",
			flagSet: goodFlagSet,
			expAddr: cli.AuthorityAddr.String(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flagSet := tc.flagSet()
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var addr string
			testFunc := func() {
				addr, err = cli.ReadFlagAuthority(flagSet)
			}
			require.NotPanics(t, testFunc, "ReadFlagAuthority")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagAuthority error")
			assert.Equal(t, tc.expAddr, addr, "ReadFlagAuthority address")
		})
	}
}

func TestReadFlagAuthorityOrDefault(t *testing.T) {
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.String(cli.FlagAuthority, "", "The authority")
		return flagSet
	}
	badFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.Int(cli.FlagAuthority, 0, "The authority")
		return flagSet
	}

	tests := []struct {
		name    string
		flagSet func() *pflag.FlagSet
		flags   []string
		def     string
		expAddr string
		expErr  string
	}{
		{
			name:    "wrong flag type, no default",
			flagSet: badFlagSet,
			expAddr: cli.AuthorityAddr.String(),
			expErr:  "trying to get string value of flag of type int",
		},
		{
			name:    "wrong flag type, with default",
			flagSet: badFlagSet,
			def:     "thedefault",
			expAddr: "thedefault",
			expErr:  "trying to get string value of flag of type int",
		},
		{
			name:    "provided, no default",
			flags:   []string{"--" + cli.FlagAuthority, "usemeinstead"},
			expAddr: "usemeinstead",
		},
		{
			name:    "provided, with default",
			flags:   []string{"--" + cli.FlagAuthority, "usemeinstead"},
			def:     "thedefault",
			expAddr: "usemeinstead",
		},
		{
			name:    "not provided, no default",
			expAddr: cli.AuthorityAddr.String(),
		},
		{
			name:    "not provided, with default",
			def:     "thedefault",
			expAddr: "thedefault",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.flagSet == nil {
				tc.flagSet = goodFlagSet
			}
			flagSet := tc.flagSet()
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var addr string
			testFunc := func() {
				addr, err = cli.ReadFlagAuthorityOrDefault(flagSet, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagAuthorityOrDefault")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagAuthorityOrDefault error")
			assert.Equal(t, tc.expAddr, addr, "ReadFlagAuthorityOrDefault address")
		})
	}
}

func TestReadAddrFlagOrFrom(t *testing.T) {
	tests := []struct {
		testName  string
		flags     []string
		clientCtx client.Context
		name      string
		expAddr   string
		expErr    string
	}{
		{
			testName: "unknown flag",
			name:     "notsetup",
			expErr:   "flag accessed but not defined: notsetup",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "flag given",
			flags:    []string{"--" + flagString, "someaddr"},
			name:     flagString,
			expAddr:  "someaddr",
		},
		{
			testName:  "using from",
			clientCtx: client.Context{FromAddress: sdk.AccAddress("FromAddress_________")},
			name:      flagString,
			expAddr:   sdk.AccAddress("FromAddress_________").String(),
		},
		{
			testName: "not provided",
			name:     flagString,
			expErr:   "no <string> provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var addr string
			testFunc := func() {
				addr, err = cli.ReadAddrFlagOrFrom(tc.clientCtx, flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadAddrFlagOrFrom")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadAddrFlagOrFrom error")
			assert.Equal(t, tc.expAddr, addr, "ReadAddrFlagOrFrom address")
		})
	}
}

func TestReadAddrFlagOrFromOrDefault(t *testing.T) {
	tests := []struct {
		testName  string
		flags     []string
		clientCtx client.Context
		name      string
		def       string
		expAddr   string
		expErr    string
	}{
		{
			testName: "unknown flag, no default",
			name:     "notsetup",
			expErr:   "flag accessed but not defined: notsetup",
		},
		{
			testName: "unknown flag, with default",
			name:     "notsetup",
			def:      "justathing",
			expAddr:  "justathing",
			expErr:   "flag accessed but not defined: notsetup",
		},
		{
			testName: "wrong flag type, no default",
			name:     flagInt,
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "wrong flag type, with default",
			name:     flagInt,
			def:      "andanotherthing",
			expAddr:  "andanotherthing",
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "flag given",
			flags:    []string{"--" + flagString, "someaddr"},
			name:     flagString,
			def:      "other",
			expAddr:  "someaddr",
		},
		{
			testName:  "using from",
			clientCtx: client.Context{FromAddress: sdk.AccAddress("FromAddress_________")},
			name:      flagString,
			def:       "other",
			expAddr:   sdk.AccAddress("FromAddress_________").String(),
		},
		{
			testName: "not provided, with default",
			name:     flagString,
			def:      "other",
			expAddr:  "other",
		},
		{
			testName: "not provided, empty default",
			name:     flagString,
			def:      "",
			expErr:   "no <string> provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var addr string
			testFunc := func() {
				addr, err = cli.ReadAddrFlagOrFromOrDefault(tc.clientCtx, flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadAddrFlagOrFromOrDefault")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadAddrFlagOrFromOrDefault error")
			assert.Equal(t, tc.expAddr, addr, "ReadAddrFlagOrFromOrDefault address")
		})
	}
}

func TestAddFlagsEnableDisable(t *testing.T) {
	expAnnotations := map[string][]string{
		mutExc: {cli.FlagEnable + " " + cli.FlagDisable},
		oneReq: {cli.FlagEnable + " " + cli.FlagDisable},
	}

	cmd := &cobra.Command{
		Use: "testing",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("the command should not have been run")
		},
	}

	cli.AddFlagsEnableDisable(cmd, "unittest")

	enableFlag := cmd.Flags().Lookup(cli.FlagEnable)
	if assert.NotNil(t, enableFlag, "The --%s flag", cli.FlagEnable) {
		expUsage := "Set the market's unittest field to true"
		actusage := enableFlag.Usage
		assert.Equal(t, expUsage, actusage, "--%s flag usage", cli.FlagEnable)
		actAnnotations := enableFlag.Annotations
		assert.Equal(t, expAnnotations, actAnnotations, "--%s flag annotations", cli.FlagEnable)
	}

	disableFlag := cmd.Flags().Lookup(cli.FlagDisable)
	if assert.NotNil(t, disableFlag, "The --%s flag", cli.FlagDisable) {
		expUsage := "Set the market's unittest field to false"
		actusage := disableFlag.Usage
		assert.Equal(t, expUsage, actusage, "--%s flag usage", cli.FlagDisable)
		actAnnotations := disableFlag.Annotations
		assert.Equal(t, expAnnotations, actAnnotations, "--%s flag annotations", cli.FlagDisable)
	}
}

func TestReadFlagsEnableDisable(t *testing.T) {
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.Bool(cli.FlagEnable, false, "Enable")
		flagSet.Bool(cli.FlagDisable, false, "Disable")
		return flagSet
	}

	tests := []struct {
		name    string
		flags   []string
		flagSet func() *pflag.FlagSet
		exp     bool
		expErr  string
	}{
		{
			name: "cannot read enable",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.Int(cli.FlagEnable, 0, "Enable")
				flagSet.Bool(cli.FlagDisable, false, "Disable")
				return flagSet
			},
			expErr: "trying to get bool value of flag of type int",
		},
		{
			name: "cannot read disable",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.Bool(cli.FlagEnable, false, "Enable")
				flagSet.Int(cli.FlagDisable, 0, "Disable")
				return flagSet
			},
			expErr: "trying to get bool value of flag of type int",
		},
		{
			name:    "enable",
			flags:   []string{"--" + cli.FlagEnable},
			flagSet: goodFlagSet,
			exp:     true,
		},
		{
			name:    "disable",
			flags:   []string{"--" + cli.FlagDisable},
			flagSet: goodFlagSet,
			exp:     false,
		},
		{
			name:    "neither",
			flagSet: goodFlagSet,
			expErr:  "exactly one of --enable or --disable must be provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flagSet := tc.flagSet()
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var act bool
			testFunc := func() {
				act, err = cli.ReadFlagsEnableDisable(flagSet)
			}
			require.NotPanics(t, testFunc, "ReadFlagsEnableDisable")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagsEnableDisable error")
			assert.Equal(t, tc.exp, act, "ReadFlagsEnableDisable bool")
		})
	}
}

func TestAddFlagsAsksBidsBools(t *testing.T) {
	expAnnotations := map[string][]string{
		mutExc: {cli.FlagAsks + " " + cli.FlagBids},
	}

	cmd := &cobra.Command{
		Use: "testing",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("the command should not have been run")
		},
	}

	cli.AddFlagsAsksBidsBools(cmd)

	asksFlag := cmd.Flags().Lookup(cli.FlagAsks)
	if assert.NotNil(t, asksFlag, "The --%s flag", cli.FlagAsks) {
		expUsage := "Limit results to only ask orders"
		actusage := asksFlag.Usage
		assert.Equal(t, expUsage, actusage, "--%s flag usage", cli.FlagAsks)
		actAnnotations := asksFlag.Annotations
		assert.Equal(t, expAnnotations, actAnnotations, "--%s flag annotations", cli.FlagAsks)
	}

	bidsFlag := cmd.Flags().Lookup(cli.FlagBids)
	if assert.NotNil(t, bidsFlag, "The --%s flag", cli.FlagBids) {
		expUsage := "Limit results to only bid orders"
		actusage := bidsFlag.Usage
		assert.Equal(t, expUsage, actusage, "--%s flag usage", cli.FlagBids)
		actAnnotations := bidsFlag.Annotations
		assert.Equal(t, expAnnotations, actAnnotations, "--%s flag annotations", cli.FlagBids)
	}
}

func TestReadFlagsAsksBidsOpt(t *testing.T) {
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.Bool(cli.FlagAsks, false, "Asks")
		flagSet.Bool(cli.FlagBids, false, "Bids")
		return flagSet
	}

	tests := []struct {
		name    string
		flags   []string
		flagSet func() *pflag.FlagSet
		expStr  string
		expErr  string
	}{
		{
			name: "cannot read asks",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.Int(cli.FlagAsks, 0, "Asks")
				flagSet.Bool(cli.FlagBids, false, "Bids")
				return flagSet
			},
			expErr: "trying to get bool value of flag of type int",
		},
		{
			name: "cannot read bids",
			flagSet: func() *pflag.FlagSet {
				flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
				flagSet.Bool(cli.FlagAsks, false, "Asks")
				flagSet.Int(cli.FlagBids, 0, "Bids")
				return flagSet
			},
			expErr: "trying to get bool value of flag of type int",
		},
		{
			name:    "asks",
			flags:   []string{"--" + cli.FlagAsks},
			flagSet: goodFlagSet,
			expStr:  "ask",
		},
		{
			name:    "bids",
			flags:   []string{"--" + cli.FlagBids},
			flagSet: goodFlagSet,
			expStr:  "bid",
		},
		{
			name:    "neither",
			flagSet: goodFlagSet,
			expStr:  "",
			expErr:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flagSet := tc.flagSet()
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var str string
			testFunc := func() {
				str, err = cli.ReadFlagsAsksBidsOpt(flagSet)
			}
			require.NotPanics(t, testFunc, "ReadFlagsAsksBidsOpt")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagsAsksBidsOpt error")
			assert.Equal(t, tc.expStr, str, "ReadFlagsAsksBidsOpt string")
		})
	}
}

func TestReadFlagOrderOrArg(t *testing.T) {
	theFlag := cli.FlagOrder
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.Uint64(theFlag, 0, "The id")
		return flagSet
	}
	badFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.String(theFlag, "", "The id")
		return flagSet
	}

	tests := []struct {
		name    string
		flags   []string
		flagSet *pflag.FlagSet
		args    []string
		expID   uint64
		expErr  string
	}{
		{
			name:    "unknown flag",
			flagSet: pflag.NewFlagSet("", pflag.ContinueOnError),
			expErr:  "flag accessed but not defined: " + theFlag,
		},
		{
			name:    "wrong flag type",
			flagSet: badFlagSet(),
			expErr:  "trying to get uint64 value of flag of type string",
		},
		{
			name:    "both flag and arg",
			flags:   []string{"--" + theFlag, "8"},
			flagSet: goodFlagSet(),
			args:    []string{"8"},
			expErr:  "cannot provide <order id> as both an arg (\"8\") and flag (--order 8)",
		},
		{
			name:    "just flag",
			flags:   []string{"--" + theFlag, "8"},
			flagSet: goodFlagSet(),
			expID:   8,
		},
		{
			name:    "just flag zero",
			flags:   []string{"--" + theFlag, "0"},
			flagSet: goodFlagSet(),
			expErr:  "no <order id> provided",
		},
		{
			name:    "just arg, bad",
			flagSet: goodFlagSet(),
			args:    []string{"8v8"},
			expErr:  "could not convert <order id> arg: strconv.ParseUint: parsing \"8v8\": invalid syntax",
		},
		{
			name:    "just arg, zero",
			flagSet: goodFlagSet(),
			args:    []string{"0"},
			expErr:  "no <order id> provided",
		},
		{
			name:    "just arg, good",
			flagSet: goodFlagSet(),
			args:    []string{"987"},
			expID:   987,
		},
		{
			name:    "neither flag nor arg",
			flagSet: goodFlagSet(),
			expErr:  "no <order id> provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var id uint64
			testFunc := func() {
				id, err = cli.ReadFlagOrderOrArg(tc.flagSet, tc.args)
			}
			require.NotPanics(t, testFunc, "ReadFlagOrderOrArg")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagOrderOrArg error")
			assert.Equal(t, int(tc.expID), int(id), "ReadFlagOrderOrArg id")
		})
	}
}

func TestReadFlagMarketOrArg(t *testing.T) {
	theFlag := cli.FlagMarket
	goodFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.Uint32(theFlag, 0, "The id")
		return flagSet
	}
	badFlagSet := func() *pflag.FlagSet {
		flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
		flagSet.String(theFlag, "", "The id")
		return flagSet
	}

	tests := []struct {
		name    string
		flags   []string
		flagSet *pflag.FlagSet
		args    []string
		expID   uint32
		expErr  string
	}{
		{
			name:    "unknown flag",
			flagSet: pflag.NewFlagSet("", pflag.ContinueOnError),
			expErr:  "flag accessed but not defined: " + theFlag,
		},
		{
			name:    "wrong flag type",
			flagSet: badFlagSet(),
			expErr:  "trying to get uint32 value of flag of type string",
		},
		{
			name:    "both flag and arg",
			flags:   []string{"--" + theFlag, "8"},
			flagSet: goodFlagSet(),
			args:    []string{"8"},
			expErr:  "cannot provide <market id> as both an arg (\"8\") and flag (--market 8)",
		},
		{
			name:    "just flag",
			flags:   []string{"--" + theFlag, "8"},
			flagSet: goodFlagSet(),
			expID:   8,
		},
		{
			name:    "just flag zero",
			flags:   []string{"--" + theFlag, "0"},
			flagSet: goodFlagSet(),
			expErr:  "no <market id> provided",
		},
		{
			name:    "just arg, bad",
			flagSet: goodFlagSet(),
			args:    []string{"8v8"},
			expErr:  "could not convert <market id> arg: strconv.ParseUint: parsing \"8v8\": invalid syntax",
		},
		{
			name:    "just arg, zero",
			flagSet: goodFlagSet(),
			args:    []string{"0"},
			expErr:  "no <market id> provided",
		},
		{
			name:    "just arg, good",
			flagSet: goodFlagSet(),
			args:    []string{"987"},
			expID:   987,
		},
		{
			name:    "neither flag nor arg",
			flagSet: goodFlagSet(),
			expErr:  "no <market id> provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var id uint32
			testFunc := func() {
				id, err = cli.ReadFlagMarketOrArg(tc.flagSet, tc.args)
			}
			require.NotPanics(t, testFunc, "ReadFlagMarketOrArg")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagMarketOrArg error")
			assert.Equal(t, int(tc.expID), int(id), "ReadFlagMarketOrArg id")
		})
	}
}

func TestReadCoinsFlag(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		expCoins sdk.Coins
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagString,
			expErr:   "",
		},
		{
			testName: "invalid coins",
			flags:    []string{"--" + flagString, "2yupcoin,nopecoin"},
			name:     flagString,
			expErr:   "error parsing --" + flagString + " as coins: invalid coin expression: \"nopecoin\"",
		},
		{
			testName: "one coin",
			flags:    []string{"--" + flagString, "2grape"},
			name:     flagString,
			expCoins: sdk.NewCoins(sdk.NewInt64Coin("grape", 2)),
		},
		{
			testName: "three coins",
			flags:    []string{"--" + flagString, "8banana,5apple,14cherry"},
			name:     flagString,
			expCoins: sdk.NewCoins(
				sdk.NewInt64Coin("apple", 5), sdk.NewInt64Coin("banana", 8), sdk.NewInt64Coin("cherry", 14),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var coins sdk.Coins
			testFunc := func() {
				coins, err = cli.ReadCoinsFlag(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadCoinsFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadCoinsFlag(%q) error", tc.name)
			assert.Equal(t, tc.expCoins.String(), coins.String(), "ReadCoinsFlag(%q) coins", tc.name)
		})
	}
}

func TestReadCoinsFlagOrDefault(t *testing.T) {
	coins := func(coinsStr string) sdk.Coins {
		rv, err := sdk.ParseCoinsNormalized(coinsStr)
		require.NoError(t, err, "ParseCoinsNormalized(%q)", coinsStr)
		return rv
	}
	tests := []struct {
		testName string
		flags    []string
		name     string
		def      sdk.Coins
		expCoins sdk.Coins
		expErr   string
	}{
		{
			testName: "unknown flag, no default",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "unknown flag, with default",
			name:     "unknown",
			def:      coins("12orange"),
			expCoins: coins("12orange"),
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type, no default",
			name:     flagInt,
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "wrong flag type, with default",
			name:     flagInt,
			def:      coins("8banana,14pear"),
			expCoins: coins("8banana,14pear"),
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "nothing provided, no default",
			name:     flagString,
			expErr:   "",
		},
		{
			testName: "nothing provided, with default",
			name:     flagString,
			def:      coins("88cherry"),
			expCoins: coins("88cherry"),
			expErr:   "",
		},
		{
			testName: "invalid coins, no default",
			flags:    []string{"--" + flagString, "2yupcoin,nopecoin"},
			name:     flagString,
			expErr:   "error parsing --" + flagString + " as coins: invalid coin expression: \"nopecoin\"",
		},
		{
			testName: "invalid coins, with default",
			flags:    []string{"--" + flagString, "2yupcoin,nopecoin"},
			name:     flagString,
			def:      coins("57lychee"),
			expCoins: coins("57lychee"),
			expErr:   "error parsing --" + flagString + " as coins: invalid coin expression: \"nopecoin\"",
		},
		{
			testName: "one coin, no default",
			flags:    []string{"--" + flagString, "2grape"},
			name:     flagString,
			expCoins: coins("2grape"),
		},
		{
			testName: "one coin, with default",
			flags:    []string{"--" + flagString, "2grape"},
			name:     flagString,
			def:      coins("6apple"),
			expCoins: coins("2grape"),
		},
		{
			testName: "three coins, no default",
			flags:    []string{"--" + flagString, "8banana,5apple,14cherry"},
			name:     flagString,
			expCoins: sdk.NewCoins(
				sdk.NewInt64Coin("apple", 5), sdk.NewInt64Coin("banana", 8), sdk.NewInt64Coin("cherry", 14),
			),
		},
		{
			testName: "three coins, with default",
			flags:    []string{"--" + flagString, "8banana,5apple,14cherry"},
			name:     flagString,
			def:      coins("15plum"),
			expCoins: coins("5apple,8banana,14cherry"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var actCoins sdk.Coins
			testFunc := func() {
				actCoins, err = cli.ReadCoinsFlagOrDefault(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadCoinsFlagOrDefault(%q, %q)", tc.name, tc.def)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadCoinsFlagOrDefault(%q, %q) error", tc.name, tc.def)
			assert.Equal(t, tc.expCoins.String(), actCoins.String(), "ReadCoinsFlagOrDefault(%q, %q) coins", tc.name, tc.def)
		})
	}
}

func TestReadReqCoinsFlag(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		expCoins sdk.Coins
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagString,
			expErr:   "missing required --" + flagString + " flag",
		},
		{
			testName: "invalid coins",
			flags:    []string{"--" + flagString, "2yupcoin,nopecoin"},
			name:     flagString,
			expErr:   "error parsing --" + flagString + " as coins: invalid coin expression: \"nopecoin\"",
		},
		{
			testName: "one coin",
			flags:    []string{"--" + flagString, "2grape"},
			name:     flagString,
			expCoins: sdk.NewCoins(sdk.NewInt64Coin("grape", 2)),
		},
		{
			testName: "three coins",
			flags:    []string{"--" + flagString, "8banana,5apple,14cherry"},
			name:     flagString,
			expCoins: sdk.NewCoins(
				sdk.NewInt64Coin("apple", 5), sdk.NewInt64Coin("banana", 8), sdk.NewInt64Coin("cherry", 14),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var coins sdk.Coins
			testFunc := func() {
				coins, err = cli.ReadReqCoinsFlag(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadReqCoinsFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadReqCoinsFlag(%q) error", tc.name)
			assert.Equal(t, tc.expCoins.String(), coins.String(), "ReadReqCoinsFlag(%q) coins", tc.name)
		})
	}
}

func TestParseCoins(t *testing.T) {
	tests := []struct {
		name     string
		coinsStr string
		expCoins sdk.Coins
		expErr   string
	}{
		{
			name:     "empty string",
			coinsStr: "",
			expCoins: nil,
			expErr:   "",
		},
		{
			name:     "one entry, bad",
			coinsStr: "bad",
			expErr:   "invalid coin expression: \"bad\"",
		},
		{
			name:     "one entry, good",
			coinsStr: "55good",
			expCoins: sdk.NewCoins(sdk.NewInt64Coin("good", 55)),
		},
		{
			name:     "three entries, first bad",
			coinsStr: "1234,555second,63third",
			expErr:   "invalid coin expression: \"1234\"",
		},
		{
			name:     "three entries, second bad",
			coinsStr: "1234first,second,55third",
			expErr:   "invalid coin expression: \"second\"",
		},
		{
			name:     "three entries, third bad",
			coinsStr: "1234first,555second,63x",
			expErr:   "invalid coin expression: \"63x\"",
		},
		{
			name:     "three entries, all good",
			coinsStr: "1234one,555two,63three",
			expCoins: sdk.NewCoins(
				sdk.NewInt64Coin("one", 1234),
				sdk.NewInt64Coin("three", 63),
				sdk.NewInt64Coin("two", 555),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var coins sdk.Coins
			var err error
			testFunc := func() {
				coins, err = cli.ParseCoins(tc.coinsStr)
			}
			require.NotPanics(t, testFunc, "ParseCoins(%q)", tc.coinsStr)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseCoins(%q) error", tc.coinsStr)
			assert.Equal(t, tc.expCoins.String(), coins.String(), "ParseCoins(%q) coins", tc.coinsStr)
		})
	}
}

func TestReadCoinFlag(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		expCoin  *sdk.Coin
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagString,
			expErr:   "",
		},
		{
			testName: "invalid coin",
			flags:    []string{"--" + flagString, "nopecoin"},
			name:     flagString,
			expErr:   "error parsing --" + flagString + " as a coin: invalid coin expression: \"nopecoin\"",
		},
		{
			testName: "zero coin",
			flags:    []string{"--" + flagString, "0zerocoin"},
			name:     flagString,
			expCoin:  &sdk.Coin{Denom: "zerocoin", Amount: sdkmath.NewInt(0)},
		},
		{
			testName: "normal coin",
			flags:    []string{"--" + flagString, "99banana"},
			name:     flagString,
			expCoin:  &sdk.Coin{Denom: "banana", Amount: sdkmath.NewInt(99)},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var coin *sdk.Coin
			testFunc := func() {
				coin, err = cli.ReadCoinFlag(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadCoinFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadCoinFlag(%q) error", tc.name)
			if !assert.Equal(t, tc.expCoin, coin, "ReadCoinFlag(%q)", tc.name) && tc.expCoin != nil && coin != nil {
				t.Logf("Expected: %q", tc.expCoin)
				t.Logf("  Actual: %q", coin)
			}
		})
	}
}

func TestReadReqCoinFlag(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		expCoin  sdk.Coin
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagString,
			expErr:   "missing required --" + flagString + " flag",
		},
		{
			testName: "invalid coin",
			flags:    []string{"--" + flagString, "nopecoin"},
			name:     flagString,
			expErr:   "error parsing --" + flagString + " as a coin: invalid coin expression: \"nopecoin\"",
		},
		{
			testName: "zero coin",
			flags:    []string{"--" + flagString, "0zerocoin"},
			name:     flagString,
			expCoin:  sdk.Coin{Denom: "zerocoin", Amount: sdkmath.NewInt(0)},
		},
		{
			testName: "normal coin",
			flags:    []string{"--" + flagString, "99banana"},
			name:     flagString,
			expCoin:  sdk.Coin{Denom: "banana", Amount: sdkmath.NewInt(99)},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var coin sdk.Coin
			testFunc := func() {
				coin, err = cli.ReadReqCoinFlag(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadReqCoinFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadReqCoinFlag(%q) error", tc.name)
			assert.Equal(t, tc.expCoin.String(), coin.String(), "ReadReqCoinFlag(%q)", tc.name)
		})
	}
}

func TestReadOrderIDsFlag(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		expIDs   []uint64
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagString,
			expErr:   "trying to get string value of flag of type uintSlice",
		},
		{
			testName: "nothing provided",
			name:     flagUintSlice,
			expErr:   "",
		},
		{
			testName: "one val",
			flags:    []string{"--" + flagUintSlice, "15"},
			name:     flagUintSlice,
			expIDs:   []uint64{15},
		},
		{
			testName: "three vals",
			flags:    []string{"--" + flagUintSlice, "42,9001,3"},
			name:     flagUintSlice,
			expIDs:   []uint64{42, 9001, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.UintSlice(flagUintSlice, nil, "A slice of uints")
			flagSet.String(flagString, "", "A string")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var ids []uint64
			testFunc := func() {
				ids, err = cli.ReadOrderIDsFlag(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadOrderIDsFlag(%q)", tc.name)
			assertEqualSlices(t, tc.expIDs, ids, orderIDStringer, "ReadOrderIDsFlag(%q)", tc.name)
		})
	}
}

func TestReadAccessGrantsFlag(t *testing.T) {
	tests := []struct {
		testName  string
		flags     []string
		name      string
		def       []exchange.AccessGrant
		expGrants []exchange.AccessGrant
		expErr    string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided, nil default",
			name:     flagStringSlice,
			expErr:   "",
		},
		{
			testName: "nothing provided, with default",
			name:     flagStringSlice,
			def: []exchange.AccessGrant{
				{Address: "someone", Permissions: []exchange.Permission{3, 4}},
			},
			expGrants: []exchange.AccessGrant{
				{Address: "someone", Permissions: []exchange.Permission{3, 4}},
			},
			expErr: "",
		},
		{
			testName: "three vals, one bad",
			flags: []string{
				"--" + flagStringSlice, "addr1:all",
				"--" + flagStringSlice, "withdraw",
				"--" + flagStringSlice, "addr2:setids+update",
			},
			name: flagStringSlice,
			expGrants: []exchange.AccessGrant{
				{
					Address:     "addr1",
					Permissions: exchange.AllPermissions(),
				},
				{
					Address:     "addr2",
					Permissions: []exchange.Permission{exchange.Permission_set_ids, exchange.Permission_update},
				},
			},
			expErr: "could not parse \"withdraw\" as an <access grant>: expected format <address>:<permissions>",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var grants []exchange.AccessGrant
			testFunc := func() {
				grants, err = cli.ReadAccessGrantsFlag(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadAccessGrantsFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadAccessGrantsFlag(%q) error", tc.name)
			assert.Equal(t, tc.expGrants, grants, "ReadAccessGrantsFlag(%q) grants", tc.name)
		})
	}
}

func TestParseAccessGrant(t *testing.T) {
	addr := "pb1v9jxgujlta047h6lta047h6lta047h6l5rpeqp" // = sdk.AccAddress("addr________________")

	tests := []struct {
		name   string
		val    string
		expAG  *exchange.AccessGrant
		expErr string
	}{
		{
			name:   "empty string",
			val:    "",
			expAG:  nil,
			expErr: "could not parse \"\" as an <access grant>: expected format <address>:<permissions>",
		},
		{
			name:   "zero colons",
			val:    "something",
			expErr: "could not parse \"something\" as an <access grant>: expected format <address>:<permissions>",
		},
		{
			name:   "two colons",
			val:    "part0:part1:part2",
			expErr: "could not parse \"part0:part1:part2\" as an <access grant>: expected format <address>:<permissions>",
		},
		{
			name:   "empty address",
			val:    ":part1",
			expErr: "invalid <access grant> \":part1\": both an <address> and <permissions> are required",
		},
		{
			name:   "empty permissions",
			val:    "part0:",
			expErr: "invalid <access grant> \"part0:\": both an <address> and <permissions> are required",
		},
		{
			name:   "unspecified",
			val:    "part0:unspecified",
			expErr: "could not parse permissions for \"part0\" from \"unspecified\": invalid permission: \"unspecified\"",
		},
		{
			name:  "all",
			val:   addr + ":all",
			expAG: &exchange.AccessGrant{Address: addr, Permissions: exchange.AllPermissions()},
		},
		{
			name: "one perm, enum name",
			val:  addr + ":PERMISSION_UPDATE",
			expAG: &exchange.AccessGrant{
				Address:     addr,
				Permissions: []exchange.Permission{exchange.Permission_update},
			},
		},
		{
			name: "one perm, simple name",
			val:  addr + ":cancel",
			expAG: &exchange.AccessGrant{
				Address:     addr,
				Permissions: []exchange.Permission{exchange.Permission_cancel},
			},
		},
		{
			name: "multiple perms, plus delim",
			val:  addr + ":Cancel+PERMISSION_SETTLE+setids",
			expAG: &exchange.AccessGrant{
				Address: addr,
				Permissions: []exchange.Permission{
					exchange.Permission_cancel,
					exchange.Permission_settle,
					exchange.Permission_set_ids,
				},
			},
		},
		{
			name: "multiple perms, dot delim",
			val:  addr + ":permissions.PERMISSION_ATTRIBUTES.withdraw",
			expAG: &exchange.AccessGrant{
				Address: addr,
				Permissions: []exchange.Permission{
					exchange.Permission_permissions,
					exchange.Permission_attributes,
					exchange.Permission_withdraw,
				},
			},
		},
		{
			name: "multiple perms, space delim",
			val:  addr + ":Set_Ids update settle permissions",
			expAG: &exchange.AccessGrant{
				Address: addr,
				Permissions: []exchange.Permission{
					exchange.Permission_set_ids,
					exchange.Permission_update,
					exchange.Permission_settle,
					exchange.Permission_permissions,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var ag *exchange.AccessGrant
			var err error
			testFunc := func() {
				ag, err = cli.ParseAccessGrant(tc.val)
			}
			require.NotPanics(t, testFunc, "ParseAccessGrant(%q)", tc.val)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseAccessGrant(%q) error", tc.val)
			assert.Equal(t, tc.expAG, ag, "ParseAccessGrant(%q) AccessGrant", tc.val)
		})
	}
}

func TestParseAccessGrants(t *testing.T) {
	addr1 := "pb1v9jxgu33ta047h6lta047h6lta047h6l0r6x5v" // = sdk.AccAddress("addr1_______________").String()
	addr2 := "pb1taskgerjxf047h6lta047h6lta047h6lrcgmd9" // = sdk.AccAddress("_addr2______________").String()
	addr3 := "pb10elxzerywge47h6lta047h6lta047h6l90x0zx" // = sdk.AccAddress("~~addr3_____________").String()

	tests := []struct {
		name      string
		vals      []string
		expGrants []exchange.AccessGrant
		expErr    string
	}{
		{
			name:   "nil",
			vals:   nil,
			expErr: "",
		},
		{
			name:   "empty",
			vals:   []string{},
			expErr: "",
		},
		{
			name:   "one, bad",
			vals:   []string{"not good"},
			expErr: "could not parse \"not good\" as an <access grant>: expected format <address>:<permissions>",
		},
		{
			name: "one, good",
			vals: []string{addr1 + ":update+permissions"},
			expGrants: []exchange.AccessGrant{{
				Address:     addr1,
				Permissions: []exchange.Permission{exchange.Permission_update, exchange.Permission_permissions},
			}},
		},
		{
			name: "three, all good",
			vals: []string{addr1 + ":settle", addr2 + ":setids", addr3 + ":permission_withdraw"},
			expGrants: []exchange.AccessGrant{
				{Address: addr1, Permissions: []exchange.Permission{exchange.Permission_settle}},
				{Address: addr2, Permissions: []exchange.Permission{exchange.Permission_set_ids}},
				{Address: addr3, Permissions: []exchange.Permission{exchange.Permission_withdraw}},
			},
		},
		{
			name: "three, first bad",
			vals: []string{":settle", addr2 + ":setids", addr3 + ":permission_withdraw"},
			expGrants: []exchange.AccessGrant{
				{Address: addr2, Permissions: []exchange.Permission{exchange.Permission_set_ids}},
				{Address: addr3, Permissions: []exchange.Permission{exchange.Permission_withdraw}},
			},
			expErr: "invalid <access grant> \":settle\": both an <address> and <permissions> are required",
		},
		{
			name: "three, second bad",
			vals: []string{addr1 + ":settle", addr2 + ":unspecified", addr3 + ":permission_withdraw"},
			expGrants: []exchange.AccessGrant{
				{Address: addr1, Permissions: []exchange.Permission{exchange.Permission_settle}},
				{Address: addr3, Permissions: []exchange.Permission{exchange.Permission_withdraw}},
			},
			expErr: "could not parse permissions for \"" + addr2 + "\" from \"unspecified\": invalid permission: \"unspecified\"",
		},
		{
			name: "three, third bad",
			vals: []string{addr1 + ":settle", addr2 + ":setids", "someaddr:"},
			expGrants: []exchange.AccessGrant{
				{Address: addr1, Permissions: []exchange.Permission{exchange.Permission_settle}},
				{Address: addr2, Permissions: []exchange.Permission{exchange.Permission_set_ids}},
			},
			expErr: "invalid <access grant> \"someaddr:\": both an <address> and <permissions> are required",
		},
		{
			name: "three, all bad",
			vals: []string{":settle", addr2 + ":unspecified", "someaddr:"},
			expErr: joinErrs(
				"invalid <access grant> \":settle\": both an <address> and <permissions> are required",
				"could not parse permissions for \""+addr2+"\" from \"unspecified\": invalid permission: \"unspecified\"",
				"invalid <access grant> \"someaddr:\": both an <address> and <permissions> are required",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expGrants == nil {
				tc.expGrants = []exchange.AccessGrant{}
			}

			var grants []exchange.AccessGrant
			var err error
			testFunc := func() {
				grants, err = cli.ParseAccessGrants(tc.vals)
			}
			require.NotPanics(t, testFunc, "ParseAccessGrants(%q)", tc.vals)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseAccessGrants(%q) error", tc.vals)
			assert.Equal(t, tc.expGrants, grants, "ParseAccessGrants(%q) grants", tc.vals)
		})
	}
}

func TestReadFlatFeeFlag(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		def      []sdk.Coin
		expCoins []sdk.Coin
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided, nil default",
			name:     flagStringSlice,
			expErr:   "",
		},
		{
			testName: "nothing provided, with default",
			name:     flagStringSlice,
			def:      []sdk.Coin{sdk.NewInt64Coin("cherry", 123)},
			expCoins: []sdk.Coin{sdk.NewInt64Coin("cherry", 123)},
			expErr:   "",
		},
		{
			testName: "three vals, one bad",
			flags:    []string{"--" + flagStringSlice, "apple,100pear", "--" + flagStringSlice, "777cherry"},
			name:     flagStringSlice,
			expCoins: []sdk.Coin{sdk.NewInt64Coin("pear", 100), sdk.NewInt64Coin("cherry", 777)},
			expErr:   "invalid coin expression: \"apple\"",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var coins []sdk.Coin
			testFunc := func() {
				coins, err = cli.ReadFlatFeeFlag(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlatFeeFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlatFeeFlag(%q) error", tc.name)
			assertEqualSlices(t, tc.expCoins, coins, sdk.Coin.String, "ReadFlatFeeFlag(%q) ratios", tc.name)
		})
	}
}

func TestParseFlatFeeOptions(t *testing.T) {
	tests := []struct {
		name     string
		vals     []string
		expCoins []sdk.Coin
		expErr   string
	}{
		{
			name:   "nil",
			vals:   nil,
			expErr: "",
		},
		{
			name:   "empty",
			vals:   []string{},
			expErr: "",
		},
		{
			name:   "one, bad",
			vals:   []string{"nope"},
			expErr: "invalid coin expression: \"nope\"",
		},
		{
			name:     "one, good",
			vals:     []string{"18banana"},
			expCoins: []sdk.Coin{sdk.NewInt64Coin("banana", 18)},
		},
		{
			name:     "one, zero",
			vals:     []string{"0durian"},
			expCoins: []sdk.Coin{sdk.NewInt64Coin("durian", 0)},
		},
		{
			name: "three, all good",
			vals: []string{"1apple", "2banana", "3cherry"},
			expCoins: []sdk.Coin{
				sdk.NewInt64Coin("apple", 1), sdk.NewInt64Coin("banana", 2), sdk.NewInt64Coin("cherry", 3),
			},
		},
		{
			name: "three, first bad",
			vals: []string{"notgonnacoin", "2banana", "3cherry"},
			expCoins: []sdk.Coin{
				sdk.NewInt64Coin("banana", 2), sdk.NewInt64Coin("cherry", 3),
			},
			expErr: "invalid coin expression: \"notgonnacoin\"",
		},
		{
			name: "three, second bad",
			vals: []string{"1apple", "12345", "3cherry"},
			expCoins: []sdk.Coin{
				sdk.NewInt64Coin("apple", 1), sdk.NewInt64Coin("cherry", 3),
			},
			expErr: "invalid coin expression: \"12345\"",
		},
		{
			name: "three, third bad",
			vals: []string{"1apple", "2banana", ""},
			expCoins: []sdk.Coin{
				sdk.NewInt64Coin("apple", 1), sdk.NewInt64Coin("banana", 2),
			},
			expErr: "invalid coin expression: \"\"",
		},
		{
			name: "three, all bad",
			vals: []string{"notgonnacoin", "12345", ""},
			expErr: joinErrs(
				"invalid coin expression: \"notgonnacoin\"",
				"invalid coin expression: \"12345\"",
				"invalid coin expression: \"\"",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expCoins == nil {
				tc.expCoins = []sdk.Coin{}
			}
			var coins []sdk.Coin
			var err error
			testFunc := func() {
				coins, err = cli.ParseFlatFeeOptions(tc.vals)
			}
			require.NotPanics(t, testFunc, "ParseFlatFeeOptions(%q)", tc.vals)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseFlatFeeOptions(%q) error", tc.vals)
			assertEqualSlices(t, tc.expCoins, coins, sdk.Coin.String, "ParseFlatFeeOptions(%q) coins", tc.vals)
		})
	}
}

func TestReadFeeRatiosFlag(t *testing.T) {
	tests := []struct {
		testName  string
		flags     []string
		name      string
		def       []exchange.FeeRatio
		expRatios []exchange.FeeRatio
		expErr    string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided, nil default",
			name:     flagStringSlice,
			expErr:   "",
		},
		{
			testName:  "nothing provided, with default",
			name:      flagStringSlice,
			def:       []exchange.FeeRatio{{Price: sdk.NewInt64Coin("apple", 500), Fee: sdk.NewInt64Coin("plum", 3)}},
			expRatios: []exchange.FeeRatio{{Price: sdk.NewInt64Coin("apple", 500), Fee: sdk.NewInt64Coin("plum", 3)}},
			expErr:    "",
		},
		{
			testName: "three vals, one bad",
			flags:    []string{"--" + flagStringSlice, "8apple:3apple,100pear:1apple", "--" + flagStringSlice, "cherry:777cherry"},
			name:     flagStringSlice,
			expRatios: []exchange.FeeRatio{
				{Price: sdk.NewInt64Coin("apple", 8), Fee: sdk.NewInt64Coin("apple", 3)},
				{Price: sdk.NewInt64Coin("pear", 100), Fee: sdk.NewInt64Coin("apple", 1)},
			},
			expErr: "cannot create FeeRatio from \"cherry:777cherry\": price: invalid coin expression: \"cherry\"",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var ratios []exchange.FeeRatio
			testFunc := func() {
				ratios, err = cli.ReadFeeRatiosFlag(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFeeRatiosFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFeeRatiosFlag(%q) error", tc.name)
			assertEqualSlices(t, tc.expRatios, ratios, exchange.FeeRatio.String, "ReadFeeRatiosFlag(%q) ratios", tc.name)
		})
	}
}

func TestParseFeeRatios(t *testing.T) {
	tests := []struct {
		name      string
		vals      []string
		expRatios []exchange.FeeRatio
		expErr    string
	}{
		{
			name:   "nil",
			vals:   nil,
			expErr: "",
		},
		{
			name:   "empty",
			vals:   []string{},
			expErr: "",
		},
		{
			name:   "one, bad",
			vals:   []string{"notaratio"},
			expErr: "cannot create FeeRatio from \"notaratio\": expected exactly one colon",
		},
		{
			name:      "one, good",
			vals:      []string{"10apple:3banana"},
			expRatios: []exchange.FeeRatio{{Price: sdk.NewInt64Coin("apple", 10), Fee: sdk.NewInt64Coin("banana", 3)}},
		},
		{
			name:      "one, zeros",
			vals:      []string{"0cherry:0durian"},
			expRatios: []exchange.FeeRatio{{Price: sdk.NewInt64Coin("cherry", 0), Fee: sdk.NewInt64Coin("durian", 0)}},
		},
		{
			name: "three, all good",
			vals: []string{"10apple:1cherry", "321banana:8grape", "66plum:7plum"},
			expRatios: []exchange.FeeRatio{
				{Price: sdk.NewInt64Coin("apple", 10), Fee: sdk.NewInt64Coin("cherry", 1)},
				{Price: sdk.NewInt64Coin("banana", 321), Fee: sdk.NewInt64Coin("grape", 8)},
				{Price: sdk.NewInt64Coin("plum", 66), Fee: sdk.NewInt64Coin("plum", 7)},
			},
		},
		{
			name: "three, first bad",
			vals: []string{"10apple", "321banana:8grape", "66plum:7plum"},
			expRatios: []exchange.FeeRatio{
				{Price: sdk.NewInt64Coin("banana", 321), Fee: sdk.NewInt64Coin("grape", 8)},
				{Price: sdk.NewInt64Coin("plum", 66), Fee: sdk.NewInt64Coin("plum", 7)},
			},
			expErr: "cannot create FeeRatio from \"10apple\": expected exactly one colon",
		},
		{
			name: "three, second bad",
			vals: []string{"10apple:1cherry", "8grape", "66plum:7plum"},
			expRatios: []exchange.FeeRatio{
				{Price: sdk.NewInt64Coin("apple", 10), Fee: sdk.NewInt64Coin("cherry", 1)},
				{Price: sdk.NewInt64Coin("plum", 66), Fee: sdk.NewInt64Coin("plum", 7)},
			},
			expErr: "cannot create FeeRatio from \"8grape\": expected exactly one colon",
		},
		{
			name: "three, third bad",
			vals: []string{"10apple:1cherry", "321banana:8grape", ""},
			expRatios: []exchange.FeeRatio{
				{Price: sdk.NewInt64Coin("apple", 10), Fee: sdk.NewInt64Coin("cherry", 1)},
				{Price: sdk.NewInt64Coin("banana", 321), Fee: sdk.NewInt64Coin("grape", 8)},
			},
			expErr: "cannot create FeeRatio from \"\": expected exactly one colon",
		},
		{
			name: "three, all bad",
			vals: []string{"10apple", "8grape", ""},
			expErr: joinErrs(
				"cannot create FeeRatio from \"10apple\": expected exactly one colon",
				"cannot create FeeRatio from \"8grape\": expected exactly one colon",
				"cannot create FeeRatio from \"\": expected exactly one colon",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expRatios == nil {
				tc.expRatios = []exchange.FeeRatio{}
			}

			var ratios []exchange.FeeRatio
			var err error
			testFunc := func() {
				ratios, err = cli.ParseFeeRatios(tc.vals)
			}
			require.NotPanics(t, testFunc, "ParseFeeRatios(%q)", tc.vals)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseFeeRatios(%q) error", tc.vals)
			assertEqualSlices(t, tc.expRatios, ratios, exchange.FeeRatio.String, "ParseFeeRatios(%q) ratios", tc.vals)
		})
	}
}

func TestReadSplitsFlag(t *testing.T) {
	tests := []struct {
		testName  string
		flags     []string
		name      string
		expSplits []exchange.DenomSplit
		expErr    string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagStringSlice,
			expErr:   "",
		},
		{
			testName: "three vals, one bad",
			flags:    []string{"--" + flagStringSlice, "apple:3,banana:80q0", "--" + flagStringSlice, "cherry:777"},
			name:     flagStringSlice,
			expSplits: []exchange.DenomSplit{
				{Denom: "apple", Split: 3},
				{Denom: "cherry", Split: 777},
			},
			expErr: "could not parse \"banana:80q0\" amount: strconv.ParseUint: parsing \"80q0\": invalid syntax",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var splits []exchange.DenomSplit
			testFunc := func() {
				splits, err = cli.ReadSplitsFlag(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadSplitsFlag(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadSplitsFlag(%q) error", tc.name)
			assertEqualSlices(t, tc.expSplits, splits, splitStringer, "ReadSplitsFlag(%q) splits", tc.name)
		})
	}
}

func TestParseSplit(t *testing.T) {
	tests := []struct {
		name     string
		val      string
		expSplit *exchange.DenomSplit
		expErr   string
	}{
		{
			name:   "empty",
			val:    "",
			expErr: "invalid denom split \"\": expected format <denom>:<amount>",
		},
		{
			name:     "no colons",
			val:      "banana",
			expSplit: nil,
			expErr:   "invalid denom split \"banana\": expected format <denom>:<amount>",
		},
		{
			name:   "two colons",
			val:    "plum:8:123",
			expErr: "invalid denom split \"plum:8:123\": expected format <denom>:<amount>",
		},
		{
			name:   "empty denom",
			val:    ":444",
			expErr: "invalid denom split \":444\": both a <denom> and <amount> are required",
		},
		{
			name:   "empty amount",
			val:    "apple:",
			expErr: "invalid denom split \"apple:\": both a <denom> and <amount> are required",
		},
		{
			name:   "invalid amount",
			val:    "apple:banana",
			expErr: "could not parse \"apple:banana\" amount: strconv.ParseUint: parsing \"banana\": invalid syntax",
		},
		{
			name:     "good, zero",
			val:      "cherry:0",
			expSplit: &exchange.DenomSplit{Denom: "cherry", Split: 0},
		},
		{
			name:     "good, 10,000",
			val:      "pear:10000",
			expSplit: &exchange.DenomSplit{Denom: "pear", Split: 10000},
		},
		{
			name:     "good, 123",
			val:      "acorn:123",
			expSplit: &exchange.DenomSplit{Denom: "acorn", Split: 123},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var split *exchange.DenomSplit
			var err error
			testFunc := func() {
				split, err = cli.ParseSplit(tc.val)
			}
			require.NotPanics(t, testFunc, "ParseSplit(%q)", tc.val)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseSplit(%q) error", tc.val)
			if !assert.Equal(t, tc.expSplit, split, "ParseSplit(%q) split", tc.val) {
				t.Logf("Expected: %s:%d", tc.expSplit.Denom, tc.expSplit.Split)
				t.Logf("  Actual: %s:%d", split.Denom, split.Split)
			}
		})
	}
}

func TestParseSplits(t *testing.T) {
	tests := []struct {
		name      string
		vals      []string
		expSplits []exchange.DenomSplit
		expErr    string
	}{
		{
			name:   "nil",
			vals:   nil,
			expErr: "",
		},
		{
			name:   "empty",
			vals:   []string{},
			expErr: "",
		},
		{
			name:   "one, bad",
			vals:   []string{"nope"},
			expErr: "invalid denom split \"nope\": expected format <denom>:<amount>",
		},
		{
			name:      "one, good",
			vals:      []string{"yup:5"},
			expSplits: []exchange.DenomSplit{{Denom: "yup", Split: 5}},
		},
		{
			name: "three, all good",
			vals: []string{"first:1", "second:22", "third:333"},
			expSplits: []exchange.DenomSplit{
				{Denom: "first", Split: 1}, {Denom: "second", Split: 22}, {Denom: "third", Split: 333},
			},
		},
		{
			name: "three, first bad",
			vals: []string{"first", "second:22", "third:333"},
			expSplits: []exchange.DenomSplit{
				{Denom: "second", Split: 22}, {Denom: "third", Split: 333},
			},
			expErr: "invalid denom split \"first\": expected format <denom>:<amount>",
		},
		{
			name: "three, second bad",
			vals: []string{"first:1", ":22", "third:333"},
			expSplits: []exchange.DenomSplit{
				{Denom: "first", Split: 1}, {Denom: "third", Split: 333},
			},
			expErr: "invalid denom split \":22\": both a <denom> and <amount> are required",
		},
		{
			name: "three, third bad",
			vals: []string{"first:1", "second:22", "third:333x"},
			expSplits: []exchange.DenomSplit{
				{Denom: "first", Split: 1}, {Denom: "second", Split: 22},
			},
			expErr: "could not parse \"third:333x\" amount: strconv.ParseUint: parsing \"333x\": invalid syntax",
		},
		{
			name: "three, all bad",
			vals: []string{"first", ":22", "third:333x"},
			expErr: joinErrs(
				"invalid denom split \"first\": expected format <denom>:<amount>",
				"invalid denom split \":22\": both a <denom> and <amount> are required",
				"could not parse \"third:333x\" amount: strconv.ParseUint: parsing \"333x\": invalid syntax",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expSplits == nil {
				tc.expSplits = []exchange.DenomSplit{}
			}

			var splits []exchange.DenomSplit
			var err error
			testFunc := func() {
				splits, err = cli.ParseSplits(tc.vals)
			}
			require.NotPanics(t, testFunc, "ParseSplits(%q)", tc.vals)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseSplits(%q) error", tc.vals)
			assertEqualSlices(t, tc.expSplits, splits, splitStringer, "ParseSplits(%q) splits", tc.vals)
		})
	}
}

func TestReadStringFlagOrArg(t *testing.T) {
	tests := []struct {
		name     string
		flags    []string
		args     []string
		flagName string
		varName  string
		expStr   string
		expErr   string
	}{
		{
			name:     "unknown flag name",
			flagName: "other",
			varName:  "nope",
			expErr:   "flag accessed but not defined: other",
		},
		{
			name:     "wrong flag type",
			flagName: flagInt,
			varName:  "number",
			expErr:   "trying to get string value of flag of type int",
		},
		{
			name:     "both flag and arg",
			flags:    []string{"--" + flagString, "flagval"},
			args:     []string{"argval"},
			flagName: flagString,
			varName:  "value",
			expErr:   "cannot provide <value> as both an arg (\"argval\") and flag (--" + flagString + " \"flagval\")",
		},
		{
			name:     "only flag",
			flags:    []string{"--" + flagString, "flagval"},
			flagName: flagString,
			varName:  "value",
			expStr:   "flagval",
		},
		{
			name:     "only arg",
			args:     []string{"argval"},
			flagName: flagString,
			varName:  "value",
			expStr:   "argval",
		},
		{
			name:     "neither flag nor arg",
			flagName: flagString,
			varName:  "value",
			expErr:   "no <value> provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var str string
			testFunc := func() {
				str, err = cli.ReadStringFlagOrArg(flagSet, tc.args, tc.flagName, tc.varName)
			}
			require.NotPanics(t, testFunc, "ReadStringFlagOrArg")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadStringFlagOrArg error")
			assert.Equal(t, tc.expStr, str, "ReadStringFlagOrArg string")
		})
	}
}

func TestReadOptStringFlagOrArg(t *testing.T) {
	tests := []struct {
		name     string
		flags    []string
		args     []string
		flagName string
		varName  string
		expStr   string
		expErr   string
	}{
		{
			name:     "unknown flag name",
			flagName: "other",
			varName:  "nope",
			expErr:   "flag accessed but not defined: other",
		},
		{
			name:     "wrong flag type",
			flagName: flagInt,
			varName:  "number",
			expErr:   "trying to get string value of flag of type int",
		},
		{
			name:     "both flag and arg",
			flags:    []string{"--" + flagString, "flagval"},
			args:     []string{"argval"},
			flagName: flagString,
			varName:  "value",
			expErr:   "cannot provide <value> as both an arg (\"argval\") and flag (--" + flagString + " \"flagval\")",
		},
		{
			name:     "only flag",
			flags:    []string{"--" + flagString, "flagval"},
			flagName: flagString,
			varName:  "value",
			expStr:   "flagval",
		},
		{
			name:     "only arg",
			args:     []string{"argval"},
			flagName: flagString,
			varName:  "value",
			expStr:   "argval",
		},
		{
			name:     "neither flag nor arg",
			flagName: flagString,
			varName:  "value",
			expErr:   "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var str string
			testFunc := func() {
				str, err = cli.ReadOptStringFlagOrArg(flagSet, tc.args, tc.flagName, tc.varName)
			}
			require.NotPanics(t, testFunc, "ReadOptStringFlagOrArg")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadOptStringFlagOrArg error")
			assert.Equal(t, tc.expStr, str, "ReadOptStringFlagOrArg string")
		})
	}
}

func TestReadTxFileFlag(t *testing.T) {
	tests := []struct {
		name string
		// setup should return the filename and the expected Tx.
		setup    func(t *testing.T) (string, *txtypes.Tx)
		flagSet  *pflag.FlagSet
		fileFlag string
		expInErr []string
	}{
		{
			name: "flag not defined",
			setup: func(t *testing.T) (string, *txtypes.Tx) {
				return "", nil
			},
			flagSet:  pflag.NewFlagSet("", pflag.ContinueOnError),
			fileFlag: "fruits",
			expInErr: []string{"flag accessed but not defined: fruits"},
		},
		{
			name: "flag not provided",
			setup: func(t *testing.T) (string, *txtypes.Tx) {
				return "", nil
			},
			fileFlag: "fruits",
		},
		{
			name: "file does not exist",
			setup: func(t *testing.T) (string, *txtypes.Tx) {
				tdir := t.TempDir()
				noSuchFile := filepath.Join(tdir, "no-such-file.json")
				return noSuchFile, nil
			},
			fileFlag: "fruits",
			expInErr: []string{"open ", "no-such-file.json", "no such file or directory"},
		},
		{
			name: "error unmarshalling file",
			setup: func(t *testing.T) (string, *txtypes.Tx) {
				tdir := t.TempDir()
				notJSON := filepath.Join(tdir, "not-json.json")
				contents := []byte("This is not\na JSON file.\n")
				writeFile(t, notJSON, contents)
				return notJSON, nil
			},
			fileFlag: "nuts",
			expInErr: []string{
				"failed to unmarshal --nuts \"", "\" contents as Tx",
				"invalid character 'T' looking for beginning of value",
			},
		},
		{
			name: "okay: one message",
			setup: func(t *testing.T) (string, *txtypes.Tx) {
				msg := &exchange.MsgMarketWithdrawRequest{
					Admin:     sdk.AccAddress("Admin_______________").String(),
					MarketId:  3,
					ToAddress: sdk.AccAddress("ToAddress___________").String(),
					Amount:    sdk.NewCoins(sdk.NewInt64Coin("apple", 15)),
				}
				tx := newTx(t, msg)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "one-message.json")
				writeFileAsJson(t, fn, tx)
				return fn, tx

			},
			fileFlag: "legumes",
		},
		{
			name: "okay: three messages",
			setup: func(t *testing.T) (string, *txtypes.Tx) {
				msg1 := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market:    exchange.Market{MarketId: 88},
				}
				msg2 := &exchange.MsgGovManageFeesRequest{
					Authority:           cli.AuthorityAddr.String(),
					MarketId:            42,
					AddFeeCreateAskFlat: []sdk.Coin{sdk.NewInt64Coin("plum", 5)},
				}
				msg3 := &exchange.MsgCancelOrderRequest{
					Signer:  cli.AuthorityAddr.String(),
					OrderId: 5555,
				}
				tx := newTx(t, msg1, msg2, msg3)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "three-messages.json")
				writeFileAsJson(t, fn, tx)
				return fn, tx
			},
			fileFlag: "berries",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			expFilename, expTx := tc.setup(t)

			if tc.flagSet == nil {
				tc.flagSet = pflag.NewFlagSet("", pflag.ContinueOnError)
				tc.flagSet.String(tc.fileFlag, "", "The file")
			}

			if len(expFilename) > 0 && len(tc.expInErr) > 0 {
				tc.expInErr = append(tc.expInErr, expFilename)
			}

			args := make([]string, 0, 2)
			if len(expFilename) > 0 {
				args = append(args, "--"+tc.fileFlag, expFilename)
			}

			err := tc.flagSet.Parse(args)
			require.NoError(t, err, "flagSet.Parse(%q)", args)

			clientCtx := newClientContextWithCodec(t)

			var actFilename string
			var actTx *txtypes.Tx
			testFunc := func() {
				actFilename, actTx, err = cli.ReadTxFileFlag(clientCtx, tc.flagSet, tc.fileFlag)
			}
			require.NotPanics(t, testFunc, "ReadTxFileFlag")

			assertions.AssertErrorContents(t, err, tc.expInErr, "ReadTxFileFlag error")
			assert.Equal(t, expFilename, actFilename, "ReadTxFileFlag filename")

			// We can't just assert that two []Any lists are equal due to some internal differences.
			// All we really care about is that they have the same types and msg contents.
			if expTx == nil {
				assert.Nil(t, actTx, "ReadTxFileFlag tx")
				return
			}
			if !assert.NotNil(t, actTx, "ReadTxFileFlag tx") {
				return
			}

			if expTx.Body == nil {
				assert.Nil(t, actTx.Body, "ReadTxFileFlag Tx.Body")
				return
			}
			if !assert.NotNil(t, actTx.Body, "ReadTxFileFlag Tx.Body") {
				return
			}

			expTypes := getAnyTypes(expTx.Body.Messages)
			actTypes := getAnyTypes(actTx.Body.Messages)
			if assert.Equal(t, expTypes, actTypes, "ReadTxFileFlag Tx.Body.Messages types") {
				for i := range expTx.Body.Messages {
					expMsg := expTx.Body.Messages[i].GetCachedValue()
					actMsg := actTx.Body.Messages[i].GetCachedValue()
					assert.Equal(t, expMsg, actMsg, "ReadTxFileFlag Tx.Body.Messages[%d] cached value", i)
				}
			}
		})
	}
}

func TestReadProposalFlag(t *testing.T) {
	tests := []struct {
		name string
		// setup should return the proposal filename and the expected Anys.
		setup    func(t *testing.T) (string, []*codectypes.Any)
		flagSet  *pflag.FlagSet
		expInErr []string
	}{
		{
			name: "err getting flag",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				return "", nil
			},
			flagSet:  pflag.NewFlagSet("", pflag.ContinueOnError),
			expInErr: []string{"flag accessed but not defined: proposal"},
		},
		{
			name: "no flag given",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				return "", nil
			},
			expInErr: nil,
		},
		{
			name: "file does not exist",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				tdir := t.TempDir()
				noSuchFile := filepath.Join(tdir, "no-such-file.json")
				return noSuchFile, nil
			},
			expInErr: []string{"open ", "no-such-file.json", "no such file or directory"},
		},
		{
			name: "cannot unmarshal contents",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				tdir := t.TempDir()
				notJSON := filepath.Join(tdir, "not-json.json")
				contents := []byte("This is not\na JSON file.\n")
				writeFile(t, notJSON, contents)
				return notJSON, nil
			},
			expInErr: []string{
				"failed to unmarshal --proposal \"", "\" contents as Tx",
				"invalid character 'T' looking for beginning of value",
			},
		},
		{
			name: "no body",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				contents := `{
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    },
    "tip": null
  },
  "signatures": []
}
`
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-body.json")
				writeFile(t, fn, []byte(contents))
				return fn, nil
			},
			expInErr: []string{"the contents of \"", "\" does not have a \"body\""},
		},
		{
			name: "no body messages",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				tx := newTx(t)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-body-messages.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"the contents of \"", "\" does not have any body messages"},
		},
		{
			name: "no submit proposals",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				msg := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market: exchange.Market{
						MarketDetails: exchange.MarketDetails{
							Name: "New Market Name",
						},
					},
				}
				tx := newTx(t, msg)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-proposals.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"no *v1.MsgSubmitProposal messages found in \""},
		},
		{
			name: "no messages in submit proposals",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				prop := newGovProp(t)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-messages-in-proposal.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"no messages found in any *v1.MsgSubmitProposal messages in \""},
		},
		{
			name: "1 message found",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				msg := &exchange.MsgMarketWithdrawRequest{
					Admin:     sdk.AccAddress("Admin_______________").String(),
					MarketId:  3,
					ToAddress: sdk.AccAddress("ToAddress___________").String(),
					Amount:    sdk.NewCoins(sdk.NewInt64Coin("apple", 15)),
				}
				prop := newGovProp(t, msg)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "one-message.json")
				writeFileAsJson(t, fn, tx)
				return fn, prop.Messages
			},
			expInErr: nil,
		},
		{
			name: "3 messages found",
			setup: func(t *testing.T) (string, []*codectypes.Any) {
				msg1 := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market:    exchange.Market{MarketId: 88},
				}
				msg2 := &exchange.MsgGovManageFeesRequest{
					Authority:           cli.AuthorityAddr.String(),
					MarketId:            42,
					AddFeeCreateAskFlat: []sdk.Coin{sdk.NewInt64Coin("plum", 5)},
				}
				msg3 := &exchange.MsgCancelOrderRequest{
					Signer:  cli.AuthorityAddr.String(),
					OrderId: 5555,
				}
				prop1 := newGovProp(t, msg1)
				prop2 := newGovProp(t, msg2, msg3)
				tx := newTx(t, prop1, prop2)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "three-messages.json")
				writeFileAsJson(t, fn, tx)
				expAnys := make([]*codectypes.Any, 0, len(prop1.Messages)+len(prop2.Messages))
				expAnys = append(expAnys, prop1.Messages...)
				expAnys = append(expAnys, prop2.Messages...)
				return fn, expAnys
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			propFN, expAnys := tc.setup(t)

			if tc.flagSet == nil {
				tc.flagSet = pflag.NewFlagSet("", pflag.ContinueOnError)
				tc.flagSet.String(cli.FlagProposal, "", "The Proposal")
			}

			if len(propFN) > 0 && len(tc.expInErr) > 0 {
				tc.expInErr = append(tc.expInErr, propFN)
			}

			args := make([]string, 0, 2)
			if len(propFN) > 0 {
				args = append(args, "--"+cli.FlagProposal, propFN)
			}

			err := tc.flagSet.Parse(args)
			require.NoError(t, err, "flagSet.Parse(%q)", args)

			clientCtx := newClientContextWithCodec(t)

			var actPropFN string
			var actAnys []*codectypes.Any
			testFunc := func() {
				actPropFN, actAnys, err = cli.ReadProposalFlag(clientCtx, tc.flagSet)
			}
			require.NotPanics(t, testFunc, "ReadProposalFlag")

			assertions.AssertErrorContents(t, err, tc.expInErr, "ReadProposalFlag error")
			assert.Equal(t, propFN, actPropFN, "ReadProposalFlag filename")
			// We can't just assert that expAnys and actAnys are equal due to some internal differences.
			// All we really care about is that they have the same types and msg contents.
			expTypes := getAnyTypes(expAnys)
			actTypes := getAnyTypes(actAnys)
			if assert.Equal(t, expTypes, actTypes, "ReadProposalFlag anys types") {
				for i := range expAnys {
					expMsg := expAnys[i].GetCachedValue()
					actMsg := actAnys[i].GetCachedValue()
					assert.Equal(t, expMsg, actMsg, "ReadProposalFlag anys[%d] cached value", i)
				}
			}
		})
	}
}

func TestReadMsgGovCreateMarketRequestFromProposalFlag(t *testing.T) {
	tests := []struct {
		name string
		// setup should return the proposal filename and the expected Msg.
		setup    func(t *testing.T) (string, *exchange.MsgGovCreateMarketRequest)
		expInErr []string
	}{
		{
			name: "error reading file",
			setup: func(t *testing.T) (string, *exchange.MsgGovCreateMarketRequest) {
				tx := newTx(t)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-body-messages.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"the contents of \"", "\" does not have any body messages"},
		},
		{
			name: "no flag given",
			setup: func(t *testing.T) (string, *exchange.MsgGovCreateMarketRequest) {
				return "", nil
			},
			expInErr: nil,
		},
		{
			name: "no msgs of interest",
			setup: func(t *testing.T) (string, *exchange.MsgGovCreateMarketRequest) {
				msg := &exchange.MsgGovManageFeesRequest{
					Authority:           cli.AuthorityAddr.String(),
					MarketId:            13,
					AddFeeCreateAskFlat: []sdk.Coin{sdk.NewInt64Coin("cherry", 5000)},
				}
				prop := newGovProp(t, msg)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"no *exchange.MsgGovCreateMarketRequest found in \""},
		},
		{
			name: "two msgs of interest",
			setup: func(t *testing.T) (string, *exchange.MsgGovCreateMarketRequest) {
				msg1 := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market:    exchange.Market{MarketDetails: exchange.MarketDetails{Name: "Some Name"}},
				}
				msg2 := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market:    exchange.Market{MarketDetails: exchange.MarketDetails{Name: "Another Name"}},
				}
				prop := newGovProp(t, msg1, msg2)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"2 *exchange.MsgGovCreateMarketRequest found in \""},
		},
		{
			name: "one msg of interest",
			setup: func(t *testing.T) (string, *exchange.MsgGovCreateMarketRequest) {
				msg := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market:    exchange.Market{MarketDetails: exchange.MarketDetails{Name: "The Only Name"}},
				}
				prop := newGovProp(t, msg)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, msg
			},
			expInErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			propFN, expected := tc.setup(t)

			if expected == nil {
				expected = &exchange.MsgGovCreateMarketRequest{}
			}

			if len(propFN) > 0 && len(tc.expInErr) > 0 {
				tc.expInErr = append(tc.expInErr, propFN)
			}

			args := make([]string, 0, 2)
			if len(propFN) > 0 {
				args = append(args, "--"+cli.FlagProposal, propFN)
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(cli.FlagProposal, "", "The Proposal")
			err := flagSet.Parse(args)
			require.NoError(t, err, "flagSet.Parse(%q)", args)

			clientCtx := newClientContextWithCodec(t)

			var actual *exchange.MsgGovCreateMarketRequest
			testFunc := func() {
				actual, err = cli.ReadMsgGovCreateMarketRequestFromProposalFlag(clientCtx, flagSet)
			}
			require.NotPanics(t, testFunc, "ReadMsgGovCreateMarketRequestFromProposalFlag")
			assertions.AssertErrorContents(t, err, tc.expInErr, "ReadMsgGovCreateMarketRequestFromProposalFlag error")
			assert.Equal(t, expected, actual, "ReadMsgGovCreateMarketRequestFromProposalFlag result")
		})
	}
}

func TestReadMsgGovManageFeesRequestFromProposalFlag(t *testing.T) {
	tests := []struct {
		name string
		// setup should return the proposal filename and the expected Msg.
		setup    func(t *testing.T) (string, *exchange.MsgGovManageFeesRequest)
		expInErr []string
	}{
		{
			name: "error reading file",
			setup: func(t *testing.T) (string, *exchange.MsgGovManageFeesRequest) {
				tx := newTx(t)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-body-messages.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"the contents of \"", "\" does not have any body messages"},
		},
		{
			name: "no flag given",
			setup: func(t *testing.T) (string, *exchange.MsgGovManageFeesRequest) {
				return "", nil
			},
			expInErr: nil,
		},
		{
			name: "no msgs of interest",
			setup: func(t *testing.T) (string, *exchange.MsgGovManageFeesRequest) {
				msg := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market:    exchange.Market{MarketDetails: exchange.MarketDetails{Name: "Some Name"}},
				}
				prop := newGovProp(t, msg)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"no *exchange.MsgGovManageFeesRequest found in \""},
		},
		{
			name: "two msgs of interest",
			setup: func(t *testing.T) (string, *exchange.MsgGovManageFeesRequest) {
				msg1 := &exchange.MsgGovManageFeesRequest{
					Authority:              cli.AuthorityAddr.String(),
					MarketId:               12,
					RemoveFeeCreateAskFlat: []sdk.Coin{sdk.NewInt64Coin("banana", 99)},
				}
				msg2 := &exchange.MsgGovManageFeesRequest{
					Authority:           cli.AuthorityAddr.String(),
					MarketId:            13,
					AddFeeCreateAskFlat: []sdk.Coin{sdk.NewInt64Coin("cherry", 5000)},
				}
				prop := newGovProp(t, msg1, msg2)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"2 *exchange.MsgGovManageFeesRequest found in \""},
		},
		{
			name: "one msg of interest",
			setup: func(t *testing.T) (string, *exchange.MsgGovManageFeesRequest) {
				msg := &exchange.MsgGovManageFeesRequest{
					Authority:                  cli.AuthorityAddr.String(),
					MarketId:                   2,
					AddFeeSellerSettlementFlat: []sdk.Coin{sdk.NewInt64Coin("fig", 8)},
				}
				prop := newGovProp(t, msg)
				tx := newTx(t, prop)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, msg
			},
			expInErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			propFN, expected := tc.setup(t)

			if expected == nil {
				expected = &exchange.MsgGovManageFeesRequest{}
			}

			if len(propFN) > 0 && len(tc.expInErr) > 0 {
				tc.expInErr = append(tc.expInErr, propFN)
			}

			args := make([]string, 0, 2)
			if len(propFN) > 0 {
				args = append(args, "--"+cli.FlagProposal, propFN)
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(cli.FlagProposal, "", "The Proposal")
			err := flagSet.Parse(args)
			require.NoError(t, err, "flagSet.Parse(%q)", args)

			clientCtx := newClientContextWithCodec(t)

			var actual *exchange.MsgGovManageFeesRequest
			testFunc := func() {
				actual, err = cli.ReadMsgGovManageFeesRequestFromProposalFlag(clientCtx, flagSet)
			}
			require.NotPanics(t, testFunc, "ReadMsgGovManageFeesRequestFromProposalFlag")
			assertions.AssertErrorContents(t, err, tc.expInErr, "ReadMsgGovManageFeesRequestFromProposalFlag error")
			assert.Equal(t, expected, actual, "ReadMsgGovManageFeesRequestFromProposalFlag result")
		})
	}
}

func TestReadMsgMarketCommitmentSettleFromFileFlag(t *testing.T) {
	coins := func(t *testing.T, str string) sdk.Coins {
		rv, err := sdk.ParseCoinsNormalized(str)
		require.NoError(t, err, "ParseCoinsNormalized(%q)", str)
		return rv
	}

	tests := []struct {
		name string
		// setup should return the filename and the expected Msg.
		setup    func(t *testing.T) (string, *exchange.MsgMarketCommitmentSettleRequest)
		expInErr []string
	}{
		{
			name: "error reading file",
			setup: func(t *testing.T) (string, *exchange.MsgMarketCommitmentSettleRequest) {
				tx := newTx(t)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-body-messages.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"the contents of \"", "\" does not have any body messages"},
		},
		{
			name: "no flag given",
			setup: func(t *testing.T) (string, *exchange.MsgMarketCommitmentSettleRequest) {
				return "", nil
			},
			expInErr: nil,
		},
		{
			name: "no msgs of interest",
			setup: func(t *testing.T) (string, *exchange.MsgMarketCommitmentSettleRequest) {
				msg := &exchange.MsgGovCreateMarketRequest{
					Authority: cli.AuthorityAddr.String(),
					Market:    exchange.Market{MarketDetails: exchange.MarketDetails{Name: "Some Name"}},
				}
				tx := newTx(t, msg)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"no *exchange.MsgMarketCommitmentSettleRequest messages found in \""},
		},
		{
			name: "two msgs of interest",
			setup: func(t *testing.T) (string, *exchange.MsgMarketCommitmentSettleRequest) {
				msg1 := &exchange.MsgMarketCommitmentSettleRequest{
					Admin:    "msg1_admin_addr",
					MarketId: 1,
					Inputs: []exchange.AccountAmount{
						{Account: "msg1_input1_addr", Amount: coins(t, "10cherry")},
					},
					Outputs: []exchange.AccountAmount{
						{Account: "msg1_output1_addr", Amount: coins(t, "11cherry")},
					},
					Fees: []exchange.AccountAmount{
						{Account: "msg1_fee1_addr", Amount: coins(t, "3orange")},
					},
					Navs: []exchange.NetAssetPrice{
						{Assets: coins(t, "1apple")[0], Price: coins(t, "2banana")[0]},
					},
					EventTag: "msg1_event_tag",
				}
				msg2 := &exchange.MsgMarketCommitmentSettleRequest{
					Admin:    "msg2_admin_addr",
					MarketId: 2,
					Inputs: []exchange.AccountAmount{
						{Account: "msg2_input1_addr", Amount: coins(t, "10cherry,5plum")},
						{Account: "msg2_input2_addr", Amount: coins(t, "1apple,12orange")},
						{Account: "msg2_input3_addr", Amount: coins(t, "4apple,4cherry,4orange,4plum")},
					},
					Outputs: []exchange.AccountAmount{
						{Account: "msg2_output1_addr", Amount: coins(t, "5apple")},
						{Account: "msg2_output2_addr", Amount: coins(t, "14cherry")},
						{Account: "msg2_output3_addr", Amount: coins(t, "16orange")},
						{Account: "msg2_output5_addr", Amount: coins(t, "9plum")},
					},
					EventTag: "msg2_event_tag",
				}
				tx := newTx(t, msg1, msg2)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, nil
			},
			expInErr: []string{"2 *exchange.MsgMarketCommitmentSettleRequest found in \""},
		},
		{
			name: "one msg of interest",
			setup: func(t *testing.T) (string, *exchange.MsgMarketCommitmentSettleRequest) {
				msg := &exchange.MsgMarketCommitmentSettleRequest{
					Admin:    "msg_admin",
					MarketId: 1,
					Inputs: []exchange.AccountAmount{
						{Account: "msg_input1_addr", Amount: coins(t, "10cherry,5plum")},
						{Account: "msg_input2_addr", Amount: coins(t, "1apple,12orange")},
						{Account: "msg_input3_addr", Amount: coins(t, "4apple,4cherry,4orange,4plum")},
					},
					Outputs: []exchange.AccountAmount{
						{Account: "msg_output1_addr", Amount: coins(t, "5apple")},
						{Account: "msg_output2_addr", Amount: coins(t, "14cherry")},
						{Account: "msg_output3_addr", Amount: coins(t, "16orange")},
						{Account: "msg_output4_addr", Amount: coins(t, "9plum")},
					},
					Fees: []exchange.AccountAmount{
						{Account: "msg_fee1_addr", Amount: coins(t, "12fig")},
						{Account: "msg_fee2_addr", Amount: coins(t, "88fig")},
					},
					Navs: []exchange.NetAssetPrice{
						{Assets: coins(t, "1apple")[0], Price: coins(t, "2banana")[0]},
						{Assets: coins(t, "1apple")[0], Price: coins(t, "3fig")[0]},
						{Assets: coins(t, "10apple")[0], Price: coins(t, "1cherry")[0]},
					},
					EventTag: "just_some_event_tag",
				}
				tx := newTx(t, msg)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-messages-of-interest.json")
				writeFileAsJson(t, fn, tx)
				return fn, msg
			},
			expInErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filename, expected := tc.setup(t)

			if expected == nil {
				expected = &exchange.MsgMarketCommitmentSettleRequest{}
			}

			if len(filename) > 0 && len(tc.expInErr) > 0 {
				tc.expInErr = append(tc.expInErr, filename)
			}

			args := make([]string, 0, 2)
			if len(filename) > 0 {
				args = append(args, "--"+cli.FlagFile, filename)
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(cli.FlagFile, "", "The File")
			err := flagSet.Parse(args)
			require.NoError(t, err, "flagSet.Parse(%q)", args)

			clientCtx := newClientContextWithCodec(t)

			var actual *exchange.MsgMarketCommitmentSettleRequest
			testFunc := func() {
				actual, err = cli.ReadMsgMarketCommitmentSettleFromFileFlag(clientCtx, flagSet)
			}
			require.NotPanics(t, testFunc, "ReadMsgMarketCommitmentSettleFromFileFlag")
			assertions.AssertErrorContents(t, err, tc.expInErr, "ReadMsgMarketCommitmentSettleFromFileFlag error")
			assert.Equal(t, expected, actual, "ReadMsgMarketCommitmentSettleFromFileFlag result")
		})
	}
}

func TestReadPaymentFromFileFlag(t *testing.T) {
	coins := func(t *testing.T, str string) sdk.Coins {
		rv, err := sdk.ParseCoinsNormalized(str)
		require.NoError(t, err, "ParseCoinsNormalized(%q)", str)
		return rv
	}
	testAddr := func(prefix string) sdk.AccAddress {
		return sdk.AccAddress(prefix + strings.Repeat("_", 20-len(prefix)))
	}
	newPayment := func(t *testing.T, sourceAddrPrefix, sourceAmount, targetAddrPrefix, targetAmount, externalID string) exchange.Payment {
		return exchange.Payment{
			Source:       testAddr(sourceAddrPrefix).String(),
			SourceAmount: coins(t, sourceAmount),
			Target:       testAddr(targetAddrPrefix).String(),
			TargetAmount: coins(t, targetAmount),
			ExternalId:   externalID,
		}
	}

	tests := []struct {
		name string
		// setup should return the filename and the expected Payment.
		setup    func(t *testing.T) (string, exchange.Payment)
		expInErr []string
	}{
		{
			name: "no flag given",
			setup: func(t *testing.T) (string, exchange.Payment) {
				return "", exchange.Payment{}
			},
			expInErr: nil,
		},
		{
			name: "error reading file",
			setup: func(t *testing.T) (string, exchange.Payment) {
				tdir := t.TempDir()
				noSuchFile := filepath.Join(tdir, "no-such-file.json")
				return noSuchFile, exchange.Payment{}
			},
			expInErr: []string{"open ", "no-such-file.json", "no such file or directory"},
		},
		{
			name: "no msgs in tx",
			setup: func(t *testing.T) (string, exchange.Payment) {
				tx := newTx(t)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "no-body-messages.json")
				writeFileAsJson(t, fn, tx)
				return fn, exchange.Payment{}
			},
			expInErr: []string{"the contents of \"", "\" does not have any body messages"},
		},
		{
			name: "no msgs of interest",
			setup: func(t *testing.T) (string, exchange.Payment) {
				msg1 := &exchange.MsgRejectPaymentRequest{
					Target:     testAddr("target1").String(),
					Source:     testAddr("source1").String(),
					ExternalId: "something",
				}
				msg2 := &exchange.MsgCommitFundsRequest{
					Account:  testAddr("account1").String(),
					MarketId: 3,
					Amount:   coins(t, "10apple"),
				}
				tx := newTx(t, msg1, msg2)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-random-msgs.json")
				writeFileAsJson(t, fn, tx)
				return fn, exchange.Payment{}
			},
			expInErr: []string{"no messages with a Payment found"},
		},
		{
			name: "two msgs: both MsgCreatePaymentRequest",
			setup: func(t *testing.T) (string, exchange.Payment) {
				msg1 := &exchange.MsgCreatePaymentRequest{
					Payment: newPayment(t, "m1_source", "1strawberry", "m1_target", "10tomato", "msg1"),
				}
				msg2 := &exchange.MsgCreatePaymentRequest{
					Payment: newPayment(t, "m2_source", "2starfruit", "m2_target", "20tangerine", "msg2"),
				}
				tx := newTx(t, msg1, msg2)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-create-msgs.json")
				writeFileAsJson(t, fn, tx)
				return fn, exchange.Payment{}
			},
			expInErr: []string{"2 messages with a Payment found"},
		},
		{
			name: "two msgs: both MsgAcceptPaymentRequest",
			setup: func(t *testing.T) (string, exchange.Payment) {
				msg1 := &exchange.MsgAcceptPaymentRequest{
					Payment: newPayment(t, "m1_source", "1strawberry", "m1_target", "10tomato", "msg1"),
				}
				msg2 := &exchange.MsgAcceptPaymentRequest{
					Payment: newPayment(t, "m2_source", "2starfruit", "m2_target", "20tangerine", "msg2"),
				}
				tx := newTx(t, msg1, msg2)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-accept-msgs.json")
				writeFileAsJson(t, fn, tx)
				return fn, exchange.Payment{}
			},
			expInErr: []string{"2 messages with a Payment found"},
		},
		{
			name: "two msgs: MsgCreatePaymentRequest MsgAcceptPaymentRequest",
			setup: func(t *testing.T) (string, exchange.Payment) {
				msg1 := &exchange.MsgCreatePaymentRequest{
					Payment: newPayment(t, "m1_source", "1strawberry", "m1_target", "10tomato", "msg1"),
				}
				msg2 := &exchange.MsgAcceptPaymentRequest{
					Payment: newPayment(t, "m2_source", "2starfruit", "m2_target", "20tangerine", "msg2"),
				}
				tx := newTx(t, msg1, msg2)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-msgs-with-payments.json")
				writeFileAsJson(t, fn, tx)
				return fn, exchange.Payment{}
			},
			expInErr: []string{"2 messages with a Payment found"},
		},
		{
			name: "one msg: MsgCreatePaymentRequest",
			setup: func(t *testing.T) (string, exchange.Payment) {
				pmt := newPayment(t, "source_c", "11strawberry", "target_c", "52tangerine", "the-msg")
				msg1 := &exchange.MsgCreatePaymentRequest{Payment: pmt}
				tx := newTx(t, msg1)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "create-payment.json")
				writeFileAsJson(t, fn, tx)
				return fn, pmt
			},
		},
		{
			name: "one msg: MsgAcceptPaymentRequest",
			setup: func(t *testing.T) (string, exchange.Payment) {
				pmt := newPayment(t, "source", "1starfruit", "target", "5tomato", "the-other-msg")
				msg1 := &exchange.MsgAcceptPaymentRequest{Payment: pmt}
				tx := newTx(t, msg1)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "create-payment.json")
				writeFileAsJson(t, fn, tx)
				return fn, pmt
			},
		},
		{
			name: "three msgs in tx, only one has a payment",
			setup: func(t *testing.T) (string, exchange.Payment) {
				msg1 := &exchange.MsgCommitFundsRequest{
					Account:  testAddr("account1").String(),
					MarketId: 3,
					Amount:   coins(t, "10apple"),
				}
				pmt := newPayment(t, "msg2_source", "22strawberry", "msg2_target", "20tangerine", "msg-of-interest")
				msg2 := &exchange.MsgAcceptPaymentRequest{Payment: pmt}
				msg3 := &exchange.MsgRejectPaymentRequest{
					Target:     testAddr("target1").String(),
					Source:     testAddr("source1").String(),
					ExternalId: "something",
				}
				tx := newTx(t, msg1, msg2, msg3)
				tdir := t.TempDir()
				fn := filepath.Join(tdir, "two-random-msgs.json")
				writeFileAsJson(t, fn, tx)
				return fn, pmt
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filename, expected := tc.setup(t)

			if len(filename) > 0 && len(tc.expInErr) > 0 {
				tc.expInErr = append(tc.expInErr, filename)
			}

			args := make([]string, 0, 2)
			if len(filename) > 0 {
				args = append(args, "--"+cli.FlagFile, filename)
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(cli.FlagFile, "", "The File")
			err := flagSet.Parse(args)
			require.NoError(t, err, "flagSet.Parse(%q)", args)

			clientCtx := newClientContextWithCodec(t)

			var actual exchange.Payment
			testFunc := func() {
				actual, err = cli.ReadPaymentFromFileFlag(clientCtx, flagSet)
			}
			require.NotPanics(t, testFunc, "ReadPaymentFromFileFlag")
			assertions.AssertErrorContents(t, err, tc.expInErr, "ReadPaymentFromFileFlag error")
			assert.Equal(t, expected, actual, "ReadPaymentFromFileFlag result")
		})
	}
}

func TestReadFlagUint32OrDefault(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string // defaults to flagUint32.
		def      uint32
		exp      uint32
		expErr   string
	}{
		{
			testName: "error getting flag",
			flags:    []string{"--" + flagString, "what"},
			name:     flagString,
			def:      3,
			exp:      3,
			expErr:   "trying to get uint32 value of flag of type string",
		},
		{
			testName: "not provided, 0 default",
			def:      0,
			exp:      0,
		},
		{
			testName: "not provided, other default",
			def:      18,
			exp:      18,
		},
		{
			testName: "provided",
			flags:    []string{"--" + flagUint32, "43"},
			def:      100,
			exp:      43,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			if len(tc.name) == 0 {
				tc.name = flagUint32
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.Uint32(flagUint32, 0, "A uint32")
			flagSet.String(flagString, "", "A string")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var act uint32
			testFunc := func() {
				act, err = cli.ReadFlagUint32OrDefault(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagUint32OrDefault")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagUint32OrDefault error")
			assert.Equal(t, tc.exp, act, "ReadFlagUint32OrDefault result")
		})
	}
}

func TestReadFlagBoolOrDefault(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string // defaults to flagBool.
		def      bool
		exp      bool
		expErr   string
	}{
		{
			testName: "error getting flag, true default",
			flags:    []string{"--" + flagString, "what"},
			name:     flagString,
			def:      true,
			exp:      true,
			expErr:   "trying to get bool value of flag of type string",
		},
		{
			testName: "error getting flag, false default",
			flags:    []string{"--" + flagString, "what"},
			name:     flagString,
			def:      false,
			exp:      false,
			expErr:   "trying to get bool value of flag of type string",
		},
		{
			testName: "not provided, false default",
			def:      false,
			exp:      false,
		},
		{
			testName: "not provided, true default",
			def:      true,
			exp:      true,
		},
		{
			testName: "provided false, true default",
			flags:    []string{"--" + flagBool + "=false"},
			def:      true,
			exp:      false,
		},
		{
			testName: "provided false, false default",
			flags:    []string{"--" + flagBool + "=false"},
			def:      false,
			exp:      false,
		},
		{
			testName: "provided true, true default",
			flags:    []string{"--" + flagBool + "=true"},
			def:      true,
			exp:      true,
		},
		{
			testName: "provided true, false default",
			flags:    []string{"--" + flagBool + "=true"},
			def:      false,
			exp:      true,
		},
		{
			testName: "provided normal, false default",
			flags:    []string{"--" + flagBool},
			def:      false,
			exp:      true,
		},
		{
			testName: "provided normal, true default",
			flags:    []string{"--" + flagBool},
			def:      true,
			exp:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			if len(tc.name) == 0 {
				tc.name = flagBool
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.Bool(flagBool, false, "A bool")
			flagSet.String(flagString, "", "A string")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var act bool
			testFunc := func() {
				act, err = cli.ReadFlagBoolOrDefault(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagBoolOrDefault")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagBoolOrDefault error")
			assert.Equal(t, tc.exp, act, "ReadFlagBoolOrDefault result")
		})
	}
}

func TestReadFlagStringSliceOrDefault(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string // defaults to flagStringSlice.
		def      []string
		exp      []string
		expErr   string
	}{
		{
			testName: "error getting flag",
			flags:    []string{"--" + flagInt, "4"},
			name:     flagInt,
			def:      []string{"eight"},
			exp:      []string{"eight"},
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "not provided, nil default",
			def:      nil,
			exp:      nil,
		},
		{
			testName: "not provided, empty default",
			def:      []string{},
			exp:      []string{},
		},
		{
			testName: "not provided, other default",
			def:      []string{"one", "two", "three", "fourteen"},
			exp:      []string{"one", "two", "three", "fourteen"},
		},
		{
			testName: "provided",
			flags:    []string{"--" + flagStringSlice, "one", "--" + flagStringSlice, "two,three"},
			def:      []string{"seven"},
			exp:      []string{"one", "two", "three"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			if len(tc.name) == 0 {
				tc.name = flagStringSlice
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "Some strings")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var act []string
			testFunc := func() {
				act, err = cli.ReadFlagStringSliceOrDefault(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagStringSliceOrDefault")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagStringSliceOrDefault error")
			assert.Equal(t, tc.exp, act, "ReadFlagStringSliceOrDefault result")
		})
	}
}

func TestReadFlagStringOrDefault(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string // defaults to flagString.
		def      string
		exp      string
		expErr   string
	}{
		{
			testName: "error getting flag",
			flags:    []string{"--" + flagInt, "7"},
			name:     flagInt,
			def:      "what",
			exp:      "what",
			expErr:   "trying to get string value of flag of type int",
		},
		{
			testName: "not provided, empty default",
			def:      "",
			exp:      "",
		},
		{
			testName: "not provided, other default",
			def:      "other",
			exp:      "other",
		},
		{
			testName: "provided",
			flags:    []string{"--" + flagString, "yayaya"},
			def:      "thedefault",
			exp:      "yayaya",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			if len(tc.name) == 0 {
				tc.name = flagString
			}

			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.String(flagString, "", "A string")
			flagSet.Int(flagInt, 0, "A uint32")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var act string
			testFunc := func() {
				act, err = cli.ReadFlagStringOrDefault(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagStringOrDefault")
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagStringOrDefault error")
			assert.Equal(t, tc.exp, act, "ReadFlagStringOrDefault result")
		})
	}
}

func TestParseAccountAmount(t *testing.T) {
	tests := []struct {
		name   string
		val    string
		exp    *exchange.AccountAmount
		expErr string
	}{
		{
			name:   "empty",
			val:    "",
			expErr: "invalid account-amount \"\": expected format <account>:<amount>",
		},
		{
			name:   "no colons",
			val:    "banana",
			expErr: "invalid account-amount \"banana\": expected format <account>:<amount>",
		},
		{
			name:   "two colons",
			val:    "plum:8:123",
			expErr: "invalid account-amount \"plum:8:123\": expected format <account>:<amount>",
		},
		{
			name:   "empty account",
			val:    ":444",
			expErr: "invalid account-amount \":444\": both an <account> and <amount> are required",
		},
		{
			name:   "empty amount",
			val:    "apple:",
			expErr: "invalid account-amount \"apple:\": both an <account> and <amount> are required",
		},
		{
			name:   "invalid amount",
			val:    "apple:banana",
			expErr: "could not parse \"apple:banana\" amount: invalid coin expression: \"banana\"",
		},
		{
			name: "good, one coin",
			val:  "cherry:1apple",
			exp:  &exchange.AccountAmount{Account: "cherry", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 1))},
		},
		{
			name: "good, two coins",
			val:  "pear:1acorn,2beachnut",
			exp: &exchange.AccountAmount{
				Account: "pear",
				Amount:  sdk.NewCoins(sdk.NewInt64Coin("acorn", 1), sdk.NewInt64Coin("beachnut", 2))},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var actual *exchange.AccountAmount
			var err error
			testFunc := func() {
				actual, err = cli.ParseAccountAmount(tc.val)
			}
			require.NotPanics(t, testFunc, "ParseAccountAmount(%q)", tc.val)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseAccountAmount(%q) error", tc.val)
			if !assert.Equal(t, tc.exp, actual, "ParseAccountAmount(%q) result", tc.val) {
				t.Logf("Expected: %s:%s", tc.exp.Account, tc.exp.Amount)
				t.Logf("  Actual: %s:%s", actual.Account, actual.Amount)
			}
		})
	}
}

func TestParseAccountAmounts(t *testing.T) {
	tests := []struct {
		name   string
		vals   []string
		exp    []exchange.AccountAmount
		expErr string
	}{
		{
			name:   "nil",
			vals:   nil,
			expErr: "",
		},
		{
			name:   "empty",
			vals:   []string{},
			expErr: "",
		},
		{
			name:   "one, bad",
			vals:   []string{"nope"},
			expErr: "invalid account-amount \"nope\": expected format <account>:<amount>",
		},
		{
			name: "one, good",
			vals: []string{"yup:5cherry"},
			exp:  []exchange.AccountAmount{{Account: "yup", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 5))}},
		},
		{
			name: "three, all good",
			vals: []string{"first:1apple", "second:2banana", "third:3cherry"},
			exp: []exchange.AccountAmount{
				{Account: "first", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 1))},
				{Account: "second", Amount: sdk.NewCoins(sdk.NewInt64Coin("banana", 2))},
				{Account: "third", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 3))},
			},
		},
		{
			name: "three, first bad",
			vals: []string{"first", "second:2banana", "third:3cherry"},
			exp: []exchange.AccountAmount{
				{Account: "second", Amount: sdk.NewCoins(sdk.NewInt64Coin("banana", 2))},
				{Account: "third", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 3))},
			},
			expErr: "invalid account-amount \"first\": expected format <account>:<amount>",
		},
		{
			name: "three, second bad",
			vals: []string{"first:1apple", ":22", "third:3cherry"},
			exp: []exchange.AccountAmount{
				{Account: "first", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 1))},
				{Account: "third", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 3))},
			},
			expErr: "invalid account-amount \":22\": both an <account> and <amount> are required",
		},
		{
			name: "three, third bad",
			vals: []string{"first:1apple", "second:2banana", "third:333x"},
			exp: []exchange.AccountAmount{
				{Account: "first", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 1))},
				{Account: "second", Amount: sdk.NewCoins(sdk.NewInt64Coin("banana", 2))},
			},
			expErr: "could not parse \"third:333x\" amount: invalid coin expression: \"333x\"",
		},
		{
			name: "three, all bad",
			vals: []string{"first", ":22", "third:333x"},
			expErr: joinErrs(
				"invalid account-amount \"first\": expected format <account>:<amount>",
				"invalid account-amount \":22\": both an <account> and <amount> are required",
				"could not parse \"third:333x\" amount: invalid coin expression: \"333x\"",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.exp == nil {
				tc.exp = []exchange.AccountAmount{}
			}

			var actual []exchange.AccountAmount
			var err error
			testFunc := func() {
				actual, err = cli.ParseAccountAmounts(tc.vals)
			}
			require.NotPanics(t, testFunc, "ParseAccountAmounts(%q)", tc.vals)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseAccountAmounts(%q) error", tc.vals)
			assertEqualSlices(t, tc.exp, actual, exchange.AccountAmount.String, "ParseAccountAmounts(%q) result", tc.vals)
		})
	}
}

func TestReadFlagAccountAmounts(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		exp      []exchange.AccountAmount
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagStringSlice,
			expErr:   "",
		},
		{
			testName: "three vals, one bad",
			flags:    []string{"--" + flagStringSlice, "first:3apple,second:80", "--" + flagStringSlice, "third:777cherry,123durian"},
			name:     flagStringSlice,
			expErr:   "could not parse \"second:80\" amount: invalid coin expression: \"80\"",
		},
		{
			testName: "three vals, all good",
			flags:    []string{"--" + flagStringSlice, "first:3apple,second:80pear", "--" + flagStringSlice, "third:777cherry,123durian"},
			name:     flagStringSlice,
			exp: []exchange.AccountAmount{
				{Account: "first", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 3))},
				{Account: "second", Amount: sdk.NewCoins(sdk.NewInt64Coin("pear", 80))},
				{Account: "third", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 777), sdk.NewInt64Coin("durian", 123))},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var actual []exchange.AccountAmount
			testFunc := func() {
				actual, err = cli.ReadFlagAccountAmounts(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadFlagAccountAmounts(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagAccountAmounts(%q) error", tc.name)
			assertEqualSlices(t, tc.exp, actual, exchange.AccountAmount.String, "ReadFlagAccountAmounts(%q) result", tc.name)
		})
	}
}

func TestReadFlagAccountAmountsOrDefault(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		def      []exchange.AccountAmount
		exp      []exchange.AccountAmount
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			def:      []exchange.AccountAmount{{Account: "alex", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 4))}},
			exp:      []exchange.AccountAmount{{Account: "alex", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 4))}},
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			def:      []exchange.AccountAmount{{Account: "blake", Amount: sdk.NewCoins(sdk.NewInt64Coin("banana", 99))}},
			exp:      []exchange.AccountAmount{{Account: "blake", Amount: sdk.NewCoins(sdk.NewInt64Coin("banana", 99))}},
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagStringSlice,
			def:      []exchange.AccountAmount{{Account: "carter", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 21))}},
			exp:      []exchange.AccountAmount{{Account: "carter", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 21))}},
			expErr:   "",
		},
		{
			testName: "three vals, one bad",
			flags:    []string{"--" + flagStringSlice, "first:3apple,second:80", "--" + flagStringSlice, "third:777cherry,123durian"},
			name:     flagStringSlice,
			def:      []exchange.AccountAmount{},
			exp:      []exchange.AccountAmount{},
			expErr:   "could not parse \"second:80\" amount: invalid coin expression: \"80\"",
		},
		{
			testName: "three vals, all good",
			flags:    []string{"--" + flagStringSlice, "first:3apple,second:80pear", "--" + flagStringSlice, "third:777cherry,123durian"},
			name:     flagStringSlice,
			def:      []exchange.AccountAmount{{Account: "ellis", Amount: sdk.NewCoins(sdk.NewInt64Coin("elderberry", 5))}},
			exp: []exchange.AccountAmount{
				{Account: "first", Amount: sdk.NewCoins(sdk.NewInt64Coin("apple", 3))},
				{Account: "second", Amount: sdk.NewCoins(sdk.NewInt64Coin("pear", 80))},
				{Account: "third", Amount: sdk.NewCoins(sdk.NewInt64Coin("cherry", 777), sdk.NewInt64Coin("durian", 123))},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var actual []exchange.AccountAmount
			testFunc := func() {
				actual, err = cli.ReadFlagAccountAmountsOrDefault(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagAccountAmountsOrDefault(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagAccountAmountsOrDefault(%q) error", tc.name)
			assertEqualSlices(t, tc.exp, actual, exchange.AccountAmount.String, "ReadFlagAccountAmountsOrDefault(%q) result", tc.name)
		})
	}
}

func TestReadFlagAccountsWithoutAmounts(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		exp      []exchange.AccountAmount
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagStringSlice,
			expErr:   "",
		},
		{
			testName: "three vals",
			flags:    []string{"--" + flagStringSlice, "first,second", "--" + flagStringSlice, "third"},
			name:     flagStringSlice,
			exp: []exchange.AccountAmount{
				{Account: "first", Amount: nil},
				{Account: "second", Amount: nil},
				{Account: "third", Amount: nil},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var actual []exchange.AccountAmount
			testFunc := func() {
				actual, err = cli.ReadFlagAccountsWithoutAmounts(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadFlagAccountsWithoutAmounts(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagAccountsWithoutAmounts(%q) error", tc.name)
			assertEqualSlices(t, tc.exp, actual, exchange.AccountAmount.String, "ReadFlagAccountsWithoutAmounts(%q) result", tc.name)
		})
	}
}

func TestParseNetAssetPrice(t *testing.T) {
	tests := []struct {
		name   string
		val    string
		exp    *exchange.NetAssetPrice
		expErr string
	}{
		{
			name:   "empty",
			val:    "",
			expErr: "invalid net-asset-price \"\": expected format <assets>:<price>",
		},
		{
			name:   "no colons",
			val:    "banana",
			expErr: "invalid net-asset-price \"banana\": expected format <assets>:<price>",
		},
		{
			name:   "two colons",
			val:    "plum:8:123",
			expErr: "invalid net-asset-price \"plum:8:123\": expected format <assets>:<price>",
		},
		{
			name:   "empty assets",
			val:    ":444plum",
			expErr: "invalid net-asset-price \":444plum\": both an <assets> and <price> are required",
		},
		{
			name:   "empty price",
			val:    "2apple:",
			expErr: "invalid net-asset-price \"2apple:\": both an <assets> and <price> are required",
		},
		{
			name:   "invalid assets",
			val:    "apple:3plum",
			expErr: "could not parse \"apple:3plum\" assets: invalid coin expression: \"apple\"",
		},
		{
			name:   "invalid price",
			val:    "1apple:plum",
			expErr: "could not parse \"1apple:plum\" price: invalid coin expression: \"plum\"",
		},
		{
			name: "good",
			val:  "57apple:91plum",
			exp:  &exchange.NetAssetPrice{Assets: sdk.NewInt64Coin("apple", 57), Price: sdk.NewInt64Coin("plum", 91)},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var actual *exchange.NetAssetPrice
			var err error
			testFunc := func() {
				actual, err = cli.ParseNetAssetPrice(tc.val)
			}
			require.NotPanics(t, testFunc, "ParseNetAssetPrice(%q)", tc.val)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseNetAssetPrice(%q) error", tc.val)
			if !assert.Equal(t, tc.exp, actual, "ParseNetAssetPrice(%q) result", tc.val) {
				t.Logf("Expected: %s:%s", tc.exp.Assets, tc.exp.Price)
				t.Logf("  Actual: %s:%s", actual.Assets, actual.Price)
			}
		})
	}
}

func TestParseNetAssetPrices(t *testing.T) {
	tests := []struct {
		name   string
		vals   []string
		exp    []exchange.NetAssetPrice
		expErr string
	}{
		{
			name:   "nil",
			vals:   nil,
			expErr: "",
		},
		{
			name:   "empty",
			vals:   []string{},
			expErr: "",
		},
		{
			name:   "one, bad",
			vals:   []string{"nope"},
			expErr: "invalid net-asset-price \"nope\": expected format <assets>:<price>",
		},
		{
			name: "one, good",
			vals: []string{"3apple:5plum"},
			exp:  []exchange.NetAssetPrice{{Assets: sdk.NewInt64Coin("apple", 3), Price: sdk.NewInt64Coin("plum", 5)}},
		},
		{
			name: "three, all good",
			vals: []string{"1apple:2peach", "3acorn:4plum", "5acai:6pear"},
			exp: []exchange.NetAssetPrice{
				{Assets: sdk.NewInt64Coin("apple", 1), Price: sdk.NewInt64Coin("peach", 2)},
				{Assets: sdk.NewInt64Coin("acorn", 3), Price: sdk.NewInt64Coin("plum", 4)},
				{Assets: sdk.NewInt64Coin("acai", 5), Price: sdk.NewInt64Coin("pear", 6)},
			},
		},
		{
			name: "three, first bad",
			vals: []string{"first", "3acorn:4plum", "5acai:6pear"},
			exp: []exchange.NetAssetPrice{
				{Assets: sdk.NewInt64Coin("acorn", 3), Price: sdk.NewInt64Coin("plum", 4)},
				{Assets: sdk.NewInt64Coin("acai", 5), Price: sdk.NewInt64Coin("pear", 6)},
			},
			expErr: "invalid net-asset-price \"first\": expected format <assets>:<price>",
		},
		{
			name: "three, second bad",
			vals: []string{"1apple:2peach", "second", "5acai:6pear"},
			exp: []exchange.NetAssetPrice{
				{Assets: sdk.NewInt64Coin("apple", 1), Price: sdk.NewInt64Coin("peach", 2)},
				{Assets: sdk.NewInt64Coin("acai", 5), Price: sdk.NewInt64Coin("pear", 6)},
			},
			expErr: "invalid net-asset-price \"second\": expected format <assets>:<price>",
		},
		{
			name: "three, third bad",
			vals: []string{"1apple:2peach", "3acorn:4plum", "third"},
			exp: []exchange.NetAssetPrice{
				{Assets: sdk.NewInt64Coin("apple", 1), Price: sdk.NewInt64Coin("peach", 2)},
				{Assets: sdk.NewInt64Coin("acorn", 3), Price: sdk.NewInt64Coin("plum", 4)},
			},
			expErr: "invalid net-asset-price \"third\": expected format <assets>:<price>",
		},
		{
			name: "three, all bad",
			vals: []string{"first", "second", "third"},
			expErr: joinErrs(
				"invalid net-asset-price \"first\": expected format <assets>:<price>",
				"invalid net-asset-price \"second\": expected format <assets>:<price>",
				"invalid net-asset-price \"third\": expected format <assets>:<price>",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.exp == nil {
				tc.exp = []exchange.NetAssetPrice{}
			}

			var actual []exchange.NetAssetPrice
			var err error
			testFunc := func() {
				actual, err = cli.ParseNetAssetPrices(tc.vals)
			}
			require.NotPanics(t, testFunc, "ParseNetAssetPrices(%q)", tc.vals)
			assertions.AssertErrorValue(t, err, tc.expErr, "ParseNetAssetPrices(%q) error", tc.vals)
			assertEqualSlices(t, tc.exp, actual, exchange.NetAssetPrice.String, "ParseNetAssetPrices(%q) result", tc.vals)
		})
	}
}

func TestReadFlagNetAssetPrices(t *testing.T) {
	tests := []struct {
		testName string
		flags    []string
		name     string
		exp      []exchange.NetAssetPrice
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagStringSlice,
			expErr:   "",
		},
		{
			testName: "three vals, one bad",
			flags:    []string{"--" + flagStringSlice, "1apple:2peach,3acai", "--" + flagStringSlice, "5acorn:6pear"},
			name:     flagStringSlice,
			expErr:   "invalid net-asset-price \"3acai\": expected format <assets>:<price>",
		},
		{
			testName: "three vals, all good",
			flags:    []string{"--" + flagStringSlice, "1apple:2peach,3acai:4plum", "--" + flagStringSlice, "5acorn:6pear"},
			name:     flagStringSlice,
			exp: []exchange.NetAssetPrice{
				{Assets: sdk.NewInt64Coin("apple", 1), Price: sdk.NewInt64Coin("peach", 2)},
				{Assets: sdk.NewInt64Coin("acai", 3), Price: sdk.NewInt64Coin("plum", 4)},
				{Assets: sdk.NewInt64Coin("acorn", 5), Price: sdk.NewInt64Coin("pear", 6)},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var actual []exchange.NetAssetPrice
			testFunc := func() {
				actual, err = cli.ReadFlagNetAssetPrices(flagSet, tc.name)
			}
			require.NotPanics(t, testFunc, "ReadFlagNetAssetPrices(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagNetAssetPrices(%q) error", tc.name)
			assertEqualSlices(t, tc.exp, actual, exchange.NetAssetPrice.String, "ReadFlagNetAssetPrices(%q) result", tc.name)
		})
	}
}

func TestReadFlagNetAssetPricesOrDefault(t *testing.T) {
	coin := func(str string) sdk.Coin {
		rv, err := sdk.ParseCoinNormalized(str)
		require.NoError(t, err, "ParseCoinNormalized(%q)", str)
		return rv
	}
	tests := []struct {
		testName string
		flags    []string
		name     string
		def      []exchange.NetAssetPrice
		exp      []exchange.NetAssetPrice
		expErr   string
	}{
		{
			testName: "unknown flag",
			name:     "unknown",
			def:      []exchange.NetAssetPrice{{Assets: coin("10apple"), Price: coin("11banana")}},
			exp:      []exchange.NetAssetPrice{{Assets: coin("10apple"), Price: coin("11banana")}},
			expErr:   "flag accessed but not defined: unknown",
		},
		{
			testName: "wrong flag type",
			name:     flagInt,
			def:      []exchange.NetAssetPrice{{Assets: coin("3apple"), Price: coin("4banana")}},
			exp:      []exchange.NetAssetPrice{{Assets: coin("3apple"), Price: coin("4banana")}},
			expErr:   "trying to get stringSlice value of flag of type int",
		},
		{
			testName: "nothing provided",
			name:     flagStringSlice,
			def:      []exchange.NetAssetPrice{{Assets: coin("12fig"), Price: coin("3date")}},
			exp:      []exchange.NetAssetPrice{{Assets: coin("12fig"), Price: coin("3date")}},
			expErr:   "",
		},
		{
			testName: "three vals, one bad",
			flags:    []string{"--" + flagStringSlice, "1apple:2peach,3acai", "--" + flagStringSlice, "5acorn:6pear"},
			name:     flagStringSlice,
			def:      []exchange.NetAssetPrice{{Assets: coin("8fig"), Price: coin("7date")}},
			exp:      []exchange.NetAssetPrice{{Assets: coin("8fig"), Price: coin("7date")}},
			expErr:   "invalid net-asset-price \"3acai\": expected format <assets>:<price>",
		},
		{
			testName: "three vals, all good",
			flags:    []string{"--" + flagStringSlice, "1apple:2peach,3acai:4plum", "--" + flagStringSlice, "5acorn:6pear"},
			name:     flagStringSlice,
			def:      []exchange.NetAssetPrice{{Assets: coin("8fig"), Price: coin("7date")}},
			exp: []exchange.NetAssetPrice{
				{Assets: coin("1apple"), Price: coin("2peach")},
				{Assets: coin("3acai"), Price: coin("4plum")},
				{Assets: coin("5acorn"), Price: coin("6pear")},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			flagSet := pflag.NewFlagSet("", pflag.ContinueOnError)
			flagSet.StringSlice(flagStringSlice, nil, "A string slice")
			flagSet.Int(flagInt, 0, "An int")
			err := flagSet.Parse(tc.flags)
			require.NoError(t, err, "flagSet.Parse(%q)", tc.flags)

			var actual []exchange.NetAssetPrice
			testFunc := func() {
				actual, err = cli.ReadFlagNetAssetPricesOrDefault(flagSet, tc.name, tc.def)
			}
			require.NotPanics(t, testFunc, "ReadFlagNetAssetPricesOrDefault(%q)", tc.name)
			assertions.AssertErrorValue(t, err, tc.expErr, "ReadFlagNetAssetPricesOrDefault(%q) error", tc.name)
			assertEqualSlices(t, tc.exp, actual, exchange.NetAssetPrice.String, "ReadFlagNetAssetPricesOrDefault(%q) result", tc.name)
		})
	}
}
