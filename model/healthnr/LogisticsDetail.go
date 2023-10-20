package healthnr

import (
	"sync"
)

// LogisticsDetail 结构体
type LogisticsDetail struct {
	// 物流商
	LogisticsName string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
	// 配送单下单时间
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 配送单确认时间
	ConfirmTime string `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 配送单取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 骑手取货时间
	FetchTime string `json:"fetch_time,omitempty" xml:"fetch_time,omitempty"`
	// 配送完成时间
	CompleteTime string `json:"complete_time,omitempty" xml:"complete_time,omitempty"`
	// 骑手姓名
	DispatcherName string `json:"dispatcher_name,omitempty" xml:"dispatcher_name,omitempty"`
	// 骑手电话
	DispatcherMobile string `json:"dispatcher_mobile,omitempty" xml:"dispatcher_mobile,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 物流状态
	LogisticsStatus int64 `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
}

var poolLogisticsDetail = sync.Pool{
	New: func() any {
		return new(LogisticsDetail)
	},
}

// GetLogisticsDetail() 从对象池中获取LogisticsDetail
func GetLogisticsDetail() *LogisticsDetail {
	return poolLogisticsDetail.Get().(*LogisticsDetail)
}

// ReleaseLogisticsDetail 释放LogisticsDetail
func ReleaseLogisticsDetail(v *LogisticsDetail) {
	v.LogisticsName = ""
	v.SendTime = ""
	v.ConfirmTime = ""
	v.CancelTime = ""
	v.FetchTime = ""
	v.CompleteTime = ""
	v.DispatcherName = ""
	v.DispatcherMobile = ""
	v.OrderId = 0
	v.LogisticsStatus = 0
	poolLogisticsDetail.Put(v)
}
