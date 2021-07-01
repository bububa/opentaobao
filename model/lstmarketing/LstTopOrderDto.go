package lstmarketing

// LstTopOrderDto 结构体
type LstTopOrderDto struct {
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 营销活动列表
	PromotionDtoList []Promotiondtolist `json:"promotion_dto_list,omitempty" xml:"promotion_dto_list>promotiondtolist,omitempty"`
}
