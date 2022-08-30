package wlbimports

// BigbagStatusResponse 结构体
type BigbagStatusResponse struct {
	// parcelOrderList
	ParcelOrderList []ParcelOrderStatusResponse `json:"parcel_order_list,omitempty" xml:"parcel_order_list>parcel_order_status_response,omitempty"`
	// handoverContentCode
	BigbagCode string `json:"bigbag_code,omitempty" xml:"bigbag_code,omitempty"`
	// 交接物状态，draft：草稿、committed：已提交、awaiting_tracking_number：等待分配运单号、awaiting_pickup：等待揽收、pickup：已揽收、pickup_failed：揽收失败、arrived：已到达、signed_normal：签收正常、signed_abnormal：签收异常、signed_failed：签收失败、canceled：已取消、cancel_failure：取消失败、canceling：取消中
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 更新时间（北京时间）
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// handoverContentId
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
}
