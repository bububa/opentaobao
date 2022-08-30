package alsc

// AlscConsumeRecordOpenInfo 结构体
type AlscConsumeRecordOpenInfo struct {
	// 支付单模型
	PaymentOpenInfos []PaymentOpenInfo `json:"payment_open_infos,omitempty" xml:"payment_open_infos>payment_open_info,omitempty"`
	// 退款单模型
	RefundOpenInfos []RefundOpenInfo `json:"refund_open_infos,omitempty" xml:"refund_open_infos>refund_open_info,omitempty"`
	// 子单模型
	SubOrderOpenInfos []SubOrderOpenInfo `json:"sub_order_open_infos,omitempty" xml:"sub_order_open_infos>sub_order_open_info,omitempty"`
	// 消费记录业务上下文
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
	// 业务单据号
	BizNo string `json:"biz_no,omitempty" xml:"biz_no,omitempty"`
	// 业务单据类型
	BizNoType string `json:"biz_no_type,omitempty" xml:"biz_no_type,omitempty"`
	// 业务方
	BizSource string `json:"biz_source,omitempty" xml:"biz_source,omitempty"`
	// 消费记录业务状态
	BizStatus string `json:"biz_status,omitempty" xml:"biz_status,omitempty"`
	// 业务子类型
	BizSubType string `json:"biz_sub_type,omitempty" xml:"biz_sub_type,omitempty"`
	// 业务类型
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 消费记录扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 业务创建时间
	GmtBizCreate string `json:"gmt_biz_create,omitempty" xml:"gmt_biz_create,omitempty"`
	// 业务最后修改时间
	GmtBizModified string `json:"gmt_biz_modified,omitempty" xml:"gmt_biz_modified,omitempty"`
	// 对侧ID
	OppositeId string `json:"opposite_id,omitempty" xml:"opposite_id,omitempty"`
	// 对侧ID类型
	OppositeIdType string `json:"opposite_id_type,omitempty" xml:"opposite_id_type,omitempty"`
	// 评价状态
	RateStatus string `json:"rate_status,omitempty" xml:"rate_status,omitempty"`
	// 消费记录类型
	RecordType string `json:"record_type,omitempty" xml:"record_type,omitempty"`
	// 物流单模型
	DeliveryOpenInfo *DeliveryOpenInfo `json:"delivery_open_info,omitempty" xml:"delivery_open_info,omitempty"`
	// 主单模型
	OrderOpenInfo *OrderOpenInfo `json:"order_open_info,omitempty" xml:"order_open_info,omitempty"`
	// 店铺模型
	ShopOpenInfo *ShopOpenInfo `json:"shop_open_info,omitempty" xml:"shop_open_info,omitempty"`
	// 订单是否可见
	Visibility bool `json:"visibility,omitempty" xml:"visibility,omitempty"`
}
