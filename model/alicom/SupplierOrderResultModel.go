package alicom

import (
	"sync"
)

// SupplierOrderResultModel 结构体
type SupplierOrderResultModel struct {
	// 业务类型：7-合约机分销、
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 结果码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 交易订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 供应商外部订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 订购结果状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolSupplierOrderResultModel = sync.Pool{
	New: func() any {
		return new(SupplierOrderResultModel)
	},
}

// GetSupplierOrderResultModel() 从对象池中获取SupplierOrderResultModel
func GetSupplierOrderResultModel() *SupplierOrderResultModel {
	return poolSupplierOrderResultModel.Get().(*SupplierOrderResultModel)
}

// ReleaseSupplierOrderResultModel 释放SupplierOrderResultModel
func ReleaseSupplierOrderResultModel(v *SupplierOrderResultModel) {
	v.BizType = ""
	v.Code = ""
	v.Desc = ""
	v.OrderNo = ""
	v.OutOrderNo = ""
	v.Success = false
	poolSupplierOrderResultModel.Put(v)
}
