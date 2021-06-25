package wdk

// CreateReverseResponse 
type CreateReverseResponse struct {

    // tp单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 请求id
    RequestId   string `json:"request_id,omitempty"`

    // 逆向单ids
    ReverseIds   []Number `json:"reverse_ids,omitempty"`

    // 门店id
    StoreId   string `json:"store_id,omitempty"`

    // 外部单号
    OutBizOrderIds   []String `json:"out_biz_order_ids,omitempty"`

}
