package omniorder

import (
	"sync"
)

// PageResult 结构体
type PageResult struct {
	// 返回数据
	OrderList []OrderDto `json:"order_list,omitempty" xml:"order_list>order_dto,omitempty"`
	// 订单列表
	Datas []TaobaoOmniDealerOdersListData `json:"datas,omitempty" xml:"datas>taobao_omni_dealer_oders_list_data,omitempty"`
	// 页码，请求的值
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码，200-成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 页大小，请求的值
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页多少条记录
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码
	CurrPage int64 `json:"curr_page,omitempty" xml:"curr_page,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPageResult = sync.Pool{
	New: func() any {
		return new(PageResult)
	},
}

// GetPageResult() 从对象池中获取PageResult
func GetPageResult() *PageResult {
	return poolPageResult.Get().(*PageResult)
}

// ReleasePageResult 释放PageResult
func ReleasePageResult(v *PageResult) {
	v.OrderList = v.OrderList[:0]
	v.Datas = v.Datas[:0]
	v.Message = ""
	v.Code = ""
	v.Total = 0
	v.PageNo = 0
	v.PageSize = 0
	v.CurrPage = 0
	v.TotalCount = 0
	v.Success = false
	poolPageResult.Put(v)
}
