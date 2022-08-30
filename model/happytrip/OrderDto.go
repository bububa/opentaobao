package happytrip

// OrderDto 结构体
type OrderDto struct {
	// 航旅交易号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 订单业务类型，各个业务自定义使用
	BizTypeDesc string `json:"biz_type_desc,omitempty" xml:"biz_type_desc,omitempty"`
	// 外系统下单时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 购买人用户id
	BuyerUserId string `json:"buyer_user_id,omitempty" xml:"buyer_user_id,omitempty"`
	// 获取购买人欢行ID的异常信息
	BuyerUserMemo string `json:"buyer_user_memo,omitempty" xml:"buyer_user_memo,omitempty"`
	// 购买人姓名
	BuyerUserName string `json:"buyer_user_name,omitempty" xml:"buyer_user_name,omitempty"`
	// 订单预计关闭时间
	CloseTime string `json:"close_time,omitempty" xml:"close_time,omitempty"`
	// 订单当前状态  -1, "未知状态"  1, "预定中"  2, "已取消"  3, "待付款"  4, "已付款"  5, "已删除"  6, "已完成"  7, "已关闭"  8, "已预订"  9, "已变更"  10, "预定失败"
	CurrentStatus string `json:"current_status,omitempty" xml:"current_status,omitempty"`
	// 订单当前状态描述 （参考current_status）
	CurrentStatusDesc string `json:"current_status_desc,omitempty" xml:"current_status_desc,omitempty"`
	// 订单过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 系统出票时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 订单一级类型描述 定义参考一级类型
	OrderClassName string `json:"order_class_name,omitempty" xml:"order_class_name,omitempty"`
	// 订单二级类型描述 定义参考一级类型
	OrderClassSecondName string `json:"order_class_second_name,omitempty" xml:"order_class_second_name,omitempty"`
	// 外系统订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外系统订单号状态
	OutOrderStatus string `json:"out_order_status,omitempty" xml:"out_order_status,omitempty"`
	// 外系统订单号状态描述
	OutOrderStatusDesc string `json:"out_order_status_desc,omitempty" xml:"out_order_status_desc,omitempty"`
	// 关联的差旅申请单的数据同步关联id
	OuterTravelHeadId string `json:"outer_travel_head_id,omitempty" xml:"outer_travel_head_id,omitempty"`
	// 支付状态  -1, "未知状态"  1, "未付款"  2, "待付款"  3, "已付款"  4, "支付超时"
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 支付状态描述
	PayStatusDesc string `json:"pay_status_desc,omitempty" xml:"pay_status_desc,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 支付类型
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 计划时间
	PlanTime string `json:"plan_time,omitempty" xml:"plan_time,omitempty"`
	// 订单来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 订单总价
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 总金额货币代码
	TotalPriceCurrencyCode string `json:"total_price_currency_code,omitempty" xml:"total_price_currency_code,omitempty"`
	// 总金额小数点位数
	TotalPriceDecimalPlaces string `json:"total_price_decimal_places,omitempty" xml:"total_price_decimal_places,omitempty"`
	// 是否B2G的标记位0不是 1是
	B2gFlag int64 `json:"b2g_flag,omitempty" xml:"b2g_flag,omitempty"`
	// 订单业务类型，各个业务自定义使用
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 订单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 线上线下标记0线下 1线上
	OnlineFlag int64 `json:"online_flag,omitempty" xml:"online_flag,omitempty"`
	// 订单一级类型 (一级类型，一级类型描述，二级类型，二级类型描述)  (-1, "未知", -1, "未知")  (4, "机票", 24, "飞猪国内机票")  (4, "机票", 47, "飞猪国际机票")  (4, "机票", 48, "飞猪国际询价单机票")  (4, "机票", 49, "GT国际机票")  (5, "酒店", 14, "飞猪国内酒店")  (5, "酒店", 34, "HRS国际酒店")  (6, "用车", 23, "国内用车")
	OrderClassId int64 `json:"order_class_id,omitempty" xml:"order_class_id,omitempty"`
	// 订单二级类型 定义参考一级类型
	OrderClassSecondId int64 `json:"order_class_second_id,omitempty" xml:"order_class_second_id,omitempty"`
	// 是否有效订单0无效，1有效
	ValidFlag int64 `json:"valid_flag,omitempty" xml:"valid_flag,omitempty"`
}
