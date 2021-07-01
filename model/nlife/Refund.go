package nlife

// Refund 结构体
type Refund struct {
	// refundTime
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// outRequestNo
	OutRequestNo string `json:"out_request_no,omitempty" xml:"out_request_no,omitempty"`
	// 退款渠道列表
	RefundBillList []FundBill `json:"refund_bill_list,omitempty" xml:"refund_bill_list>fund_bill,omitempty"`
	// 退货的商品，逗号分隔元素，商品和数量冒号分隔
	RefundGoods string `json:"refund_goods,omitempty" xml:"refund_goods,omitempty"`
}
