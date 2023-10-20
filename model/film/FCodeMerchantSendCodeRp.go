package film

// FcodeMerchantSendCodeRp 结构体
type FcodeMerchantSendCodeRp struct {
	// 码对外信息描述列表
	FCodeMerchantInfoList []FcodeMerchantVo `json:"f_code_merchant_info_list,omitempty" xml:"f_code_merchant_info_list>fcode_merchant_vo,omitempty"`
}
