package mos

// OnsiteRefundRequest 结构体
type OnsiteRefundRequest struct {
	// 业务扩展参数，json格式
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 商户操作员编号
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 商户退款请求号。标识一次退款请求，同一笔交易多次退款需要保证唯一
	OutRequestNo string `json:"out_request_no,omitempty" xml:"out_request_no,omitempty"`
	// 退款的原因说明
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 商户门店编号。可以是喵街内的商户门店ID，也可以是商户系统内自己的门店ID，其取值的含义由store_id_type定义
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 商户门店编号的类型。取值：miaojie和out。如果取值为miaojie，则store_id的取值为商户门店在喵街中的编号；如果取值为out，则store_id的取值为商户自己的编号
	StoreIdType string `json:"store_id_type,omitempty" xml:"store_id_type,omitempty"`
	// 商户机具终端编号
	TerminalId string `json:"terminal_id,omitempty" xml:"terminal_id,omitempty"`
	// 喵街商户号
	MjShopId string `json:"mj_shop_id,omitempty" xml:"mj_shop_id,omitempty"`
	// appId。兼容老退款
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 订单号。可能为外部订单号，也可能为喵街订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 退款来源.1：商家退款，2：用户主动退款，3：过期退款，4：线下POS退款
	RefundSource string `json:"refund_source,omitempty" xml:"refund_source,omitempty"`
	// 退款金额。需要退款的金额，该金额不能大于订单金额（同一笔交易多次退款时累计退款金额不能超过订单金额），单位为分
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}
