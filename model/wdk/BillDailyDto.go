package wdk

// BillDailyDto 结构体
type BillDailyDto struct {
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 订单来源
	OrderFrom string `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 明细单据唯一结算标识
	SettleBizId string `json:"settle_biz_id,omitempty" xml:"settle_biz_id,omitempty"`
	// 账单日期
	BillDate string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// 其它费用
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 商家应收金额，正号代表收入，负号代表支出
	SettleAmount *BigDecimal `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
	// 商品总金额
	SkuAmount *BigDecimal `json:"sku_amount,omitempty" xml:"sku_amount,omitempty"`
	// 计费基数
	ChargeBaseAmount *BigDecimal `json:"charge_base_amount,omitempty" xml:"charge_base_amount,omitempty"`
	// 商户应收包装费总额，正号代表收入，负号代表支出
	PackageAmount *BigDecimal `json:"package_amount,omitempty" xml:"package_amount,omitempty"`
	// 商户应收总运费，正号代表收入，负号代表支出
	SendAmount *BigDecimal `json:"send_amount,omitempty" xml:"send_amount,omitempty"`
	// 商户补贴总金额，正负号存在渠道差异
	MerchantSubsidyAmount *BigDecimal `json:"merchant_subsidy_amount,omitempty" xml:"merchant_subsidy_amount,omitempty"`
	// 平台营销补贴费用，正号代表收入，负号代表支出
	PlatSubsidyAmount *BigDecimal `json:"plat_subsidy_amount,omitempty" xml:"plat_subsidy_amount,omitempty"`
	// 品牌营销补贴费用，正号代表收入，负号代表支出
	BrandSubsidyAmount *BigDecimal `json:"brand_subsidy_amount,omitempty" xml:"brand_subsidy_amount,omitempty"`
	// 代理商营销补贴费用，正号代表收入，负号代表支出
	AgentSubsidyAmount *BigDecimal `json:"agent_subsidy_amount,omitempty" xml:"agent_subsidy_amount,omitempty"`
	// 技术服务费，正号代表收入，负号代表支出
	ChannelCommissionAmount *BigDecimal `json:"channel_commission_amount,omitempty" xml:"channel_commission_amount,omitempty"`
	// 基础物流费，正号代表收入，负号代表支出
	BaseLogisticsAmount *BigDecimal `json:"base_logistics_amount,omitempty" xml:"base_logistics_amount,omitempty"`
	// 增值服务费，正号代表收入，负号代表支出
	AddedValueAmount *BigDecimal `json:"added_value_amount,omitempty" xml:"added_value_amount,omitempty"`
	// 其它服务费，正号代表收入，负号代表支出
	OtherFeeAmount *BigDecimal `json:"other_fee_amount,omitempty" xml:"other_fee_amount,omitempty"`
	// 手续费，正号代表收入，负号代表支出
	HandleFeeAmount *BigDecimal `json:"handle_fee_amount,omitempty" xml:"handle_fee_amount,omitempty"`
	// 平台红包补贴
	PlatformVoucherSubsidyFee *BigDecimal `json:"platform_voucher_subsidy_fee,omitempty" xml:"platform_voucher_subsidy_fee,omitempty"`
	// 商家承担红包补贴费用
	MerchantVoucherSubsidyFee *BigDecimal `json:"merchant_voucher_subsidy_fee,omitempty" xml:"merchant_voucher_subsidy_fee,omitempty"`
	// 平台包装费
	PlatPackageFee *BigDecimal `json:"plat_package_fee,omitempty" xml:"plat_package_fee,omitempty"`
	// 用户实付金额
	UserPayAmount *BigDecimal `json:"user_pay_amount,omitempty" xml:"user_pay_amount,omitempty"`
}
