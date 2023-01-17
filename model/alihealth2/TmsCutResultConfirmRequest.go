package alihealth2

// TmsCutResultConfirmRequest 结构体
type TmsCutResultConfirmRequest struct {
	// 回传数组
	ConfirmItems []TmsCutResultConfirmItem `json:"confirm_items,omitempty" xml:"confirm_items>tms_cut_result_confirm_item,omitempty"`
	// 扩展参数
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
}
