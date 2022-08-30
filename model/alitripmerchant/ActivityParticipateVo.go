package alitripmerchant

// ActivityParticipateVo 结构体
type ActivityParticipateVo struct {
	// 奖品名称
	GoodName string `json:"good_name,omitempty" xml:"good_name,omitempty"`
	// 奖品图片
	GoodImage string `json:"good_image,omitempty" xml:"good_image,omitempty"`
	// 奖品发放类型
	GoodSendType int64 `json:"good_send_type,omitempty" xml:"good_send_type,omitempty"`
	// 活动参与状态，中奖，未中奖
	ParticipateStatus int64 `json:"participate_status,omitempty" xml:"participate_status,omitempty"`
	// 奖品类型
	GoodType int64 `json:"good_type,omitempty" xml:"good_type,omitempty"`
	// 奖品id
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 奖品数量
	GoodAmount int64 `json:"good_amount,omitempty" xml:"good_amount,omitempty"`
	// 参与结果记录id
	RecordId int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
}
