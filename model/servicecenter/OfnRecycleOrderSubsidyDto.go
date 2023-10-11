package servicecenter

// OfnRecycleOrderSubsidyDto 结构体
type OfnRecycleOrderSubsidyDto struct {
	// 旧品主订单号
	OldOrderId string `json:"old_order_id,omitempty" xml:"old_order_id,omitempty"`
	// 新品主订单号
	NewOrderId string `json:"new_order_id,omitempty" xml:"new_order_id,omitempty"`
	// 新品补贴状态
	SubsidyStatus string `json:"subsidy_status,omitempty" xml:"subsidy_status,omitempty"`
	// 关联的新机单状态
	NewOrderPayStatus string `json:"new_order_pay_status,omitempty" xml:"new_order_pay_status,omitempty"`
	// 关联新机单店铺名称
	NewOrderShopName string `json:"new_order_shop_name,omitempty" xml:"new_order_shop_name,omitempty"`
	// 拆机费结算使用的关联新机单ID
	KitchenApplianceSettleNewOrderId string `json:"kitchen_appliance_settle_new_order_id,omitempty" xml:"kitchen_appliance_settle_new_order_id,omitempty"`
	// 旧品款预付款
	OldItemPreRedPacket *ExternalPreRedPacketDto `json:"old_item_pre_red_packet,omitempty" xml:"old_item_pre_red_packet,omitempty"`
	// 旧品款尾款
	OldItemEndRedPacket *ExternalEndRedPacketDto `json:"old_item_end_red_packet,omitempty" xml:"old_item_end_red_packet,omitempty"`
	// 旧品补贴红包
	OldItemSubsidyRedPacket *ExternalRedPacketDto `json:"old_item_subsidy_red_packet,omitempty" xml:"old_item_subsidy_red_packet,omitempty"`
	// 回收商活动（限时活动）信息
	SupplyActivity *ActivitySummaryDto `json:"supply_activity,omitempty" xml:"supply_activity,omitempty"`
	// 新品卖家补贴金额，单位：分
	NewItemSellerBenefitAmount int64 `json:"new_item_seller_benefit_amount,omitempty" xml:"new_item_seller_benefit_amount,omitempty"`
	// 新品平台补贴金额，单位：分
	NewItemPlatformBenefitAmount int64 `json:"new_item_platform_benefit_amount,omitempty" xml:"new_item_platform_benefit_amount,omitempty"`
	// 新品回收商补贴金额，单位：分
	NewItemSupplyBenefitAmount int64 `json:"new_item_supply_benefit_amount,omitempty" xml:"new_item_supply_benefit_amount,omitempty"`
}
