package vo

import "github.com/starter-go/dns-server-starter/app/web/dto"

// Example ... VO
type Example struct {
	Base

	Items []*dto.Example `json:"items"`
}
