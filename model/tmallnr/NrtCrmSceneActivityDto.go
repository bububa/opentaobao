package tmallnr

// NrtCrmSceneActivityDto 结构体
type NrtCrmSceneActivityDto struct {
	// 下挂模板DTO
	NrtCrmBenefitList []NrtCrmBenefitDto `json:"nrt_crm_benefit_list,omitempty" xml:"nrt_crm_benefit_list>nrt_crm_benefit_dto,omitempty"`
	// 有价礼包类型Type
	TmpType string `json:"tmp_type,omitempty" xml:"tmp_type,omitempty"`
	// 有价礼包商品价格
	ReservePrice int64 `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 有价礼包商品id
	SellGiftItemId int64 `json:"sell_gift_item_id,omitempty" xml:"sell_gift_item_id,omitempty"`
	// 有价礼包活动状态
	SceneStatus int64 `json:"scene_status,omitempty" xml:"scene_status,omitempty"`
	// 有价礼包ID
	SceneActivityId int64 `json:"scene_activity_id,omitempty" xml:"scene_activity_id,omitempty"`
}
