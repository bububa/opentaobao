package alitripmerchant

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryResponse struct {
	// 兑换结果
	Content []DerbyVoucherCardRedeemHistoryVo `json:"content,omitempty" xml:"content>derby_voucher_card_redeem_history_vo,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
