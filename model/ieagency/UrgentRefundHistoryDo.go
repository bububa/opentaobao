package ieagency

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
