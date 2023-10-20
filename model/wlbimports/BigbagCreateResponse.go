package wlbimports

import (
	"sync"
)

// BigbagCreateResponse 结构体
type BigbagCreateResponse struct {
	// 大包Code
	BigbagCode string `json:"bigbag_code,omitempty" xml:"bigbag_code,omitempty"`
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
}

var poolBigbagCreateResponse = sync.Pool{
	New: func() any {
		return new(BigbagCreateResponse)
	},
}

// GetBigbagCreateResponse() 从对象池中获取BigbagCreateResponse
func GetBigbagCreateResponse() *BigbagCreateResponse {
	return poolBigbagCreateResponse.Get().(*BigbagCreateResponse)
}

// ReleaseBigbagCreateResponse 释放BigbagCreateResponse
func ReleaseBigbagCreateResponse(v *BigbagCreateResponse) {
	v.BigbagCode = ""
	v.BigbagId = 0
	poolBigbagCreateResponse.Put(v)
}
