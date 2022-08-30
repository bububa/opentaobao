package btrip

// BtripAccountPrestoreRs 结构体
type BtripAccountPrestoreRs struct {
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 账户信息返回
	Module *BtripAccountPrestoreRs `json:"module,omitempty" xml:"module,omitempty"`
	// 错误编码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 预存账户
	BtripAccountPrestoreRs *BtripAccountDetailRs `json:"btrip_account_prestore_rs,omitempty" xml:"btrip_account_prestore_rs,omitempty"`
	// 月结
	BtripCrediAmountRs *BtripCrediAmountRs `json:"btrip_credi_amount_rs,omitempty" xml:"btrip_credi_amount_rs,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
