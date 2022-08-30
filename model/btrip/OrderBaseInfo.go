package btrip

// OrderBaseInfo 结构体
type OrderBaseInfo struct {
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 联系人
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 第三方申请单id
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 企业名称
	CorpName string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 申请单id
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 第三方行程id
	ThirdpartItineraryId string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// 申请事由
	BtripTitle string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	// 行程id
	ItineraryId string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// 用户所在部门id
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 用户所在部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 三方企业id
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 行程类型。0:单程，1:往返，2:中转
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}
