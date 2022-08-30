package alitripmerchant

// ActivityDrawUserGoodsVo 结构体
type ActivityDrawUserGoodsVo struct {
	// 奖品图片
	Image string `json:"image,omitempty" xml:"image,omitempty"`
	// 奖品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 记录ID
	RecordId string `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 奖品ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 活动ID
	OfferId int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 奖品状态
	GoodsStatus int64 `json:"goods_status,omitempty" xml:"goods_status,omitempty"`
	// 奖品邮寄信息
	ActivityDrawGoodsPostInfo *ActivityDrawGoodsPostInfoVo `json:"activity_draw_goods_post_info,omitempty" xml:"activity_draw_goods_post_info,omitempty"`
}
