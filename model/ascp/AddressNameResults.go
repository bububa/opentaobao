package ascp

import (
	"sync"
)

// AddressNameResults 结构体
type AddressNameResults struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 错误原因
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}

var poolAddressNameResults = sync.Pool{
	New: func() any {
		return new(AddressNameResults)
	},
}

// GetAddressNameResults() 从对象池中获取AddressNameResults
func GetAddressNameResults() *AddressNameResults {
	return poolAddressNameResults.Get().(*AddressNameResults)
}

// ReleaseAddressNameResults 释放AddressNameResults
func ReleaseAddressNameResults(v *AddressNameResults) {
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.ErrorMessage = ""
	poolAddressNameResults.Put(v)
}
