package btrip

// MultiCabinClassInfo 结构体
type MultiCabinClassInfo struct {
	// 退改签规则列表
	FlightRuleList []FlightRule `json:"flight_rule_list,omitempty" xml:"flight_rule_list>flight_rule,omitempty"`
	// 必填项,剩余座位数
	RemainedSeatCount string `json:"remained_seat_count,omitempty" xml:"remained_seat_count,omitempty"`
	// 必填项,舱位代码
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱位名称，公务舱
	ClassName string `json:"class_name,omitempty" xml:"class_name,omitempty"`
	// 必填项,舱等
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 子舱位等级展示用名称，超级经济舱
	CabinClassName string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// 必填项,折扣
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 优惠金额
	PromotionPrice string `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// {&#34;key&#34;:&#34;value&#34;} 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 加密参数,包含agentId 登下单需要用到的参数信息
	OrderParams string `json:"order_params,omitempty" xml:"order_params,omitempty"`
	// 退改签描述
	ClassRule string `json:"class_rule,omitempty" xml:"class_rule,omitempty"`
	// 标准/非标准产品
	ProductTypeDesc string `json:"product_type_desc,omitempty" xml:"product_type_desc,omitempty"`
	// 子舱位代码
	ChildCabin string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
	// 商品id
	OtaItemId string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// 退改签规则列表json
	FlightRuleListStr string `json:"flight_rule_list_str,omitempty" xml:"flight_rule_list_str,omitempty"`
	// 必填项,票面价（分）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 销售价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 产品类型
	ProductType int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 行程单类型
	InvoiceType int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 燃油费
	OilPrice int64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// 基建费
	BuildPrice int64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// 总价=销售价+基建+燃油费
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 是否协议价
	IsProtocol bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
}
