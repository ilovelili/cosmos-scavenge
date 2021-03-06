package types

const (
	// ModuleName is the name of the module
	ModuleName = "scavenge"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for queriers
	QuerierRoute = ModuleName

	ScavengePrefix = "sk-"
	CommitPrefix   = "ck-"
)
