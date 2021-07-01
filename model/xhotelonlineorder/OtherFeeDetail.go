package xhotelonlineorder

// OtherFeeDetail 结构体
type OtherFeeDetail struct {
	// 无
	OtherFeeInfos []OtherFeeInfo `json:"other_fee_infos,omitempty" xml:"other_fee_infos>other_fee_info,omitempty"`
}
