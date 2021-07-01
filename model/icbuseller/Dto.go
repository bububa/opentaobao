package icbuseller

// Dto 结构体
type Dto struct {
	// 订单类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 订单编号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 订单状态
	CurrentStatus string `json:"current_status,omitempty" xml:"current_status,omitempty"`
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 状态更新时间
	FireTime string `json:"fire_time,omitempty" xml:"fire_time,omitempty"`
	// 成交价格
	TransactionPrice string `json:"transaction_price,omitempty" xml:"transaction_price,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 成交单价
	TransactionUnitPrice string `json:"transaction_unit_price,omitempty" xml:"transaction_unit_price,omitempty"`
	// 支付通道
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 订单货币类型
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 退款金额
	RefundPrice string `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	// 订单标题
	OrderTitle string `json:"order_title,omitempty" xml:"order_title,omitempty"`
	// 规格编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 购买者阿里id
	BuyerAliId int64 `json:"buyer_ali_id,omitempty" xml:"buyer_ali_id,omitempty"`
	// 购买者登陆账号
	BuyerLoginId string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
	// 服务分类id
	ServiceCategoryId int64 `json:"service_category_id,omitempty" xml:"service_category_id,omitempty"`
	// 服务分类名称
	ServiceCategory string `json:"service_category,omitempty" xml:"service_category,omitempty"`
	// 用户购买sku
	ServiceSkuLabel string `json:"service_sku_label,omitempty" xml:"service_sku_label,omitempty"`
}
