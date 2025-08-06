package dxo

import "github.com/starter-go/libgorm"

// DataGroupInfo ...
type DataGroupInfo interface {
	Prototypes() []any

	SetTableNamePrefix(prefix string)
}

// DatabaseAgent ...
type DatabaseAgent interface {
	libgorm.Agent
}
