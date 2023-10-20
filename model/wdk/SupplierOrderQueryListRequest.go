package wdk

import (
	"sync"
)

// SupplierOrderQueryListRequest 结构体
type SupplierOrderQueryListRequest struct {
	// 淘宝订单号
	TbBizIdList []int64 `json:"tb_biz_id_list,omitempty" xml:"tb_biz_id_list>int64,omitempty"`
	// 盒马订单号
	BizIdList []int64 `json:"biz_id_list,omitempty" xml:"biz_id_list>int64,omitempty"`
	// 商场code
	SourceMerchantCode string `json:"source_merchant_code,omitempty" xml:"source_merchant_code,omitempty"`
	// 订单渠道来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}

var poolSupplierOrderQueryListRequest = sync.Pool{
	New: func() any {
		return new(SupplierOrderQueryListRequest)
	},
}

// GetSupplierOrderQueryListRequest() 从对象池中获取SupplierOrderQueryListRequest
func GetSupplierOrderQueryListRequest() *SupplierOrderQueryListRequest {
	return poolSupplierOrderQueryListRequest.Get().(*SupplierOrderQueryListRequest)
}

// ReleaseSupplierOrderQueryListRequest 释放SupplierOrderQueryListRequest
func ReleaseSupplierOrderQueryListRequest(v *SupplierOrderQueryListRequest) {
	v.TbBizIdList = v.TbBizIdList[:0]
	v.BizIdList = v.BizIdList[:0]
	v.SourceMerchantCode = ""
	v.OrderFrom = 0
	poolSupplierOrderQueryListRequest.Put(v)
}
