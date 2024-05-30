package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ResourceList: []Resource{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in resource
	resourceIndexMap := make(map[string]struct{})

	for _, elem := range gs.ResourceList {
		index := string(ResourceKey(elem.Index))
		if _, ok := resourceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for resource")
		}
		resourceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
