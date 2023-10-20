package ascp

import (
	"sync"
)

// AlibabaDchainAoxiangWmsDeliveryorderCreateDetail 结构体
type AlibabaDchainAoxiangWmsDeliveryorderCreateDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolAlibabaDchainAoxiangWmsDeliveryorderCreateDetail = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangWmsDeliveryorderCreateDetail)
	},
}

// GetAlibabaDchainAoxiangWmsDeliveryorderCreateDetail() 从对象池中获取AlibabaDchainAoxiangWmsDeliveryorderCreateDetail
func GetAlibabaDchainAoxiangWmsDeliveryorderCreateDetail() *AlibabaDchainAoxiangWmsDeliveryorderCreateDetail {
	return poolAlibabaDchainAoxiangWmsDeliveryorderCreateDetail.Get().(*AlibabaDchainAoxiangWmsDeliveryorderCreateDetail)
}

// ReleaseAlibabaDchainAoxiangWmsDeliveryorderCreateDetail 释放AlibabaDchainAoxiangWmsDeliveryorderCreateDetail
func ReleaseAlibabaDchainAoxiangWmsDeliveryorderCreateDetail(v *AlibabaDchainAoxiangWmsDeliveryorderCreateDetail) {
	v.Items = v.Items[:0]
	poolAlibabaDchainAoxiangWmsDeliveryorderCreateDetail.Put(v)
}
