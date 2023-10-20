package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyOrderListQueryResponse 结构体
type AlitripMerchantGalaxyOrderListQueryResponse struct {
	// 查询结果
	OrderDtos []OrderDto `json:"order_dtos,omitempty" xml:"order_dtos>order_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 每页的数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页的数量
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 总页数
	TotalPageNum int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否有下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
}

var poolAlitripMerchantGalaxyOrderListQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyOrderListQueryResponse)
	},
}

// GetAlitripMerchantGalaxyOrderListQueryResponse() 从对象池中获取AlitripMerchantGalaxyOrderListQueryResponse
func GetAlitripMerchantGalaxyOrderListQueryResponse() *AlitripMerchantGalaxyOrderListQueryResponse {
	return poolAlitripMerchantGalaxyOrderListQueryResponse.Get().(*AlitripMerchantGalaxyOrderListQueryResponse)
}

// ReleaseAlitripMerchantGalaxyOrderListQueryResponse 释放AlitripMerchantGalaxyOrderListQueryResponse
func ReleaseAlitripMerchantGalaxyOrderListQueryResponse(v *AlitripMerchantGalaxyOrderListQueryResponse) {
	v.OrderDtos = v.OrderDtos[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageSize = 0
	v.Size = 0
	v.PageNo = 0
	v.TotalPageNum = 0
	v.TotalCount = 0
	v.Success = false
	v.HasNextPage = false
	poolAlitripMerchantGalaxyOrderListQueryResponse.Put(v)
}
