package moscm

// RmaDto 结构体
type RmaDto struct {
	// 商品明细
	RmaItems []RmaItemDto `json:"rma_items,omitempty" xml:"rma_items>rma_item_dto,omitempty"`
	// 快递单号
	ExpressNo string `json:"express_no,omitempty" xml:"express_no,omitempty"`
	// 快递公司名称
	ExpressName string `json:"express_name,omitempty" xml:"express_name,omitempty"`
	// 单据类型，可选值：GOODRETURN（退货）, GOODEXCHANGE（换货）,仅退款（&#34;RETURN&#34;）;
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 单据状态，可选值：CREATED(&#34;已创建&#34;),  INBOUND(&#34;已收货同意退款&#34;),  NOTINBOUND(&#34;已收货不同意退款&#34;),REFUNDED(&#34;已退款&#34;),COMPLETED(&#34;已完成&#34;),Obsolete(&#34;已作废&#34;);
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 最后更新时间
	LastUpdated string `json:"last_updated,omitempty" xml:"last_updated,omitempty"`
	// 退款时间
	RefundDate string `json:"refund_date,omitempty" xml:"refund_date,omitempty"`
	// 创建时间
	DateCreated string `json:"date_created,omitempty" xml:"date_created,omitempty"`
	// 退/换货原因详情
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 退/换货原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 商品寄回方式，可选值：RETURNTOSTORE（到店退）, EXPRESSDELIVERY(邮寄退)
	ReturnTheWay string `json:"return_the_way,omitempty" xml:"return_the_way,omitempty"`
	// 运费，单位:元
	Freight string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 金额，单位:元
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// 单据号
	RmaNumber string `json:"rma_number,omitempty" xml:"rma_number,omitempty"`
	// 退款资金信息
	Refunds string `json:"refunds,omitempty" xml:"refunds,omitempty"`
	// 供应商专柜Id
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 银泰专柜Id
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
}
