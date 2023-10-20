package wdk

// TxdBillDetailBo 结构体
type TxdBillDetailBo struct {
	// 平台补贴
	PlatformSubsidyFee string `json:"platform_subsidy_fee,omitempty" xml:"platform_subsidy_fee,omitempty"`
	// 配送费
	DistributionFee string `json:"distribution_fee,omitempty" xml:"distribution_fee,omitempty"`
	// 手续费
	HandleFee string `json:"handle_fee,omitempty" xml:"handle_fee,omitempty"`
	// 技术服务费
	TechnicalServiceFee string `json:"technical_service_fee,omitempty" xml:"technical_service_fee,omitempty"`
	// 运费
	TransportFee string `json:"transport_fee,omitempty" xml:"transport_fee,omitempty"`
	// 计费基数（不含运费）,实付金额-运费
	ChargeBase string `json:"charge_base,omitempty" xml:"charge_base,omitempty"`
	// 实付金额（支付宝收款金额，即货款）
	PayAmount string `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 交易渠道(APP/POS)
	TradeChannel string `json:"trade_channel,omitempty" xml:"trade_channel,omitempty"`
	// 订单号,正向展示主单号，逆向展示退款单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 单据类型（正/逆向）
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 经营店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 经营店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 财务组织名称
	SettleCompanyName string `json:"settle_company_name,omitempty" xml:"settle_company_name,omitempty"`
	// 财务组织编码
	SettleCompanyCode string `json:"settle_company_code,omitempty" xml:"settle_company_code,omitempty"`
	// 商家名称
	MerchantName string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 业务日期,记录实际业务发生时间
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 账单日期, 账单出具日，按自然日
	BillDate string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// 账单编号
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 商家补贴
	MerchantSubsidyFee string `json:"merchant_subsidy_fee,omitempty" xml:"merchant_subsidy_fee,omitempty"`
	// 品牌商补贴
	BrandSubsidyFee string `json:"brand_subsidy_fee,omitempty" xml:"brand_subsidy_fee,omitempty"`
	// 商家应收金额
	ReceivableAmount string `json:"receivable_amount,omitempty" xml:"receivable_amount,omitempty"`
	// 原正向单号
	PTradeId string `json:"p_trade_id,omitempty" xml:"p_trade_id,omitempty"`
	// 结算公司名称
	SrcSettleCompanyName string `json:"src_settle_company_name,omitempty" xml:"src_settle_company_name,omitempty"`
	// 平台红包补贴
	PlatformVoucherSubsidyFee float64 `json:"platform_voucher_subsidy_fee,omitempty" xml:"platform_voucher_subsidy_fee,omitempty"`
	// 商家承担红包补贴费用
	MerchantVoucherSubsidyFee float64 `json:"merchant_voucher_subsidy_fee,omitempty" xml:"merchant_voucher_subsidy_fee,omitempty"`
	// 商家配送费补贴
	MerchantSendSubsidyFee float64 `json:"merchant_send_subsidy_fee,omitempty" xml:"merchant_send_subsidy_fee,omitempty"`
	// 渠道配送费补贴
	PlatSendSubsidyFee float64 `json:"plat_send_subsidy_fee,omitempty" xml:"plat_send_subsidy_fee,omitempty"`
	// 申诉赔付金
	DisputeFee float64 `json:"dispute_fee,omitempty" xml:"dispute_fee,omitempty"`
}
