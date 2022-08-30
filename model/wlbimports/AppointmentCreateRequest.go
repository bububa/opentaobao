package wlbimports

// AppointmentCreateRequest 结构体
type AppointmentCreateRequest struct {
	// 预约大包信息列表
	HandoverContentSynopsisList []OrderHandoverContentSynopsisDto `json:"handover_content_synopsis_list,omitempty" xml:"handover_content_synopsis_list>order_handover_content_synopsis_dto,omitempty"`
	// 时区
	ZoneOffSet string `json:"zone_off_set,omitempty" xml:"zone_off_set,omitempty"`
	// 接收仓资源名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 揽收方式：TRUCK（卡车）
	PickupType string `json:"pickup_type,omitempty" xml:"pickup_type,omitempty"`
	// 接收仓资源编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 发货人信息
	SenderInfo *ContactInfoRequest `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 收件人
	ReceiverInfo *ContactInfoRequest `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 揽收信息
	PickupInfo *PickupInfo `json:"pickup_info,omitempty" xml:"pickup_info,omitempty"`
}
