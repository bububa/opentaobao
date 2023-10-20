package tmallnr

import (
	"sync"
)

// ResultData 结构体
type ResultData struct {
	// 服务范围，支持多服务范围返回
	Ranges []NrServiceRangeResponseDto `json:"ranges,omitempty" xml:"ranges>nr_service_range_response_dto,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolResultData = sync.Pool{
	New: func() any {
		return new(ResultData)
	},
}

// GetResultData() 从对象池中获取ResultData
func GetResultData() *ResultData {
	return poolResultData.Get().(*ResultData)
}

// ReleaseResultData 释放ResultData
func ReleaseResultData(v *ResultData) {
	v.Ranges = v.Ranges[:0]
	v.SellerId = 0
	v.StoreId = 0
	poolResultData.Put(v)
}
