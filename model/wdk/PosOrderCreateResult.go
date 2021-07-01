package wdk

// PosOrderCreateResult 结构体
type PosOrderCreateResult struct {
	// 结果msg
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 结果码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// mainOrderId
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// outOrderId
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// subOrderDOList
	SubOrderDOList []PosSubOrderResult `json:"sub_order_d_o_list,omitempty" xml:"sub_order_d_o_list>pos_sub_order_result,omitempty"`
}
