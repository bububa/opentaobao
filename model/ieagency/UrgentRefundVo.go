package ieagency

import (
	"sync"
)

// UrgentRefundVo 结构体
type UrgentRefundVo struct {
	// 催退历史记录
	UrgentRefundHistory []UrgentRefundHistoryDo `json:"urgent_refund_history,omitempty" xml:"urgent_refund_history>urgent_refund_history_do,omitempty"`
	// 最新要求商家处理时间
	LatestUrgentToSellerFinishTime string `json:"latest_urgent_to_seller_finish_time,omitempty" xml:"latest_urgent_to_seller_finish_time,omitempty"`
	// 最新平台承若用户完成时间
	LatestUrgentToBuyerPromiseTime string `json:"latest_urgent_to_buyer_promise_time,omitempty" xml:"latest_urgent_to_buyer_promise_time,omitempty"`
	// 最新催退时间
	LatestRequestDate string `json:"latest_request_date,omitempty" xml:"latest_request_date,omitempty"`
	// 催退次数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 是否触发了强制退款
	ForceRefund bool `json:"force_refund,omitempty" xml:"force_refund,omitempty"`
}

var poolUrgentRefundVo = sync.Pool{
	New: func() any {
		return new(UrgentRefundVo)
	},
}

// GetUrgentRefundVo() 从对象池中获取UrgentRefundVo
func GetUrgentRefundVo() *UrgentRefundVo {
	return poolUrgentRefundVo.Get().(*UrgentRefundVo)
}

// ReleaseUrgentRefundVo 释放UrgentRefundVo
func ReleaseUrgentRefundVo(v *UrgentRefundVo) {
	v.UrgentRefundHistory = v.UrgentRefundHistory[:0]
	v.LatestUrgentToSellerFinishTime = ""
	v.LatestUrgentToBuyerPromiseTime = ""
	v.LatestRequestDate = ""
	v.Count = 0
	v.ForceRefund = false
	poolUrgentRefundVo.Put(v)
}
