package entity

import "github.com/starter-go/dns-server-starter/app/data/dxo"

// GetDataGroupInfo  ...
func GetDataGroupInfo() dxo.DataGroupInfo {
	return new(dgInfo)
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string

type dgInfo struct{}

func (inst *dgInfo) Prototypes() []any {

	list := make([]any, 0)

	list = append(list, new(Example))

	return list
}

func (inst *dgInfo) SetTableNamePrefix(prefix string) {
	if theTableNamePrefix == "" {
		theTableNamePrefix = prefix
	}
}

////////////////////////////////////////////////////////////////////////////////

// TableName 。。。
func (Example) TableName() string {
	return theTableNamePrefix + "example"
}
