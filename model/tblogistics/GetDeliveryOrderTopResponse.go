package tblogistics

// GetDeliveryOrderTopResponse 结构体
type GetDeliveryOrderTopResponse struct {
	// 商品清单
	ItemList []ItemTopDto `json:"item_list,omitempty" xml:"item_list>item_top_dto,omitempty"`
	// 业务流水号
	CwOrderId string `json:"cw_order_id,omitempty" xml:"cw_order_id,omitempty"`
	// 物流单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 取件码
	PickupCode string `json:"pickup_code,omitempty" xml:"pickup_code,omitempty"`
	// ERP单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 配送员姓名
	DeliveryName string `json:"delivery_name,omitempty" xml:"delivery_name,omitempty"`
	// 配送员电话
	DeliveryPhone string `json:"delivery_phone,omitempty" xml:"delivery_phone,omitempty"`
	// 扩展
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 业务类型，INSTANT_ONLINE：同城配送-在线下单
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 状态，INIT(初始化),CREATED(已创建),ACCEPTED(骑手已接单),SHIPPING(配送中),SUCCESS(妥投/签收),CANCELED(已取消),ABNORMAL(配送异常),REFUSED(拒收)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝订主单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 发货联系人
	Sender *SenderTopDto `json:"sender,omitempty" xml:"sender,omitempty"`
	// 选择的物流资源
	SelectedResource *SelectedResourceTopDto `json:"selected_resource,omitempty" xml:"selected_resource,omitempty"`
	// 商品总价（原价），单位分，默认：商品单价总和
	TotalItemValue int64 `json:"total_item_value,omitempty" xml:"total_item_value,omitempty"`
	// 商品实付总价（总价），单位分
	TotalItemActualValue int64 `json:"total_item_actual_value,omitempty" xml:"total_item_actual_value,omitempty"`
	// 总重量，单位KG
	TotalWeight int64 `json:"total_weight,omitempty" xml:"total_weight,omitempty"`
}
