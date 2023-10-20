package omniorder

import (
	"sync"
)

// StoreConsignedResponse 结构体
type StoreConsignedResponse struct {
	// shortId
	ShortId string `json:"short_id,omitempty" xml:"short_id,omitempty"`
	// mailNo
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// gotCode
	GotCode string `json:"got_code,omitempty" xml:"got_code,omitempty"`
}

var poolStoreConsignedResponse = sync.Pool{
	New: func() any {
		return new(StoreConsignedResponse)
	},
}

// GetStoreConsignedResponse() 从对象池中获取StoreConsignedResponse
func GetStoreConsignedResponse() *StoreConsignedResponse {
	return poolStoreConsignedResponse.Get().(*StoreConsignedResponse)
}

// ReleaseStoreConsignedResponse 释放StoreConsignedResponse
func ReleaseStoreConsignedResponse(v *StoreConsignedResponse) {
	v.ShortId = ""
	v.MailNo = ""
	v.GotCode = ""
	poolStoreConsignedResponse.Put(v)
}
