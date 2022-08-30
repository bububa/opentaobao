package aliexpress

// AeopWarehouseAbnormalOrderResult 结构体
type AeopWarehouseAbnormalOrderResult struct {
	// 解决方案_仓编码
	SchemeCode string `json:"scheme_code,omitempty" xml:"scheme_code,omitempty"`
	// 解决方案名称
	SolutionName string `json:"solution_name,omitempty" xml:"solution_name,omitempty"`
	// 国内快递公司
	DomesticLogisticCompanyName string `json:"domestic_logistic_company_name,omitempty" xml:"domestic_logistic_company_name,omitempty"`
	// 状态变更时间
	GmtStatusUpdate string `json:"gmt_status_update,omitempty" xml:"gmt_status_update,omitempty"`
	// 异常原因
	AbnormalReason string `json:"abnormal_reason,omitempty" xml:"abnormal_reason,omitempty"`
	// 异常编码
	AbnormalCode string `json:"abnormal_code,omitempty" xml:"abnormal_code,omitempty"`
	// 运单号
	IntlTrackingNo string `json:"intl_tracking_no,omitempty" xml:"intl_tracking_no,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 揽收仓名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 解决方案编码
	SolutionCode string `json:"solution_code,omitempty" xml:"solution_code,omitempty"`
	// 揽收仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 物流详情链接
	LogisticDetailUrl string `json:"logistic_detail_url,omitempty" xml:"logistic_detail_url,omitempty"`
	// 订单状态名称
	OrderStatusName string `json:"order_status_name,omitempty" xml:"order_status_name,omitempty"`
	// 取消状态编码
	CancelStatusCode string `json:"cancel_status_code,omitempty" xml:"cancel_status_code,omitempty"`
	// 支付状态编码
	PaymentStatusCode string `json:"payment_status_code,omitempty" xml:"payment_status_code,omitempty"`
	// 退回运单号
	UndExpressMailNo string `json:"und_express_mail_no,omitempty" xml:"und_express_mail_no,omitempty"`
	// 禁运审核结果
	ForbiddenAuditResult string `json:"forbidden_audit_result,omitempty" xml:"forbidden_audit_result,omitempty"`
	// 订单状态编码
	OrderStatusCode string `json:"order_status_code,omitempty" xml:"order_status_code,omitempty"`
	// 取消状态名称
	CancelStatusName string `json:"cancel_status_name,omitempty" xml:"cancel_status_name,omitempty"`
	// 支付状态名称
	PaymentStatusName string `json:"payment_status_name,omitempty" xml:"payment_status_name,omitempty"`
	// 国内运单号
	DomesticLogisticTrackingNo string `json:"domestic_logistic_tracking_no,omitempty" xml:"domestic_logistic_tracking_no,omitempty"`
	// 交易单号
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 不可达保险
	UnreachableInsured bool `json:"unreachable_insured,omitempty" xml:"unreachable_insured,omitempty"`
	// 高货值保险
	HighValueInsure bool `json:"high_value_insure,omitempty" xml:"high_value_insure,omitempty"`
	// 逆袭高货值保险
	ReHighValueInsure bool `json:"re_high_value_insure,omitempty" xml:"re_high_value_insure,omitempty"`
}
