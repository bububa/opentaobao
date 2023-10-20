package omniorder

import (
	"sync"
)

// CommissionResultQuery 结构体
type CommissionResultQuery struct {
	// 分佣导购
	CommissionEmployeeName string `json:"commission_employee_name,omitempty" xml:"commission_employee_name,omitempty"`
	// 分佣导购所属门店
	CommissionStoreId string `json:"commission_store_id,omitempty" xml:"commission_store_id,omitempty"`
	// 订单支付的开始时间
	OrderPayTimeStart string `json:"order_pay_time_start,omitempty" xml:"order_pay_time_start,omitempty"`
	// 分佣的结束时间
	CommissionTimeEnd string `json:"commission_time_end,omitempty" xml:"commission_time_end,omitempty"`
	// 分佣的开始时间
	CommissionTimeStart string `json:"commission_time_start,omitempty" xml:"commission_time_start,omitempty"`
	// 订单支付的结束时间
	OrderPayTimeEnd string `json:"order_pay_time_end,omitempty" xml:"order_pay_time_end,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 订单类型，1:旗舰店，23：线下门店
	BizOrderType int64 `json:"biz_order_type,omitempty" xml:"biz_order_type,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
}

var poolCommissionResultQuery = sync.Pool{
	New: func() any {
		return new(CommissionResultQuery)
	},
}

// GetCommissionResultQuery() 从对象池中获取CommissionResultQuery
func GetCommissionResultQuery() *CommissionResultQuery {
	return poolCommissionResultQuery.Get().(*CommissionResultQuery)
}

// ReleaseCommissionResultQuery 释放CommissionResultQuery
func ReleaseCommissionResultQuery(v *CommissionResultQuery) {
	v.CommissionEmployeeName = ""
	v.CommissionStoreId = ""
	v.OrderPayTimeStart = ""
	v.CommissionTimeEnd = ""
	v.CommissionTimeStart = ""
	v.OrderPayTimeEnd = ""
	v.PageSize = 0
	v.BizOrderType = 0
	v.PageNo = 0
	poolCommissionResultQuery.Put(v)
}
