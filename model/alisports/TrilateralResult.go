package alisports

// TrilateralResult 结构体
type TrilateralResult struct {
	// 系统自动生成
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 系统自动生成
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回alipayId以及补助金额
	Value *SubsidyAmountVo `json:"value,omitempty" xml:"value,omitempty"`
	// 系统自动生成
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
