package ascp

// OmsComplaintWorkCardStatusModifyParameter 结构体
type OmsComplaintWorkCardStatusModifyParameter struct {
	// 投诉工单id
	ComplaintWorkcardId string `json:"complaint_workcard_id,omitempty" xml:"complaint_workcard_id,omitempty"`
	// BFC履行单号
	WdsCoordinationOrderId string `json:"wds_coordination_order_id,omitempty" xml:"wds_coordination_order_id,omitempty"`
	// 工单回复(处理中需要填写)      * 完结内容(已完结需要填写)
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 状态0 - 待收款   1 - 待补充举证    2 - 处理中  3 - 已完结
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
