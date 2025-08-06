package dto

import "github.com/starter-go/dns-server-starter/app/data/dxo"

// Example ... DTO
type Example struct {
	ID dxo.ExampleID `json:"id"`
	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
