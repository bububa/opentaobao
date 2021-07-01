package wdk

// ItemPoolActivityElementOperateRequest 结构体
type ItemPoolActivityElementOperateRequest struct {
	// 商品元素列表
	SkuElements []ItemPoolSkuActivityElementDto `json:"sku_elements,omitempty" xml:"sku_elements>item_pool_sku_activity_element_dto,omitempty"`
	// 同城零售活动id
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 操作人id
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 操作人名称
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 商品所属分组序号，默认单个分组则填1
	GroupNumber int64 `json:"group_number,omitempty" xml:"group_number,omitempty"`
}
