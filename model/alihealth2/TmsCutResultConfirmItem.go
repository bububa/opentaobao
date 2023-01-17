package alihealth2

// TmsCutResultConfirmItem 结构体
type TmsCutResultConfirmItem struct {
	// 运单号
	TmsCodes []string `json:"tms_codes,omitempty" xml:"tms_codes>string,omitempty"`
	// 扩展参数
	ExtendItem string `json:"extend_item,omitempty" xml:"extend_item,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 拦截状态 2 拦截成功 3拦截失败
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
