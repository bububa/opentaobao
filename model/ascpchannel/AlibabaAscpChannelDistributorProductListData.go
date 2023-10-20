package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelDistributorProductListData 结构体
type AlibabaAscpChannelDistributorProductListData struct {
	// 产品列表
	Products []Products `json:"products,omitempty" xml:"products>products,omitempty"`
	// 返回值总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolAlibabaAscpChannelDistributorProductListData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelDistributorProductListData)
	},
}

// GetAlibabaAscpChannelDistributorProductListData() 从对象池中获取AlibabaAscpChannelDistributorProductListData
func GetAlibabaAscpChannelDistributorProductListData() *AlibabaAscpChannelDistributorProductListData {
	return poolAlibabaAscpChannelDistributorProductListData.Get().(*AlibabaAscpChannelDistributorProductListData)
}

// ReleaseAlibabaAscpChannelDistributorProductListData 释放AlibabaAscpChannelDistributorProductListData
func ReleaseAlibabaAscpChannelDistributorProductListData(v *AlibabaAscpChannelDistributorProductListData) {
	v.Products = v.Products[:0]
	v.TotalCount = 0
	poolAlibabaAscpChannelDistributorProductListData.Put(v)
}
