package alsc

// OrderInfo 结构体
type OrderInfo struct {
	// 订单生效时间（外卖单）
	ActiveTime string `json:"active_time,omitempty" xml:"active_time,omitempty"`
	// 渠道
	BizChannel string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 商品信息
	ItemList []Item `json:"item_list,omitempty" xml:"item_list>item,omitempty"`
	// 订单备注,长度不超过256
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// pos点餐：POS_ORDER  扫码点餐：CODE_ORDER  线上点餐到店（美团）:MEITUAN_ONLINE  线上点餐到店（口碑）:KOUBEI_ONLINE  线上点餐到店（其他）:OTHER_ONLINE  到家外卖（美团）:MEITUAN_TAKEOUT  到家外卖（饿了么）:ELEME_TAKEOUT  其他:OTHER
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 订单状态:  WAIT_PAY("WAIT_PAY", "订单待支付"),  PAID("PAID", "订单已支付"),  SUCCESS("SUCCESS", "订单完成"),  CLOSED("CLOSED", "已关闭"),  REFUND("REFUND", "订单已退款");
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 外部买家基础信息
	OutBuyer *OrderUser `json:"out_buyer,omitempty" xml:"out_buyer,omitempty"`
	// 外部卖家基础信息
	OutSeller *OrderUser `json:"out_seller,omitempty" xml:"out_seller,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 门店地址
	ShopAddress string `json:"shop_address,omitempty" xml:"shop_address,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 履约方式
	PerformanceWay string `json:"performance_way,omitempty" xml:"performance_way,omitempty"`
	// 订单实收金额，单位分
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// 营业日，格式例如：“yyyy-mm-dd”
	BusinessDate string `json:"business_date,omitempty" xml:"business_date,omitempty"`
	// 业态：业务方直接传自己的业态code或id
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 业态：业务方直接传自己的业态文字描述，例如储值卡充值，储值卡售卡，押金单等
	BusinessTypeDesc string `json:"business_type_desc,omitempty" xml:"business_type_desc,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 设备ip
	DeviceIp string `json:"device_ip,omitempty" xml:"device_ip,omitempty"`
	// 取餐号
	DinnerNumber string `json:"dinner_number,omitempty" xml:"dinner_number,omitempty"`
	// 就餐人数
	DinnerPersons int64 `json:"dinner_persons,omitempty" xml:"dinner_persons,omitempty"`
	// 就餐形式:FOR_HERE堂食  TAKE_OUT外带  SEND外送  TO_GO自提  OTHER其他
	DinnerType string `json:"dinner_type,omitempty" xml:"dinner_type,omitempty"`
	// 订单渠道 HANDHELD_DEVICES手持设备 SAMPLE_ORDER扫码点餐 POS POS端 MSTORE 移动门店 THIRDPARTY第三方 OTHER 其他
	OrderChannel string `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 下单时间
	OrderTime int64 `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 收银员id
	OutCashierId string `json:"out_cashier_id,omitempty" xml:"out_cashier_id,omitempty"`
	// 收银员名称
	OutCashierName string `json:"out_cashier_name,omitempty" xml:"out_cashier_name,omitempty"`
	// 配送员id
	OutDeliveryId string `json:"out_delivery_id,omitempty" xml:"out_delivery_id,omitempty"`
	// 配送员名称
	OutDeliveryName string `json:"out_delivery_name,omitempty" xml:"out_delivery_name,omitempty"`
	// 订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 销售员id
	OutSalesmanId string `json:"out_salesman_id,omitempty" xml:"out_salesman_id,omitempty"`
	// 销售员名称
	OutSalesmanName string `json:"out_salesman_name,omitempty" xml:"out_salesman_name,omitempty"`
	// 卖家id
	OutSellerId string `json:"out_seller_id,omitempty" xml:"out_seller_id,omitempty"`
	// 店铺id
	OutStoreId string `json:"out_store_id,omitempty" xml:"out_store_id,omitempty"`
	// 外部会员id，业务方自己的会员id，如果有vipUserId的订单，这个字段可以为空
	OutVipUserId string `json:"out_vip_user_id,omitempty" xml:"out_vip_user_id,omitempty"`
	// 服务员id
	OutWaiterId string `json:"out_waiter_id,omitempty" xml:"out_waiter_id,omitempty"`
	// 服务员名称
	OutWaiterName string `json:"out_waiter_name,omitempty" xml:"out_waiter_name,omitempty"`
	// 订单优惠金额 不包含抹零金额，减钱为‘-’，加钱为‘+’，单位分
	PromoFee int64 `json:"promo_fee,omitempty" xml:"promo_fee,omitempty"`
	// 关联订单号，如反结账的原始业务订单号
	RelatedOutOrderNo string `json:"related_out_order_no,omitempty" xml:"related_out_order_no,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 整单抹零金额，单位分
	ScrapFee int64 `json:"scrap_fee,omitempty" xml:"scrap_fee,omitempty"`
	// 状态：SUCCESS已完成 REFUND已退款 REFUSED已拒绝 CANCELLED已取消 ANTI_SEELEMENT已反结账 CHARGED已挂账 CANCELLED_DEBT已销账
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 桌号
	TableNumber string `json:"table_number,omitempty" xml:"table_number,omitempty"`
	// 订单总金额，单位分
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 会员id，口碑crm的会员id，有vipUserId可以不传outVipUserId
	VipUserId string `json:"vip_user_id,omitempty" xml:"vip_user_id,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 附加费总金额，如餐盒，
	AttachTotalFee int64 `json:"attach_total_fee,omitempty" xml:"attach_total_fee,omitempty"`
	// 服务费总金额
	ServiceTotalFee int64 `json:"service_total_fee,omitempty" xml:"service_total_fee,omitempty"`
	// 税费总金额
	TaxTotalFee int64 `json:"tax_total_fee,omitempty" xml:"tax_total_fee,omitempty"`
	// 外部组织id，辰森和客如云直接填out_store_id
	OutOrgId string `json:"out_org_id,omitempty" xml:"out_org_id,omitempty"`
}
