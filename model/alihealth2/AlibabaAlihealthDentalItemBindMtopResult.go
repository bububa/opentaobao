package alihealth2

// AlibabaAlihealthDentalItemBindMtopResult 结构体
type AlibabaAlihealthDentalItemBindMtopResult struct {
	// model
	StoreItemRelVoList []StoreItemRelVo `json:"store_item_rel_vo_list,omitempty" xml:"store_item_rel_vo_list>store_item_rel_vo,omitempty"`
	// 200
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
