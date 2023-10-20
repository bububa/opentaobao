package trade

import (
	"sync"
)

// CainiaoCntecSupplierOrderServiceResult 结构体
type CainiaoCntecSupplierOrderServiceResult struct {
	// 订单列表
	OrderList []SupplierOrder `json:"order_list,omitempty" xml:"order_list>supplier_order,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 分页游标
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 是否有下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoCntecSupplierOrderServiceResult = sync.Pool{
	New: func() any {
		return new(CainiaoCntecSupplierOrderServiceResult)
	},
}

// GetCainiaoCntecSupplierOrderServiceResult() 从对象池中获取CainiaoCntecSupplierOrderServiceResult
func GetCainiaoCntecSupplierOrderServiceResult() *CainiaoCntecSupplierOrderServiceResult {
	return poolCainiaoCntecSupplierOrderServiceResult.Get().(*CainiaoCntecSupplierOrderServiceResult)
}

// ReleaseCainiaoCntecSupplierOrderServiceResult 释放CainiaoCntecSupplierOrderServiceResult
func ReleaseCainiaoCntecSupplierOrderServiceResult(v *CainiaoCntecSupplierOrderServiceResult) {
	v.OrderList = v.OrderList[:0]
	v.ErrorMsg = ""
	v.ErrCode = ""
	v.PageSize = 0
	v.TotalCount = 0
	v.PageIndex = 0
	v.HasNextPage = false
	v.Success = false
	poolCainiaoCntecSupplierOrderServiceResult.Put(v)
}
