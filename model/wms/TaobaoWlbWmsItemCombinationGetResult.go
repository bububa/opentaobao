package wms

// TaobaowlbwmsitemcombinationgetResult 结构体
type TaobaowlbwmsitemcombinationgetResult struct {
	// 子货品列表
	SubItemList []SubItemList `json:"sub_item_list,omitempty" xml:"sub_item_list>sub_item_list,omitempty"`
	// 错误编码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 错误信息
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}
