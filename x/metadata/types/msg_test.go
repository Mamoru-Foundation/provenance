package types

import (
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddScopeRoute(t *testing.T) {
	var scope = NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		[]string{"data_owner"},
		[]string{"data_accessor"},
		"value_owner",
	)
	var msg = NewMsgAddScopeRequest(scope)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_scope_request")
	yaml := `scope:
  scope_id: scope1qzxcpvj6czy5g354dews3nlruxjsahhnsp
  specification_id: scopespec1qs30c9axgrw5669ft0kffe6h9gysfe58v3
  owneraddress:
  - data_owner
  dataaccess:
  - data_accessor
  valueowneraddress: value_owner
`
	require.Equal(t, yaml, msg.String())

	require.Equal(t, "{\"type\":\"provenance/metadata/AddScopeRequest\",\"value\":{\"scope\":{\"data_access\":[\"data_accessor\"],\"owner_address\":[\"data_owner\"],\"scope_id\":\"scope1qzxcpvj6czy5g354dews3nlruxjsahhnsp\",\"specification_id\":\"scopespec1qs30c9axgrw5669ft0kffe6h9gysfe58v3\",\"value_owner_address\":\"value_owner\"}}}", string(msg.GetSignBytes()))
}

func TestAddScopeValidation(t *testing.T) {
	var scope = NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		[]string{"data_owner"},
		[]string{"data_accessor"},
		"value_owner",
	)
	var msg = NewMsgAddScopeRequest(scope)
	err := msg.ValidateBasic()
	require.Panics(t, func() { msg.GetSigners() }, "panics due to invalid addresses")
	require.Error(t, err, "invalid addresses")
	require.Equal(t, "invalid owner on scope: decoding bech32 failed: invalid index of 1", err.Error())

	msg.Scope = NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		[]string{},
		[]string{},
		"",
	)
	err = msg.ValidateBasic()
	require.Error(t, err, "no owners")
	require.Equal(t, "scope must have at least one owner", err.Error())

	msg.Scope = NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		[]string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"},
		[]string{},
		"",
	)
	err = msg.ValidateBasic()
	require.NoError(t, err, "valid add scope request")
	requiredSigners := msg.GetSigners()
	require.Equal(t, 1, len(requiredSigners))
	hex, err := hex.DecodeString("85EA54E8598B27EC37EAEEEEA44F1E78A9B5E671")
	require.NoError(t, err)
	require.Equal(t, sdk.AccAddress(hex), requiredSigners[0])
}
