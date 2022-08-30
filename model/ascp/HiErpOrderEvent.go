package ascp

// HiErpOrderEvent 结构体
type HiErpOrderEvent struct {
	// 货品明细
	HiErpItemDtoList []HiErpItemDto `json:"hi_erp_item_dto_list,omitempty" xml:"hi_erp_item_dto_list>hi_erp_item_dto,omitempty"`
	// 主单信息
	HiErpOrderDto *HiErpOrderDto `json:"hi_erp_order_dto,omitempty" xml:"hi_erp_order_dto,omitempty"`
	// 收件人信息
	HiErpReceiverDto *HiErpReceiverDto `json:"hi_erp_receiver_dto,omitempty" xml:"hi_erp_receiver_dto,omitempty"`
}
