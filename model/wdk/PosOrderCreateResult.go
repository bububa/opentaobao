package wdk

// PosOrderCreateResult 
type PosOrderCreateResult struct {

    // 结果msg
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 结果码
    ReturnCode   string `json:"return_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // mainOrderId
    MainOrderId   int64 `json:"main_order_id,omitempty"`

    // outOrderId
    OutOrderId   string `json:"out_order_id,omitempty"`

    // subOrderDOList
    SubOrderDOList   []PosSubOrderResult `json:"sub_order_d_o_list,omitempty"`

}