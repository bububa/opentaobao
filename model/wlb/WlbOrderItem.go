package wlb

// WlbOrderItem 结构体
type WlbOrderItem struct {
	// 物流宝订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 订单商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 订单商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 物流宝订单确认状态：<br/>NO_NEED_CONFIRM--不需要确认<br/>WAIT_CONFIRM--待确认<br/>CONFIRMED--已确认
	ConfirmStatus string `json:"confirm_status,omitempty" xml:"confirm_status,omitempty"`
	// 订单商品用户昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// (1)其它: OTHER (2)淘宝交易: TAOBAO (3)调拨: ALLOCATION (4)盘点:CHECK (5)销售采购:PRUCHASE(6)其它交易：OTHER_TRADE
	OrderSubType string `json:"order_sub_type,omitempty" xml:"order_sub_type,omitempty"`
	// 订单号
	OrderSubCode string `json:"order_sub_code,omitempty" xml:"order_sub_code,omitempty"`
	// 子交易号
	OrderSub2code string `json:"order_sub_2code,omitempty" xml:"order_sub_2code,omitempty"`
	// 货主nick
	ProviderTpNick string `json:"provider_tp_nick,omitempty" xml:"provider_tp_nick,omitempty"`
	// 订单商品备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 批次号备注
	BatchRemark string `json:"batch_remark,omitempty" xml:"batch_remark,omitempty"`
	// 订单商品图片url
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// 外部实体类型
	ExtEntityType string `json:"ext_entity_type,omitempty" xml:"ext_entity_type,omitempty"`
	// 外部实体id
	ExtEntityId string `json:"ext_entity_id,omitempty" xml:"ext_entity_id,omitempty"`
	// INVENTORY_TYPE_SELL 可销库存<br/>INVENTORY_TYPE_IMPERFECTIONS 残次库存<br/>INVENTORY_TYPE_FREEZE 冻结库存<br/>INVENTORY_TYPE_ON_PASSAGE 在途库存<br/>INVENTORY_TYPE_ENGINE_DAMAGE 机损<br/>INVENTORY_TYPE_BOX_DAMAGE 箱损
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 计划数量
	PlanQuantity int64 `json:"plan_quantity,omitempty" xml:"plan_quantity,omitempty"`
	// 实际数量
	RealQuantity int64 `json:"real_quantity,omitempty" xml:"real_quantity,omitempty"`
	// 物流宝订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单商品id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单商品用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 物流宝订单商品的物流宝商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货主id
	ProviderTpId int64 `json:"provider_tp_id,omitempty" xml:"provider_tp_id,omitempty"`
	// 商品价格
	ItemPrice int64 `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 商品发布版本号
	PublishVersion int64 `json:"publish_version,omitempty" xml:"publish_version,omitempty"`
	// 第一位标示是否为赠品
	Flag int64 `json:"flag,omitempty" xml:"flag,omitempty"`
}
