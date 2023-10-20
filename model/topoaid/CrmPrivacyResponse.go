package topoaid

import (
	"sync"
)

// CrmPrivacyResponse 结构体
type CrmPrivacyResponse struct {
	// omid
	Omid string `json:"omid,omitempty" xml:"omid,omitempty"`
	// ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
}

var poolCrmPrivacyResponse = sync.Pool{
	New: func() any {
		return new(CrmPrivacyResponse)
	},
}

// GetCrmPrivacyResponse() 从对象池中获取CrmPrivacyResponse
func GetCrmPrivacyResponse() *CrmPrivacyResponse {
	return poolCrmPrivacyResponse.Get().(*CrmPrivacyResponse)
}

// ReleaseCrmPrivacyResponse 释放CrmPrivacyResponse
func ReleaseCrmPrivacyResponse(v *CrmPrivacyResponse) {
	v.Omid = ""
	v.Ouid = ""
	poolCrmPrivacyResponse.Put(v)
}
