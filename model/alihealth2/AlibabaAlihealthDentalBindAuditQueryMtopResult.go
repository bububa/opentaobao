package alihealth2

// AlibabaAlihealthDentalBindAuditQueryMtopResult 结构体
type AlibabaAlihealthDentalBindAuditQueryMtopResult struct {
	// model
	StoreitemrelvoList []StoreItemRelVo `json:"storeitemrelvo_list,omitempty" xml:"storeitemrelvo_list>store_item_rel_vo,omitempty"`
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
