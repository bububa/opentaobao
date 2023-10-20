package pur

import (
	"sync"
)

// MallReceivePrResponseData 结构体
type MallReceivePrResponseData struct {
	// 采购商城申请单ID
	PurReqId string `json:"pur_req_id,omitempty" xml:"pur_req_id,omitempty"`
	// 下单成功后跳转地址
	RedirectUrl string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
}

var poolMallReceivePrResponseData = sync.Pool{
	New: func() any {
		return new(MallReceivePrResponseData)
	},
}

// GetMallReceivePrResponseData() 从对象池中获取MallReceivePrResponseData
func GetMallReceivePrResponseData() *MallReceivePrResponseData {
	return poolMallReceivePrResponseData.Get().(*MallReceivePrResponseData)
}

// ReleaseMallReceivePrResponseData 释放MallReceivePrResponseData
func ReleaseMallReceivePrResponseData(v *MallReceivePrResponseData) {
	v.PurReqId = ""
	v.RedirectUrl = ""
	poolMallReceivePrResponseData.Put(v)
}
