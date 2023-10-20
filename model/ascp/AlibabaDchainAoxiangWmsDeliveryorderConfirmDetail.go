package ascp

import (
	"sync"
)

// AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail 结构体
type AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail struct {
	// 商品
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolAlibabaDchainAoxiangWmsDeliveryorderConfirmDetail = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail)
	},
}

// GetAlibabaDchainAoxiangWmsDeliveryorderConfirmDetail() 从对象池中获取AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail
func GetAlibabaDchainAoxiangWmsDeliveryorderConfirmDetail() *AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail {
	return poolAlibabaDchainAoxiangWmsDeliveryorderConfirmDetail.Get().(*AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail)
}

// ReleaseAlibabaDchainAoxiangWmsDeliveryorderConfirmDetail 释放AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail
func ReleaseAlibabaDchainAoxiangWmsDeliveryorderConfirmDetail(v *AlibabaDchainAoxiangWmsDeliveryorderConfirmDetail) {
	v.Items = v.Items[:0]
	poolAlibabaDchainAoxiangWmsDeliveryorderConfirmDetail.Put(v)
}
