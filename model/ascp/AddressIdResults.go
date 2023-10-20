package ascp

import (
	"sync"
)

// AddressIdResults 结构体
type AddressIdResults struct {
	// 错误的地址ID
	AddressId string `json:"address_id,omitempty" xml:"address_id,omitempty"`
	// 错误原因
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}

var poolAddressIdResults = sync.Pool{
	New: func() any {
		return new(AddressIdResults)
	},
}

// GetAddressIdResults() 从对象池中获取AddressIdResults
func GetAddressIdResults() *AddressIdResults {
	return poolAddressIdResults.Get().(*AddressIdResults)
}

// ReleaseAddressIdResults 释放AddressIdResults
func ReleaseAddressIdResults(v *AddressIdResults) {
	v.AddressId = ""
	v.ErrorMessage = ""
	poolAddressIdResults.Put(v)
}
