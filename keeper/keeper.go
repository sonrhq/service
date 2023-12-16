package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sonrhq/service"
	modulev1 "github.com/sonrhq/service/api/module/v1"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec
	db           modulev1.StateStore

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// referenced keepers
	identityKeeper service.IdentityKeeper
	bankKeeper     service.BankKeeper
	groupKeeper    service.GroupKeeper

	// state management
	ormtable.Schema
	CollSchema collections.Schema
	Params     collections.Item[service.Params]
	Counter    collections.Map[string, uint64]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, identityKeeper service.IdentityKeeper,
	bankKeeper service.BankKeeper, groupKeeper service.GroupKeeper,
	authority string) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}
	db, err := ormdb.NewModuleDB(serviceSchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}

	store, err := modulev1.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:            cdc,
		addressCodec:   addressCodec,
		authority:      authority,
		Params:         collections.NewItem(sb, service.ParamsKey, "params", codec.CollValue[service.Params](cdc)),
		Counter:        collections.NewMap(sb, service.CounterKey, "counter", collections.StringKey, collections.Uint64Value),
		db:             store,
		identityKeeper: identityKeeper,
		bankKeeper:     bankKeeper,
		groupKeeper:    groupKeeper,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.CollSchema = schema
	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}
