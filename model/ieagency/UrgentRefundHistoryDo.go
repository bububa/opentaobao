package ieagency

import (
	"sync"
)

// UrgentRefundHistoryDo 结构体
type UrgentRefundHistoryDo struct {
	// 催退时间
	RequestDate string `json:"request_date,omitempty" xml:"request_date,omitempty"`
	// 要求商家处理时间
	UrgentToSellerFinishTime string `json:"urgent_to_seller_finish_time,omitempty" xml:"urgent_to_seller_finish_time,omitempty"`
	// 平台承若用户完成时间
	UrgentToBuyerPromiseTime string `json:"urgent_to_buyer_promise_time,omitempty" xml:"urgent_to_buyer_promise_time,omitempty"`
	// 第num次催退
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}

var poolUrgentRefundHistoryDo = sync.Pool{
	New: func() any {
		return new(UrgentRefundHistoryDo)
	},
}

// GetUrgentRefundHistoryDo() 从对象池中获取UrgentRefundHistoryDo
func GetUrgentRefundHistoryDo() *UrgentRefundHistoryDo {
	return poolUrgentRefundHistoryDo.Get().(*UrgentRefundHistoryDo)
}

// ReleaseUrgentRefundHistoryDo 释放UrgentRefundHistoryDo
func ReleaseUrgentRefundHistoryDo(v *UrgentRefundHistoryDo) {
	v.RequestDate = ""
	v.UrgentToSellerFinishTime = ""
	v.UrgentToBuyerPromiseTime = ""
	v.Num = 0
	poolUrgentRefundHistoryDo.Put(v)
}
