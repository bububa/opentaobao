package wdk

// BillDetailDto 结构体
type BillDetailDto struct {
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道
	OrderFrom string `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 明细单据唯一结算标识
	SettleBizId string `json:"settle_biz_id,omitempty" xml:"settle_biz_id,omitempty"`
	// 账单日期
	BillDate string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// 其它费用
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 账单类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 渠道订单号
	ChannelOrderId string `json:"channel_order_id,omitempty" xml:"channel_order_id,omitempty"`
	// 商家应收金额，正号代表收入，负号代表支出
	SettleAmount float64 `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
	// 商品总金额
	SkuAmount float64 `json:"sku_amount,omitempty" xml:"sku_amount,omitempty"`
	// 计费基数
	ChargeBaseAmount float64 `json:"charge_base_amount,omitempty" xml:"charge_base_amount,omitempty"`
	// 商户应收包装费总额，正号代表收入，负号代表支出
	PackageAmount float64 `json:"package_amount,omitempty" xml:"package_amount,omitempty"`
	// 商户应收总运费，正号代表收入，负号代表支出
	SendAmount float64 `json:"send_amount,omitempty" xml:"send_amount,omitempty"`
	// 商户补贴总金额，正负号存在渠道差异
	MerchantSubsidyAmount float64 `json:"merchant_subsidy_amount,omitempty" xml:"merchant_subsidy_amount,omitempty"`
	// 平台营销补贴费用，正号代表收入，负号代表支出
	PlatSubsidyAmount float64 `json:"plat_subsidy_amount,omitempty" xml:"plat_subsidy_amount,omitempty"`
	// 品牌营销补贴费用，正号代表收入，负号代表支出
	BrandSubsidyAmount float64 `json:"brand_subsidy_amount,omitempty" xml:"brand_subsidy_amount,omitempty"`
	// 代理商营销补贴费用，正号代表收入，负号代表支出
	AgentSubsidyAmount float64 `json:"agent_subsidy_amount,omitempty" xml:"agent_subsidy_amount,omitempty"`
	// 技术服务费，正号代表收入，负号代表支出
	ChannelCommissionAmount float64 `json:"channel_commission_amount,omitempty" xml:"channel_commission_amount,omitempty"`
	// 基础物流费，正号代表收入，负号代表支出
	BaseLogisticsAmount float64 `json:"base_logistics_amount,omitempty" xml:"base_logistics_amount,omitempty"`
	// 增值服务费，正号代表收入，负号代表支出
	AddedValueAmount float64 `json:"added_value_amount,omitempty" xml:"added_value_amount,omitempty"`
	// 其它服务费，正号代表收入，负号代表支出
	OtherFeeAmount float64 `json:"other_fee_amount,omitempty" xml:"other_fee_amount,omitempty"`
	// 手续费，正号代表收入，负号代表支出
	HandleFeeAmount float64 `json:"handle_fee_amount,omitempty" xml:"handle_fee_amount,omitempty"`
	// 平台红包补贴
	PlatformVoucherSubsidyFee float64 `json:"platform_voucher_subsidy_fee,omitempty" xml:"platform_voucher_subsidy_fee,omitempty"`
	// 商家承担红包补贴费用
	MerchantVoucherSubsidyFee float64 `json:"merchant_voucher_subsidy_fee,omitempty" xml:"merchant_voucher_subsidy_fee,omitempty"`
	// 平台包装费
	PlatPackageFee float64 `json:"plat_package_fee,omitempty" xml:"plat_package_fee,omitempty"`
	// 用户实付金额
	UserPayAmount float64 `json:"user_pay_amount,omitempty" xml:"user_pay_amount,omitempty"`
}
