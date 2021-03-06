package smartstore

// TmallPopupstoreItemDiscountPriceResultDto 结构体
type TmallPopupstoreItemDiscountPriceResultDto struct {
	// 错误码code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 实际结果
	ResultList string `json:"result_list,omitempty" xml:"result_list,omitempty"`
	// 数据条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
